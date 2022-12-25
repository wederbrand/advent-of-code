package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strconv"
)

type point struct {
	x      int
	y      int
	height int
}

func (p point) visible(m map[string]point, maxX int, maxY int) bool {

	up := p.lowerBetween(m, p.height, 0, p.y-1, p.x, p.x)
	down := p.lowerBetween(m, p.height, p.y+1, maxY, p.x, p.x)
	left := p.lowerBetween(m, p.height, p.y, p.y, 0, p.x-1)
	right := p.lowerBetween(m, p.height, p.y, p.y, p.x+1, maxX)

	return up || down || left || right
}

func (p point) lowerBetween(m map[string]point, h int, minY int, maxY int, minX int, maxX int) bool {
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			n, found := m[key(x, y)]
			if !found {
				continue
			}
			if h > n.height {
				// still visible
			} else {
				return false
			}
		}
	}
	return true
}

func (p point) score(m map[string]point) int {
	// up
	score := 0
	y := p.y - 1
	for {
		n, found := m[key(p.x, y)]
		if !found {
			break
		}
		score++
		if n.height >= p.height {
			break
		}
		y--
	}
	total := score

	// down
	score = 0
	y = p.y + 1
	for {
		n, found := m[key(p.x, y)]
		if !found {
			break
		}
		score++
		if n.height >= p.height {
			break
		}
		y++
	}
	total *= score

	// left
	score = 0
	x := p.x - 1
	for {
		n, found := m[key(x, p.y)]
		if !found {
			break
		}
		score++
		if n.height >= p.height {
			break
		}
		x--
	}
	total *= score

	// right
	score = 0
	x = p.x + 1
	for {
		n, found := m[key(x, p.y)]
		if !found {
			break
		}
		score++
		if n.height >= p.height {
			break
		}
		x++
	}
	total *= score

	return total
}

func key(x int, y int) string {
	return strconv.Itoa(x) + ":" + strconv.Itoa(y)
}

func main() {
	inFile := util.GetFileContents("2022/08/input.txt", "\n")

	m := make(map[string]point, 0)

	maxY := 0
	maxX := 0
	for y, s := range inFile {
		for x, tree := range s {
			if y > maxY {
				maxY = y
			}
			if x > maxX {
				maxX = x
			}
			h, _ := strconv.Atoi(string(tree))
			p := point{
				x:      x,
				y:      y,
				height: h,
			}
			m[key(x, y)] = p
		}
	}

	visible := 0
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			b := m[key(x, y)].visible(m, maxX, maxY)
			if b {
				visible++
			}
		}
	}
	fmt.Println("part 1: ", visible)

	mostScore := 0
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			p := m[key(x, y)]
			score := p.score(m)
			if score > mostScore {
				mostScore = score
			}
		}
	}

	fmt.Println("part 2: ", mostScore)

}
