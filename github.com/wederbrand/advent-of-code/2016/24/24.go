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
	inFile := GetFileContents("2016/23/input.txt", "\n")

	instructions := make([]string, 0)
	for _, line := range inFile {
		instructions = append(instructions, line)
	}

	registers := make(map[string]int)
	registers["a"] = 12
	doIt(instructions, registers)
	fmt.Println("Part 1: ", registers["a"], "in", time.Since(start))

	// registers = make(map[string]int)
	// registers["c"] = 1
	// doIt(instructions, registers)
	// fmt.Println("Part 2: ", registers["a"], "in", time.Since(start))
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
		default:
			panic("Unknown instruction")
		}
	}
}
