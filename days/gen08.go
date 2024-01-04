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

// Day08 is for genuary 8
var Day08 = Day{
	ImageWidth:  800,
	ImageHeight: 800,
	VideoWidth:  400,
	VideoHeight: 400,
	VideoTime:   5,
	RenderFrame: Day08Render,
	Target:      target.Video,
}

// Day08Render is for genuary 8
// Chaotic system.
//
//revive:disable-next-line:unused-parameter
func Day08Render(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	context.Save()
	context.TranslateCenter()

	random.RandSeed()
	x := 0.1
	y := 0.1
	a := random.FloatRange(-2, 2)
	b := random.FloatRange(-2, 2)
	c := random.FloatRange(-2, 2)
	d := random.FloatRange(-2, 2)
	a = 1.6 + math.Cos(blmath.Tau*percent)*0.1
	b = 0.9871 + math.Sin(blmath.Tau*percent)*0.1
	c = 1.8942 + math.Cos(blmath.Tau*percent)*0.01
	d = 0.9011 + math.Sin(blmath.Tau*percent)*0.01
	scale := 100.0

	for i := 0; i < 100000; i++ {
		x1 := math.Sin(x+c)*a - math.Cos(y+d)*b
		y1 := math.Sin(y+b)*c + math.Cos(x+a)*d
		context.FillRectangle(x1*scale, y1*scale, 1, 1)
		x, y = x1, y1
	}

	context.Restore()
	context.FillText(fmt.Sprintf("a: %f0.2", a), 20, 20)
	context.FillText(fmt.Sprintf("b: %f0.2", b), 20, 35)
	context.FillText(fmt.Sprintf("c: %f0.2", c), 20, 50)
	context.FillText(fmt.Sprintf("c: %f0.2", d), 20, 65)
	util.Stampit(context, "genuary2024 day 8. Chaotic system")
}
