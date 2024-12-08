package chart

import (
	"fmt"
	"math"
	"strings"
)

type Dir [2]int

var N = Dir{0, -1}
var S = Dir{0, +1}
var E = Dir{+1, 0}
var W = Dir{-1, 0}

var NW = Dir{-1, -1}
var NE = Dir{+1, -1}
var SW = Dir{-1, +1}
var SE = Dir{+1, +1}

var UP = Dir{0, -1}
var DOWN = Dir{0, +1}
var RIGHT = Dir{+1, 0}
var LEFT = Dir{-1, 0}

var UPLEFT = Dir{-1, -1}
var UPRIGHT = Dir{+1, -1}
var DOWNLEFT = Dir{-1, +1}
var DOWNRIGHT = Dir{+1, +1}

var ALL = [4]Dir{UP, RIGHT, DOWN, LEFT}
var ALL_AND_DIAG = [8]Dir{UP, UPRIGHT, RIGHT, DOWNRIGHT, DOWN, DOWNLEFT, LEFT, UPLEFT}

func (d Dir) Left() Dir {
	switch d {
	case UP:
		return LEFT
	case LEFT:
		return DOWN
	case DOWN:
		return RIGHT
	case RIGHT:
		return UP
	default:
		panic("hoho")
	}
}

func (d Dir) Right() Dir {
	switch d {
	case UP:
		return RIGHT
	case LEFT:
		return UP
	case DOWN:
		return LEFT
	case RIGHT:
		return DOWN
	default:
		panic("hoho")
	}
}

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

func (c Coord) AllBut(last Coord, m Chart, but string) ([]Coord, []string) {
	out := make([]Coord, 0)
	outString := make([]string, 0)
	for _, dir := range ALL {
		next := c.Move(dir)
		if next == last {
			continue
		}
		s2 := m[next]

		if !strings.Contains(but, s2) {
			out = append(out, next)
			outString = append(outString, s2)
		}
	}

	return out, outString
}

type Chart map[Coord]string

func MakeChart(in []string, ignored string) Chart {
	m := make(Chart)

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

func Find(m Chart, target string) (bool, Coord) {
	for k, v := range m {
		if v == target {
			return true, k
		}
	}
	return false, Coord{0, 0}
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
