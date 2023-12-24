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

// GetDay returns the day based on the passed string.
func GetDay(day string) Day {
	return daymap[day]
}

var daymap = map[string]Day{
	"day01": Day01,
	"day02": Day02,
	"day03": Day03,
	"day04": Day04,
	"day05": Day05,
	"day06": Day06,
	"day07": Day07,
	"day08": Day08,
	"day09": Day09,
	"day10": Day10,
	"day11": Day11,
	"day12": Day12,
	"day13": Day13,
	"day14": Day14,
	"day15": Day15,
	"day16": Day16,
	"day17": Day17,
	"day18": Day18,
	"day19": Day19,
	"day20": Day20,
	"day21": Day21,
	"day22": Day22,
	"day23": Day23,
	"day24": Day24,
	"day25": Day25,
	"day26": Day26,
	"day27": Day27,
	"day28": Day28,
	"day29": Day29,
	"day30": Day30,
	"day31": Day31,
}
