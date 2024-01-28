// Package main renders an image, gif or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/easing"
	"github.com/bit101/bitlib/geom"
	"github.com/bit101/bitlib/noise"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
)

func main() {
	render.MixAV("out/output.mp4", "../sound/gen31.mp3", "final13.mp4")
	// act1()
	// act2()
	// act3()
	// act4()
	// act5()
	// act6()
	// act7()
	// act8()
}

func act1() {
	program := render.NewProgram(640, 480, 30)
	program.AddSceneWithFrames(scene1, 120)
	program.AddSceneWithFrames(scene2, 30)
	program.AddSceneWithFrames(scene3, 30)
	program.RenderVideo("out/frames", "out/act1.mp4")
	render.PlayVideo("out/act1.mp4")
}

func act2() {
	program := render.NewProgram(640, 480, 30)
	program.AddSceneWithFrames(scene4, 30)
	for i := 0; i < 4; i++ {
		program.AddSceneWithFrames(scene5, 15)
		program.AddSceneWithFrames(scene6, 15)
		program.AddSceneWithFrames(scene7, 15)
		program.AddSceneWithFrames(scene8, 15)
	}
	program.RenderVideo("out/frames", "out/act2.mp4")
	render.PlayVideo("out/act2.mp4")
}

func act3() {
	program := render.NewProgram(640, 480, 30)
	program.AddSceneWithFrames(scene9, 30)
	program.AddSceneWithFrames(scene10, 60)
	program.RenderVideo("out/frames", "out/act3.mp4")
	render.PlayVideo("out/act3.mp4")
}

func act4() {
	program := render.NewProgram(640, 480, 30)
	program.AddSceneWithFrames(scene11, 30)
	program.AddSceneWithFrames(scene12, 90)
	program.AddSceneWithFrames(scene13, 30)
	program.RenderVideo("out/frames", "out/act4.mp4")
	render.PlayVideo("out/act4.mp4")
}

func act5() {
	program := render.NewProgram(640, 480, 30)
	program.AddSceneWithFrames(scene14, 30)
	for i := 0; i < 8; i++ {
		program.AddSceneWithFrames(scene15, 15)
		program.AddSceneWithFrames(scene16, 15)
	}
	program.RenderVideo("out/frames", "out/act5.mp4")
	render.PlayVideo("out/act5.mp4")
}

func act6() {
	program := render.NewProgram(640, 480, 30)
	for i := 0; i < 8; i++ {
		program.AddSceneWithFrames(scene17, 15)
	}
	program.AddSceneWithFrames(scene18, 60)
	for i := 0; i < 4; i++ {
		program.AddSceneWithFrames(scene19, 30)
		program.AddSceneWithFrames(scene20, 30)
	}
	for i := 0; i < 4; i++ {
		program.AddSceneWithFrames(scene21, 15)
	}
	program.RenderVideo("out/frames", "out/act6.mp4")
	render.PlayVideo("out/act6.mp4")
}

func act7() {
	program := render.NewProgram(640, 480, 30)
	program.AddSceneWithFrames(scene22, 60)
	program.AddSceneWithFrames(scene23, 60)
	program.RenderVideo("out/frames", "out/act7.mp4")
	render.PlayVideo("out/act7.mp4")
}

func act8() {
	program := render.NewProgram(640, 480, 30)
	program.AddSceneWithFrames(scene24, 120)
	program.AddSceneWithFrames(scene25, 120)
	program.RenderVideo("out/frames", "out/act8.mp4")
	render.PlayVideo("out/act8.mp4")
}

var points geom.PointList

const pointSize = 1.0

func setup(context *cairo.Context) {
	context.BlackOnWhite()
	context.Save()
	context.TranslateCenter()
	random.Seed(0)
	points = geom.PoissonDiskSampling(800, 800, 4, 30)
	points.Translate(-400, -400)
}

//revive:disable-next-line:unused-parameter
func scene1(context *cairo.Context, width, height, percent float64) {
	setup(context)
	points.Scale(2, 2)

	total := percent * float64(len(points))
	for i := 0.0; i < total; i++ {
		p := points[int(i)]
		context.FillCircle(p.X, p.Y, pointSize)
	}
	context.Restore()
}

//revive:disable-next-line:unused-parameter
func scene2(context *cairo.Context, width, height, percent float64) {
	setup(context)

	points.Scale(2, 2)
	rand := blmath.LoopSin(percent, 0, 30)

	points.Randomize(rand, rand)
	context.Points(points, pointSize)

	context.Restore()
}

//revive:disable-next-line:unused-parameter
func scene3(context *cairo.Context, width, height, percent float64) {
	setup(context)

	s := easing.SineEaseOut(percent, 2, 1)
	points.Scale(s, s)
	context.Points(points, pointSize)

	context.Restore()
}

//revive:disable-next-line:unused-parameter
func scene4(context *cairo.Context, width, height, percent float64) {
	setup(context)
	context.Points(points, pointSize)
	context.Restore()
}

//revive:disable-next-line:unused-parameter
func scene5(context *cairo.Context, width, height, percent float64) {
	setup(context)

	bulge(-160, -120, percent)
	context.Points(points, pointSize)
	context.Restore()
}

//revive:disable-next-line:unused-parameter
func scene6(context *cairo.Context, width, height, percent float64) {
	setup(context)

	bulge(160, 120, percent)
	context.Points(points, pointSize)
	context.Restore()
}

//revive:disable-next-line:unused-parameter
func scene7(context *cairo.Context, width, height, percent float64) {
	setup(context)

	bulge(-160, 120, percent)
	context.Points(points, pointSize)
	context.Restore()
}

//revive:disable-next-line:unused-parameter
func scene8(context *cairo.Context, width, height, percent float64) {
	setup(context)

	bulge(160, -120, percent)
	context.Points(points, pointSize)
	context.Restore()
}
func bulge(x, y, percent float64) {
	c := geom.NewPoint(x, y)
	size := 200.0
	offset := 40.0

	for _, p := range points {
		dist := p.Distance(c)
		if dist < size {
			angle := c.AngleTo(p)
			t := blmath.LoopSin(percent, 0, offset*(1-dist/size))
			p.Translate(math.Cos(angle)*t, math.Sin(angle)*t)
		}
	}
}

//revive:disable-next-line:unused-parameter
func scene9(context *cairo.Context, width, height, percent float64) {
	setup(context)

	s := 0.004
	offset := 10 * percent
	for _, p := range points {
		n := noise.Simplex3(p.X*s, p.Y*s, 0) * blmath.Tau
		p.X += math.Cos(n) * offset
		p.Y += math.Sin(n) * offset
	}
	context.Points(points, pointSize)
	context.Restore()
}

//revive:disable-next-line:unused-parameter
func scene10(context *cairo.Context, width, height, percent float64) {
	setup(context)

	s := 0.004
	offset := 10.0
	z := blmath.Lerp(percent, 0, 1)
	for _, p := range points {
		n := noise.Simplex3(p.X*s, p.Y*s, z) * blmath.Tau
		p.X += math.Cos(n) * offset
		p.Y += math.Sin(n) * offset
	}

	context.Points(points, pointSize)
	context.Restore()
}

//revive:disable-next-line:unused-parameter
func scene11(context *cairo.Context, width, height, percent float64) {
	setup(context)

	s := 0.004
	offset := 10.0
	z := 1.0
	for _, p := range points {
		n := noise.Simplex3(p.X*s, p.Y*s, z) * blmath.Tau
		p.X += math.Cos(n) * offset
		p.Y += math.Sin(n) * offset
	}
	scale := blmath.Lerp(percent, 1, 3)
	points.Scale(scale, scale)

	context.Points(points, pointSize)
	context.Restore()
}

//revive:disable-next-line:unused-parameter
func scene12(context *cairo.Context, width, height, percent float64) {
	setup(context)
	points.Rotate(percent * blmath.Tau)

	s := 0.004
	offset := 10.0
	angle := math.Pi * (1 - percent)

	y := math.Cos(angle) * 0.5
	z := 1.0
	for _, p := range points {
		n := noise.Simplex3(p.X*s, (p.Y+y)*s, z) * blmath.Tau
		p.X += math.Cos(n) * offset
		p.Y += math.Sin(n) * offset
	}
	scale := blmath.Lerp(percent, 3, 0.8)
	points.Scale(scale, scale)

	context.Points(points, pointSize)
	context.Restore()
}

//revive:disable-next-line:unused-parameter
func scene13(context *cairo.Context, width, height, percent float64) {
	setup(context)

	s := 0.004
	offset := 10.0
	z := 1.0
	for _, p := range points {
		n := noise.Simplex3(p.X*s, p.Y*s, z) * blmath.Tau
		p.X += math.Cos(n) * offset
		p.Y += math.Sin(n) * offset
	}
	scale := blmath.Lerp(percent, 0.8, 1)
	points.Scale(scale, scale)

	context.Points(points, pointSize)
	context.Restore()
}

//revive:disable-next-line:unused-parameter
func scene14(context *cairo.Context, width, height, percent float64) {
	setup(context)

	s := blmath.Lerp(percent, 0.004, 0.01)
	offset := blmath.Lerp(percent, 10, 5)
	z := 1.0
	for _, p := range points {
		n := noise.Simplex3(p.X*s, p.Y*s, z) * blmath.Tau
		p.X += math.Cos(n) * offset
		p.Y += math.Sin(n) * offset
	}

	context.Points(points, pointSize)
	context.Restore()
}

//revive:disable-next-line:unused-parameter
func scene15(context *cairo.Context, width, height, percent float64) {
	setup(context)

	s := 0.01
	offset := 5.0
	z := 1.0
	for _, p := range points {
		n := noise.Simplex3(p.X*s, p.Y*s, z) * blmath.Tau
		p.X += math.Cos(n) * offset
		p.Y += math.Sin(n) * offset
	}

	bulge(-160, 0, percent)
	context.Points(points, pointSize)
	context.Restore()
}

//revive:disable-next-line:unused-parameter
func scene16(context *cairo.Context, width, height, percent float64) {
	setup(context)

	s := 0.01
	offset := 5.0
	z := 1.0
	for _, p := range points {
		n := noise.Simplex3(p.X*s, p.Y*s, z) * blmath.Tau
		p.X += math.Cos(n) * offset
		p.Y += math.Sin(n) * offset
	}

	bulge(160, 0, percent)
	context.Points(points, pointSize)
	context.Restore()
}

//revive:disable-next-line:unused-parameter
func scene17(context *cairo.Context, width, height, percent float64) {
	setup(context)

	s := 0.01
	offset := blmath.LoopSin(percent, 5, 2)
	z := 1.0
	for _, p := range points {
		n := noise.Simplex3(p.X*s, p.Y*s, z) * blmath.Tau
		p.X += math.Cos(n) * offset
		p.Y += math.Sin(n) * offset
	}

	context.Points(points, pointSize)
	context.Restore()
}

//revive:disable-next-line:unused-parameter
func scene18(context *cairo.Context, width, height, percent float64) {
	setup(context)

	s := 0.01
	offset := 5.0
	z := blmath.Lerp(percent, 1.0, 0.0)
	for _, p := range points {
		n := noise.Simplex3(p.X*s, p.Y*s, z) * blmath.Tau
		p.X += math.Cos(n) * offset
		p.Y += math.Sin(n) * offset
	}

	context.Points(points, pointSize)
	context.Restore()
}

//revive:disable-next-line:unused-parameter
func scene19(context *cairo.Context, width, height, percent float64) {
	setup(context)

	s := 0.01
	offset := 5.0
	z := 0.0
	for _, p := range points {
		n := noise.Simplex3(p.X*s, p.Y*s, z) * blmath.Tau
		p.X += math.Cos(n) * offset
		p.Y += math.Sin(n) * offset
	}
	bulge(blmath.Lerp(percent, -width/2, width/2), math.Sin(percent*blmath.Tau)*100, percent)

	context.Points(points, pointSize)
	context.Restore()
}

//revive:disable-next-line:unused-parameter
func scene20(context *cairo.Context, width, height, percent float64) {
	setup(context)

	s := 0.01
	offset := 5.0
	z := 0.0
	for _, p := range points {
		n := noise.Simplex3(p.X*s, p.Y*s, z) * blmath.Tau
		p.X += math.Cos(n) * offset
		p.Y += math.Sin(n) * offset
	}
	bulge(blmath.Lerp(percent, width/2, -width/2), math.Sin(percent*blmath.Tau)*100, percent)

	context.Points(points, pointSize)
	context.Restore()
}

//revive:disable-next-line:unused-parameter
func scene21(context *cairo.Context, width, height, percent float64) {
	setup(context)

	s := 0.01
	offset := 5.0
	z := 0.0
	for _, p := range points {
		n := noise.Simplex3(p.X*s, p.Y*s, z) * blmath.Tau
		p.X += math.Cos(n) * offset
		p.Y += math.Sin(n) * offset
	}
	bulge(0, 0, percent)

	context.Points(points, pointSize)
	context.Restore()
}

//revive:disable-next-line:unused-parameter
func scene22(context *cairo.Context, width, height, percent float64) {
	setup(context)

	s := 0.01
	offset := 5.0
	z := 0.0
	for _, p := range points {
		n := noise.Simplex3(p.X*s, p.Y*s, z) * blmath.Tau
		p.X += math.Cos(n) * offset
		p.Y += math.Sin(n) * offset
		p.X = easing.ElasticEaseOut(percent, p.X, p.X*0.25)
		p.Y = easing.ElasticEaseOut(percent, p.Y, p.Y*0.25)
	}

	context.Points(points, pointSize)
	context.Restore()
}

//revive:disable-next-line:unused-parameter
func scene23(context *cairo.Context, width, height, percent float64) {
	setup(context)

	s := 0.01
	offset := 5.0
	z := 0.0
	for _, p := range points {
		n := noise.Simplex3(p.X*s, p.Y*s, z) * blmath.Tau
		p.X += math.Cos(n) * offset
		p.Y += math.Sin(n) * offset
		p.X = easing.ElasticEaseOut(percent, p.X*0.25, p.X)
		p.Y = easing.ElasticEaseOut(percent, p.Y*0.25, p.Y)
	}

	context.Points(points, pointSize)
	context.Restore()
}

//revive:disable-next-line:unused-parameter
func scene24(context *cairo.Context, width, height, percent float64) {
	setup(context)

	s := blmath.Lerp(percent, 0.01, 0.001)
	offset := blmath.Lerp(percent, 5, 0)
	z := 0.0
	for _, p := range points {
		n := noise.Simplex3(p.X*s, p.Y*s, z) * blmath.Tau
		p.X += math.Cos(n) * offset
		p.Y += math.Sin(n) * offset
	}
	points.Rotate(-percent * blmath.Tau)

	context.Points(points, pointSize)
	context.Restore()
}

//revive:disable-next-line:unused-parameter
func scene25(context *cairo.Context, width, height, percent float64) {
	setup(context)
	points.Rotate(-percent * blmath.Tau)

	scale := blmath.Lerp(percent, 1, 5)
	points.Scale(scale, scale)
	ps := blmath.Lerp(percent, pointSize, 0)

	context.Points(points, ps)
	context.Restore()
}
