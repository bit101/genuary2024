// Package days contains genuary 2024 code for each day.
package days

import (
	"github.com/bit101/bitlib/blmath"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/target"
)

// Day02 is for genuary 2
var Day02 = Day{
	ImageWidth:  800,
	ImageHeight: 800,
	VideoWidth:  400,
	VideoHeight: 400,
	VideoTime:   2,
	RenderFrame: Day02Render,
	Target:      target.Image,
}

// Day02Render is for genuary 2
// In the style of Vera Molnár (1924-2023).
//
//revive:disable-next-line:unused-parameter
func Day02Render(context *cairo.Context, width, height, percent float64) {
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
