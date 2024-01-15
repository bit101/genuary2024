// Package main renders an image, gif or video
package main

import (
	"fmt"

	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/genuary2024/days"
)

func main() {

	day := "day19"
	theDay := days.GetDay(day)

	switch theDay.Target {
	case target.Image:
		outfile := fmt.Sprintf("out/%s.png", day)
		render.Image(theDay.ImageWidth, theDay.ImageHeight, outfile, theDay.RenderFrame, 0.0)
		render.ViewImage(outfile)
		break

	case target.Video:
		outfile := fmt.Sprintf("out/%s.mp4", day)
		seconds := theDay.VideoTime
		fps := 60
		w := theDay.VideoWidth
		h := theDay.VideoHeight
		render.Frames(w, h, seconds*fps, "out/frames", theDay.RenderFrame)
		render.ConvertToVideo("out/frames", outfile, w, h, fps, seconds)
		render.PlayVideo(outfile)
		break
	}
}
