// Package days contains genuary 2024 code for each day.
package days

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/target"
)

// Day23 is for genuary 23
var Day23 = Day{
	ImageWidth:  30 * 32,
	ImageHeight: 30 * 32,
	VideoWidth:  30 * 32,
	VideoHeight: 30 * 32,
	VideoTime:   1,
	RenderFrame: Day23Render,
	Target:      target.Video,
}

// Day23Render is for genuary 23
// 32×32.
//
//revive:disable-next-line:unused-parameter
func Day23Render(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	context.SetLineWidth(0.75)
	size := 30.0
	for i := 0.0; i < 32; i++ {
		for j := 0.0; j < 32; j++ {
			x := i * size
			y := j * size
			d := math.Hypot(x-width/2, y-height/2)

			context.Save()
			r := blmath.Map(math.Sin(d*0.01-percent*blmath.Tau), -1, 1, -0.5, 0.5)
			context.Translate(x, y)
			context.ConcentricRects(0, 0, size, size, 2.5, r, true)
			context.Restore()
		}
	}
	context.Stroke()
}
