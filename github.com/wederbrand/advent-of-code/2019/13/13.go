package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/2019/computer"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"strconv"
	"time"
)

const BLOCK = "2"
const PADDLE = 3
const BALL = 4

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2019/13/input.txt", "\n")

	m := chart.Chart{}
	output := make([]int, 0)
	out := func(i int) {
		output = append(output, i)
		if len(output) == 3 {
			x := output[0]
			y := output[1]
			tile := output[2]

			m[chart.Coord{x, y}] = strconv.Itoa(tile)
			output = nil
		}
	}

	computer := NewComputer(inFile, nil, out)
	computer.Run()

	part1 := 0
	part1 = getCount(m, BLOCK)
	fmt.Println("part1: ", part1, "in", time.Since(start))

	part2 := 0
	ball := chart.Coord{0, 0}
	paddle := chart.Coord{0, 0}
	clear(m)

	output = make([]int, 0)
	out = func(i int) {
		output = append(output, i)
		if len(output) == 3 {
			x := output[0]
			y := output[1]
			tile := output[2]
			if x == -1 && y == 0 {
				part2 = tile
			}
			if tile == BALL {
				ball.X = x
				ball.Y = y
			} else if tile == PADDLE {
				paddle.X = x
				paddle.Y = y
			}

			m[chart.Coord{x, y}] = strconv.Itoa(tile)

			output = nil
		}
	}

	in := func() int {
		if paddle.X < ball.X {
			return 1
		} else if paddle.X > ball.X {
			return -1
		} else {
			return 0
		}
	}

	computer = NewComputer(inFile, in, out)
	computer.SetMemory(0, 2)
	computer.Run()

	fmt.Println("part2: ", part2, "in", time.Since(start))
}

func getCount(m chart.Chart, tile string) int {
	result := 0
	for _, s := range m {
		if s == tile {
			result++
		}
	}
	return result
}
