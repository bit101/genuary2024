// Package days contains genuary 2024 code for each day.
package days

import (
	"math"

	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/target"
)

// Day16 is for genuary 16
var Day16 = Day{
	ImageWidth:  800,
	ImageHeight: 800,
	VideoWidth:  100,
	VideoHeight: 100,
	VideoTime:   10020 / 60,
	RenderFrame: Day16Render,
	Target:      target.Video,
}

// Day16Render is for genuary 16
// Draw 10000 of something.
// 10,000 strange attractors
// Each drawn with 10,000 points
// On a 10,000 pixel (100x100) canvas
// Video length is 10,000 seconds (2:46:40)
//
//revive:disable-next-line:unused-parameter
func Day16Render(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	context.Save()
	context.TranslateCenter()

	random.RandSeed()
	x := 0.1
	y := 0.1
	a := random.FloatRange(-2, 2)
	b := random.FloatRange(-2, 2)
	c := random.FloatRange(-2, 2)
	d := random.FloatRange(-2, 2)
	scale := 20.0

	for i := 0; i < 10000; i++ {
		x1 := math.Sin(x+c)*a - math.Cos(y+d)*b
		y1 := math.Sin(y+b)*c + math.Cos(x+a)*d
		context.FillRectangle(x1*scale, y1*scale, 1, 1)
		x, y = x1, y1
	}

	context.Restore()
}
