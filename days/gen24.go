// Package days contains genuary 2024 code for each day.
package days

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/easing"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/iso"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/genuary2024/util"
)

// Day24 is for genuary 24
var Day24 = Day{
	ImageWidth:  800,
	ImageHeight: 800,
	VideoWidth:  400,
	VideoHeight: 400,
	VideoTime:   10,
	RenderFrame: Day24Render,
	Target:      target.Video,
}

// Day24Render is for genuary 24
// Impossible objects (undecided geometry).
//
//revive:disable-next-line:unused-parameter
func Day24Render(context *cairo.Context, width, height, percent float64) {
	context.ClearGray(0.8)
	context.Save()
	context.TranslateCenter()
	if percent < 0.25 {
		p := blmath.Map(percent, 0, 0.25, 0, 1)
		context.ClearGray(easing.QuadraticEaseOut(p, 0.8, 0.2))
		context.Rotate(easing.ElasticEaseOut(p, 0, math.Pi))
	} else if percent < 0.5 {
		context.ClearGray(0.2)
		context.Rotate(math.Pi)
	} else if percent < 0.75 {
		p := blmath.Map(percent, 0.5, 0.75, 0, 1)
		context.ClearGray(easing.QuadraticEaseOut(p, 0.2, 0.8))
		context.Rotate(easing.ElasticEaseOut(p, math.Pi, math.Pi*2))
	}
	box := iso.NewBox(100, 100, 100)
	box.Position(-100, 50, 0)
	box.Render(context)
	box.Position(100, 50, 0)
	box.Render(context)
	box.Position(0, 50, -150)
	box.Render(context)

	context.Restore()
	util.Stampit(context, "genuary2024 day 24 Impossible objects")
}
