package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2016/12/input.txt", "\n")

	instructions := make([]string, 0)
	for _, line := range inFile {
		instructions = append(instructions, line)
	}

	registers := make(map[string]int)
	doIt(instructions, registers)
	fmt.Println("Part 1: ", registers["a"], "in", time.Since(start))

	registers = make(map[string]int)
	registers["c"] = 1
	doIt(instructions, registers)
	fmt.Println("Part 2: ", registers["a"], "in", time.Since(start))
}

func doIt(instructions []string, registers map[string]int) {
	pointer := 0
	for pointer < len(instructions) {
		instruction := instructions[pointer]
		split := strings.Split(instruction, " ")
		switch split[0] {
		case "cpy":
			value, err := strconv.Atoi(split[1])
			if err != nil {
				value = registers[split[1]]
			}
			registers[split[2]] = value
			pointer++
		case "inc":
			registers[split[1]]++
			pointer++
		case "dec":
			registers[split[1]]--
			pointer++
		case "jnz":
			value, err := strconv.Atoi(split[1])
			if err != nil {
				value = registers[split[1]]
			}
			if value != 0 {
				pointer += Atoi(split[2])
			} else {
				pointer++
			}
		default:
			panic("Unknown instruction")
		}
	}
}
