package main

import (
	"fmt"
	"time"

	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
)

func main() {
	start := time.Now()
	input := GetFileContents("2025/07/input.txt", "\n")

	part1, part2 := doIt(input)
	fmt.Println("Part 1:", part1, "in", time.Since(start))
	fmt.Println("Part 2:", part2, "in", time.Since(start))
}

func doIt(input []string) (int, int) {
	beams := make([]int, len(input[0])+2)
	splits := 0
	for _, s := range input {
		for x, r := range s {
			if r == 'S' {
				beams[x]++
			} else if r == '^' && beams[x] > 0 {
				beams[x-1] += beams[x]
				beams[x+1] += beams[x]
				beams[x] = 0
				splits++
			}
		}
	}

	timelines := 0
	for _, beam := range beams {
		timelines += beam
	}

	return splits, timelines
}
