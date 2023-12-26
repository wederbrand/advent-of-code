package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2019/03/input.txt", "\n")

	m := Chart{}

	wire1 := strings.Split(inFile[0], ",")
	wire2 := strings.Split(inFile[1], ",")

	port := Coord{0, 0}
	pos := port
	m[pos] = "o"
	totalSteps := 0
	for _, s := range wire1 {
		var d Dir
		switch s[0] {
		case 'U':
			d = UP
		case 'L':
			d = LEFT
		case 'R':
			d = RIGHT
		case 'D':
			d = DOWN
		}
		steps := util.Atoi(s[1:])
		for i := 0; i < steps; i++ {
			totalSteps++
			pos = pos.Move(d)
			_, found := m[pos]
			if !found {
				m[pos] = strconv.Itoa(totalSteps)
			}
		}
	}

	pos = port
	cross := make([]Coord, 0)
	part2 := math.MaxInt
	totalSteps = 0
	for _, s := range wire2 {
		var d Dir
		switch s[0] {
		case 'U':
			d = UP
		case 'L':
			d = LEFT
		case 'R':
			d = RIGHT
		case 'D':
			d = DOWN
		}
		steps := util.Atoi(s[1:])
		for i := 0; i < steps; i++ {
			totalSteps++
			pos = pos.Move(d)
			otherSteps, found := m[pos]
			if found {
				part2 = min(part2, util.Atoi(otherSteps)+totalSteps)
				cross = append(cross, pos)
			}
		}
	}

	part1 := math.MaxInt
	for _, coord := range cross {
		part1 = min(part1, Manhattan(port, coord))
	}

	fmt.Println("part1: ", part1, "in", time.Since(start))
	fmt.Println("part2: ", part2, "in", time.Since(start))
}
