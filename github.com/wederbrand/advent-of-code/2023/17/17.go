package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/priorityqueue"
	"time"
)

type crucible struct {
	pos  Coord
	last Dir
	same int
	loss int
}

type seenKey struct {
	pos Coord
	dir Dir
	cnt int
}

func main() {
	startTimer := time.Now()
	inFile := util.GetFileContents("2023/17/input.txt", "\n")

	m := MakeChart(inFile, ".")
	part1 := doIt(m, 1, 3)
	fmt.Println("part1: ", part1, "in", time.Since(startTimer))
	startTimer = time.Now()
	part2 := doIt(m, 4, 10)
	fmt.Println("part2: ", part2, "in", time.Since(startTimer))
}

func doIt(m Chart, minStraight int, maxStraight int) int {
	q := priorityqueue.NewQueue()
	start := crucible{Coord{0, 0}, RIGHT, 1, 0}
	state := &priorityqueue.State{start, 0}
	q.Add(state)

	_, maxC := GetChartMaxes(m)

	seen := make(map[seenKey]int)
	for q.HasNext() {
		next := q.Next()
		c := next.Data.(crucible)

		// find illegal moves
		sk := seenKey{c.pos, c.last, c.same}
		_, found := seen[sk]
		if found {
			// we've been here before
			continue
		}
		seen[sk] = c.loss

		if c.pos == maxC {
			// done
			return c.loss
		}

		if c.same < maxStraight {
			queueNext(c, c.last, c.same+1, m, q)
		}
		if c.same >= minStraight {
			queueNext(c, c.last.Left(), 1, m, q)
			queueNext(c, c.last.Right(), 1, m, q)
		}
	}

	panic("ho ho")
}

func queueNext(c crucible, nextDir Dir, same int, m Chart, q *priorityqueue.Queue) {
	nextPos := Coord{c.pos.X + nextDir[0], c.pos.Y + nextDir[1]}
	nextHeat, found := m[nextPos]
	if found {
		newHeat := c.loss + util.Atoi(nextHeat)
		nextCrucible := crucible{nextPos, nextDir, same, newHeat}
		q.Add(&priorityqueue.State{nextCrucible, newHeat})
	}
}
