package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	x          int
	y          int
	price      int
	totalPrice int
}

type queue struct {
	points []*point
}

func newQueue() *queue {
	q := new(queue)
	q.points = make([]*point, 0)

	return q
}

func (q *queue) add(p *point) {
	// add point to queue and sort it by total price
	q.points = append(q.points, p)
	sort.Slice(q.points, func(i, j int) bool {
		return q.points[i].totalPrice < q.points[j].totalPrice
	})
}

func (q *queue) empty() bool {
	return len(q.points) == 0
}

func (q *queue) dequeue() *point {
	p := q.points[0]
	q.points = q.points[1:len(q.points)]
	return p
}

func key(x int, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}

func main() {
	readFile, err := os.ReadFile("2021/15/2021_15.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	m := make(map[string]*point)
	maxX := 0
	maxY := 0
	for y, s := range inFile {
		split := strings.Split(s, "")
		for x, s2 := range split {
			price, _ := strconv.Atoi(s2)
			p := new(point)
			p.x = x
			p.y = y
			p.price = price
			p.totalPrice = math.MaxInt

			m[key(x, y)] = p
			if maxX < x {
				maxX = x
			}
			if maxY < y {
				maxY = y
			}
		}
	}

	newM := make(map[string]*point)
	for _, v := range m {
		for dy := 0; dy < 5; dy++ {
			for dx := 0; dx < 5; dx++ {
				dist := dy + dx
				p := new(point)
				p.price = (v.price-1+dist)%9 + 1
				p.x = v.x + dx*(maxX+1)
				p.y = v.y + dy*(maxY+1)
				p.totalPrice = math.MaxInt

				newM[key(p.x, p.y)] = p
			}
		}
	}
	m = newM

	part1Exit := m[key(maxX, maxY)]
	part2Exit := m[key((maxX+1)*5-1, (maxY+1)*5-1)]

	start := m[key(0, 0)]
	start.price = 0
	start.totalPrice = start.price

	q := newQueue()
	q.add(start)

	for !q.empty() {
		p := q.dequeue()
		if p == part1Exit {
			fmt.Println("part 1", p.totalPrice)
		}
		if p == part2Exit {
			fmt.Println("part 2", p.totalPrice)
			os.Exit(0)
		}
		// move in all directions
		dx := [4]int{0, 0, -1, 1}
		dy := [4]int{-1, 1, 0, 0}

		for i := 0; i < 4; i++ {
			x := p.x + dx[i]
			y := p.y + dy[i]
			p2, found := m[key(x, y)]
			if !found {
				// can't go there
				continue
			}

			if p2.totalPrice > p.totalPrice+p2.price {
				// if the new total price is cheaper update it and put on queue
				p2.totalPrice = p.totalPrice + p2.price
				q.add(p2)
			}
		}
	}
}
