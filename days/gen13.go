// Package days contains genuary 2024 code for each day.
package days

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/genuary2024/util"
)

// Day13 is for genuary 13
var Day13 = Day{
	ImageWidth:  800,
	ImageHeight: 800,
	VideoWidth:  400,
	VideoHeight: 400,
	VideoTime:   4,
	RenderFrame: Day13Render,
	Target:      target.Video,
}

// Day13Render is for genuary 13
// Wobbly function day.
//
//revive:disable-next-line:unused-parameter
func Day13Render(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	context.SetLineWidth(0.5)

	scenes := 6.0
	percent *= scenes

	for s := 0.0; s < scenes; s++ {
		if percent < 1.0 {
			renderScene(context, width, height, s, percent)
			return
		}
		percent -= 1.0
	}
}

func renderScene(context *cairo.Context, width, height, s, percent float64) {
	for t := 0.0; t < 6; t++ {
		x := blmath.Map(t, 0, 5, 50, 350)
		if t == s {
			x1 := elasticEaseOut(percent, -40, 0) + x
			context.MoveTo(x, 0)
			context.QuadraticCurveThrough(x1, height/2, x, height)
		} else {
			context.MoveTo(x, 0)
			context.LineTo(x, height)
		}
		context.Stroke()
	}
	util.Stampit(context, "genuary2024 day 13. Wobbly function day")
}

func elasticEaseOut(t, start, end float64) float64 {
	t = math.Sin(-50*math.Pi/2*(t+1))*math.Pow(2, -10*t) + 1
	return start + (end-start)*t
}
