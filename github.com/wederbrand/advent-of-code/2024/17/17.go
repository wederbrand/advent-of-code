package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"math"
	"slices"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2024/17/input.txt", "\n")

	var a, b, c int
	fmt.Sscanf(inFile[0], "Register A: %d", &a)
	fmt.Sscanf(inFile[1], "Register B: %d", &b)
	fmt.Sscanf(inFile[2], "Register C: %d", &c)

	orig := strings.Split(inFile[4], ": ")[1]
	program := strings.Split(orig, ",")

	p1 := runProgram(program, a, b, c)
	part1 := strings.Join(p1, ",")
	fmt.Println("Part 1:", part1, "in", time.Since(start))

	facit := slices.Clone(program)
	slices.Reverse(facit)

	part2 := tryIt(program, facit, 0, 0)
	fmt.Println("Part 2:", part2, "in", time.Since(start))
}

/**
 * The program was reverse engineered to understand that it looked at three bits of A in each iteration
 * It also calculated a C, some offset into the "higher" parts of A making it clear it was smart
 * to do it reversed, one output at a time.
 */
func tryIt(program []string, facit []string, i int, a int) int {
	for j := 0; j <= 0b111; j++ {
		a2 := a + j
		out := runProgram(program, a2, 0, 0)
		if out[0] == facit[len(out)-1] {
			if i+1 == len(program) {
				// done
				return a2
			} else {
				it := tryIt(program, facit, i+1, (a+j)<<3)
				if it == 0 {
					// nope, try next
				} else {
					return it
				}
			}
		}
	}
	return 0
}

func dv(inst string, a int, b int, c int) int {
	denom := combo(inst, a, b, c)
	return a / int(math.Pow(2, float64(denom)))
}

func bst(inst string, a int, b int, c int) int {
	return combo(inst, a, b, c) % 8
}

func runProgram(program []string, a int, b int, c int) []string {
	i := 0
	output := make([]int, 0)
	for i < len(program) {
		switch program[i] {
		case "0":
			a = dv(program[i+1], a, b, c)
		case "6":
			b = dv(program[i+1], a, b, c)
		case "7":
			c = dv(program[i+1], a, b, c)
		case "1":
			b ^= Atoi(program[i+1])
		case "2":
			b = bst(program[i+1], a, b, c)
		case "3":
			if a != 0 {
				i = Atoi(program[i+1])
				i -= 2
			}
		case "4":
			b ^= c
		case "5":
			output = append(output, combo(program[i+1], a, b, c)%8)
		}

		i += 2
	}

	part1 := make([]string, len(output))
	for i, v := range output {
		part1[i] = fmt.Sprintf("%d", v)
	}

	return part1
}

func combo(s string, a int, b int, c int) int {
	switch s {
	case "4":
		return a
	case "5":
		return b
	case "6":
		return c
	case "7":
		panic("7 is reserved")
	default:
		return Atoi(s)
	}
}
