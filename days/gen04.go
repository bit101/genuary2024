// Package days contains genuary 2024 code for each day.
package days

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/noise"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/genuary2024/util"
)

// Day04 is for genuary 4
var Day04 = Day{
	ImageWidth:  800,
	ImageHeight: 800,
	VideoWidth:  400,
	VideoHeight: 400,
	VideoTime:   10,
	RenderFrame: Day04Render,
	Target:      target.Video,
}

var data []byte

const (
	imgWidth  = 236.0
	imgHeight = 236.0
)

func init() {
	surface, _ := cairo.NewSurfaceFromPNG("days/assets/mouse.png")
	data, _ = surface.GetData()
}

// Day04Render is for genuary 4
// Pixels.
//
//revive:disable-next-line:unused-parameter
func Day04Render(context *cairo.Context, width, height, percent float64) {
	context.ClearWhite()
	scale := blmath.LerpSin(percent, 0.0001, 0.01)
	offset := blmath.LerpSin(percent, 0, 100)
	// for i := 0; i < 100000; i++ {
	// 	x := random.FloatRange(0, imgWidth)
	// 	y := random.FloatRange(0, imgHeight)
	for x := 0.0; x < imgWidth; x++ {
		for y := 0.0; y < imgHeight; y++ {
			t := noise.Simplex3(x*scale, y*scale, 0) * blmath.Tau
			x1, y1 := x+math.Cos(t)*offset, y+math.Sin(t)*offset
			x1 = blmath.Wrap(x1, 0, imgWidth-1)
			y1 = blmath.Wrap(y1, 0, imgHeight-1)
			b, g, r, a := getPixel(x1, y1)
			context.SetSourceRGBA(
				float64(r)/255,
				float64(g)/255,
				float64(b)/255,
				float64(a)/255,
			)
			x2 := blmath.Map(x, 0, float64(imgWidth), 0, width)
			y2 := blmath.Map(y, 0, float64(imgHeight), 0, height)
			context.FillRectangle(x2, y2, 2, 2)
		}
	}

	util.Stampit(context, "genuary2024 day 4. Pixels")
}

// func getPixel(x, y float64) (float64, float64, float64, float64) {
func getPixel(x, y float64) (byte, byte, byte, byte) {
	i := (int(y)*imgWidth + int(x)) * 4
	return data[i], data[i+1], data[i+2], data[i+3]
}
