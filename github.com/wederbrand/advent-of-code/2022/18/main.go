package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"math"
	"time"
)

type drop struct {
	x, y, z int
}

func (d drop) key() string {
	return key(d.x, d.y, d.z)
}

func key(x, y, z int) string {
	return fmt.Sprintf("%d:%d:%d", x, y, z)
}

type queue struct {
	states []*drop
}

func newQueue() *queue {
	q := new(queue)
	q.states = make([]*drop, 0)

	return q
}

func (q *queue) add(s *drop) {
	// add drop to queue, no need to sort
	q.states = append(q.states, s)
}

func (q *queue) dequeue() *drop {
	p := q.states[0]
	q.states = q.states[1:len(q.states)]
	return p
}

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2022/18/input.txt", "\n")

	// move in all directions
	dx := []int{-1, 1, 0, 0, 0, 0}
	dy := []int{0, 0, -1, 1, 0, 0}
	dz := []int{0, 0, 0, 0, -1, 1}

	part1 := 0
	world := make(map[string]*drop)
	minX, minY, minZ := math.MaxInt, math.MaxInt, math.MaxInt
	maxX, maxY, maxZ := math.MinInt, math.MinInt, math.MinInt
	for _, s := range inFile {
		d := drop{}
		fmt.Sscanf(s, "%d,%d,%d", &d.x, &d.y, &d.z)

		minX = min(minX, d.x)
		minY = min(minY, d.y)
		minZ = min(minZ, d.z)
		maxX = max(maxX, d.x)
		maxY = max(maxY, d.y)
		maxZ = max(maxZ, d.z)

		world[d.key()] = &d
		part1 += 6
		for i := 0; i < 6; i++ {
			k := key(d.x+dx[i], d.y+dy[i], d.z+dz[i])
			_, found := world[k]
			if found {
				part1 -= 2
			}
		}
	}

	fmt.Println("part1:", part1, "in", time.Since(start))

	outside := make(map[string]*drop)

	q := newQueue()
	// start at minX, minY, minZ
	q.add(&drop{
		x: minX,
		y: minY,
		z: minZ,
	})

	part2 := 0
	for len(q.states) > 0 {
		d := q.dequeue()
		// check all directions
		for i := 0; i < 6; i++ {
			d2 := &drop{
				x: d.x + dx[i],
				y: d.y + dy[i],
				z: d.z + dz[i],
			}
			if d2.x < minX-1 || d2.x > maxX+1 || d2.y < minY-1 || d2.y > maxY+1 || d2.z < minZ-1 || d2.z > maxZ+1 {
				continue
			}

			_, found := world[d2.key()]
			if found {
				// if "that other shape" is hit don't add to the queue
				// and increment the counter
				part2 += 1
				continue
			}

			_, found = outside[d2.key()]
			if !found {
				// if not already in the outside add it, and queue it
				outside[d2.key()] = d2
				q.add(d2)
			}
		}
	}

	// 2070 is too high
	fmt.Println("part2:", part2, "in", time.Since(start))
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
