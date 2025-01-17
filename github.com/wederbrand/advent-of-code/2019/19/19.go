package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/2019/computer"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"time"
)

const SIZE = 100

func main() {
	start := time.Now()
	inFile := GetFileContents("2019/19/input.txt", "\n")

	m := getMap(inFile)

	part1 := 0
	for y := 0; y <= 50; y++ {
		for x := 0; x <= 50; x++ {
			c := Coord{X: x, Y: y}
			if m[c] == "#" {
				part1++
			}
		}
	}

	// PrintChart(m)
	fmt.Println("part1: ", part1, "in", time.Since(start))

	testBottomRowMax := SIZE
	for {
		fitted := testIt(inFile, testBottomRowMax)
		if fitted == nil {
			testBottomRowMax *= 2
		} else {
			break
		}
	}

	testBottomRowMin := testBottomRowMax / 2
	testBottomRowMid := (testBottomRowMax + testBottomRowMin) / 2
	for {
		fitted := testIt(inFile, testBottomRowMid)
		if fitted != nil {
			testBottomRowMax = testBottomRowMid
		} else {
			testBottomRowMin = testBottomRowMid
		}
		testBottomRowMid = (testBottomRowMax + testBottomRowMin) / 2
		if testBottomRowMax == testBottomRowMin+1 {
			break
		}
	}

	// max is the first line that works
	corner := testIt(inFile, testBottomRowMax)
	fmt.Println("part2: ", corner.X*10000+corner.Y, "in", time.Since(start))
}

func testIt(inFile []string, row int) *Coord {
	x := 0
	for {
		x++
		if getCoord(inFile, Coord{X: x, Y: row}) == 1 {
			break
		}
	}

	lowerLeft := Coord{X: x, Y: row}
	upperLeft := Coord{X: x, Y: row - (SIZE - 1)}
	lowerRight := Coord{X: x + (SIZE - 1), Y: row}
	upperRight := Coord{X: x + (SIZE - 1), Y: row - (SIZE - 1)}

	if getCoord(inFile, lowerLeft) == 1 && getCoord(inFile, upperLeft) == 1 && getCoord(inFile, lowerRight) == 1 && getCoord(inFile, upperRight) == 1 {
		return &upperLeft
	} else {
		return nil
	}
}

func getMap(inFile []string) Chart {
	m := Chart{}
	// for part1 we need 50x50
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			c := Coord{X: x, Y: y}
			if getCoord(inFile, c) == 1 {
				m[c] = "#"
			} else {
				m[c] = "."
			}
		}
	}
	return m
}

func getCoord(inFile []string, c Coord) int {
	input := []int{c.X, c.Y}
	in := func() int {
		i := input[0]
		input = input[1:]
		return i
	}

	output := 0
	out := func(i int) {
		output = i
	}

	computer := NewComputer(inFile, in, out)
	computer.Run()
	return output
}
