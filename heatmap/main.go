package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"strconv"
)

func main() {
	fName := flag.String("f", "data.csv", "CSV File to load")
	flag.Parse()

	// Read Data
	hm := NewHeatMapFromCSVFile(*fName)

	xSize, ySize := 640, 480
	img := CreateImage(xSize, ySize)

	DrawGrid(img, hm)

	SaveImage(img)
}

type HeatMap [][]int

func NewHeatMapFromStringMatrix(bs [][]string) HeatMap {
	// Validate

	// Convert [][]string to [][]int
	hm := [][]int{}

	for j := 0; j < len(bs); j++ {
		currRow := []int{}
		for i := 0; i < len(bs[j]); i++ {
			// Make sure width is constant

			// Make sure the string value is actually an int

			// For now, input will be floats, so do string -> float -> int
			// until I fix heatmap to work with floats instead of ints (should be simple)
			vf, err := strconv.ParseFloat(bs[j][i], 64)
			if err != nil {
				panic("failed to parse string into integer")
			}
			v := int(vf)

			// v, err := strconv.Atoi(bs[j][i])
			// if err != nil {
			// 	panic("failed to parse string into integer")
			// }

			// Add to current row
			currRow = append(currRow, v)
		}

		hm = append(hm, currRow)
	}

	return HeatMap(hm)
}

func NewHeatMapFromCSVFile(fName string) HeatMap {
	f, err := os.Open(fName)
	if err != nil {
		panic("failed to open csv file")
	}

	r := csv.NewReader(f)
	bs, err := r.ReadAll()
	if err != nil {
		panic("failed to read csv file")
	}

	return NewHeatMapFromStringMatrix(bs)
}

func (hm HeatMap) Get(i, j int) int {
	return hm[j][i]
}

func (hm HeatMap) Width() int {
	return len(hm[0])
}

func (hm HeatMap) Height() int {
	return len(hm)
}

func (hm HeatMap) MaxValue() int {
	max := hm.Get(0, 0)
	for i := 0; i < hm.Width(); i++ {
		for j := 0; j < hm.Height(); j++ {
			if tmp := hm.Get(i, j); tmp > max {
				max = tmp
			}
		}
	}
	return max
}

func (hm HeatMap) Validate() error {
	height := len(hm)
	if height == 0 {
		return errors.New("HeatMap is empty")
	}

	width := len(hm[0])
	if width == 0 {
		return errors.New("HeatMap is empty")
	}

	for j := 1; j < height; j++ {
		if len(hm[j]) != width {
			return errors.New("HeatMap must have rows of same size")
		}
	}

	return nil
}

func DrawGrid(m *image.RGBA, hm HeatMap) {
	if err := hm.Validate(); err != nil {
		log.Fatal(err)
	}

	hmXSize, hmYSize := len(hm[0]), len(hm)
	imgWidth, imgHeight := m.Bounds().Size().X, m.Bounds().Size().Y
	gridXSize, gridYSize := imgWidth/hmXSize, imgHeight/hmYSize

	pDelta := Point{gridXSize, gridYSize}

	hmvc := NewHeatMapValueToColorConverter(hm)

	for i := 0; i < hmXSize; i++ {
		for j := 0; j < hmYSize; j++ {
			p := Point{i * gridXSize, j * gridYSize}

			// TODO: If it's the last row/column, stretch pDelta to the end of the image
			// This will avoid the unpainted strip that will occur otherwise

			DrawRectByDelta(m, p, pDelta, hmvc.ValueToColor(hm.Get(i, j)))
		}
	}
}

type HeatMapValueToColorConverter struct {
	hm  HeatMap
	max int
}

func NewHeatMapValueToColorConverter(hm HeatMap) *HeatMapValueToColorConverter {
	hmvc := &HeatMapValueToColorConverter{
		hm:  hm,
		max: hm.MaxValue(),
	}

	return hmvc
}

func (hmvc HeatMapValueToColorConverter) ValueToColor(v int) color.RGBA {
	// Scale value to be between 0 and hmvc.max

	m := 255.0 / float64(hmvc.max)
	cVal := uint8(m * float64(v))

	return color.RGBA{cVal, cVal, cVal, 255}
}

// Drawing Utils

type Point struct {
	X, Y int
}

func (p0 Point) Add(p1 Point) Point {
	return Point{p0.X + p1.X, p0.Y + p1.Y}
}

func DrawRect(m *image.RGBA, p0, p1 Point, c color.Color) {
	for i := p0.X; i < p1.X; i++ {
		for j := p0.Y; j < p1.Y; j++ {
			m.Set(i, j, c)
		}
	}
}

func DrawRectByDelta(m *image.RGBA, p, d Point, c color.Color) {
	DrawRect(m, p, p.Add(d), c)
}

func DrawSquare(m *image.RGBA, p Point, d int, c color.Color) {
	DrawRectByDelta(m, p, Point{d, d}, c)
}

// Image Utils

func CreateImage(sizeX, sizeY int) *image.RGBA {
	rect := image.Rect(0, 0, sizeX, sizeY)
	return image.NewRGBA(rect)
}

func SaveImage(m *image.RGBA) {
	outF, err := os.Create("./myImage.jpeg")
	if err != nil {
		panic(err)
	}

	err = jpeg.Encode(outF, m, nil)
	if err != nil {
		panic(err)
	}
}
