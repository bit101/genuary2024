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

// Day06 is for genuary 6
var Day06 = Day{
	ImageWidth:  600,
	ImageHeight: 400,
	VideoWidth:  600,
	VideoHeight: 400,
	VideoTime:   2,
	RenderFrame: Day06Render,
	Target:      target.Video,
}

// Day06Render is for genuary 6
// Screensaver.
//
//revive:disable-next-line:unused-parameter
func Day06Render(context *cairo.Context, width, height, percent float64) {
	w := width - 80.0
	h := height - 15.0
	margin := 20.0
	screenWidth := w - margin
	screenHeight := h - margin
	context.ClearGray(0.15)
	context.Superellipse(w/2, h/2, screenWidth/2, screenHeight/2, 8.0)
	context.Clip()

	res := 2.0
	xscale := 0.0005
	yscale := 0.1
	simplexScale := 0.005
	offset := blmath.LerpSin(percent, 0, 20)

	for x := 0.0; x < width; x += res {
		for y := 0.0; y < height; y += res {
			n := noise.Simplex2(x*simplexScale, y*simplexScale) * blmath.Tau

			x1 := x + math.Cos(n)*offset
			y1 := y + math.Sin(n)*offset
			g := noise.Perlin2(x1*xscale, y1*yscale+blmath.LerpSin(percent+0.125, 0, -10))
			g = blmath.Map(g, -1, 1, 0, 1)
			context.SetSourceGray(g)
			context.FillRectangle(x, y, res, res)
		}
	}
	g := cairo.CreateLinearGradient(120, 120, 0, 0)
	g.AddColorStopRGBA(0, 1, 1, 1, 0)
	g.AddColorStopRGBA(1, 1, 1, 1, 0.9)
	context.SetSource(g)
	context.FillSuperellipse(w/2, h/2, screenWidth/2-4, screenHeight/2-4, 8.0)

	context.ResetClip()
	context.SetSourceGray(0.1)
	context.FillCircle(width-45, 80, 30)
	context.FillCircle(width-45, 160, 30)
	context.FillCircle(width-60, 230, 10)
	context.FillCircle(width-30, 230, 10)
	context.SetSourceGray(0.7)
	context.SetLineWidth(4)
	context.Ray(width-45, 80, blmath.LerpSin(percent, -math.Pi*1.3, math.Pi*0.3), 20, 10)
	context.Ray(width-45, 160, blmath.LerpSin(percent+0.125, -math.Pi*1.3, math.Pi*0.3), 20, 10)

	context.SetLineWidth(3)
	context.SetSourceGray(0)
	for y := 270.0; y < 340; y += 7 {
		context.StrokeLine(width-75, y, width-15, y)
	}

	util.Stampit(context, "genuary2024 day y. Screensaver")
}
