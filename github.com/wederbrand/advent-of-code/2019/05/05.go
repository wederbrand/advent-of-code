package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2019/04/input.txt", "\n")

	for _, s := range inFile {
	}

	part1 := 0
	fmt.Println("part1: ", part1, "in", time.Since(start))
}
