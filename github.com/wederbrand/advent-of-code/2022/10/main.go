package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
)

func main() {
	inFile := util.GetFileContents("2022/10/input.txt", "\n")

	cycle := 0
	x := 1
	history := make([]int, 0)
	for _, instruction := range inFile {
		cycle++
		if instruction == "noop" {
			history = append(history, x)
		} else {
			var add int
			fmt.Sscanf(instruction, "addx %d", &add)
			history = append(history, x)

			cycle++
			history = append(history, x)
			x += add
		}
	}

	sum := 0
	sum += 20 * history[20-1]
	sum += 60 * history[60-1]
	sum += 100 * history[100-1]
	sum += 140 * history[140-1]
	sum += 180 * history[180-1]
	sum += 220 * history[220-1]

	fmt.Println("part 1:", sum)
	fmt.Println("part 2:")

	for line := 0; line < 6; line++ {
		for crt := 0; crt < 40; crt++ {
			index := line*40 + crt

			sprite := history[index] % 40
			if sprite-1 == crt || sprite == crt || sprite+1 == crt {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
