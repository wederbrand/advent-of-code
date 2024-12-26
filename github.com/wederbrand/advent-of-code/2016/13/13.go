package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/priorityqueue"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	startPos := Coord{X: 1, Y: 1}
	endPos := Coord{X: 31, Y: 39}
	input := 1358

	part1, part2 := walkIt(startPos, endPos, input)

	fmt.Println("Part 1: ", part1, "in", time.Since(start))
	fmt.Println("Part 2: ", part2, "in", time.Since(start))
}

func walkIt(start Coord, end Coord, input int) (int, int) {
	q := priorityqueue.NewQueue()
	q.Add(&priorityqueue.State{Data: start})

	seen := make(map[Coord]int)
	seen[start] = 0
	for q.HasNext() {
		s := q.Next()
		c := s.Data.(Coord)

		if c == end {
			sub50 := 0
			for _, i := range seen {
				if i <= 50 {
					sub50++
				}
			}
			return s.Priority, sub50
		}

		for _, dir := range ALL {
			next := c.Move(dir)
			if next.X < 0 || next.Y < 0 {
				continue
			}

			if isWall(next.X, next.Y, input) {
				continue
			}

			_, found := seen[next]
			if found {
				// we have been here before
				continue
			}

			seen[next] = s.Priority + 1
			nextState := priorityqueue.State{Data: next, Priority: s.Priority + 1}
			q.Add(&nextState)
		}
	}
	panic("No path found")
}

func isWall(x int, y int, input int) bool {
	value := input + x*x + 3*x + 2*x*y + y + y*y
	binary := strconv.FormatInt(int64(value), 2)
	if strings.Count(binary, "1")%2 == 0 {
		return false
	} else {
		return true
	}
}
