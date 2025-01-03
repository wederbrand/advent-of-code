package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	input := 29000000

	p1 := doIt(input, 10, input)
	fmt.Println("Part 1: ", p1, "in", time.Since(start))

	p2 := doIt(input, 11, 50)
	fmt.Println("Part 2: ", p2, "in", time.Since(start))
}

func doIt(input int, factor int, maxHouses int) int {
	houses := make([]int, 1000000) // one million is enough
	lowestHouse := math.MaxInt
	for elf := 1; elf < len(houses); elf++ {
		if elf > lowestHouse {
			return lowestHouse
		}
		for houseNumber := elf; houseNumber < len(houses) && houseNumber <= elf*maxHouses; houseNumber += elf {
			houses[houseNumber] += elf * factor
			if houses[houseNumber] >= input {
				if houseNumber < lowestHouse {
					lowestHouse = houseNumber
				}
			}
		}
	}

	panic("No solution found")
}
