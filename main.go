// Package main renders an image, gif or video
package main

import (
	"fmt"

	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/genuary2024/days"
)

func main() {

	day := "day01"
	theDay := daymap[day]

	switch theDay.Target {
	case target.Image:
		outfile := fmt.Sprintf("out/%s.png", day)
		render.Image(theDay.ImageWidth, theDay.ImageHeight, outfile, theDay.RenderFrame, 0.0)
		render.ViewImage(outfile)
		break

	case target.Video:
		outfile := fmt.Sprintf("out/%s.mp4", day)
		seconds := theDay.VideoTime
		fps := 30
		w := theDay.VideoWidth
		h := theDay.VideoHeight
		render.Frames(w, h, seconds*fps, "out/frames", theDay.RenderFrame)
		render.ConvertToVideo("out/frames", outfile, w, h, fps, seconds)
		render.PlayVideo(outfile)
		break
	}
}

var daymap = map[string]days.Day{
	"day01": days.Day01,
	"day02": days.Day02,
	"day03": days.Day03,
	"day04": days.Day04,
	"day05": days.Day05,
	"day06": days.Day06,
	"day07": days.Day07,
	"day08": days.Day08,
	"day09": days.Day09,
	"day10": days.Day10,
	"day11": days.Day11,
	"day12": days.Day12,
	"day13": days.Day13,
	"day14": days.Day14,
	"day15": days.Day15,
	"day16": days.Day16,
	"day17": days.Day17,
	"day18": days.Day18,
	"day19": days.Day19,
	"day20": days.Day20,
	"day21": days.Day21,
	"day22": days.Day22,
	"day23": days.Day23,
	"day24": days.Day24,
	"day25": days.Day25,
	"day26": days.Day26,
	"day27": days.Day27,
	"day28": days.Day28,
	"day29": days.Day29,
	"day30": days.Day30,
	"day31": days.Day31,
}
