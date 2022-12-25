package main

import (
	"errors"
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"math"
	"time"
)

type elf struct {
	x   int
	y   int
	pX  int
	pY  int
	pOk bool
}

func newElf(x int, y int) *elf {
	p := new(elf)
	p.x = x
	p.y = y

	return p
}

func (e *elf) key() string {
	return key(e.x, e.y)
}

func (e *elf) propose(dir int) error {
	free := true
	for i := 0; i < 8; i++ {
		x2 := e.x + dx[i]
		y2 := e.y + dy[i]
		_, found := elves[key(x2, y2)]
		if found {
			free = false
			break
		}
	}
	if free {
		return errors.New("all free")
	}
	for i := 0; i < 4; i++ {
		// check for directions
		found := false
		for j := -1; j <= 1; j++ {
			// check 3 sub directions
			x2 := e.x + dx[dir+j]
			y2 := e.y + dy[dir+j]
			_, found = elves[key(x2, y2)]
			if found {
				break
			}
		}
		if !found {
			e.pX = e.x + dx[dir]
			e.pY = e.y + dy[dir]
			return nil
		}

		dir = nextDir[dir]
	}
	return errors.New("no proposal")
}

func (e *elf) move() {
	e.x = e.pX
	e.y = e.pY
}

func key(x int, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

// in clockwise order
// starting AND ending with NW
// N 1
// E 3
// S 5
// W 7
var dx = [9]int{-1, 0, 1, 1, 1, 0, -1, -1, -1}
var dy = [9]int{-1, -1, -1, 0, 1, 1, 1, 0, -1}
var nextDir = [9]int{0, 5, 0, 1, 0, 7, 0, 3, 0}

var elves = make(map[string]*elf)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2022/23/input.txt", "\n")

	for fileY, s := range inFile {
		for fileX, r := range s {
			if r == '#' {
				p := newElf(fileX, fileY)
				elves[p.key()] = p
			}
		}
	}

	direction := 1 // North
	i := 0
	for {
		// all propose
		// proposal is returned and kept in a map with proposal as key and a list of elves as value
		proposals := make(map[string][]*elf)
		for _, e := range elves {
			err := e.propose(direction)
			if err == nil {
				proposals[key(e.pX, e.pY)] = append(proposals[key(e.pX, e.pY)], e)
			}
		}
		if len(proposals) == 0 {
			// if list of proposals is empty we're done
			break
		}

		// iterate proposals.
		for _, elfList := range proposals {
			if len(elfList) == 1 {
				// move
				elfList[0].pOk = true
			} else {
				for _, e := range elfList {
					e.pOk = false
				}
				elfList[0].pOk = false
			}
		}

		next := make(map[string]*elf)
		for _, e := range elves {
			if e.pOk {
				// move
				e.move()
			}

			next[e.key()] = e
		}
		elves = next
		// update directionIndex
		direction = nextDir[direction]
		if i == 9 {
			part1 := countIt()
			fmt.Println("part1", part1, "in", time.Since(start))
		}
		i++
	}

	fmt.Println("part2", i+1, "in", time.Since(start))
}

func countIt() (sum int) {
	minX := math.MaxInt
	minY := math.MaxInt
	maxX := math.MinInt
	maxY := math.MinInt

	for _, e := range elves {
		minX = min(minX, e.x)
		minY = min(minY, e.y)
		maxX = max(maxX, e.x)
		maxY = max(maxY, e.y)
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			_, found := elves[key(x, y)]
			if !found {
				sum += 1
			}
		}
	}
	return
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
