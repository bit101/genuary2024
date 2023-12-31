// Package days contains genuary 2024 code for each day.
package days

import (
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/target"
)

// Day07 is for genuary 7
var Day07 = Day{
	ImageWidth:  400,
	ImageHeight: 400,
	VideoWidth:  400,
	VideoHeight: 400,
	VideoTime:   10,
	RenderFrame: Day07Render,
	Target:      target.Video,
}

// Day07Render is for genuary 7
// Progress bar / indicator / loading animation.
//
//revive:disable-next-line:unused-parameter
func Day07Render(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()

}
