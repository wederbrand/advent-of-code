package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2024/15/input.txt", "\n")
	chartData1 := make([]string, 0)
	chartData2 := make([]string, 0)

	spaceIndex := 0
	for i, s := range inFile {
		if s == "" {
			spaceIndex = i
			break
		}

		chartData1 = append(chartData1, s)
		// for part 2 duplicate all
		line := ""
		for _, r := range s {
			if r == '.' {
				line += ".."
			} else if r == '#' {
				line += "##"
			} else if r == 'O' {
				line += "[]"
			} else if r == '@' {
				line += "@."
			} else {
				panic("hoho")
			}
		}
		chartData2 = append(chartData2, line)
	}

	moves := strings.Join(inFile[spaceIndex+1:], "")

	c1 := MakeChart(chartData1, "")
	doIt(c1, moves)
	part1 := scoreIt(c1)
	fmt.Println("Part 1:", part1, "in", time.Since(start))

	c2 := MakeChart(chartData2, "")
	doIt(c2, moves)
	part2 := scoreIt(c2)
	fmt.Println("Part 2:", part2, "in", time.Since(start))
}

func doIt(c Chart, moves string) {
	robot := Coord{}

	for coord, s := range c {
		if s == "@" {
			robot = coord
			break
		}
	}

	for _, r := range moves {
		d := UP
		switch r {
		case '^':
			d = UP
		case 'v':
			d = DOWN
		case '<':
			d = LEFT
		case '>':
			d = RIGHT
		}

		if pushIt(c, robot, d, false) {
			pushIt(c, robot, d, true)
			robot = robot.Move(d)
		}
	}
}

func scoreIt(c Chart) int {
	score := 0
	for coord, s := range c {
		if s == "O" || s == "[" {
			score += coord.X + coord.Y*100
		}
	}
	return score
}

func pushIt(c Chart, object Coord, d Dir, move bool) bool {
	next := object.Move(d, 1)
	if c[next] == "#" {
		// wall, that's the end of it
		return false
	}

	if c[next] == "." {
		// empty space, no more pushing
		if move {
			c[object], c[next] = c[next], c[object]
		}
		return true
	}

	if c[next] == "O" || ((c[next] == "[" || c[next] == "]") && (d == LEFT || d == RIGHT)) {
		// push the food first
		if pushIt(c, next, d, move) {
			if move {
				c[object], c[next] = c[next], c[object]
			}
			return true
		} else {
			return false
		}
	}

	if (c[next] == "[" || c[next] == "]") && (d == UP || d == DOWN) {
		// also push the othe half
		theOtherHalf := next.Move(RIGHT)
		if c[next] == "]" {
			theOtherHalf = next.Move(LEFT)
		}
		if pushIt(c, next, d, move) && pushIt(c, theOtherHalf, d, move) {
			if move {
				c[object], c[next] = c[next], c[object]
			}
			return true
		} else {
			return false
		}
	}

	panic("ho ho")
}
