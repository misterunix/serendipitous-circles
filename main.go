package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	gd "github.com/misterunix/cgo-gd"
	"github.com/misterunix/colorworks/hsl"
)

var randomSource rand.Source
var rnd *rand.Rand

func main() {

	fmt.Println("Starting...")

	stwo := math.Sqrt(2)
	var width, height float64
	width = 512
	height = 512

	maxl := stwo * float64(width)

	rseed := time.Now().UnixNano()
	randomSource = rand.NewSource(rseed)
	rnd = rand.New(randomSource)

	//for k := 0; k < 35; k++ {

	ibuffer0 := gd.CreateTrueColor(int(width), int(height))
	bgColor := ibuffer0.ColorAllocate(0x0, 0x0, 0x0)
	ibuffer0.Fill(int(width/2.0), int(height/2.0), bgColor)
	//c1 := ibuffer0.ColorAllocateAlpha(0xee, 0xee, 0xee, 1)

	var x, y, xnew, ynew uint16

	x = uint16(rnd.Intn(65535))
	y = uint16(rnd.Intn(65535))

	nmax := float64(width) / 2
	nmin := -(float64(width) / 2)
	omax := 65535.0
	omin := 0.0
	for i := 0; i < 100000; i++ {

		plotx := ((float64(x)-omin)/(omax-omin))*(nmax-nmin) + nmin
		ploty := ((float64(y)-omin)/(omax-omin))*(nmax-nmin) + nmin

		ppx := plotx + float64(width)/2
		ppy := ploty + float64(height)/2

		d := math.Sqrt(ppx*ppx + ppy*ppy)
		hue := (d / maxl) * 360.0
		r, g, b := hsl.HSLtoRGB(hue, .8, .8)
		c2 := ibuffer0.ColorAllocate(int(r), int(g), int(b))
		ibuffer0.SetPixel(int(ppx), int(ppy), c2)

		xnew = x - y/2
		ynew = y + uint16(float64(xnew)/2.01)

		if xnew > uint16(65535) {
			xnew -= uint16(65535)
		}
		if xnew < 0 {
			xnew += uint16(65535)
		}
		if ynew > uint16(65535) {
			ynew -= uint16(65535)
		}
		if ynew < 0 {
			ynew += uint16(65535)
		}

		//ibuffer0.Line(x, y, xnew, ynew, c1)

		x = xnew
		y = ynew
	}

	pngfilename := fmt.Sprintf("images/%05d.png", 0)
	//ibuffer0.Png(pngfilename)

	ibuffer1 := gd.CreateTrueColor(int(width*2), int(height*2))

	ibuffer0.Copy(ibuffer1, int(width), int(height), 0, 0, int(width), int(height))                          // lr
	ibuffer0.CopyRotated(ibuffer1, int(width+(width/2)), int(height/2), 0, 0, int(width), int(height), 90)   // ur
	ibuffer0.CopyRotated(ibuffer1, int(width/2), int(height+(height/2)), 0, 0, int(width), int(height), 270) // ll
	ibuffer0.CopyRotated(ibuffer1, int(width/2), int(height/2), 0, 0, int(width), int(height), 180)          // ul

	ibuffer1.Png(pngfilename)

}
