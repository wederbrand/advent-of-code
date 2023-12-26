package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2019/01/input.txt", "\n")

	part1 := 0
	part2 := 0
	for _, s := range inFile {
		mass := util.Atoi(s)
		part1 += (mass / 3) - 2
		part2 += getFuel(mass)

	}

	fmt.Println("part1: ", part1, "in", time.Since(start))
	fmt.Println("part2: ", part2, "in", time.Since(start))
}

func getFuel(mass int) int {
	fuel := (mass / 3) - 2
	if fuel <= 0 {
		return 0
	} else {
		return fuel + getFuel(fuel)
	}
}
