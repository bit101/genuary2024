// Package days contains genuary 2024 code for each day.
package days

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/easing"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/genuary2024/util"
)

// Day19 is for genuary 19
var Day19 = Day{
	ImageWidth:  800,
	ImageHeight: 800,
	VideoWidth:  400,
	VideoHeight: 400,
	VideoTime:   10,
	RenderFrame: Day19Render,
	Target:      target.Video,
}

const (
	count = 2000
)

var (
	boids = []*boid{}
)

type boid struct {
	x, y, vx, vy float64
}

func (b *boid) update(heading float64) {
	b.x += b.vx
	b.y += b.vy
	if b.x > 400 {
		b.x -= 400
	} else if b.x < 0 {
		b.x += 400
	}
	if b.y > 400 {
		b.y -= 400
	} else if b.y < 0 {
		b.y += 400
	}
	b.vx += math.Cos(heading) * 0.01
	b.vy += math.Sin(heading) * 0.01
}

func init() {
	for i := 0; i < count; i++ {
		boids = append(boids, &boid{
			x:  random.FloatRange(0, 400),
			y:  random.FloatRange(0, 400),
			vx: random.FloatRange(-1, 1),
			vy: random.FloatRange(-1, 1),
		})
	}
}

// Day19Render is for genuary 19
// Flocking.
//
//revive:disable-next-line:unused-parameter
func Day19Render(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	if percent < 0.1 {
		percent = blmath.Map(percent, 0, 0.1, 0, 1)
		scene1(context, width, height, percent)
		util.Stampit(context, "genuary2024 day 19 Flocking")
		return
	}
	if percent < 0.9 {
		percent = blmath.Map(percent, 0.1, 0.9, 0, 1)
		scene2(context, width, height, percent)
		util.Stampit(context, "genuary2024 day 19 Flocking")
		return
	}
	percent = blmath.Map(percent, 0.9, 1.0, 0, 1)
	scene3(context, width, height, percent)
	util.Stampit(context, "genuary2024 day 19 Flocking")
}

func scene1(context *cairo.Context, width, height, percent float64) {
	random.Seed(0)
	for _, b := range boids {
		x := easing.SineEaseOut(percent, random.FloatRange(0, -1600), b.x)
		context.FillCircle(x, b.y, 1)
		b.update(percent * blmath.Tau)
	}
}

func scene2(context *cairo.Context, width, height, percent float64) {
	for _, b := range boids {
		context.FillCircle(b.x, b.y, 1)
		b.update(percent * blmath.Tau)
	}
}

func scene3(context *cairo.Context, width, height, percent float64) {
	random.Seed(0)
	for _, b := range boids {
		x := easing.SineEaseIn(percent, b.x, width+random.FloatRange(0, 1600))
		context.FillCircle(x, b.y, 1)
		b.update(percent * blmath.Tau)
	}
}
