// Package days contains genuary 2024 code for each day.
package days

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/geom"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/genuary2024/util"
)

// Day02 is for genuary 2
var Day02 = Day{
	ImageWidth:  800,
	ImageHeight: 800,
	VideoWidth:  400,
	VideoHeight: 400,
	VideoTime:   2,
	RenderFrame: Day02Render,
	Target:      target.Video,
}

// Day02Render is for genuary 2
// No palettes. Generative colors, procedural colors, emergent colors.
//
//revive:disable-next-line:unused-parameter
func Day02Render(context *cairo.Context, width, height, percent float64) {
	red := geom.NewPoint(width/2+math.Cos(0)*width*0.4, height/2+math.Sin(0)*width*0.4)
	green := geom.NewPoint(width/2+math.Cos(blmath.Tau/3)*width*0.4, height/2+math.Sin(blmath.Tau/3)*width*0.4)
	blue := geom.NewPoint(width/2+math.Cos(blmath.Tau*2/3)*width*0.4, height/2+math.Sin(blmath.Tau*2/3)*width*0.4)
	res := 1.0
	offset := percent * blmath.Tau
	for x := 0.0; x < width; x += res {
		for y := 0.0; y < height; y += res {
			dr := math.Hypot(x-red.X, y-red.Y)
			dg := math.Hypot(x-green.X, y-green.Y)
			db := math.Hypot(x-blue.X, y-blue.Y)
			context.SetSourceRGB(
				dr/width+math.Sin(dr*0.015+offset)*0.2,
				dg/width+math.Sin(dg*0.02-offset)*0.2,
				db/width+math.Sin(db*0.025+offset)*0.2,
			)
			context.FillRectangle(x, y, res, res)
		}
	}
	util.Stampit(context, "genuary2024 day 2. No palettes")
}
