package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/2019/computer"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2019/05/input.txt", "\n")
	computer := NewComputer(inFile)
	go computer.Run()
	computer.GetInput() <- 1
	out := 0
	for i := range computer.GetOutput() {
		out = i
	}
	fmt.Println("part1: ", out, "in", time.Since(start))

	computer = NewComputer(inFile)
	go computer.Run()
	computer.GetInput() <- 5
	out = <-computer.GetOutput()
	fmt.Println("part2: ", out, "in", time.Since(start))
}
