package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	input := 3001330

	p1 := part1(input)
	fmt.Println("Part 1: ", p1, "in", time.Since(start))

	p2 := part2(input)
	fmt.Println("Part 2: ", p2, "in", time.Since(start))
}

func part1(input int) int {
	elfs := make([]int, input)
	for i := 0; i < input; i++ {
		// you get a present, you get a present, everyone gets a present!
		elfs[i] = i + 1
	}

	for len(elfs) > 1 {
		elfs = append(elfs[2:], elfs[0])
	}
	return elfs[0]
}

func part2(input int) int {
	elfs := make([]int, input)
	for i := 0; i < input; i++ {
		// you get a present, you get a present, everyone gets a present!
		elfs[i] = i + 1
	}

	for len(elfs) > 1 {
		target := len(elfs) / 2
		elfs = append(append(elfs[1:target], elfs[target+1:]...), elfs[0])
	}
	return elfs[0]
}
