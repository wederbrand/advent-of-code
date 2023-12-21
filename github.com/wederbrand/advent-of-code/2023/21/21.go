package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"strconv"
	"time"
)

var allDir = [4]Dir{UP, RIGHT, DOWN, LEFT}

type PosStep struct {
	c     Coord
	steps int
}

func main() {
	startTimer := time.Now()
	inFile := util.GetFileContents("2023/21/input.txt", "\n")

	m := MakeChart(inFile, "")
	for c, s := range m {
		if s == "S" {
			m[c] = "."
		}
	}

	_, maxC := GetChartMaxes(m)

	result := 0
	steps := 26501365
	mirrors := steps / (maxC.X + 1)

	oddBox := doOneFindOdd(Coord{65, 65}, m, 131, 0, false)
	evenBox := doOneFindOdd(Coord{65, 65}, m, 131+1, 1, false)

	result += oddBox * (mirrors - 1) * (mirrors - 1)
	result += evenBox * (mirrors) * (mirrors)

	normalBorderSteps := 131 + 65 - 1
	topRightNormalBorder := doOneFindOdd(Coord{0, 130}, m, normalBorderSteps, 0, false)
	topLeftNormalBorder := doOneFindOdd(Coord{130, 130}, m, normalBorderSteps, 0, false)
	bottomRightNormalBorder := doOneFindOdd(Coord{0, 0}, m, normalBorderSteps, 0, false)
	bottomLeftNormalBorder := doOneFindOdd(Coord{130, 0}, m, normalBorderSteps, 0, false)

	totalNormalBorder := topLeftNormalBorder + topRightNormalBorder + bottomLeftNormalBorder + bottomRightNormalBorder
	result += totalNormalBorder * (mirrors - 1)

	bonusBorderSteps := 65
	topRightBonusBorder := doOneFindOdd(Coord{0, 130}, m, bonusBorderSteps, 1, false)
	topLeftBonusBorder := doOneFindOdd(Coord{130, 130}, m, bonusBorderSteps, 1, false)
	bottomRightBonusBorder := doOneFindOdd(Coord{0, 0}, m, bonusBorderSteps, 1, false)
	bottomLeftBonusBorder := doOneFindOdd(Coord{130, 0}, m, bonusBorderSteps, 1, false)

	totalBonusBorder := topLeftBonusBorder + topRightBonusBorder + bottomLeftBonusBorder + bottomRightBonusBorder
	result += totalBonusBorder * (mirrors)

	cornerSteps := 131
	rightCorner := doOneFindOdd(Coord{0, 65}, m, cornerSteps, 1, false)
	leftCorner := doOneFindOdd(Coord{130, 65}, m, cornerSteps, 1, false)
	bottomCorner := doOneFindOdd(Coord{65, 0}, m, cornerSteps, 1, false)
	topCorner := doOneFindOdd(Coord{65, 130}, m, cornerSteps, 1, false)

	totalCorner := rightCorner + leftCorner + topCorner + bottomCorner

	result += totalCorner

	fmt.Println("part2: ", result, "in", time.Since(startTimer))
}

func doOneFindOdd(c Coord, m Chart, maxSteps int, oddEven int, printIt bool) int {
	visited := doOne(c, m, maxSteps, oddEven)
	if printIt {
		printVisited(visited)
	}
	result := 0
	for _, i := range visited {
		if i%2 == 1 {
			result++
		}
	}
	return result
}

func printVisited(v map[Coord]int) {
	m := make(map[Coord]string)

	for coord, i := range v {
		m[coord] = strconv.Itoa(i % 2)
	}

	PrintChart(m)
}

func doOne(start Coord, m Chart, maxSteps int, oddEven int) map[Coord]int {
	visited := make(map[Coord]int)
	visited[start] = oddEven

	q := make([]PosStep, 0)
	q = append(q, PosStep{start, oddEven})

	for len(q) > 0 {
		s := q[0]
		q = q[1:]

		if s.steps > maxSteps {
			continue
		}

		for _, dir := range allDir {
			nextC := s.c.Move(dir)
			foundLot := m[nextC]
			_, visitedBefore := visited[nextC]
			if foundLot == "." && !visitedBefore {
				// lot
				visited[nextC] = s.steps + 1
				q = append(q, PosStep{nextC, s.steps + 1})
			}
		}
	}

	return visited
}
