package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"math"

	//. "github.com/wederbrand/advent-of-code/github.com/wederbrand/priorityqueue"

	"time"
)

type Robot struct {
	c Coord
	v Coord
}

func main() {
	start := time.Now()
	inFile := GetFileContents("2024/14/input.txt", "\n")

	robots := make([]*Robot, 0)
	for _, s := range inFile {
		r := Robot{}
		fmt.Sscanf(s, "p=%d,%d v=%d,%d", &r.c.X, &r.c.Y, &r.v.X, &r.v.Y)
		robots = append(robots, &r)
	}

	const MaxWidth = 101
	const MaxHeight = 103

	doIt(robots, 100, MaxWidth, MaxHeight)
	p1 := scoreIt(robots, MaxWidth, MaxHeight)
	fmt.Println("Part 1:", p1, "in", time.Since(start))

	i := 100
	minTotal := math.MaxInt32
	minIndex := 0
	mapAsString := ""
	for i < MaxWidth*MaxHeight {
		doIt(robots, 1, MaxWidth, MaxHeight)
		i++

		total := 0
		for _, r1 := range robots {
			for _, r2 := range robots {
				total += Manhattan(r1.c, r2.c)
			}
		}
		if total < minTotal {
			minTotal = total
			minIndex = i
			mapAsString = asString(robots)
		}
	}

	c := FromString(mapAsString)
	PrintChart(c)
	fmt.Println("Part 2:", minIndex, "in", time.Since(start))
}

func doIt(robots []*Robot, moves int, maxX int, maxY int) {
	for _, r := range robots {
		r.c.X += moves * r.v.X
		r.c.Y += moves * r.v.Y

		r.c.X %= maxX
		if r.c.X < 0 {
			r.c.X += maxX
		}
		r.c.Y %= maxY
		if r.c.Y < 0 {
			r.c.Y += maxY
		}
	}
}

func scoreIt(robots []*Robot, maxX int, maxY int) int {
	border := Coord{maxX / 2, maxY / 2}

	score := [4]int{0, 0, 0, 0}
	for _, r := range robots {
		if r.c.X < border.X && r.c.Y < border.Y {
			score[0]++
		} else if r.c.X > border.X && r.c.Y < border.Y {
			score[1]++
		} else if r.c.X < border.X && r.c.Y > border.Y {
			score[2]++
		} else if r.c.X > border.X && r.c.Y > border.Y {
			score[3]++
		} else {
			// on the border
		}
	}

	return score[0] * score[1] * score[2] * score[3]
}

func asString(robots []*Robot) string {
	c := Chart{}
	for _, r := range robots {
		c[r.c] = "#"
	}

	return AsString(c)
}
