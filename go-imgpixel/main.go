package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"strconv"

	_ "image/jpeg"
	_ "image/png"
)

func main() {
	fName := os.Args[1]

	xstr, ystr := os.Args[2], os.Args[3]
	x, _ := strconv.Atoi(xstr)
	y, _ := strconv.Atoi(ystr)

	f, err := os.Open(fName)
	if err != nil {
		log.Fatal(err)
	}

	m, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	r, g, b, _ := m.At(x, y).RGBA()
	fmt.Println(r, g, b)
}
