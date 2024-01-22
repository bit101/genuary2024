// Package days contains genuary 2024 code for each day.
package days

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/sdf"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/target"
)

// Day29 is for genuary 29
var Day29 = Day{
	ImageWidth:  400,
	ImageHeight: 400,
	VideoWidth:  400,
	VideoHeight: 400,
	VideoTime:   2,
	RenderFrame: Day29Render,
	Target:      target.Video,
}

// Day29Render is for genuary 29
// Signed Distance Functions (if we keep trying once per year, eventually we will be good at it!).
//
//revive:disable-next-line:unused-parameter
func Day29Render(context *cairo.Context, width, height, percent float64) {
	context.SetLineWidth(0.25)
	context.BlackOnWhite()
	context.Save()
	vals := []float64{}

	for x := 0.0; x < width; x++ {
		for y := 0.0; y < height; y++ {

			a := sdf.Circle(x, y, 100, 200, 100)
			b := sdf.Circle(x, y, 300, 200, 100)
			c := sdf.Box(x, y, 200, 0, 100, 50)
			d := sdf.Box(x, y, 200, 400, 100, 50)
			v := sdf.Min(a, b, c, d)
			vals = append(vals, v)

		}
	}
	min, max := blmath.MinMaxFloats(vals)
	// vals = blmath.NormalizeFloats(vals)
	// fmt.Println(vals)
	newVals := make([]float64, len(vals))
	for i, v := range vals {
		if v < 0 {
			newVals[i] = blmath.Map(v, min, 0, 0, 0.5)
		} else {
			newVals[i] = blmath.Map(v, 0, max, 0.5, 1)
		}
	}
	// for i, v := range newVals {
	for i := 0; i < len(vals); i += 10 {

		v := vals[i]
		x := float64(i / int(width))
		y := float64(i % int(width))
		a := blmath.Map(v, min, max, 0, blmath.Tau*4+blmath.LerpSin(percent, 0, blmath.Tau*4))
		context.MoveTo(x, y)
		context.LineTo(x+math.Cos(a)*10, y+math.Sin(a)*10)

	}
	context.Stroke()

	// context.FillRectangle(x, y, 1, 1)
	context.Restore()
}
