package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/priorityqueue"
	"math"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2024/18/input.txt", "\n")

	m := Chart{}

	c := Coord{0, 0}
	exit := Coord{70, 70}
	minC := Coord{0, 0}
	maxC := Coord{70, 70}

	for i, s := range inFile {
		var x, y int
		fmt.Sscanf(s, "%d,%d", &x, &y)
		m[Coord{x, y}] = strconv.Itoa(i)
	}

	steps := walkIt(c, exit, minC, maxC, m, 1024)
	fmt.Println("Part 1:", steps, "in", time.Since(start))

	wasBlocked := func(testSteps int) bool {
		steps := walkIt(c, exit, minC, maxC, m, testSteps)
		if steps == -1 {
			return true
		}
		return false
	}
	_, first := BinarySearch(0, len(inFile), wasBlocked)

	fmt.Println("Part 2:", inFile[first], "in", time.Since(start))
}

func walkIt(start Coord, exit Coord, minC Coord, maxC Coord, m Chart, testSteps int) int {
	q := priorityqueue.NewQueue()
	q.Add(&priorityqueue.State{Data: start, Priority: 0})

	seen := make(map[Coord]bool)
	for q.HasNext() {
		s := q.Next()
		c := s.Data.(Coord)
		if c == exit {
			return s.Priority
		}

		if seen[c] {
			continue
		}
		seen[c] = true

		for _, dir := range ALL {
			next := c.Move(dir)
			stringVal, found := m[next]
			intVal := math.MaxInt
			if found {
				intVal = Atoi(stringVal)
			}
			if intVal <= testSteps || next.X < minC.X || next.Y < minC.Y || next.X > maxC.X || next.Y > maxC.Y {
				continue
			}
			q.Add(&priorityqueue.State{Data: next, Priority: s.Priority + 1})
		}
	}

	return -1
}
