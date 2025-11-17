package main

import (
	"fmt"
	"time"

	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
)

func main() {
	start := time.Now()
	input := GetFileContents("2025/01/input.txt", "\n")

	dial := 50
	part1 := 0
	part2 := 0
	for _, s := range input {
		dir := 1
		if s[0] == 'L' {
			dir = -1
		}
		steps := Atoi(s[1:])

		part2 += steps / 100
		steps = steps % 100

		future := dial + dir*steps

		if future < 0 {
			future += 100
			if dial != 0 {
				part2++
			}
		} else if future > 100 {
			future -= 100
			part2++
		}

		if future == 0 || future == 100 {
			part1++
			part2++
			future = 0
		}

		dial = future
	}

	fmt.Println("Part 1:", part1, "in", time.Since(start)) // 1052
	fmt.Println("Part 2:", part2, "in", time.Since(start)) // 6295
}
