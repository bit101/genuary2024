// Package days contains genuary 2024 code for each day.
package days

import (
	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/easing"
	"github.com/bit101/bitlib/geom"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/genuary2024/util"
)

// Day22 is for genuary 22
var Day22 = Day{
	ImageWidth:  400,
	ImageHeight: 400,
	VideoWidth:  400,
	VideoHeight: 400,
	VideoTime:   10,
	RenderFrame: Day22Render,
	Target:      target.Video,
}

var lines = []*geom.Segment{}
var points = geom.NewPointList()

func init() {
	points = geom.PoissonDiskSampling(400, 400, 10, 100)
	points = points.Cull(cir)
}

// Day22Render is for genuary 22
// Point - line - plane.
//
//revive:disable-next-line:unused-parameter
func Day22Render(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	if percent < 0.35 {
		percent = blmath.Map(percent, 0, 0.35, 0, 1)
		scene1_22(context, width, height, percent)
	} else if percent < 0.80 {
		percent = blmath.Map(percent, 0.35, 0.80, 0, 1)
		scene2_22(context, width, height, percent)
	} else if percent < 0.90 {
		percent = blmath.Map(percent, 0.80, 0.90, 0, 1)
		scene3_22(context, width, height, percent)
	} else {
		percent = blmath.Map(percent, 0.90, 1, 0, 1)
		scene4_22(context, width, height, percent)
	}
	util.Stampit(context, "genuary2024 day 22 Point-line-plane")

}

func cir(p *geom.Point) bool {
	radius := 180.0
	center := geom.NewPoint(200, 200)
	return p.Distance(center) < radius
}

func scene1_22(context *cairo.Context, width, height, percent float64) {
	end := int(percent * float64(len(points)))
	for i := 0; i < end; i++ {
		p := points[i]
		context.FillPoint(p.X, p.Y, 1)
	}
}

func scene2_22(context *cairo.Context, width, height, percent float64) {
	context.Points(points, 1)
	end := int(percent * float64(len(points)))
	for i := 0; i < end-1; i++ {
		p0 := points[i]
		for j := i + 1; j < end; j++ {
			p1 := points[j]
			dist := p0.Distance(p1)
			if dist < 20 {
				context.SetLineWidth(1 - dist/20)
				context.StrokeLine(p0.X, p0.Y, p1.X, p1.Y)
			}
		}
	}
}

func scene3_22(context *cairo.Context, width, height, percent float64) {
	context.SetSourceRGBA(0, 0, 0, 1-percent)
	context.Points(points, 1)
	end := len(points)
	for i := 0; i < end-1; i++ {
		p0 := points[i]
		for j := i + 1; j < end; j++ {
			p1 := points[j]
			dist := p0.Distance(p1)
			if dist < 20 {
				context.SetLineWidth(1 - dist/20)
				context.StrokeLine(p0.X, p0.Y, p1.X, p1.Y)
			}
		}
	}
	hull := geom.ConvexHull(points)
	context.SetSourceRGBA(0, 0, 0, percent)
	context.FillPath(hull)
}

func scene4_22(context *cairo.Context, width, height, percent float64) {
	context.SetSourceBlack()
	hull := geom.ConvexHull(points)
	for _, p := range hull {
		p.X = easing.LinearEase(percent, p.X, 165)
		p.Y = easing.LinearEase(percent, p.Y, 165)
	}
	context.FillPath(hull)
}
