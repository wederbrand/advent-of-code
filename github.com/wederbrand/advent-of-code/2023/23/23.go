package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"golang.org/x/exp/maps"
	"math"
	"slices"
	"time"
)

type Node struct {
	c       Coord
	targets map[Coord]int
}

var allNodes = make(map[Coord]Node)

func main() {
	startTimer := time.Now()
	inFile := util.GetFileContents("2023/23/input.txt", "\n")

	m := MakeChart(inFile, "")
	minC, maxC := GetChartMaxes(m)
	start := Node{Coord{minC.X + 1, minC.Y}, make(map[Coord]int)}
	end := Coord{maxC.X - 1, maxC.Y}
	fmt.Println("parsing:", time.Since(startTimer))
	startTimer = time.Now()

	part2 := solveIt(start, end, m, true)
	fmt.Println("part2: ", part2, "in", time.Since(startTimer))
}

func solveIt(start Node, end Coord, m Chart, ignoreArrows bool) int {
	buildGraph(&start, start.c.Move(DOWN), end, m, start.c, ignoreArrows)
	result := traverse(start, end, make([]Coord, 0))
	maps.Clear(allNodes)
	return result
}

func traverse(start Node, end Coord, seen []Coord) int {
	maxSteps := math.MinInt
	for t, steps := range start.targets {
		if slices.Contains(seen, t) {
			// been here, don't go back
		} else {
			if t == end {
				return steps
			}

			targetNode := allNodes[t]
			newSeen := slices.Clone(seen)
			newSeen = append(newSeen, start.c)
			i := steps + traverse(targetNode, end, newSeen)
			maxSteps = max(maxSteps, i)
		}
	}
	return maxSteps
}

func buildGraph(startNode *Node, startPos Coord, end Coord, m Chart, last Coord, ignoreArrows bool) {
	steps := 1
	c := startPos

	// walk this path until reaching a cross road
	// as we go increment a counter
	// or end
	out, outString := c.AllBut(last, m, "#")
	_ = outString
	for len(out) == 1 {
		steps++
		last, c = c, out[0]
		if c == end {
			startNode.targets[end] = steps
			return
		}
		out, outString = c.AllBut(last, m, "#")
	}

	if len(out) == 0 {
		// dead end
		return
	}

	newNode := Node{c, make(map[Coord]int)}
	startNode.targets[c] = steps

	_, found := allNodes[c]
	if found {
		return
	}
	allNodes[c] = *startNode

	for _, nextPos := range out {
		buildGraph(&newNode, nextPos, end, m, newNode.c, false)
	}
	buildGraph(&newNode, newNode.c, end, m, last, false)
}
