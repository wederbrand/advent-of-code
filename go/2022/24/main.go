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
	dx    int
	dy    int
	xwarp int
	ywarp int
}

func newPoint(x int, y int, dx int, dy int, minX int, maxX int, minY int, maxY int) *point {
	p := new(point)
	p.x = x
	p.y = y
	p.dx = dx
	p.dy = dy
	p.xwarp = maxX - minX + 1
	p.ywarp = maxY - minY + 1

	return p
}

func (p *point) key() string {
	return key(p.x, p.y)
}

func (p *point) getPoint(time int) *point {
	p2 := point{}
	p2.x = (p.x + p.dx*time) % p.xwarp
	p2.y = (p.y + p.dy*time) % p.ywarp

	for p2.x < 0 {
		p2.x += p.xwarp
	}
	for p2.x > maxX {
		p2.x -= p.xwarp
	}
	for p2.y < 0 {
		p2.y += p.ywarp
	}
	for p2.y > maxY {
		p2.y -= p.ywarp
	}

	return &p2
}

func key(x int, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

type state struct {
	// I can get here at this price
	p     point
	steps int
	score int
}

type queue struct {
	states []*state
}

func newQueue() *queue {
	q := new(queue)
	q.states = make([]*state, 0)

	return q
}

func (q *queue) add(s *state) {
	// add state to queue and sort it by total price
	if len(q.states) == 0 {
		q.states = append(q.states, s)
	} else {
		for i, s2 := range q.states {
			if s.score < s2.score {
				// insert here
				q.states = append(q.states, nil)   // make room for copy
				copy(q.states[i+1:], q.states[i:]) // make room at index i, overwriting the nil
				q.states[i] = s                    // insert s
				return
			}
		}
		q.states = append(q.states, s)
	}
}

func (q *queue) dequeue() *state {
	p := q.states[0]
	q.states = q.states[1:len(q.states)]
	return p
}

// move in all directions, including wait
var dx = [5]int{0, 0, 0, -1, 1}
var dy = [5]int{0, -1, 1, 0, 0}
var minX int
var minY int
var maxX int
var maxY int

var debug = false

func main() {
	start := time.Now()
	readFile, err := os.ReadFile("24/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")
	minX = 0
	minY = 0
	maxX = len(inFile[0]) - 3 // two sides and the zero-index
	maxY = len(inFile) - 3    // two sides and the zero-index
	blizzards := make(map[string]*point)

	for y, s := range inFile {
		for x, r := range s {
			switch r {
			case '>':
				b := newPoint(x-1, y-1, 1, 0, minX, maxX, minY, maxY)
				blizzards[b.key()] = b
			case '<':
				b := newPoint(x-1, y-1, -1, 0, minX, maxX, minY, maxY)
				blizzards[b.key()] = b
			case 'v':
				b := newPoint(x-1, y-1, 0, 1, minX, maxX, minY, maxY)
				blizzards[b.key()] = b
			case '^':
				b := newPoint(x-1, y-1, 0, -1, minX, maxX, minY, maxY)
				blizzards[b.key()] = b
			}
		}
	}

	m := make(map[string]int)
	for _, p := range blizzards {
		p2 := p.getPoint(0)
		m[p2.key()]++
	}

	printIt(m, 0)

	m = make(map[string]int)
	for _, p := range blizzards {
		p2 := p.getPoint(1)
		m[p2.key()]++
	}

	printIt(m, 1)

	q := newQueue()

	// initial state (pretend the first step is made
	initialState := state{
		p:     point{x: 0, y: 0},
		steps: 1, // to pretend this is after the first step
		score: maxX + maxY,
	}
	q.add(&initialState)

	// exit point (pretending the next step can always happen)
	exit := point{
		x: maxX,
		y: maxY,
	}

	seen := make(map[state]bool)
	seen[initialState] = true
	for len(q.states) > 0 {
		s := q.dequeue()
		fmt.Println("q:", len(q.states), s.steps, s.score, s.p.x, s.p.y)
		// move time
		t := s.steps + 1
		if s.p == exit {
			fmt.Println("part1", t, "in", time.Since(start))
			break
		}
		// create the map according to current time
		m := make(map[string]int)
		for _, p := range blizzards {
			p2 := p.getPoint(t)
			m[p2.key()]++
		}

		printIt(m, t)

		// try all. up, down, left, right and wait, if possible
		for i := 0; i < 5; i++ {
			p := s.p
			p.x += dx[i]
			p.y += dy[i]

			if p.x < minX || p.x > maxX || p.y < minY || p.y > maxY {
				continue
			}

			_, found := m[p.key()]
			if !found {
				s2 := state{
					p:     p,
					steps: s.steps + 1,
					score: s.steps + 1 + maxX - p.x + maxY - p.y,
				}
				if seen[s2] == false {
					seen[s2] = true
					q.add(&s2)

				}
			}

		}
	}
}

func printIt(m map[string]int, t int) {
	if !debug {
		return
	}
	fmt.Println("time", t)
	fmt.Println("########")
	for y := 0; y <= maxY; y++ {
		fmt.Print("#")
		for x := 0; x <= maxX; x++ {
			v, found := m[key(x, y)]
			if !found {
				fmt.Print(".")
			} else {
				fmt.Print(v)
			}
		}
		fmt.Println("#")
	}
	fmt.Println("########")
}
