package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/priorityqueue"
	"slices"
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

	var currentChart Chart
	for i, s := range inFile {
		if i == 1024 {
			path := walkIt(c, exit, minC, maxC, m)

			fmt.Println("Part 1:", len(path.steps), "in", time.Since(start))

			currentChart = Chart{}
			for _, step := range path.steps {
				currentChart[step] = "O"
			}
		}
		var x, y int
		fmt.Sscanf(s, "%d,%d", &x, &y)
		nextByte := Coord{x, y}
		m[nextByte] = "#"

		if i > 1024 {
			if currentChart[nextByte] == "O" {
				// This is where we used to walk, take a new route
				path := walkIt(c, exit, minC, maxC, m)

				if len(path.steps) == 0 {
					fmt.Println("Part 2:", s, "in", time.Since(start))
					break
				}

				currentChart = Chart{}
				for _, step := range path.steps {
					currentChart[step] = "O"
				}
			}
		}
	}
}

type Path struct {
	steps []Coord
}

func walkIt(start Coord, exit Coord, minC Coord, maxC Coord, m Chart) Path {
	q := priorityqueue.NewQueue()
	q.Add(&priorityqueue.State{Data: Path{[]Coord{start}}, Priority: 0})

	seen := make(map[Coord]bool)
	for q.HasNext() {
		s := q.Next()
		p := s.Data.(Path)
		c := p.steps[len(p.steps)-1]

		if c == exit {
			return p
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
			newPath := Path{slices.Clone(p.steps)}
			newPath.steps = append(newPath.steps, next)
			q.Add(&priorityqueue.State{Data: newPath, Priority: s.Priority + 1})
		}
	}

	return Path{[]Coord{}}
}
