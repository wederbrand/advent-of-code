package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"strconv"
	"time"
)

func main() {
	startTimer := time.Now()
	inFile := util.GetFileContents("2023/18/input.txt", "\n")

	pos := Coord{0, 0}

	var direction string
	var distance int
	var color string
	area := 0
	border := 0
	for _, s := range inFile {
		fmt.Sscanf(s, "%s %d (#%s)", &direction, &distance, &color)
		dist64, _ := strconv.ParseInt(color[0:5], 16, 64)
		distance = int(dist64)
		switch color[5:6] {
		case "3":
			direction = "U"
		case "1":
			direction = "D"
		case "2":
			direction = "L"
		case "0":
			direction = "R"
		}

		border += distance
		oldPos := pos

		switch direction {
		case "U":
			pos = pos.Move(UP, distance)
		case "D":
			pos = pos.Move(DOWN, distance)
		case "L":
			pos = pos.Move(LEFT, distance)
		case "R":
			pos = pos.Move(RIGHT, distance)
		}

		area += (pos.X + oldPos.X) * (pos.Y - oldPos.Y) / 2 // shoelace
	}
	if area < 0 {
		area *= -1
	}

	// The area calculated by shoelace is wrong and has the wrong borders.
	// Luckily, all the internal points are correct, so using pick's theorem
	// we can get the internal points.
	// And we already have the border that gives the final result.
	internalArea := area - border/2 + 1
	part2 := internalArea + border
	fmt.Println("part2: ", part2, "in", time.Since(startTimer))
}
