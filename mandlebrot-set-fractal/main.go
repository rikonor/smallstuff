package main

import (
	"fmt"
	"image"
	"math/cmplx"
	"os"
	"time"

	"image/color"
	"image/jpeg"
)

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

// Fractal Utils

var maxNumOfIterations = 250

// Given a point (x, y) check whethr it is part of the set or not
func CheckPoint(x, y float64) (numOfIterations int, isInSet bool) {
	z := complex(0, 0)
	c := complex(x, y)

	for {
		numOfIterations++
		if numOfIterations == maxNumOfIterations {
			break
		}

		z = z*z + c

		if cmplx.Abs(z) >= 2 {
			return numOfIterations, false
		}
	}

	return numOfIterations, true
}

func main() {
	start := time.Now()

	hist := make(map[int]int)

	xMax := 1600
	yMax := xMax * 2 / 3

	m := CreateImage(xMax, yMax)

	// Edit image
	// Create Fractal and set it to the image

	// Iterate over all points
	for x := 0; x < xMax; x++ {
		for y := 0; y < yMax; y++ {
			xMinRange, xMaxRange := -2.5, 1.0
			yMinRange, yMaxRange := -1.0, 1.0

			xReal := ((xMaxRange-xMinRange)/float64(xMax))*float64(x) + xMinRange
			yReal := ((yMaxRange-yMinRange)/float64(yMax))*float64(y) + yMinRange

			numOfIters, isInSet := CheckPoint(xReal, yReal)

			// Update histogram
			hist[numOfIters]++

			// Color all points outside of set
			if !isInSet {
				// Get Color for point - scale based on numOfIters
				minColor := 0
				maxColor := 255
				colorIntensity := uint8(float64(numOfIters*maxColor)/float64(maxNumOfIterations) + float64(minColor))
				m.Set(x, y, color.RGBA{colorIntensity, colorIntensity, colorIntensity, 255})
			}
		}
	}

	// Save image
	SaveImage(m)

	// Print histogram
	for i := 0; i <= 1000; i++ {
		if count, ok := hist[i]; ok {
			fmt.Printf("%d: %d\n", i, count)
		}
	}

	fmt.Println("Took", time.Now().Sub(start))
}
