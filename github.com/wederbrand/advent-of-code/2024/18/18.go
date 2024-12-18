package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/priorityqueue"
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
		if i == 1024 {
			steps := walkIt(c, exit, minC, maxC, m)

			fmt.Println("Part 1:", steps, "in", time.Since(start))
		}
		var x, y int
		fmt.Sscanf(s, "%d,%d", &x, &y)
		m[Coord{x, y}] = "#"

		if i > 1024 {
			steps := walkIt(c, exit, minC, maxC, m)

			if steps == -1 {
				fmt.Println("Part 2:", s, "in", time.Since(start))
				break
			}
		}
	}
}

func walkIt(start Coord, exit Coord, minC Coord, maxC Coord, m Chart) int {
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
			if m[next] == "#" || next.X < minC.X || next.Y < minC.Y || next.X > maxC.X || next.Y > maxC.Y {
				continue
			}
			q.Add(&priorityqueue.State{Data: next, Priority: s.Priority + 1})
		}
	}

	return -1
}
