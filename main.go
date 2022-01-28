package main

import (
	"fmt"
	"math/rand"
	"time"

	gd "github.com/misterunix/cgo-gd"
)

var randomSource rand.Source
var rnd *rand.Rand

func main() {

	fmt.Println("Starting...")

	width := 256
	height := 256

	rseed := time.Now().UnixNano()
	randomSource = rand.NewSource(rseed)
	rnd = rand.New(randomSource)

	ibuffer0 := gd.CreateTrueColor(width, height)
	bgColor := ibuffer0.ColorAllocate(0x0, 0x0, 0x0)
	ibuffer0.Fill(width/2, height/2, bgColor)
	c1 := ibuffer0.ColorAllocate(0xee, 0xee, 0xee)

	var x, y, xnew, ynew int16

	x = int16(width) / 2
	y = int16(height) / 2

	for i := 0; i < 10000000; i++ {

		ibuffer0.SetPixel(int(x&0x7FFF), int(y&0x7FFF), c1)

		xnew = x - y/2
		ynew = y + xnew/2

		if xnew > int16(width) {
			xnew -= int16(width)
		}
		if xnew < 0 {
			xnew += int16(width)
		}
		if ynew > int16(height) {
			ynew -= int16(height)
		}
		if ynew < 0 {
			ynew += int16(height)
		}

		//ibuffer0.Line(x, y, xnew, ynew, c1)

		x = xnew
		y = ynew
	}

	pngfilename := fmt.Sprintf("test.png")
	ibuffer0.Png(pngfilename)

}
