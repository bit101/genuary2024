// Package days contains genuary 2024 code for each day.
package days

import (
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/target"
)

// Day11 is for genuary 11
var Day11 = Day{
	ImageWidth:  800,
	ImageHeight: 800,
	VideoWidth:  400,
	VideoHeight: 800,
	VideoTime:   10,
	RenderFrame: Day11Render,
	Target:      target.Video,
}

// Day11Render is for genuary 11
// In the style of Anni Albers (1899-1994).
// https://www.theguardian.com/artanddesign/2019/jan/18/anni-albers-intersecting
//
//revive:disable-next-line:unused-parameter
func Day11Render(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	context.SetLineWidth(1)
	w := 20.0
	context.SetSourceGray(0.95)
	for t := 0.0; t < width; t += w * 4 {
		context.FillRectangle(t, 0, w*2, height)
	}
	context.SetSourceBlack()

	random.Seed(0)
	for xx := w; xx < width; xx += w * 2 {
		xc := xx
		vx := 0.0
		x := xc + random.FloatRange(-w, w)
		k := 0.1
		damp := 0.91
		for y := height; y >= 0; y-- {
			// avert your eyes, young programmers
			if random.WeightedBool(0.005) {
				// randomly go left or right
				if random.Boolean() {
					// but don't go off the left edge
					if xc > w {
						xc -= w * 2
					} else {
						xc += w * 2
					}
				} else {
					// and don't go off the right edge
					if xc < width-w {
						xc += w * 2
					} else {
						xc -= w * 2
					}
				}
			}
			x += vx
			vx += (xc - x) * k
			vx *= damp

			if y > height-height*percent {
				context.LineTo(x, y)
			}
		}
		context.SetSourceHSV(xx/width*360, 1, 1)
		context.Stroke()
	}
}
