// Package days contains genuary 2024 code for each day.
package days

import (
	"fmt"
	"math"

	"github.com/bit101/bitlib/blmath"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/genuary2024/util"
)

// Day20 is for genuary 20
var Day20 = Day{
	ImageWidth:  800,
	ImageHeight: 800,
	VideoWidth:  400,
	VideoHeight: 400,
	VideoTime:   15,
	RenderFrame: Day20Render,
	Target:      target.Video,
}

var (
	buffer = cairo.NewSurface(400, 400)
	bc     = cairo.NewContext(buffer)
)

// Day20Render is for genuary 20
// Generative typography.
//
//revive:disable-next-line:unused-parameter
func Day20Render(context *cairo.Context, width, height, percent float64) {
	bc.ClearBlack()
	bc.SetFontSize(430)
	bc.SetSourceWhite()
	bc.FillText(fmt.Sprintf("%d", int(percent*10)), 80, 360)
	data, _ := buffer.GetData()
	xres := 5.0
	yres := 5.0

	context.BlackOnWhite()
	context.SetLineWidth(0.5)
	for y := 0.0; y < width; y += yres {
		for x := 0.0; x < height; x += xres {
			yy := y + math.Sin(x*0.2+percent*10*blmath.Tau)*4
			// yy := y + random.FloatRange(0, 2)
			if hit(data, x, y) {
				yy = yy - blmath.LerpSin(percent*10, 0, 40)
			}
			context.LineTo(x, yy)
		}
		context.LineTo(width, y)
		context.LineTo(width+2, height+2)
		context.LineTo(-2, height+2)
		context.SetSourceWhite()
		context.FillPreserve()
		context.SetSourceBlack()
		context.Stroke()
	}
	util.Stampit(context, "genuary2024, day 20, Generative typography")
}

func hit(data []byte, x, y float64) bool {
	xx := int(x)
	yy := int(y)
	index := (yy*400 + xx) * 4
	return data[index] > 0
}
