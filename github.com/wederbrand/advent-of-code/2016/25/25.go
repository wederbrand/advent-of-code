package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2016/25/input.txt", "\n")

	for i := 0; i < math.MaxInt; i++ {
		if i%1000 == 0 {
			fmt.Println(i)
		}

		instructions := make([]string, 0)
		for _, line := range inFile {
			instructions = append(instructions, line)
		}

		registers := make(map[string]int)
		registers["a"] = i
		if doIt(instructions, registers) {
			fmt.Println("Part 1: ", i, "in", time.Since(start))
			break
		}
	}
}

func doIt(instructions []string, registers map[string]int) bool {
	expectedOut := 0
	correctCount := 0
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
			if _, err := strconv.Atoi(split[2]); err == nil {
				// split[2] is a number and we can't copy to a number
				pointer++
				continue
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
				value2, err := strconv.Atoi(split[2])
				if err != nil {
					value2 = registers[split[2]]
				}

				pointer += value2
			} else {
				pointer++
			}
		case "tgl":
			value, err := strconv.Atoi(split[1])
			if err != nil {
				value = registers[split[1]]
			}

			if pointer+value >= len(instructions) {
				pointer++
				continue
			}
			inst := instructions[pointer+value]
			args := strings.Split(inst, " ")
			if len(args) == 2 {
				// one argument
				if args[0] == "inc" {
					instructions[pointer+value] = "dec " + args[1]
				} else {
					instructions[pointer+value] = "inc " + args[1]
				}
			} else if len(args) == 3 {
				// two arguments
				if args[0] == "jnz" {
					instructions[pointer+value] = "cpy " + args[1] + " " + args[2]
				} else {
					instructions[pointer+value] = "jnz " + args[1] + " " + args[2]
				}
			} else {
				panic("Unknown instruction")
			}
			pointer++
		case "out":
			value, err := strconv.Atoi(split[1])
			if err != nil {
				value = registers[split[1]]
			}
			if value != expectedOut {
				return false
			}
			expectedOut = 1 - expectedOut

			correctCount++
			if correctCount == 10 {
				return true
			}
			pointer++
		default:
			panic("Unknown instruction")
		}
	}

	panic("No solution found")
}
