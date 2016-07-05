package main

import (
	"bytes"
	"image"
	"image/color/palette"
	"image/gif"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func convertToPaletted(img image.Image) *image.Paletted {
	imgPltd := image.NewPaletted(
		img.Bounds(),
		palette.Plan9,
	)

	sz := img.Bounds().Size()
	x := sz.X
	y := sz.Y

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			imgPltd.Set(i, j, img.At(i, j))
		}
	}

	return imgPltd
}

func convertToPaletted2(img image.Image) *image.Paletted {
	buf := bytes.Buffer{}
	gif.Encode(&buf, img, nil)
	tmpimg, err := gif.Decode(&buf)
	if err != nil {
		log.Fatalln(err)
	}

	return tmpimg.(*image.Paletted)
}

func main() {

	fNames, err := filepath.Glob("*.jpeg")
	if err != nil {
		log.Fatalln(err)
	}
	if len(fNames) == 0 {
		log.Fatalln("No files found..")
	}

	sort.Strings(fNames)

	outGif := &gif.GIF{}
	for _, name := range fNames {
		f, err := os.Open(name)
		defer f.Close()
		if err != nil {
			log.Fatalln(err)
		}

		img, err := jpeg.Decode(f)
		if err != nil {
			log.Fatalln(err)
		}

		// outGif.Image = append(outGif.Image, convertToPaletted(img))
		outGif.Image = append(outGif.Image, convertToPaletted2(img))
		outGif.Delay = append(outGif.Delay, 15)
	}

	f, _ := os.OpenFile("out.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, outGif)
}
