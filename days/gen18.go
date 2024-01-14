// Package days contains genuary 2024 code for each day.
package days

import (
	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/genuary2024/util"
)

// Day18 is for genuary 18
var Day18 = Day{
	ImageWidth:  800,
	ImageHeight: 800,
	VideoWidth:  400,
	VideoHeight: 400,
	VideoTime:   5,
	RenderFrame: Day18Render,
	Target:      target.Video,
}

// Day18Render is for genuary 18
// Bauhaus.
//
//revive:disable-next-line:unused-parameter
func Day18Render(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	context.SetLineWidth(0.5)
	random.Seed(0)

	for i := 0; i < 500; i++ {
		drawPlane(context, width, height, percent)
	}

}

func drawPlane(context *cairo.Context, width, height, percent float64) {

	x := random.FloatRange(0, width)
	y := random.FloatRange(0, height)
	w := random.FloatRange(20, 100)
	h := random.FloatRange(20, 100)
	w *= blmath.LerpSin(percent+random.Float(), 0.2, 1.2)
	h *= blmath.LerpSin(percent+random.Float(), 1.2, 0.2)
	m := cairo.NewMatrix()
	m.InitIdentity()
	m.X0 = x - w/2
	m.Y0 = y - h/2
	if random.Boolean() {
		m.Yx = -0.5
	} else {
		m.Yx = 0.5
	}

	context.Save()
	context.Transform(*m)
	context.Rectangle(0, 0, w, h)
	context.SetSourceGray(random.FloatRange(0.35, 1))
	if random.WeightedBool(0.03) {
		context.SetSourceRGB(1, 0, 0)
	}
	context.FillPreserve()
	context.SetSourceBlack()
	context.Stroke()

	context.Restore()
	util.Stampit(context, "genuary 2024 day 18 Bauhaus")
}
