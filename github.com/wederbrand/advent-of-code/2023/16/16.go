package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/priorityqueue"
	"math"
	"time"
)

type Beam struct {
	c   Coord
	dir Dir
}

func main() {
	startTimer := time.Now()
	inFile := util.GetFileContents("2023/16/input.txt", "\n")
	m := MakeChart(inFile, ".")

	start := Beam{c: Coord{0, 0}, dir: RIGHT}
	energized := doIt(m, start)
	fmt.Println("part1: ", len(energized), "in", time.Since(startTimer))

	minC, maxC := GetChartMaxes(m)
	part2 := math.MinInt
	for x := minC.X; x < maxC.X; x++ {
		start = Beam{Coord{x, minC.Y}, DOWN}
		energized = doIt(m, start)
		part2 = max(part2, len(energized))

		start = Beam{Coord{x, maxC.Y}, UP}
		energized = doIt(m, start)
		part2 = max(part2, len(energized))
	}
	for y := minC.Y; y < maxC.Y; y++ {
		start = Beam{Coord{minC.X, y}, RIGHT}
		energized = doIt(m, start)
		part2 = max(part2, len(energized))

		start = Beam{Coord{maxC.X, y}, LEFT}
		energized = doIt(m, start)
		part2 = max(part2, len(energized))
	}
	fmt.Println("part2: ", part2, "in", time.Since(startTimer))
}

func doIt(m Chart, start Beam) map[Coord]bool {
	q := priorityqueue.NewQueue()
	state := &priorityqueue.State{start, 0}
	q.Add(state)

	minC, maxC := GetChartMaxes(m)

	seen := make(map[Beam]bool)
	energized := make(map[Coord]bool)
	for q.HasNext() {
		state = q.Next()
		beam := state.Data.(Beam)

		if beam.c.X < minC.X || beam.c.X > maxC.X || beam.c.Y < minC.Y || beam.c.Y > maxC.Y {
			// do nothing
			continue
		}

		_, found := seen[beam]
		if found {
			// we've been here before
			continue
		}
		seen[beam] = true
		energized[beam.c] = true

		tile, found := m[beam.c]
		if !found {
			// if empty queue next tile and don't change direction
			move(beam, beam.dir, q)
		} else {
			if tile == "-" {
				if beam.dir == RIGHT || beam.dir == LEFT {
					move(beam, beam.dir, q)
				} else {
					move(beam, RIGHT, q)
					move(beam, LEFT, q)
				}
			} else if tile == "|" {
				if beam.dir == UP || beam.dir == DOWN {
					move(beam, beam.dir, q)
				} else {
					move(beam, UP, q)
					move(beam, DOWN, q)
				}
			} else if tile == "/" && beam.dir == UP {
				move(beam, RIGHT, q)
			} else if tile == "/" && beam.dir == DOWN {
				move(beam, LEFT, q)
			} else if tile == "/" && beam.dir == RIGHT {
				move(beam, UP, q)
			} else if tile == "/" && beam.dir == LEFT {
				move(beam, DOWN, q)
			} else if tile == "\\" && beam.dir == UP {
				move(beam, LEFT, q)
			} else if tile == "\\" && beam.dir == DOWN {
				move(beam, RIGHT, q)
			} else if tile == "\\" && beam.dir == RIGHT {
				move(beam, DOWN, q)
			} else if tile == "\\" && beam.dir == LEFT {
				move(beam, UP, q)
			}
		}
	}
	return energized
}

func move(beam Beam, dir Dir, q *priorityqueue.Queue) {
	beam.c = Coord{beam.c.X + dir[0], beam.c.Y + dir[1]}
	beam.dir = dir
	q.Add(&priorityqueue.State{beam, 0})
}
