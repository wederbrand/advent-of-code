package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"math"
	"time"
)

type Row struct {
	pixel []int
}

type Layer struct {
	rows []Row
}

func createLayers(input string, width int, height int) []Layer {
	result := make([]Layer, 0)
	for len(input) > 0 {
		l := Layer{}
		for i := 0; i < height; i++ {
			r := Row{}
			for j := 0; j < width; j++ {
				r.pixel = append(r.pixel, util.Atoi(string(input[0])))
				input = input[1:]
			}
			l.rows = append(l.rows, r)

		}
		result = append(result, l)
	}

	return result
}

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2019/08/input.txt", "\n")

	layers := createLayers(inFile[0], 25, 6)

	minZeros := math.MaxInt
	var minLayer Layer
	for _, layer := range layers {
		zeros := 0
		for _, row := range layer.rows {
			for _, p := range row.pixel {
				if p == 0 {
					zeros++
				}
			}
		}
		if zeros < minZeros {
			minZeros = zeros
			minLayer = layer
		}
	}

	cnt := make(map[int]int)
	for _, row := range minLayer.rows {
		for _, p := range row.pixel {
			cnt[p]++
		}
	}

	part1 := cnt[1] * cnt[2]
	fmt.Println("part1: ", part1, "in", time.Since(start))

	fmt.Println("part2")
	c := chart.Chart{}

	for y := -1; y < 7; y++ {
		for x := -1; x < 26; x++ {
			c[chart.Coord{x, y}] = "#"
		}
	}

	for y := 0; y < 6; y++ {
		for x := 0; x < 25; x++ {
			for _, layer := range layers {
				p := layer.rows[y].pixel[x]
				if p == 0 {
					c[chart.Coord{x, y}] = "#"
					break
				} else if p == 1 {
					c[chart.Coord{x, y}] = " "
					break
				}
			}
		}
	}

	chart.PrintChart(c)
}
