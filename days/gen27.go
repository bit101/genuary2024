// Package days contains genuary 2024 code for each day.
package days

import (
	"sort"

	"github.com/bit101/bitlib/blcolor"
	"github.com/bit101/bitlib/geom"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/iso"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/genuary2024/util"
)

// Day27 is for genuary 27
var Day27 = Day{
	ImageWidth:  800,
	ImageHeight: 800,
	VideoWidth:  400,
	VideoHeight: 400,
	VideoTime:   2,
	RenderFrame: Day27Render,
	Target:      target.Image,
}

// Day27Render is for genuary 27
// Code for one hour. At the one hour mark, you’re done.
//
//revive:disable-next-line:unused-parameter
func Day27Render(context *cairo.Context, width, height, percent float64) {
	context.ClearGray(1.3)
	boxes := []*iso.Box{}

	for i := 0; i < 40; i++ {
		box := makeBox()
		box.Position(random.FloatRange(0, width), random.FloatRange(0, height), 0)
		boxes = append(boxes, box)
	}
	sort.Slice(boxes, func(i, j int) bool {
		return boxes[i].Y < boxes[j].Y
	})
	for _, box := range boxes {
		box.Render(context)
	}
	util.Stampit(context, "genuary2024, day 27, Code for one hour")
}

func makeBox() *iso.Box {
	box := iso.NewBox(100, 100, random.FloatRange(0, 300))
	hue := random.FloatRange(0, 360)
	box.Left = PoissonSideRender(blcolor.White(), blcolor.HSV(hue, 1, 0.5), 3, 1)
	box.Right = PoissonSideRender(blcolor.White(), blcolor.HSV(hue, 1, 0.5), 4, 0.75)
	box.Top = PoissonSideRender(blcolor.White(), blcolor.HSV(hue, 1, 0.5), 5, 0.75)
	return box
}

func PoissonSideRender(bgColor, dotColor blcolor.Color, radius, pointSize float64) iso.SideRender {
	return func(context *cairo.Context, x, y, w, h float64) {
		context.SetLineWidth(0.25)
		context.SetSourceColor(bgColor)
		context.Rectangle(x, y, w, h)
		context.FillPreserve()
		context.SetSourceColor(dotColor)
		context.Stroke()

		context.SetSourceColor(dotColor)
		points := geom.PoissonDiskSampling(w, h, radius, 30)
		points.Translate(x, y)
		context.Points(points, pointSize)
	}
}
