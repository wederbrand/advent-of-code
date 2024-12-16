package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/priorityqueue"
	"math"
	"time"
)

type state struct {
	// I can get here at this price
	c     Coord
	d     Dir
	score int
	path  []Coord
}

type seenKey struct {
	c Coord
	d Dir
}

func main() {
	start := time.Now()
	inFile := GetFileContents("2024/16/input.txt", "\n")

	m := MakeChart(inFile, ".")

	c := Coord{}
	e := Coord{}
	for coord, s := range m {
		if s == "S" {
			c = coord
		}
		if s == "E" {
			e = coord
		}
	}

	p1, p2 := doIt(c, e, m)
	fmt.Println("Part 1:", p1, "in", time.Since(start))
	fmt.Println("Part 2:", p2, "in", time.Since(start))
}

func doIt(c Coord, e Coord, m Chart) (int, int) {
	q := priorityqueue.NewQueue()
	initial := priorityqueue.State{
		Data:     state{c, E, 0, []Coord{c}},
		Priority: 0,
	}

	q.Add(&initial)

	paths := make(Chart)
	paths[c] = "O"
	paths[e] = "O"

	seen := make(map[seenKey]int)
	minFound := math.MaxInt
	for q.HasNext() {
		next := q.Next()
		s := next.Data.(state)

		if s.score > minFound {
			// this path is already longer than the shortest path found
			break
		}

		key := seenKey{s.c, s.d}
		seenBefore, found := seen[key]
		if found && seenBefore < s.score {
			continue
		}
		seen[key] = s.score

		if s.c == e {
			minFound = s.score
			for _, coord := range s.path {
				paths[coord] = "O"
			}
		}

		// queue all possible moves
		// move forward
		if m[s.c.Move(s.d)] != "#" {
			newPath := CloneSlice(s.path)
			newPath = append(newPath, s.c.Move(s.d))
			forward := state{s.c.Move(s.d), s.d, next.Data.(state).score + 1, newPath}
			q.Add(&priorityqueue.State{Data: forward, Priority: forward.score})
		}

		// turn left if there is no wall there
		if m[s.c.Move(s.d.Left())] != "#" {
			left := state{s.c, s.d.Left(), next.Data.(state).score + 1000, s.path}
			q.Add(&priorityqueue.State{Data: left, Priority: left.score})
		}

		// turn right if there is no wall there
		if m[s.c.Move(s.d.Right())] != "#" {
			right := state{s.c, s.d.Right(), next.Data.(state).score + 1000, s.path}
			q.Add(&priorityqueue.State{Data: right, Priority: right.score})
		}
	}
	for coord, s := range paths {
		m[coord] = s
	}
	return minFound, len(paths)
}
