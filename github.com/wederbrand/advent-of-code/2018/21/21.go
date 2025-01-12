package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"time"
)

var instructions = map[string]Inst{
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

type regPointer *[6]int

var start = time.Now()

func main() {
	inFile := GetFileContents("2018/21/input.txt", "\n")

	instructionPointer := int(inFile[0][4] - '0')
	infile := inFile[1:]
	registers := [6]int{16134795, 0, 0, 0, 0, 0}
	doIt(registers, instructionPointer, infile)

	// register 0 is found from reverse engineering the input code
	// and run the program to pointer 28 so see the value needed in register 0
	fmt.Println("Part 1:", 16134795, "in", time.Since(start))

	registers = [6]int{0, 0, 0, 0, 0, 0}
	doIt(registers, instructionPointer, infile)
	fmt.Println("Part 2:", "in", time.Since(start))
}

func doIt(registers [6]int, instructionPointer int, infile []string) {
	seen := make(map[int]bool)

	pointer := registers[instructionPointer]
	for pointer < len(infile) {
		instruction := infile[pointer]
		var a, b, c int
		var inst string
		fmt.Sscanf(instruction, "%s %d %d %d", &inst, &a, &b, &c)

		if pointer == 28 {
			if seen[registers[3]] {
				break
			}
			seen[registers[3]] = true

			fmt.Println("The last outputted value when it stops is the right one. It takes 40 minutes and size 14250. Value: ", registers[3], "size", len(seen), "in", time.Since(start))
		}

		registers[instructionPointer] = pointer
		instructions[inst](a, b, c, &registers)
		pointer = registers[instructionPointer]
		pointer++
	}
}

type Inst func(a, b, c int, registers regPointer)

func addr(a, b, c int, registers regPointer) {
	registers[c] = registers[a] + registers[b]
}

func addi(a, b, c int, registers regPointer) {
	registers[c] = registers[a] + b
}

func mulr(a, b, c int, registers regPointer) {
	registers[c] = registers[a] * registers[b]
}

func muli(a, b, c int, registers regPointer) {
	registers[c] = registers[a] * b
}

func banr(a, b, c int, registers regPointer) {
	registers[c] = registers[a] & registers[b]
}

func bani(a, b, c int, registers regPointer) {
	registers[c] = registers[a] & b
}

func borr(a, b, c int, registers regPointer) {
	registers[c] = registers[a] | registers[b]
}

func bori(a, b, c int, registers regPointer) {
	registers[c] = registers[a] | b
}

func setr(a, b, c int, registers regPointer) {
	registers[c] = registers[a]
}

func seti(a, b, c int, registers regPointer) {
	registers[c] = a
}

func gtir(a, b, c int, registers regPointer) {
	if a > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func gtri(a, b, c int, registers regPointer) {
	if registers[a] > b {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func gtrr(a, b, c int, registers regPointer) {
	if registers[a] > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqir(a, b, c int, registers regPointer) {
	if a == registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqri(a, b, c int, registers regPointer) {
	if registers[a] == b {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqrr(a, b, c int, registers regPointer) {
	if registers[a] == registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}

}
