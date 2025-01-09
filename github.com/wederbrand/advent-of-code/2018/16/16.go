package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"golang.org/x/exp/maps"
	"slices"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2018/16/input.txt", "\n\n\n")

	instructions := map[string]Inst{
		"addr": addr,
		"addi": addi,
		"mulr": mulr,
		"muli": muli,
		"banr": banr,
		"bani": bani,
		"borr": borr,
		"bori": bori,
		"setr": setr,
		"seti": seti,
		"gtir": gtir,
		"gtri": gtri,
		"gtrr": gtrr,
		"eqir": eqir,
		"eqri": eqri,
		"eqrr": eqrr,
	}

	mapping := map[int][]string{}
	for i := 0; i < len(instructions); i++ {
		mapping[i] = maps.Keys(instructions)
	}

	part1 := 0
	split := strings.Split(inFile[0], "\n")
	for i := 0; i < len(split); i += 4 {
		before := [4]int{}
		fmt.Sscanf(split[i], "Before: [%d, %d, %d, %d]", &before[0], &before[1], &before[2], &before[3])
		inst := [4]int{}
		fmt.Sscanf(split[i+1], "%d %d %d %d", &inst[0], &inst[1], &inst[2], &inst[3])
		after := [4]int{}
		fmt.Sscanf(split[i+2], "After: [%d, %d, %d, %d]", &after[0], &after[1], &after[2], &after[3])

		count := 0
		for n, f := range instructions {
			registers := before
			f(inst[1], inst[2], inst[3], &registers)
			if registers == after {
				count++
			} else {
				mapping[inst[0]] = slices.DeleteFunc(mapping[inst[0]], func(s string) bool {
					return s == n
				})
			}
		}

		if count >= 3 {
			part1++
		}
	}
	fmt.Println("Part 1:", part1, "in", time.Since(start))

	for {
		done := true
		for i, v := range mapping {
			if len(v) == 1 {
				theOne := v[0]
				for j := range mapping {
					if i != j {
						mapping[j] = slices.DeleteFunc(mapping[j], func(s string) bool {
							return s == theOne
						})
					}
				}
			} else {
				done = false
			}
		}
		if done {
			break
		}
	}

	registers := [4]int{}
	split = strings.Split(inFile[1], "\n")
	for _, line := range split {
		inst := [4]int{}
		fmt.Sscanf(line, "%d %d %d %d", &inst[0], &inst[1], &inst[2], &inst[3])
		instructions[mapping[inst[0]][0]](inst[1], inst[2], inst[3], &registers)
	}

	part2 := registers[0]
	fmt.Println("Part 2:", part2, "in", time.Since(start))
}

type Inst func(a, b, c int, registers *[4]int)

func addr(a, b, c int, registers *[4]int) {
	registers[c] = registers[a] + registers[b]
}

func addi(a, b, c int, registers *[4]int) {
	registers[c] = registers[a] + b
}

func mulr(a, b, c int, registers *[4]int) {
	registers[c] = registers[a] * registers[b]
}

func muli(a, b, c int, registers *[4]int) {
	registers[c] = registers[a] * b
}

func banr(a, b, c int, registers *[4]int) {
	registers[c] = registers[a] & registers[b]
}

func bani(a, b, c int, registers *[4]int) {
	registers[c] = registers[a] & b
}

func borr(a, b, c int, registers *[4]int) {
	registers[c] = registers[a] | registers[b]
}

func bori(a, b, c int, registers *[4]int) {
	registers[c] = registers[a] | b
}

func setr(a, b, c int, registers *[4]int) {
	registers[c] = registers[a]
}

func seti(a, b, c int, registers *[4]int) {
	registers[c] = a
}

func gtir(a, b, c int, registers *[4]int) {
	if a > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func gtri(a, b, c int, registers *[4]int) {
	if registers[a] > b {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func gtrr(a, b, c int, registers *[4]int) {
	if registers[a] > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqir(a, b, c int, registers *[4]int) {
	if a == registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqri(a, b, c int, registers *[4]int) {
	if registers[a] == b {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqrr(a, b, c int, registers *[4]int) {
	if registers[a] == registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}

}
