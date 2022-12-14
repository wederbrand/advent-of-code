package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type point struct {
	x, y int
}

func (p point) key() string {
	return key(p.x, p.y)
}

func newPoint(x int, y int) point {
	p := point{
		x: x,
		y: y,
	}
	return p
}

func key(x int, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

func main() {
	readFile, err := os.ReadFile("14/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	m := make(map[string]*point)
	bottom := math.MinInt
	for _, s := range inFile {
		corners := strings.Split(s, "->")
		var currentCorner *point
		for _, corner := range corners {
			p := point{}
			fmt.Sscanf(corner, "%d,%d", &p.x, &p.y)
			m[p.key()] = &p

			// draw the line
			if currentCorner != nil {
				startX, endX, startY, endY := getPath(currentCorner, &p)
				if endY > bottom {
					bottom = endY
				}
				for x := startX; x <= endX; x++ {
					for y := startY; y <= endY; y++ {
						_, found := m[key(x, y)]
						if !found {
							p2 := newPoint(x, y)
							m[p2.key()] = &p2
						}
					}
				}
			}
			currentCorner = &p
		}
	}

	// create a bottom
	maxSide := bottom + 2 + 1
	for x := -maxSide; x < maxSide; x++ {
		p := newPoint(500+x, bottom+2)
		m[p.key()] = &p
	}

	grains := 0
	for fillOne(m, bottom, 500, 0) {
		grains++
	}

	fmt.Println("part 2:", grains)
}

func fillOne(m map[string]*point, bottom int, entryX int, entryY int) bool {
	grain := newPoint(entryX, entryY)
	_, found := m[grain.key()]
	if found {
		return false
	}

	for {
		_, down := m[key(grain.x, grain.y+1)]
		if !down {
			grain.y++
			continue
		}

		_, downleft := m[key(grain.x-1, grain.y+1)]
		if !downleft {
			grain.x--
			grain.y++
			continue
		}

		_, downright := m[key(grain.x+1, grain.y+1)]
		if !downright {
			grain.x++
			grain.y++
			continue
		}

		// settle
		m[grain.key()] = &grain
		return true
	}
}

func getPath(p1 *point, p2 *point) (sx int, ex int, sy int, ey int) {
	sx, sy = p1.x, p1.y
	ex, ey = p2.x, p2.y

	if sx > ex {
		sx, ex = ex, sx
	}
	if sy > ey {
		sy, ey = ey, sy
	}

	return
}
