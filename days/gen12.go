// Package days contains genuary 2024 code for each day.
package days

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/geom"
	"github.com/bit101/bitlib/noise"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/target"
)

// Day12 is for genuary 12
var Day12 = Day{
	ImageWidth:  800,
	ImageHeight: 800,
	VideoWidth:  400,
	VideoHeight: 400,
	VideoTime:   10,
	RenderFrame: Day12Render,
	Target:      target.Video,
}

// Day12Render is for genuary 12
// Lava lamp.
//
//revive:disable-next-line:unused-parameter
func Day12Render(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	context.SetLineWidth(0.25)
	random.Seed(0)

	for r := width/2 - 20; r > 0; r -= 2 {
		blob(context, width/2, height/2, r, percent)
	}
}

func blob(context *cairo.Context, xc, yc, radius, percent float64) {
	scale := 0.005
	offset := 20.0
	points := geom.NewPointList()
	a := percent * blmath.Tau
	for t := 0.0; t < blmath.Tau; t += 0.8 {
		x := xc + math.Cos(t)*radius
		y := yc + math.Sin(t)*radius
		n := noise.Simplex3(x*scale, y*scale+math.Cos(a)*0.2, math.Sin(a)*0.2) * blmath.Tau
		x += math.Cos(n) * offset
		y += math.Sin(n) * offset

		points.AddXY(x, y)
	}
	// context.SetSourceWhite()
	context.SetSourceHSV(-percent*360+radius/200*90, 1, 1)
	context.MultiLoop(points)
	context.FillPreserve()
	context.SetSourceBlack()
	context.StrokeMultiLoop(points)
}
