package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/2019/computer"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2019/02/input.txt", "\n")
	computer := NewComputer(inFile, nil, nil)

	part1 := doIt(computer, 12, 2)
	fmt.Println("part1: ", part1, "in", time.Since(start))

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			computer = NewComputer(inFile, nil, nil)
			result := doIt(computer, i, j)
			if result == 19690720 {
				part2 := 100*i + j
				fmt.Println("part2: ", part2, "in", time.Since(start))
				return
			}
		}
	}
}

func doIt(computer Computer, a int, b int) int {
	computer.SetMemory(1, a)
	computer.SetMemory(2, b)
	computer.Run()
	return computer.GetMemory(0)
}
