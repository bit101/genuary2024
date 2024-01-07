// Package days contains genuary 2024 code for each day.
package days

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/genuary2024/util"
)

// Day10 is for genuary 10
var Day10 = Day{
	ImageWidth:  800,
	ImageHeight: 800,
	VideoWidth:  400,
	VideoHeight: 400,
	VideoTime:   2,
	RenderFrame: Day10Render,
	Target:      target.Video,
}

// Day10Render is for genuary 10
// Hexagonal.
//
//revive:disable-next-line:unused-parameter
func Day10Render(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	context.SetLineWidth(0.5)
	radius := 57.8

	sin60r := math.Sin(math.Pi/3.0) * radius
	xInc := 2.0 * sin60r
	yInc := radius * 1.5
	offset := 0.0
	random.Seed(0)

	for y := 0.0; y < height+yInc; y += yInc {
		for x := 0.0; x < width+xInc; x += xInc {
			r := random.FloatRange(-0.2, 0.2)
			a := random.Angle() + percent*blmath.Tau
			context.Save()
			context.Translate(x+offset, y)
			context.Rotate(math.Pi / 2)
			drawHex(context, 0, 0, radius, a, blmath.LerpSin(percent, -r, r))
			context.Restore()
		}
		if offset == 0 {
			offset = sin60r
		} else {
			offset = 0
		}
	}
	util.Stampit(context, "genuary2024 day 10. Hexagonal")
}

func drawHex(context *cairo.Context, x, y, radius, angle, rotation float64) {
	t := 0.0
	dx := math.Cos(angle) * 1.3
	dy := math.Sin(angle) * 1.3
	for r := radius; r > 2; r -= 2 {
		context.StrokePolygon(x, y, r, 6, t)
		t += rotation
		x += dx
		y += dy
	}
}
