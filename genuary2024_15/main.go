// Package main renders an image, gif or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/easing"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
)

func main() {

	renderTarget := target.Video

	switch renderTarget {
	case target.Image:
		render.Image(800, 800, "out/out.png", scene1, 0.0)
		render.ViewImage("out/out.png")
		break

	case target.Video:
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 30)
		program.AddSceneWithFrames(scene2, 30)
		program.AddSceneWithFrames(scene3, 30)
		program.AddSceneWithFrames(scene4, 30)
		program.AddSceneWithFrames(scene5, 25)
		program.AddSceneWithFrames(scene6, 25)
		program.AddSceneWithFrames(scene7, 25)
		program.AddSceneWithFrames(scene8, 25)
		program.AddSceneWithFrames(scene9, 25)
		program.AddSceneWithFrames(scene10, 25)
		program.AddSceneWithFrames(scene11, 25)
		program.AddSceneWithFrames(scene12, 30)
		program.AddSceneWithFrames(scene13, 30)

		program.Render("out/frames")
		render.ConvertToVideo("out/frames", "out/out.mp4", program.Width, program.Height, program.FPS, program.Seconds())
		render.PlayVideo("out/out.mp4")
		break
	}
}

//revive:unused-parameter
func scene1(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	x := width / 2
	y := easing.ElasticEaseOut(percent, height, height/2)
	context.FillRectangle(x-4, y-4, 8, 8)
}

func scene2(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	w := easing.QuinticEaseInOut(percent, 8, 320)
	x := width/2 - w/2
	y := height/2 - 4
	context.FillRectangle(x, y, w, 8)
}

func scene3(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	w := 320.0
	h := easing.QuinticEaseOut(percent, 8, 320)
	lw := easing.QuinticEaseOut(percent, 4, 1)
	context.SetLineWidth(lw)
	x := width/2 - w/2
	y := height/2 - h/2
	context.StrokeRectangle(x, y, w, h)
}

func scene4(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	drawRect(context, width, height)
	context.SetFontSize(40)
	context.SetSourceGray(easing.QuinticEaseIn(percent, 1, 0))
	text := "genuary"
	for t := 0; t < 7; t++ {
		x := blmath.Map(float64(t), 0, 6, 60, 310)
		context.FillText(string(text[t]), x, 80)
	}
}

func scene5(context *cairo.Context, width, height, percent float64) {
	fall(context, width, height, percent, 0)
}

func scene6(context *cairo.Context, width, height, percent float64) {
	fall(context, width, height, percent, 1)
}

func scene7(context *cairo.Context, width, height, percent float64) {
	fall(context, width, height, percent, 2)
}

func scene8(context *cairo.Context, width, height, percent float64) {
	fall(context, width, height, percent, 3)
}

func scene9(context *cairo.Context, width, height, percent float64) {
	fall(context, width, height, percent, 4)
}

func scene10(context *cairo.Context, width, height, percent float64) {
	fall(context, width, height, percent, 5)
}

func scene11(context *cairo.Context, width, height, percent float64) {
	fall(context, width, height, percent, 6)
}

func fall(context *cairo.Context, width, height, percent float64, index int) {
	context.BlackOnWhite()
	drawRect(context, width, height)
	context.SetFontSize(40)
	text := "genuary"
	for t := 0; t < 7; t++ {
		x := blmath.Map(float64(t), 0, 6, 60, 310)
		if t > index {
			y := 80.0
			context.FillText(string(text[t]), x, y)
		} else if t < index {
			y := 360.0
			context.FillText(string(text[t]), x, y)
		}
	}

	x := blmath.Map(float64(index), 0, 6, 60, 310)
	y := easing.BounceEaseOut(percent, 80, 360)
	context.FillText(string(text[index]), x, y)
}

func scene12(context *cairo.Context, width, height, percent float64) {
	context.ClearWhite()
	drawRect(context, width, height)
	context.SetFontSize(40)
	context.SetSourceGray(easing.SineEaseIn(percent, 0, 1))
	text := "genuary"
	random.Seed(0)
	for t := 0; t < 7; t++ {
		x := blmath.Map(float64(t), 0, 6, 60, 310)
		y := 360.0
		a := random.Angle()
		x = easing.SineEaseIn(percent, x, x+math.Cos(a)*200)
		y = easing.SineEaseIn(percent, y, y+math.Sin(a)*200)
		context.FillText(string(text[t]), x, y)
	}
}

func scene13(context *cairo.Context, width, height, percent float64) {
	context.ClearWhite()
}

func drawRect(context *cairo.Context, width, height float64) {
	w := 320.0
	h := 320.0
	x := width/2 - w/2
	y := height/2 - h/2
	context.StrokeRectangle(x, y, w, h)
}
