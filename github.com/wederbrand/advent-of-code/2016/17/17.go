package main

import (
	"crypto/md5"
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/priorityqueue"
	"time"
)

func main() {
	start := time.Now()
	input := "veumntbg"

	startC := Coord{0, 0}
	endC := Coord{3, 3}

	paths := walkIt(startC, endC, input)

	fmt.Println("Part 1: ", paths[0], "in", time.Since(start))
	part2 := 0
	for _, p := range paths {
		if len(p) > part2 {
			part2 = len(p)
		}
	}
	fmt.Println("Part 2: ", part2, "in", time.Since(start))
}

type dirState struct {
	current Coord
	dirs    string
}

func walkIt(start Coord, end Coord, input string) []string {
	q := priorityqueue.NewQueue()
	q.Add(&priorityqueue.State{Data: dirState{current: start, dirs: ""}})

	foundPaths := make([]string, 0)
	for q.HasNext() {
		s := q.Next()
		ds := s.Data.(dirState)
		c := ds.current

		if c == end {
			foundPaths = append(foundPaths, ds.dirs)
			continue
		}

		slask := input + ds.dirs
		sum := fmt.Sprintf("%x", md5.Sum([]byte(slask)))
		upHash := rune(sum[0])
		downHash := rune(sum[1])
		leftHash := rune(sum[2])
		rightHash := rune(sum[3])

		up := c.Move(UP)
		if up.Y >= 0 {
			// not a wall
			if upHash >= 'b' && upHash <= 'f' {
				// open
				nextState := priorityqueue.State{Data: dirState{
					current: up,
					dirs:    ds.dirs + "U",
				}, Priority: s.Priority + 1}
				q.Add(&nextState)
			}
		}

		down := c.Move(DOWN)
		if down.Y < 4 {
			// not a wall
			if downHash >= 'b' && downHash <= 'f' {
				// open
				nextState := priorityqueue.State{Data: dirState{
					current: down,
					dirs:    ds.dirs + "D",
				}, Priority: s.Priority + 1}
				q.Add(&nextState)
			}
		}

		left := c.Move(LEFT)
		if left.X >= 0 {
			// not a wall
			if leftHash >= 'b' && leftHash <= 'f' {
				// open
				nextState := priorityqueue.State{Data: dirState{
					current: left,
					dirs:    ds.dirs + "L",
				}, Priority: s.Priority + 1}
				q.Add(&nextState)
			}
		}

		right := c.Move(RIGHT)
		if right.X < 4 {
			// not a wall
			if rightHash >= 'b' && rightHash <= 'f' {
				// open
				nextState := priorityqueue.State{Data: dirState{
					current: right,
					dirs:    ds.dirs + "R",
				}, Priority: s.Priority + 1}
				q.Add(&nextState)
			}
		}
	}
	return foundPaths
}
