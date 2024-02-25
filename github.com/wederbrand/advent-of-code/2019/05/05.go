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

	inFunc := func() int {
		return 1
	}

	out := 0
	outFunc := func(o int) {
		out = o
	}
	computer := NewComputer(inFile, inFunc, outFunc)
	computer.Run()
	fmt.Println("part1: ", out, "inFunc", time.Since(start))

	inFunc = func() int {
		return 5
	}

	computer = NewComputer(inFile, inFunc, outFunc)
	computer.Run()
	fmt.Println("part2: ", out, "inFunc", time.Since(start))
}
