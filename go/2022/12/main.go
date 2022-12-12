package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

type point struct {
	x, y   int
	height int
	end    bool
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
		r = 'a'
	}
	if r == 'E' {
		p.end = true
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
	var part1 point
	starts := make([]point, 0)

	for y, s := range inFile {
		for x, r := range s {
			p := newPoint(r, x, y)
			if r == 'S' {
				starts = append(starts, p)
				part1 = p
			}
			if r == 'a' {
				starts = append(starts, p)
			}
			world[p.key()] = p
		}
	}

	fmt.Println("part 1:", getDistance(world, part1), "in", time.Since(start))

	min := math.MaxInt
	for _, start := range starts {
		distance := getDistance(world, start)
		if distance < min {
			min = distance
		}
	}

	fmt.Println("part 2:", min, "in", time.Since(start))
}

func getDistance(world map[string]point, start point) int {
	q := newQueue()
	visited := make(map[string]int)
	currentState := state{
		p:     start,
		price: 0,
	}
	q.add(&currentState)
	visited[start.key()] = 0

	for len(q.states) > 0 {
		s := q.dequeue()
		// fmt.Println("testing", s, len(q.states))
		if s.p.end {
			return s.price
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
			if p2.key() != key(x, y) {
				panic("weird")
			}

			if p2.height-s.p.height > 1 {
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
				//fmt.Println("queued", newState)
			}
		}
	}

	return math.MaxInt
}
