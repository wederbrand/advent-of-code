package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"

	"regexp"
	"time"
	"unicode/utf8"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2018/12/input.txt", "\n")

	initialMatcher := regexp.MustCompile(`^initial state: (.+)$`)
	lineMatcher := regexp.MustCompile(`^(.+) => (.)$`)

	initialMatch := initialMatcher.FindStringSubmatch(inFile[0])
	inputPots := initialMatch[1]

	pots := MakeChart([]string{inputPots}, ".")

	mapper := make(map[string]bool)
	for _, s := range inFile[2:] {
		lineMatch := lineMatcher.FindStringSubmatch(s)
		r, _ := utf8.DecodeRuneInString(lineMatch[2])
		if r == '#' {
			mapper[lineMatch[1]] = true
		}
	}

	for i := 0; i < 20; i++ {
		pots = doIt(pots, mapper)
	}

	part1 := sum(pots)
	fmt.Println("Part 1:", part1, "in", time.Since(start))

	// from observed printouts it gets stable after a while
	stableAfter := 200
	for i := 20; i < stableAfter; i++ {
		pots = doIt(pots, mapper)
	}
	sum1 := sum(pots)
	pots = doIt(pots, mapper)
	sum2 := sum(pots)
	inc := sum2 - sum1
	part2 := sum2 + (50000000000-stableAfter-1)*inc

	fmt.Println("Part 2:", part2, "in", time.Since(start))
}

func sum(pots Chart) int {
	result := 0
	for c := range pots {
		result += c.X
	}
	return result
}

func doIt(pots Chart, mapper map[string]bool) Chart {
	newPots := Chart{}
	minC, maxC := GetChartMaxes(pots)
	for j := minC.X - 4; j < maxC.X+4; j++ {
		neighbours := getNeighbours(pots, j)
		if mapper[neighbours] {
			newPots[Coord{X: j}] = "#"
		}
	}

	return newPots
}

func getNeighbours(pots Chart, x int) string {
	neighbours := ""
	for i := x - 2; i <= x+2; i++ {
		if pots[Coord{X: i}] == "#" {
			neighbours += "#"
		} else {
			neighbours += "."
		}
	}
	return neighbours
}
