// Package days contains genuary 2024 code for each day.
package days

import (
	"fmt"
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/genuary2024/util"
)

// Day09 is for genuary 9
var Day09 = Day{
	ImageWidth:  800,
	ImageHeight: 800,
	VideoWidth:  400,
	VideoHeight: 400,
	VideoTime:   8,
	RenderFrame: Day09Render,
	Target:      target.Video,
}

// Day09Render is for genuary 9
// ASCII.
//
//revive:disable-next-line:unused-parameter
func Day09Render(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	context.SetFontSize(20)
	context.SetLineWidth(0.2)
	context.Save()
	context.TranslateCenter()

	random.RandSeed()
	x := 0.1
	y := 0.1
	a := 1.6 + math.Cos(blmath.Tau*percent)*0.1
	b := 0.9871 + math.Sin(blmath.Tau*percent)*0.1
	c := 1.8942 + math.Cos(blmath.Tau*percent)*0.01
	d := 0.9011 + math.Sin(blmath.Tau*percent)*0.01
	scale := 100.0

	for i := 0; i < 1000; i++ {
		x1 := math.Sin(x+c)*a - math.Cos(y+d)*b
		y1 := math.Sin(y+b)*c + math.Cos(x+a)*d
		context.LineTo(x1*scale, y1*scale)
		x, y = x1, y1
	}
	context.Stroke()
	for i := 0; i < 100; i++ {
		x1 := math.Sin(x+c)*a - math.Cos(y+d)*b
		y1 := math.Sin(y+b)*c + math.Cos(x+a)*d
		context.FillText(fmt.Sprintf("%d", i), x1*scale, y1*scale)
		x, y = x1, y1
	}

	context.Restore()
	util.Stampit(context, "genuary2024 day 9. ASCII")

}
