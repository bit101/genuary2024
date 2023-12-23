// Package days contains genuary 2024 code for each day.
package days

import (
	"github.com/bit101/bitlib/blmath"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/target"
)

// Day03 is for genuary 3
var Day03 = Day{
	ImageWidth:  800,
	ImageHeight: 800,
	VideoWidth:  400,
	VideoHeight: 400,
	VideoTime:   2,
	RenderFrame: Day03Render,
	Target:      target.Video,
}

// Day03Render is for genuary 3
//
//revive:disable-next-line:unused-parameter
func Day03Render(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	context.Save()
	context.TranslateCenter()
	context.DrawAxes(0.25)
	r := blmath.LerpSin(percent, 50, width/2)
	sphere := cairo.NewSphere(0, 0, r, 1, 0, 0)
	sphere.SetShadowColor(0.2, 0, 0)
	sphere.Draw(context)
	context.Restore()
}
