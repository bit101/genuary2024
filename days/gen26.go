// Package days contains genuary 2024 code for each day.
package days

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/geom"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/genuary2024/util"
)

// Day26 is for genuary 26
var Day26 = Day{
	ImageWidth:  800,
	ImageHeight: 800,
	VideoWidth:  400,
	VideoHeight: 400,
	VideoTime:   10,
	RenderFrame: Day26Render,
	Target:      target.Video,
}

type branch struct {
	x, y, angle, vr float64
	points          geom.PointList
}

func (b *branch) update() {
	b.x += math.Cos(b.angle) * 5
	b.y += math.Sin(b.angle) * 5
	b.points.AddXY(b.x, b.y)
	b.vr += random.FloatRange(-0.02, 0.02)
	b.angle += b.vr
	b.vr *= 0.98
}

var branches = []*branch{}

func init() {
	b := branch{
		x:      Day26.VideoWidth * 0.9,
		y:      Day26.VideoHeight,
		angle:  -math.Pi / 2,
		vr:     0.0,
		points: geom.NewPointList(),
	}
	b.points.AddXY(b.x, b.y)
	branches = append(branches, &b)
}

// Day26Render is for genuary 26
// Grow a seed.
//
//revive:disable-next-line:unused-parameter
func Day26Render(context *cairo.Context, width, height, percent float64) {
	if percent < 0.9 {
		percent = blmath.Map(percent, 0, 0.9, 0, 1)
		scene1_26(context, width, height, percent)
	} else {
		percent = blmath.Map(percent, 0.9, 1, 0, 1)
		scene2_26(context, width, height, percent)
	}
	context.SetSourceBlack()
	util.Stampit(context, "genuary2024, day26, Grow a seed")
}

func scene1_26(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	for _, b := range branches {
		b.points.Translate(-1, 0)
		context.StrokePath(b.points, false)
		b.update()
	}
	if random.WeightedBool(0.05) {
		branchIndex := random.IntRange(0, len(branches))
		theBranch := branches[branchIndex]
		pointIndex := random.IntRange(0, len(theBranch.points))
		point := theBranch.points[pointIndex]
		b := branch{
			x:      point.X,
			y:      point.Y,
			angle:  -math.Pi / 2,
			vr:     0.0,
			points: geom.NewPointList(),
		}
		b.points.AddXY(b.x, b.y)
		branches = append(branches, &b)

	}

}
func scene2_26(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	context.SetSourceGray(percent)
	for _, b := range branches {
		b.points.Translate(-1, 0)
		context.StrokePath(b.points, false)
		b.update()
	}
}
