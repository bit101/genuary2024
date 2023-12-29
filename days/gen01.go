// Package days contains genuary 2024 code for each day.
package days

import (
	"log"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/easing"
	"github.com/bit101/bitlib/geom"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/genuary2024/util"
)

// Day01 is for genuary 2
var Day01 = Day{
	ImageWidth:  800,
	ImageHeight: 800,
	VideoWidth:  400,
	VideoHeight: 400,
	VideoTime:   5,
	RenderFrame: Day01Render,
	Target:      target.Video,
}

const numPoints = 50000

var (
	pointsA geom.PointList
	pointsB geom.PointList
	pointsC geom.PointList
	pointsD geom.PointList
)

func init() {
	random.Seed(0)
	width := Day01.VideoWidth
	height := Day01.VideoHeight

	surfaceA := cairo.NewSurface(width, height)
	contextA := cairo.NewContext(surfaceA)
	contextA.SetFontSize(100)
	contextA.SetSourceRGB(0.2, 0.2, 0.2)
	contextA.FillText("genuary", 5, 100)
	dataA, err := contextA.Surface.GetData()
	if err != nil {
		log.Fatal(err)
	}

	surfaceB := cairo.NewSurface(width, height)
	contextB := cairo.NewContext(surfaceB)
	contextB.SetFontSize(160)
	contextB.SetSourceRGB(0.2, 0.2, 0.2)
	contextB.FillText("2024", 14, 350)
	dataB, err := contextB.Surface.GetData()
	if err != nil {
		log.Fatal(err)
	}

	pointsA = geom.NewPointList()
	pointsB = geom.NewPointList()
	pointsC = geom.NewPointList()
	pointsD = geom.NewPointList()

	for i := 0; i < numPoints; i++ {
		pointsA.Add(geom.RandomPointInRect(50, 50, width-100, height-100))
	}

	i := 0
	for i < numPoints {
		p := geom.RandomPointInRect(0, 0, width-1, height-1)
		x, y := p.Coords()
		index := (int(y)*int(width) + int(x)) * 4
		if dataA[index] > 0 {
			pointsB.Add(p)
			i++
		}
	}

	for i := 0; i < numPoints; i++ {
		pointsC.Add(geom.RandomPointInCircle(width/2, height/2, width*0.4, true))
	}

	i = 0
	for i < numPoints {
		p := geom.RandomPointInRect(0, 0, width-1, height-1)
		x, y := p.Coords()
		index := (int(y)*int(width) + int(x)) * 4
		if dataB[index] > 0 {
			pointsD.Add(p)
			i++
		}
	}
}

// Day01Render is for genuary 1
// Particles, lots of them.
//
//revive:disable-next-line:unused-parameter
func Day01Render(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	ease := easing.QuarticEaseInOut
	if percent < 0.25 {
		percent *= 4
		for i := 0; i < numPoints; i++ {
			t := ease(percent, 0, 1)
			pA := pointsA[i]
			pB := pointsB[i]
			p := geom.LerpPoint(t, pA, pB)
			p.Randomize(-5, 5)
			context.FillPoint(p.X, p.Y, 0.5)
		}
	} else if percent < 0.5 {
		percent = blmath.Map(percent, 0.25, 0.5, 0, 1)
		for i := 0; i < numPoints; i++ {
			t := ease(percent, 0, 1)
			pA := pointsB[i]
			pB := pointsC[i]
			p := geom.LerpPoint(t, pA, pB)
			p.Randomize(-5, 5)
			context.FillPoint(p.X, p.Y, 0.5)
		}
	} else if percent < 0.75 {
		percent = blmath.Map(percent, 0.5, 0.75, 0, 1)
		for i := 0; i < numPoints; i++ {
			t := ease(percent, 0, 1)
			pA := pointsC[i]
			pB := pointsD[i]
			p := geom.LerpPoint(t, pA, pB)
			p.Randomize(-5, 5)
			context.FillPoint(p.X, p.Y, 0.5)
		}
	} else {
		percent = blmath.Map(percent, 0.75, 1, 0, 1)
		for i := 0; i < numPoints; i++ {
			t := ease(percent, 0, 1)
			pA := pointsD[i]
			pB := pointsA[i]
			p := geom.LerpPoint(t, pA, pB)
			p.Randomize(-5, 5)
			context.FillPoint(p.X, p.Y, 0.5)
		}
	}
	util.Stampit(context, "genuary2024 day 1. Particles, lots of them")
}
