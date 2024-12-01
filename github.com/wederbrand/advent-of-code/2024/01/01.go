package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"slices"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2024/01/input.txt", "\n")

	list1 := make([]int, 0)
	list2 := make([]int, 0)
	counts := make(map[int]int)

	for _, line := range inFile {
		var a, b int
		fmt.Sscanf(line, "%d   %d", &a, &b)
		list1 = append(list1, a)
		list2 = append(list2, b)
		counts[b]++
	}

	slices.Sort(list1)
	slices.Sort(list2)

	part1 := 0
	part2 := 0
	for i := range list1 {
		a := list1[i]
		b := list2[i]
		part1 += util.IntAbs(a - b)
		part2 += a * counts[a]
	}

	fmt.Println("Part 1:", part1, "in", time.Since(start))
	fmt.Println("Part 2:", part2, "in", time.Since(start))
}
