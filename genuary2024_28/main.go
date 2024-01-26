// Package main renders an image, gif or video
package main

import (
	"log"
	"math"

	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
)

func main() {
	// encode()
	decode()
}

var input *cairo.Surface
var index = 0.0
var encodedData []byte
var inited = false

func decode() {
	input, _ = cairo.NewSurfaceFromPNG("encoded.png")
	encodedData, _ = input.GetData()
	renderTarget := target.Video

	switch renderTarget {
	case target.Image:
		render.Image(400, 400, "out/out.png", scene1, 0.0)
		render.ViewImage("out/out.png")
		break

	case target.Video:
		program := render.NewProgram(400, 400, 60)
		program.AddSceneWithFrames(scene1, 100*100)
		program.RenderVideo("out/frames", "out/out.mp4")
		render.PlayVideo("out/out.mp4")
		break
	}
}

//revive:disable-next-line:unused-parameter
func scene1(context *cairo.Context, width, height, percent float64) {
	if !inited {
		context.WhiteOnBlack()
		inited = true
		context.FillText("pixel", 8, 140)
		context.FillText("value", 7, 155)
	}
	context.SetAntialias(cairo.AntialiasNone)
	context.Save()
	context.Scale(10, 10)
	pattern := cairo.CreatePatternForSurface(input)
	pattern.SetFilter(cairo.FilterNearest)
	m := cairo.NewMatrix()
	m.InitTranslate(index, 0)
	pattern.SetMatrix(m)
	context.SetSource(pattern)

	context.FillRectangle(0, 0, 400, 8)
	context.Restore()

	context.SetSourceRGB(1, 0, 0)
	context.SetLineWidth(1)
	context.StrokeRectangle(1, 1, 9, 79)

	context.Save()
	context.Translate(width/2-100, height/2-100)
	col := int(index)
	var b byte
	n := 0
	for row := 0; row < 8; row++ {
		i := (row*100*100 + col) * 4
		value := encodedData[i]
		if value == 0 {
			b += (1 << n)
		}
		n++
	}
	// fmt.Println(b)
	g := float64(b) / 255.0
	context.SetSourceGray(g)
	x := math.Mod(index, 100)
	y := math.Floor(index / 100)
	context.FillRectangle(x*2, y*2, 2, 2)

	index++
	context.Restore()
	context.SetSourceGray(g)
	context.FillRectangle(10, 100, 20, 20)
}

func encode() {
	sbw, err := cairo.NewSurfaceFromPNG("sbw.png")
	if err != nil {
		log.Fatal(err)
	}
	data, err := sbw.GetData()
	if err != nil {
		log.Fatal(err)
	}
	surface := cairo.NewSurface(100*100, 8)
	context := cairo.NewContext(surface)
	context.ClearWhite()
	col := 0.0
	res := 1
	for y := 0; y < 100; y += res {
		for x := 0; x < 100; x += res {
			index := (y*100 + x) * 4
			value := data[index]
			bits := getBits(value)
			row := 0.0
			for _, bit := range bits {
				if bit {
					context.FillRectangle(col, row, 1, 1)
				}
				row++
			}
			col++
		}
	}
	surface.WriteToPNG("encoded.png")
	render.ViewImage("encoded.png")
}

func getBits(value byte) [8]bool {
	bits := [8]bool{}
	bits[0] = value&1 > 0
	bits[1] = value&2 > 0
	bits[2] = value&4 > 0
	bits[3] = value&8 > 0
	bits[4] = value&16 > 0
	bits[5] = value&32 > 0
	bits[6] = value&64 > 0
	bits[7] = value&128 > 0
	return bits

}
