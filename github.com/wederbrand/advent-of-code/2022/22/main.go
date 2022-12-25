package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"time"
)

const size = 50

type point struct {
	x     int
	y     int
	value string
}

func newPoint(x int, y int, value string) *point {
	p := new(point)
	p.x = x
	p.y = y
	p.value = value

	return p
}

func (p point) key() string {
	return key(p.x, p.y)
}

func key(x int, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

func walk(steps int, x *int, y *int, dx *int, dy *int, m map[string]*point) {
	//fmt.Println("walking", steps, *dx, *dy)
	for i := 0; i < steps; i++ {
		nextX := *x + *dx
		nextY := *y + *dy
		nextDx := *dx // copy of value
		nextDy := *dy // copy of value
		p, found := m[key(nextX, nextY)]

		if !found {
			// out of bounds, warp
			nextX = *x
			nextY = *y
			warp(&nextX, &nextY, &nextDx, &nextDy)
			p = m[key(nextX, nextY)]
		}

		if p.value == "#" {
			// stone, walk no more
			return
		} else {
			p.value = "X"
			// no stone, walk
			*x = p.x
			*y = p.y
			*dx = nextDx
			*dy = nextDy
		}
	}
	//fmt.Println(*dx, *dy)
}

// move one step and change x, y, dx and dy
func warp(x *int, y *int, dx *int, dy *int) {
	if *x == (2*size-1) && *y >= size && *y <= (2*size-1) && *dx == 1 {
		// C right -> B bottom
		*dx = 0
		*dy = -1
		*x = 2*size + *y - size
		*y = size - 1
		return
	}

	if *y == (size-1) && *x >= (2*size) && *x <= (3*size-1) && *dy == 1 {
		// B bottom -> C left
		*dx = -1
		*dy = 0
		*y = size + *x - 2*size
		*x = 2*size - 1
		return
	}

	if *x == (3*size-1) && *y >= (0*size) && *y <= (1*size-1) && *dx == 1 {
		// B right -> E right
		*dx = -1
		*dy = 0
		*x = (2*size - 1)
		*y = (3*size - 1) - *y
		return
	}

	if *x == (size-1) && *y >= (3*size) && *y <= (4*size-1) && *dx == 1 {
		// F right -> E bottom
		*dx = 0
		*dy = -1
		*x = size + *y - 3*size
		*y = (3*size - 1)
		return
	}

	if *y == (3*size-1) && *x >= (1*size) && *x <= (2*size-1) && *dy == 1 {
		// E bottom -> F left
		*dx = -1
		*dy = 0
		*y = 3*size + *x - size
		*x = size - 1
		return
	}

	if *x == (2*size-1) && *y >= (2*size) && *y <= (3*size-1) && *dx == 1 {
		// E right -> B right
		*dx = -1
		*dy = 0
		*x = (3*size - 1)
		*y = 0 + (3*size - 1) - *y
		return
	}

	if *y == (4*size-1) && *x >= (0*size) && *x <= (1*size-1) && *dy == 1 {
		// F bottom -> B top
		*dx = 0
		*dy = 1
		*x = 2*size + *x
		*y = 0
		return
	}

	if *y == (2*size) && *x >= (0*size) && *x <= (1*size-1) && *dy == -1 {
		// D top -> C left
		*dx = 1
		*dy = 0
		*y = size + *x
		*x = size
		return
	}

	if *x == size && *y >= (1*size) && *y <= (2*size-1) && *dx == -1 {
		// C left -> D top
		*dx = 0
		*dy = 1
		*x = 0 + *y - size
		*y = 2 * size
		return
	}

	if *x == 0 && *y >= (2*size) && *y <= (3*size-1) && *dx == -1 {
		// D left -> A left
		*dx = 1
		*dy = 0
		*x = size
		*y = (3*size - 1) - *y
		return
	}

	if *y == 0 && *x >= (1*size) && *x <= (2*size-1) && *dy == -1 {
		// A top -> F left
		*dx = 1
		*dy = 0
		*y = 3*size + *x - size
		*x = 0
		return
	}

	if *y == 0 && *x >= (2*size) && *x <= (3*size-1) && *dy == -1 {
		// B top -> F bottom
		*dx = 0
		*dy = -1
		*x = *x - 2*size
		*y = (4*size - 1)
		return
	}

	if *x == 0 && *y >= (3*size) && *y <= (4*size-1) && *dx == -1 {
		// F left -> A top
		*dx = 0
		*dy = 1
		*x = size + *y - 3*size
		*y = 0
		return
	}

	if *x == size && *y >= (0*size) && *y <= (1*size-1) && *dx == -1 {
		// A left -> D left
		*dx = 1
		*dy = 0
		*x = 0
		*y = 2*size + size - 1 - *y
		return
	}

	fmt.Println("unhandled")
}

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2022/22/input.txt", "\n")

	m := make(map[string]*point)
	var x, y int
	var instructions string
	for fileY, s := range inFile {
		if s == "" {
			instructions = inFile[fileY+1]
			break
		}
		for fileX, r := range s {
			if r == ' ' {
				continue
			}
			if x == 0 {
				x = fileX
			}
			p := newPoint(fileX, fileY, string(r))
			m[p.key()] = p
		}
	}

	// follow instructions
	nbrSteps := 0
	dx := 1
	dy := 0

	//printIt(m)

	for _, r := range instructions {
		switch r {
		case 'R':
			// walk before turning
			walk(nbrSteps, &x, &y, &dx, &dy, m)
			//printIt(m)
			right(&dx, &dy)
			nbrSteps = 0
		case 'L':
			// walk before turning
			walk(nbrSteps, &x, &y, &dx, &dy, m)
			//printIt(m)
			left(&dx, &dy)
			nbrSteps = 0
		default:
			// number, append to current number
			nbrSteps *= 10
			nbrSteps += int(r - '0')
		}
	}
	// walk if we have more steps to walk
	walk(nbrSteps, &x, &y, &dx, &dy, m)
	//printIt(m)

	part2 := 1000 * (y + 1)
	part2 += 4 * (x + 1)
	part2 += direction(dx, dy)

	// not  46324
	// not  46325, it's too low
	// not 126049, it's too low
	fmt.Println("part2:", part2, "in", time.Since(start))
}

func direction(dx int, dy int) int {
	switch {
	case dx == 1 && dy == 0:
		return 0
	case dx == 0 && dy == 1:
		return 1
	case dx == -1 && dy == 0:
		return 2
	case dx == 0 && dy == -1:
		return 3
	}
	panic("weird direction")
	return -1
}

func printIt(m map[string]*point) {
	fmt.Println()
	for y := -1; y < 200; y++ {
		for x := -1; x < 150; x++ {
			p, found := m[key(x, y)]
			if !found {
				fmt.Print(" ")
			} else {
				fmt.Print(p.value)
			}
		}
		fmt.Println()
	}
}

func right(dx *int, dy *int) {
	switch {
	case *dx == 1 && *dy == 0:
		*dx, *dy = 0, 1
	case *dx == 0 && *dy == 1:
		*dx, *dy = -1, 0
	case *dx == -1 && *dy == 0:
		*dx, *dy = 0, -1
	case *dx == 0 && *dy == -1:
		*dx, *dy = 1, 0
	}

}

func left(dx *int, dy *int) {
	switch {
	case *dx == 1 && *dy == 0:
		*dx, *dy = 0, -1
	case *dx == 0 && *dy == 1:
		*dx, *dy = 1, 0
	case *dx == -1 && *dy == 0:
		*dx, *dy = 0, 1
	case *dx == 0 && *dy == -1:
		*dx, *dy = -1, 0
	}

}
