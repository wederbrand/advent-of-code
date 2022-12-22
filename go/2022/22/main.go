package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

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
		p, found := m[key(nextX, nextY)]

		if !found {
			// out of bounds, warp
			warp(x, y, dx, dy)
			p = m[key(*x, *y)]
		}

		if p.value == "#" {
			// stone, walk no more
			return
		} else {
			p.value = "X"
			// no stone, walk
			*x = p.x
			*y = p.y
		}
	}
	//fmt.Println(*dx, *dy)
}

// move one step and change x, y, dx and dy
func warp(x *int, y *int, dx *int, dy *int) {
	if *x == 99 && *y >= 50 && *y <= 99 && *dx == 1 {
		// C right -> B bottom
		*dx = 0
		*dy = -1
		*x = 100 + *y - 50
		*y = 49
		return
	}

	if *y == 49 && *x >= 100 && *x <= 150 && *dy == 1 {
		// B bottom -> C left
		*dx = -1
		*dy = 0
		*y = 50 + *x - 100
		*x = 99
		return
	}

	if *x == 149 && *y >= 0 && *y <= 49 && *dx == 1 {
		// B right -> F left
		*dx = -1
		*dy = 0
		*x = 49
		*y = 150 + 49 - *y
		return
	}

	if *x == 49 && *y >= 150 && *y <= 199 && *dx == 1 {
		// F right -> E bottom
		*dx = 0
		*dy = -1
		*x = 50 + *y - 150
		*y = 149
		return
	}

	if *y == 149 && *x >= 50 && *x <= 99 && *dy == 1 {
		// E bottom -> F left
		*dx = -1
		*dy = 0
		*y = 150 + *x - 50
		*x = 49
		return
	}

	if *x == 99 && *y >= 100 && *y <= 149 && *dx == 1 {
		// E right -> B left
		*dx = -1
		*dy = 0
		*x = 149
		*y = 0 + 150 - *y
		return
	}

	if *y == 199 && *x >= 0 && *x <= 49 && *dy == 1 {
		// F bottom -> B top
		*dx = 0
		*dy = 1
		*x = 100 + *x
		*y = 0 // <- error
		return
	}

	if *y == 100 && *x >= 0 && *x <= 49 && *dy == -1 {
		// D top -> C left
		*dx = 1
		*dy = 0
		*y = 50 + *x
		*x = 50
		return
	}

	if *x == 50 && *y >= 50 && *y <= 99 && *dx == -1 {
		// C left -> D top
		*dx = 0
		*dy = 1
		*x = 0 + *y - 50
		*y = 100
		return
	}

	if *x == 0 && *y >= 100 && *y <= 149 && *dx == -1 {
		// D left -> A left
		*dx = 1
		*dy = 0
		*x = 50
		*y = 149 - *y
		return
	}

	if *y == 0 && *x >= 50 && *x <= 99 && *dy == -1 {
		// A top -> F left
		*dx = 1
		*dy = 0
		*y = 150 + *x - 50
		*x = 0
		return
	}

	if *y == 0 && *x >= 100 && *x <= 149 && *dy == -1 {
		// B top -> F bottom
		*dx = 0
		*dy = -1
		*x = *x - 100
		*y = 199
		return
	}

	if *x == 0 && *y >= 150 && *y <= 199 && *dx == -1 {
		// F left -> A top
		*dx = 0
		*dy = 1
		*x = 50 + *y - 150 // <- error
		*y = 0             // <- error
		return
	}

	if *x == 50 && *y >= 0 && *y <= 49 && *dx == -1 {
		// A left -> D left
		*dx = 1
		*dy = 0
		*x = 0
		*y = 100 + 49 - *y
		return
	}

	fmt.Println("unhandled")
}

func main() {
	start := time.Now()
	readFile, err := os.ReadFile("22/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	inFile := strings.Split(string(readFile), "\n")

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

	printIt(m)

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
	printIt(m)

	part2 := 1000 * (y + 1)
	part2 += 4 * (x + 1)
	switch {
	case dx == 1 && dy == 0:
		part2 += 0
	case dx == 0 && dy == 1:
		part2 += 1
	case dx == -1 && dy == 0:
		part2 += 2
	case dx == 0 && dy == -1:
		part2 += 3
	}

	// not  46324
	// not  46325, it's too low
	// not 126049, it's too low
	fmt.Println("part2:", part2, "in", time.Since(start))
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
