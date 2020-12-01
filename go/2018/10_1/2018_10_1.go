package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	x int
	y int
	dx int
	dy int
}

func (p *Point) incTime(t int) Point {
	p.x += t*p.dx
	p.y += t*p.dy
	return *p
}

func main() {
	file, err := os.Open("10_1/2018_10.input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	lineMatcher := regexp.MustCompile(`^position=<\s*(-*)(\d+),\s*(-*)(\d+)> velocity=<\s*(-*)(\d+),\s*(-*)(\d+)>$`)

	points := make([]Point, 0)

	for scanner.Scan() {
		lineMatch := lineMatcher.FindStringSubmatch(scanner.Text())
		x, _:=strconv.Atoi(lineMatch[2])
		if lineMatch[1] == "-" {
			x *= -1
		}
		y, _:=strconv.Atoi(lineMatch[4])
		if lineMatch[3] == "-" {
			y *= -1
		}
		dx, _:=strconv.Atoi(lineMatch[6])
		if lineMatch[5] == "-" {
			dx *= -1
		}
		dy, _:=strconv.Atoi(lineMatch[8])
		if lineMatch[7] == "-" {
			dy *= -1
		}
		points = append(points, Point{y, x, dy, dx})
	}

	time := 0
	dist := math.MaxInt64

	for {
		time++
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
		
		localDist := maxX-minX
		if localDist < 0 {
			localDist *= -1
		}

		if localDist < dist {
			dist = localDist
		} else {
			time--
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
	fmt.Println(time)

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
