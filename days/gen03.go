// Package days contains genuary 2024 code for each day.
package days

import (
	"github.com/bit101/bitlib/blmath"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/genuary2024/util"
)

// Day03 is for genuary 3
var Day03 = Day{
	ImageWidth:  480,
	ImageHeight: 360,
	VideoWidth:  480,
	VideoHeight: 360,
	VideoTime:   2,
	RenderFrame: Day03Render,
	// RenderFrame: GenImage,
	Target: target.Video,
}

var img, _ = cairo.NewSurfaceFromPNG("days/assets/droste_source.png")

// GenImage generates the image to be used for the final animation
// Droste effect.
//
//revive:disable-next-line:unused-parameter
func GenImage(context *cairo.Context, width, height, percent float64) {
	context.SetSourceSurface(img, 0, 0)
	context.FillRectangle(0, 0, width, height)

	pattern := context.GetSource()
	m := cairo.NewMatrix()

	for i := 0; i < 5; i++ {
		m.Scale(1/0.35, 1/0.35)
		m.Rotate(0.07)
		m.Translate(-155, -40)
		pattern.SetMatrix(m)

		// context.SetSourceBlack()
		context.Translate(155, 40)
		context.Rotate(-0.07)
		context.Scale(0.35, 0.35)
		context.FillRectangle(0, 0, width, height)
	}
}

var droste, _ = cairo.NewSurfaceFromPNG("days/assets/droste.png")

// Day03Render is for genuary 3
// Droste effect.
//
//revive:disable-next-line:unused-parameter
func Day03Render(context *cairo.Context, width, height, percent float64) {
	context.SetSourceSurface(droste, 0, 0)
	context.FillRectangle(0, 0, width, height)

	pattern := context.GetSource()
	m := cairo.NewMatrix()

	scale := blmath.Lerp(percent, 1, 0.35)
	rotation := blmath.Lerp(percent, 0, -0.07)
	tx := blmath.Lerp(percent, 0, 155)
	ty := blmath.Lerp(percent, 0, 40)

	m.Translate(tx, ty)
	m.Rotate(rotation)
	m.Scale(scale, scale)
	pattern.SetMatrix(m)

	context.FillRectangle(0, 0, width, height)

	util.Stampit(context, "genuary2024 day 3. Droste effect")
}
