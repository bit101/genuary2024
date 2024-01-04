// Package days contains genuary 2024 code for each day.
package days

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/noise"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/genuary2024/util"
)

// Day07 is for genuary 7
var Day07 = Day{
	ImageWidth:  400,
	ImageHeight: 400,
	VideoWidth:  400,
	VideoHeight: 420,
	VideoTime:   10,
	RenderFrame: Day07Render,
	Target:      target.Video,
}

var (
	scale  = 0.01
	offset = 0.0
)

// Day07Render is for genuary 7
// Progress bar / indicator / loading animation.
//
//revive:disable-next-line:unused-parameter
func Day07Render(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	context.SetLineWidth(0.2)
	w := width - 40
	top := 10.0
	scale = 0.01
	offset = 0.0

	for top < height-40 {
		context.Save()
		context.Translate(20, top)
		// context.StrokeRectangle(0, 0, w, 10)
		y := 0.0

		// stroke
		for x := 0.0; x < w; x++ {
			context.LineTo(distort(x, y))
		}
		x := w
		context.LineTo(distort(x, y))
		y = 10.0
		context.LineTo(distort(x, y))

		for x := w; x >= 0; x-- {
			context.LineTo(distort(x, y))
		}
		x = 0.0
		y = 0.0
		context.LineTo(distort(x, y))
		context.Stroke()

		// fill
		for x := 0.0; x < w*percent; x++ {
			context.LineTo(distort(x, y))
		}
		x = w * percent
		context.LineTo(distort(x, y))
		y = 10.0
		context.LineTo(distort(x, y))

		for x := w * percent; x >= 0; x-- {
			context.LineTo(distort(x, y))
		}
		x = 0.0
		y = 0.0
		context.LineTo(distort(x, y))
		context.Fill()

		context.Restore()
		top += 40
		offset += 1.5
		scale += 0.001
	}

	util.Stampit(context, "genuary2024 day 7. Progress bar")
}

func distort(x, y float64) (float64, float64) {
	t := noise.Simplex2(x*scale, y*scale) * blmath.Tau
	return x + math.Cos(t)*offset, y + math.Sin(t)*offset
}
