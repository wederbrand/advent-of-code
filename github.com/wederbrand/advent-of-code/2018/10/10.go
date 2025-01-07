package main

import (
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"time"

	"fmt"
	"math"
	"regexp"
	"strconv"
)

type Point struct {
	x  int
	y  int
	dx int
	dy int
}

func (p *Point) incTime(t int) Point {
	p.x += t * p.dx
	p.y += t * p.dy
	return *p
}

func main() {
	startTime := time.Now()
	inFile := GetFileContents("2018/10/input.txt", "\n")

	lineMatcher := regexp.MustCompile(`^position=<\s*(-*)(\d+),\s*(-*)(\d+)> velocity=<\s*(-*)(\d+),\s*(-*)(\d+)>$`)

	points := make([]Point, 0)

	for _, s := range inFile {
		lineMatch := lineMatcher.FindStringSubmatch(s)
		x, _ := strconv.Atoi(lineMatch[2])
		if lineMatch[1] == "-" {
			x *= -1
		}
		y, _ := strconv.Atoi(lineMatch[4])
		if lineMatch[3] == "-" {
			y *= -1
		}
		dx, _ := strconv.Atoi(lineMatch[6])
		if lineMatch[5] == "-" {
			dx *= -1
		}
		dy, _ := strconv.Atoi(lineMatch[8])
		if lineMatch[7] == "-" {
			dy *= -1
		}
		points = append(points, Point{y, x, dy, dx})
	}

	clock := 0
	dist := math.MaxInt64

	for {
		clock++
		minX := math.MaxInt64
		maxX := math.MinInt64
		minY := math.MaxInt64
		maxY := math.MinInt64

		for i, p := range points {
			points[i] = p.incTime(1)
			if p.x < minX {
				minX = p.x
			}
			if p.x > maxX {
				maxX = p.x
			}
			if p.y < minY {
				minY = p.y
			}
			if p.y > maxY {
				maxY = p.y
			}
		}

		localDist := maxX - minX
		if localDist < 0 {
			localDist *= -1
		}

		if localDist < dist {
			dist = localDist
		} else {
			clock--
			for i, p := range points {
				points[i] = p.incTime(-1)
				if p.x < minX {
					minX = p.x
				}
				if p.x > maxX {
					maxX = p.x
				}
				if p.y < minY {
					minY = p.y
				}
				if p.y > maxY {
					maxY = p.y
				}
			}
			break
		}
	}

	drawPoints(points)
	fmt.Println("Part 2:", clock, "in", time.Since(startTime))
}

func drawPoints(points []Point) {
	const size = 1000
	var draw [size][size]string
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			draw[i][j] = " "
		}
	}
	for _, point := range points {
		draw[point.x+size/2][point.y+size/2] = "X"
	}
	for _, row := range draw {
		fmt.Println(row)
	}
}
