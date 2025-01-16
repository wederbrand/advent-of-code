package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/2019/computer"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2019/17/input.txt", "\n")

	inputString := make([]string, 0)
	current := ""
	out := func(i int) {
		if i == 10 {
			inputString = append(inputString, current)
			current = ""
		} else {
			current += string(rune(i))
		}
	}

	computer := NewComputer(inFile, nil, out)
	computer.Run()

	m := MakeChart(inputString, "")

	part1 := 0
	minC, maxC := GetChartMaxes(m)
	for y := minC.Y; y <= maxC.Y; y++ {
		for x := minC.X; x <= maxC.X; x++ {
			c := Coord{x, y}
			if m[c] != "#" {
				continue
			}
			count := 0
			for _, dir := range ALL {
				if m[c.Move(dir)] == "#" {
					count++
				}
			}
			if count == 4 {
				part1 += x * y
			}
		}
	}

	fmt.Println("part1: ", part1, "in", time.Since(start))

	// hand drawn path finding for part2
	A := "L,12,R,4,R,4,L,6"
	B := "L,12,R,4,R,4,R,12"
	C := "L,10,L,6,R,4"
	path := "A,B,A,C,A,B,C,B,C,A"

	part2 := 0
	out = func(i int) {
		if i > 256 {
			part2 = i
		}
	}

	input := []int{}
	for _, r := range path {
		input = append(input, int(r))
	}
	input = append(input, 10)
	for _, r := range A {
		input = append(input, int(r))
	}
	input = append(input, 10)
	for _, r := range B {
		input = append(input, int(r))
	}
	input = append(input, 10)
	for _, r := range C {
		input = append(input, int(r))
	}
	input = append(input, 10)
	input = append(input, int('n'))
	input = append(input, 10)

	in := func() int {
		if len(input) == 0 {
			return 0
		}
		i := input[0]
		input = input[1:]
		return i
	}

	inFile[0] = "2" + inFile[0][1:]
	computer = NewComputer(inFile, in, out)
	computer.Run()

	fmt.Println("part2: ", part2, "in", time.Since(start))
}
