// Package days contains genuary 2024 code for each day.
package days

import (
	"github.com/bit101/bitlib/geom"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/genuary2024/util"
)

// Day05 is for genuary 5
var Day05 = Day{
	ImageWidth:  800,
	ImageHeight: 800,
	VideoWidth:  400,
	VideoHeight: 400,
	VideoTime:   9,
	RenderFrame: Day05Render,
	Target:      target.Video,
}

// Day05Render is for genuary 5
// In the style of Vera Molnár (1924-2023).
//
//revive:disable-next-line:unused-parameter
func Day05Render(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	context.SetLineWidth(0.5)

	random.Seed(0)
	r := percent * 20.0

	path := geom.NewPointList()
	for x := 0.0; x < width; x += 5 {
		path.AddXY(x, 0)
	}
	path.AddXY(width, 0)

	for i := 0; i < 100; i++ {
		context.StrokePath(path, false)
		path.Translate(0, 10)
		path.Randomize(r, r)
	}

	context.Save()
	context.SetSourceRGBA(1, 1, 1, 0.75)
	context.FillRectangle(0, height-20, width, 20)
	util.Stampit(context, "genuary2024 day 5. In the style of Vera Molnar")
	context.Restore()
}
