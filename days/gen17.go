// Package days contains genuary 2024 code for each day.
package days

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/genuary2024/util"
)

// Day17 is for genuary 17
var Day17 = Day{
	ImageWidth:  400,
	ImageHeight: 400,
	VideoWidth:  400,
	VideoHeight: 400,
	VideoTime:   5,
	RenderFrame: Day17Render,
	Target:      target.Video,
}

// Day17Render is for genuary 17
// Inspired by Islamic art.
//
//revive:disable-next-line:unused-parameter
func Day17Render(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	context.SetLineWidth(0.75)
	context.Save()
	context.TranslateCenter()
	n := 16.0

	for t := 0.0; t < 1.0; t += 0.02 {
		context.Rotate(0.02)
		p := percent - t/3

		angle := blmath.Tau / n
		r1 := blmath.LerpSin(p, 20, 190)
		r2 := blmath.LerpSin(p+0.15, 20, 190)

		for i := 0.0; i < n; i++ {
			x := math.Cos(angle) * r1
			y := math.Sin(angle) * r1
			context.LineTo(x, y)
			angle += blmath.Tau / n

			x = math.Cos(angle) * r2
			y = math.Sin(angle) * r2
			context.LineTo(x, y)
			angle += blmath.Tau / n
		}
		context.Stroke()
	}
	context.Restore()

	util.Stampit(context, "genuary 2024 day 17 Inspired by Islamic Art")
}
