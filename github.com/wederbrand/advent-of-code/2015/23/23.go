package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	instructions := GetFileContents("2015/23/input.txt", "\n")

	for i := range instructions {
		instructions[i] = strings.Replace(instructions[i], ",", "", -1)
	}

	reg := make(map[string]int)
	doIt(instructions, reg)
	fmt.Println("Part 1: ", reg["b"], "in", time.Since(start))

	reg = make(map[string]int)
	reg["a"] = 1
	doIt(instructions, reg)
	fmt.Println("Part 2: ", reg["b"], "in", time.Since(start))
}

func doIt(instructions []string, reg map[string]int) {
	pointer := 0

	for pointer < len(instructions) {
		inst := instructions[pointer]
		parts := strings.Split(inst, " ")
		switch parts[0] {
		case "hlf":
			reg[parts[1]] /= 2
			pointer++
		case "tpl":
			reg[parts[1]] *= 3
			pointer++
		case "inc":
			reg[parts[1]]++
			pointer++
		case "jmp":
			pointer += Atoi(parts[1])
		case "jie":
			if reg[parts[1]]%2 == 0 {
				pointer += Atoi(parts[2])
			} else {
				pointer++
			}
		case "jio":
			if reg[parts[1]] == 1 {
				pointer += Atoi(parts[2])
			} else {
				pointer++
			}
		}
	}
}
