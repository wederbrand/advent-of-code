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

func walk(steps int, x *int, y *int, dx int, dy int, m map[string]*point) {
	for i := 0; i < steps; i++ {
		nextX := *x + dx
		nextY := *y + dy
		p, found := m[key(nextX, nextY)]

		if !found {
			// out of bounds, warp
			p = warp(*x, *y, dx, dy, m)
		}

		if p.value == "#" {
			// stone, walk no more
			return
		} else {
			// no stone, walk
			*x = p.x
			*y = p.y
		}
	}
}

func warp(x int, y int, dx int, dy int, m map[string]*point) (result *point) {
	var last *point
	for found := true; found; last, found = m[key(x, y)] {
		// reverse direction to warp around
		// last found is returned
		x -= dx
		y -= dy
		result = last
	}

	return result
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
	dx := [4]int{1, 0, -1, 0}
	dy := [4]int{0, 1, 0, -1}
	dxyIndex := 0 // right

	for _, r := range instructions {
		switch r {
		case 'R':
			// walk before turning
			walk(nbrSteps, &x, &y, dx[dxyIndex], dy[dxyIndex], m)
			dxyIndex += 1
			dxyIndex = ((dxyIndex % 4) + 4) % 4
			nbrSteps = 0
		case 'L':
			// walk before turning
			walk(nbrSteps, &x, &y, dx[dxyIndex], dy[dxyIndex], m)
			dxyIndex -= 1
			dxyIndex = ((dxyIndex % 4) + 4) % 4
			nbrSteps = 0
		default:
			// number, append to current number
			nbrSteps *= 10
			nbrSteps += int(r - '0')
		}
	}
	// walk if we have more steps to walk
	walk(nbrSteps, &x, &y, dx[dxyIndex], dy[dxyIndex], m)

	part1 := 1000 * (y + 1)
	part1 += 4 * (x + 1)
	part1 += dxyIndex % 4

	fmt.Println("part1:", part1, "in", time.Since(start))
}
