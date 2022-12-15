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
	x, y int
	t    string
	d    int
}

func (p point) key() string {
	return key(p.x, p.y)
}

func (p point) getManhattan(b point) int {
	return int(math.Abs(float64(p.x-b.x)) + math.Abs(float64(p.y-b.y)))
}

func (p point) sees(other point) bool {
	return p.getManhattan(other) <= p.d
}

func (p point) getBorder(offset int) (border []point) {
	for dy := -p.d - offset; dy <= +p.d+offset; dy++ {
		ddx := p.d - int(math.Abs(float64(dy)))
		b1 := point{
			x: p.x + -ddx - offset,
			y: p.y + dy,
		}
		border = append(border, b1)

		b2 := point{
			x: p.x + +ddx + offset,
			y: p.y + dy,
		}
		border = append(border, b2)
	}

	return
}

func key(x int, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

func main() {
	start := time.Now()
	readFile, err := os.ReadFile("15/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	scanners := make(map[string]point, 0)
	beacons := make(map[string]point, 0)
	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
	for _, in := range inFile {
		s := point{t: "S"}
		b := point{t: "B"}
		fmt.Sscanf(in, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.x, &s.y, &b.x, &b.y)

		d := s.getManhattan(b)
		s.d = d
		b.d = d // not needed?
		scanners[s.key()] = s
		beacons[b.key()] = b

		if s.x-d < minX {
			minX = s.x
		}
		if s.x+d > maxX {
			maxX = s.x
		}
		if s.y-d < minY {
			minY = s.y
		}
		if s.y+d > maxY {
			maxY = s.y
		}
	}

	cnt := 0
	y := 2000000
	for x := minX - 1; x < maxX+1; x++ {
		for _, scanner := range scanners {
			p := point{x: x, y: y}
			_, fs := scanners[p.key()]
			_, fb := beacons[p.key()]
			if fs || fb {
				continue
			}
			if scanner.sees(p) {
				cnt++
				break
			}
		}
		// check all
	}

	fmt.Println("part1: ", cnt, "in", time.Since(start))

	start = time.Now()
scanners:
	for _, s1 := range scanners {
	boarder:
		for _, p := range s1.getBorder(1) {
			if p.x < 0 || p.x > 4000000 || p.y < 0 || p.y > 4000000 {
				continue
			}
			if p.x < minX || p.x > maxX || p.y < minY || p.y > maxY {
				continue
			}
			for _, s2 := range scanners {
				if s2.sees(p) {
					continue boarder
				}
			}
			// found it!
			freq := p.x*4000000 + p.y
			fmt.Println("part2:", freq, p, "in", time.Since(start))
			break scanners
		}
	}
}
