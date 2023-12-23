// Package days contains genuary 2024 code for each day.
package days

import "github.com/bit101/blcairo/render"

// Day is the structure of a single day experiment.
type Day struct {
	ImageWidth  float64
	ImageHeight float64
	VideoWidth  float64
	VideoHeight float64
	VideoTime   int
	RenderFrame render.FrameFunc
	Target      int
}
