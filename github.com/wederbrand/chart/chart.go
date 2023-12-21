package chart

import (
	"fmt"
	"math"
	"strings"
)

type Dir [2]int

var N = Dir{0, -1}
var UP = Dir{0, -1}
var S = Dir{0, +1}
var DOWN = Dir{0, +1}
var E = Dir{+1, 0}
var RIGHT = Dir{+1, 0}
var W = Dir{-1, 0}
var LEFT = Dir{-1, 0}

type Coord struct {
	X int
	Y int
}

func (c Coord) Move(dir Dir, length ...int) Coord {
	l := 1
	if len(length) == 1 {
		l = length[0]
	}
	return Coord{c.X + dir[0]*l, c.Y + dir[1]*l}
}

type Chart map[Coord]string

func MakeChart(in []string, ignored string) Chart {
	m := make(map[Coord]string)

	for y, s := range in {
		for x, r := range s {
			if strings.ContainsRune(ignored, r) {
				continue
			}
			m[Coord{x, y}] = string(r)
		}
	}

	return m
}

func Manhattan(a Coord, b Coord) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}

func RotateClockWise(in Chart) Chart {
	// For my reversed Y clockwise is the same a normal counterclockwise
	// 90° counterclockwise rotation: (𝑥,𝑦) becomes (−𝑦,𝑥)

	out := make(Chart)

	for key, value := range in {
		out[Coord{-key.Y, key.X}] = value
	}

	return out
}

func RotateCounterClockWise(in Chart) Chart {
	// For my reversed Y counterclockwise is the same a normal clockwise
	// 90° clockwise rotation: (𝑥,𝑦) becomes (𝑦,-𝑥)

	out := make(Chart)

	for key, value := range in {
		out[Coord{key.Y, -key.X}] = value
	}

	return out
}

func GetChartMaxes(m Chart) (minC Coord, maxC Coord) {
	minX := math.MaxInt
	minY := math.MaxInt
	maxX := math.MinInt
	maxY := math.MinInt
	for k := range m {
		minX = min(minX, k.X)
		minY = min(minY, k.Y)
		maxX = max(maxX, k.X)
		maxY = max(maxY, k.Y)
	}
	return Coord{minX, minY}, Coord{maxX, maxY}
}

func PrintChart(m Chart) {
	minC, maxC := GetChartMaxes(m)
	for y := minC.Y; y <= maxC.Y; y++ {
		for x := minC.X; x <= maxC.X; x++ {
			s, found := m[Coord{x, y}]
			if !found {
				fmt.Print(".")
			} else {
				fmt.Print(s)
			}
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println()
}

func ChartAsString(m Chart) string {
	out := ""
	minC, maxC := GetChartMaxes(m)
	for y := minC.Y; y <= maxC.Y; y++ {
		for x := minC.X; x <= maxC.X; x++ {
			s, found := m[Coord{x, y}]
			if !found {
				out += "."
			} else {
				out += s
			}
		}
		out += "|"
	}

	return out
}
