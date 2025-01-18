package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/2019/computer"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2019/21/input.txt", "\n")
	part1 := program(inFile, []string{
		"NOT C J",
		"AND D J",
		"NOT A T",
		"OR T J",
		"WALK",
	})

	fmt.Println("part1: ", part1, "in", time.Since(start))

	part2 := program(inFile, []string{
		"NOT C J",
		"AND D J",
		"AND H J",
		"NOT B T",
		"AND D T",
		"OR T J",
		"NOT A T",
		"OR T J",
		"RUN",
	})
	fmt.Println("part2: ", part2, "in", time.Since(start))
}

func program(inFile []string, instructions []string) int {
	input := make([]int, 0)
	for _, instruction := range instructions {
		for _, c := range instruction {
			input = append(input, int(c))
		}
		input = append(input, 10)
	}

	in := func() int {
		i := input[0]
		input = input[1:]
		return i
	}

	output := 0
	out := func(i int) {
		if i > 256 {
			output = i
		} else {
			fmt.Print(string(i))
		}
	}

	computer := NewComputer(inFile, in, out)
	computer.Run()
	return output
}
