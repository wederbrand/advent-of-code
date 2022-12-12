package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type point struct {
	x, y   int
	height int
	part1  bool
	part2  bool
}

func (p point) key() string {
	return key(p.x, p.y)
}

func newPoint(r rune, x int, y int) point {
	p := point{
		x: x,
		y: y,
	}
	if r == 'S' {
		p.part1 = true
		r = 'a'
	}
	if r == 'a' {
		p.part2 = true
	}
	if r == 'E' {
		r = 'z'
	}
	p.height = int(r - 'a')
	return p
}

func key(x int, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

type state struct {
	// I can get here at this price
	p     point
	price int
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
			if s.price < s2.price {
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

func main() {
	start := time.Now()
	readFile, err := os.ReadFile("12/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	world := make(map[string]point)
	var summit point

	for y, s := range inFile {
		for x, r := range s {
			p := newPoint(r, x, y)
			if r == 'E' {
				summit = p
			}
			world[p.key()] = p
		}
	}

	part1, part2 := getDistance(world, summit)
	fmt.Println("part 1", part1)
	fmt.Println("part 2", part2)
	fmt.Println("time", time.Since(start))
}

func getDistance(world map[string]point, summit point) (part1 int, part2 int) {
	q := newQueue()
	visited := make(map[string]int)
	currentState := state{
		p:     summit,
		price: 0,
	}
	q.add(&currentState)
	visited[summit.key()] = 0

	for len(q.states) > 0 {
		s := q.dequeue()
		if s.p.part2 && part2 == 0 {
			part2 = s.price
		}
		if s.p.part1 {
			part1 = s.price
			return
		}

		// move in all directions
		dx := [4]int{0, 0, -1, 1}
		dy := [4]int{-1, 1, 0, 0}

		for i := 0; i < 4; i++ {
			x := s.p.x + dx[i]
			y := s.p.y + dy[i]
			p2, found := world[key(x, y)]
			if !found {
				// can't go there
				continue
			}

			if s.p.height-p2.height > 1 {
				// can't go there
				continue
			}

			oldPrice, found := visited[p2.key()]
			if !found || oldPrice > (s.price+1) {
				newState := state{
					p:     p2,
					price: s.price + 1,
				}
				q.add(&newState)
				visited[newState.p.key()] = newState.price
			}
		}
	}
	return
}
