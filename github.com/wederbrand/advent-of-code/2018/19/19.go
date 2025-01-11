package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	part1 := cheatIt(0)
	fmt.Println("Part 1:", part1, "in", time.Since(start))

	part2 := cheatIt(1)
	fmt.Println("Part 2:", part2, "in", time.Since(start))
}

// this was manually translated from the instructions, then optimized
// it finds all the divisors of r1 and sums them up
// r1 is calculated once; then the loop is run
func cheatIt(r0 int) int {
	r1 := 0

	if r0 == 0 {
		r1 = 998
		r0 = 0
	} else {
		r1 = 10551398
		r0 = 0
	}

	for r2 := 1; r2 <= r1; r2++ {
		if r1%r2 == 0 {
			r0 += r1 / r2
		}
	}

	return r0
}
