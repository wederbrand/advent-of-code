package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
)

type point struct {
	x     int
	y     int
	trace map[string]bool
}

func newPoint() point {
	p := point{
		x:     0,
		y:     0,
		trace: make(map[string]bool),
	}
	p.trace[p.key()] = true
	return p
}

func (p point) key() string {
	return key(p.x, p.y)
}

func key(x int, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

func main() {
	inFile := util.GetFileContents("2022/09/input.txt", "\n")

	rope := make([]point, 10)
	for i := 0; i < len(rope); i++ {
		rope[i] = newPoint()
	}

	for _, s := range inFile {
		var dir string
		var steps int
		fmt.Sscanf(s, "%s %d", &dir, &steps)
		//fmt.Println("moving", steps, dir)
		for i := 0; i < steps; i++ {
			// head moves
			switch dir {
			case "U":
				rope[0].y--
			case "D":
				rope[0].y++
			case "L":
				rope[0].x--
			case "R":
				rope[0].x++
			}

			// tails follow
			for i2 := range rope {
				if i2 == 0 {
					continue
				}
				follow(&rope[i2-1], &rope[i2])
			}

		}
	}

	fmt.Println("part 1:", len(rope[1].trace))
	fmt.Println("part 2:", len(rope[9].trace))
}

func follow(head *point, tail *point) {
	dx := head.x - tail.x
	dy := head.y - tail.y

	if dx == 0 && dy == 0 {
		return
	}
	if (dx == 1 || dx == -1) && (dy == 1 || dy == -1) {
		return
	}

	switch {
	// vertical
	case dx == 0:
		if dy == -2 {
			tail.y--
		} else if dy == 2 {
			tail.y++
		}
	// horizontal
	case dy == 0:
		if dx == -2 {
			tail.x--
		} else if dx == 2 {
			tail.x++
		}
	default:
		// diagonal
		if dy < 0 {
			tail.y--
		} else if dy > 0 {
			tail.y++
		}

		if dx < 0 {
			tail.x--
		} else if dx > 0 {
			tail.x++
		}
	}
	tail.trace[tail.key()] = true
}
