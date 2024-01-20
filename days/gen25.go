// Package days contains genuary 2024 code for each day.
package days

import (
	"github.com/bit101/bitlib/geom"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/genuary2024/util"
)

// Day25 is for genuary 25
var Day25 = Day{
	ImageWidth:  800,
	ImageHeight: 800,
	VideoWidth:  400,
	VideoHeight: 400,
	VideoTime:   2,
	RenderFrame: Day25Render,
	Target:      target.Image,
}

// Day25Render is for genuary 25
// If you like generative art, you probably have some photos on your phone of cool looking patterns, textures, shapes or things that you’ve seen. You might have even thought, “I should try to recreate this with code”. Today is the day.
//
//revive:disable-next-line:unused-parameter
func Day25Render(context *cairo.Context, width, height, percent float64) {
	g := cairo.CreateLinearGradient(0, 0, 0, height)
	g.AddColorStopRGB(0, 0, 0, 1)
	g.AddColorStopRGB(1, 0, 1, 0)
	context.SetSource(g)
	context.Paint()

	context.SetSourceWhite()
	points := geom.PoissonDiskSampling(width, height, 30, 5)
	for _, p := range points {
		context.Save()

		context.Translate(p.X, p.Y)
		pc := p.Y / height
		s := random.FloatRange(0.8, 1.3)
		context.Scale(s, s)

		g := cairo.CreateLinearGradient(0, 6*(pc), 0, -6*(1-pc))
		g.AddColorStopRGBA(0, 0, 0, 1, 0.75)
		g.AddColorStopRGBA(1, 0, 1, 0, 0.5)
		context.SetSource(g)
		context.FillEllipse(0, 0, 10, 12)

		g = cairo.CreateRadialGradient(-1, -9, 0, 0, 0, 10)
		g.AddColorStopRGBA(0, 1, 1, 1, 0.55)
		g.AddColorStopRGBA(1, 1, 1, 1, 0)
		context.SetSource(g)
		context.FillEllipse(0, 0, 10, 12)

		context.Restore()
	}
	util.Stampit(context, "genuary2024, day 25, recreate photo pattern/texture/shape")
}
