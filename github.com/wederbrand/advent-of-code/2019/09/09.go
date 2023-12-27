package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/2019/computer"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2019/09/input.txt", "\n")

	computer := NewComputer(inFile)
	go computer.Run()
	computer.GetInput() <- 1
	part1 := 0
	for out := range computer.GetOutput() {
		part1 = out
	}
	fmt.Println("part1: ", part1, "in", time.Since(start))

	computer = NewComputer(inFile)
	go computer.Run()
	computer.GetInput() <- 2
	part2 := 0
	for out := range computer.GetOutput() {
		part2 = out
	}
	fmt.Println("part2: ", part2, "in", time.Since(start))
}
