// Package days contains genuary 2024 code for each day.
package days

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/target"
)

// Day14 is for genuary 14
var Day14 = Day{
	ImageWidth:  150,
	ImageHeight: 100,
	VideoWidth:  16,
	VideoHeight: 16,
	VideoTime:   1,
	RenderFrame: Day14Render,
	Target:      target.Image,
}

// Day14Render is for genuary 14
// Less than 1KB artwork.
//
//revive:disable-next-line:unused-parameter
func Day14Render(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	// in the end, compress with:
	// convert -gaussian-blur 0.01 -quality 24% day14.png day14.jpg
	// should come out at 999 bytes

	iter := 16.0

	for x := 0.0; x < width; x++ {
		for y := 0.0; y < height; y++ {
			r := blmath.Map(x, 0, width, -2, 1)
			i := blmath.Map(y, 0, height, -1, 1)
			z := complex(0, 0)
			c := complex(r, i)

			for t := 0.0; t < iter; t++ {
				z = z*z + c
				if math.Hypot(real(z), imag(z)) > 2 {
					g := t / iter
					context.SetSourceGray(g)
					context.FillRectangle(x, y, 1, 1)
					break
				}
			}
		}
	}
}
