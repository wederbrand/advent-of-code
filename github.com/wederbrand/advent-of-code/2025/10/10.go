package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
	"time"

	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"

	// This requires the Z3 solver
	// CGO_CFLAGS=-I/opt/homebrew/include
	// CGO_LDFLAGS=-L/opt/homebrew/lib
	"github.com/aclements/go-z3/z3"
)

type Machine struct {
	buttons [][]int
	target  []bool
	joltage []int
}

func (m *Machine) setTarget(input string) {
	input = input[1 : len(input)-1]
	m.target = make([]bool, len(input))
	for i, r := range input {
		if r == '#' {
			m.target[i] = true
		}
	}
}

func (m *Machine) setJoltage(input string) {
	m.joltage = make([]int, 0)
	input = input[1 : len(input)-1]
	for _, s := range strings.Split(input, ",") {
		m.joltage = append(m.joltage, Atoi(s))
	}
}

func (m *Machine) addButton(input string) {
	input = input[1 : len(input)-1]

	button := make([]int, 0)
	for _, s := range strings.Split(input, ",") {
		atoi := Atoi(s)
		button = append(button, atoi)
	}

	m.buttons = append(m.buttons, button)
}

func (m *Machine) p1() int {
	minPresses := math.MaxInt

	combos := booleanCombos(len(m.buttons))

	for _, combo := range combos {
		outcome := make([]bool, len(m.target))
		for i, button := range m.buttons {
			if combo[i] {
				for _, v := range button {
					outcome[v] = !outcome[v]
				}
			}
		}
		if slices.Equal(m.target, outcome) {
			count := 0
			for _, pressed := range combo {
				if pressed {
					count++
				}
			}
			if count < minPresses {
				minPresses = count
			}
		}
	}

	return minPresses
}

func booleanCombos(n int) [][]bool {
	if n <= 0 {
		return [][]bool{}
	}

	total := 1 << n // 2^n
	result := make([][]bool, total)

	for mask := 0; mask < total; mask++ {
		comb := make([]bool, n)
		for bit := 0; bit < n; bit++ {
			comb[bit] = (mask & (1 << bit)) != 0
		}
		result[mask] = comb
	}

	return result
}

func (m *Machine) ashamed(target []int) int {
	ctx := z3.NewContext(nil)
	solver := z3.NewSolver(ctx)

	intSort := ctx.IntSort()
	I := func(i int) z3.Int {
		return ctx.FromInt(int64(i), intSort).(z3.Int)
	}

	// each variable is the number of times button i is pressed
	variables := make([]z3.Int, len(m.buttons))
	for i := range m.buttons {
		intConst := ctx.IntConst(fmt.Sprintf("b%d", i))
		variables[i] = intConst
		solver.Assert(intConst.GE(I(0))) // all variables >= 0
	}

	for targetIndex, targetValue := range target {
		sum := I(0)
		for buttonIndex, button := range m.buttons {
			for _, value := range button {
				if value != targetIndex {
					continue
				}
				sum = sum.Add(variables[buttonIndex])
			}
		}
		solver.Assert(sum.Eq(I(targetValue)))
	}

	// Add restriction on total number of presses
	totalPresses := ctx.IntConst("totalPresses")
	sumPresses := I(0)
	for _, v := range variables {
		sumPresses = sumPresses.Add(v)
	}
	solver.Assert(totalPresses.Eq(sumPresses)) // we can never press more than the total increases

	minumum := math.MaxInt
	for {
		if sat, _ := solver.Check(); sat {
			model := solver.Model()

			val := model.Eval(totalPresses, true)
			intVal, _ := val.(z3.Int)
			asInt64, _, _ := intVal.AsInt64()
			if int(asInt64) < minumum {
				minumum = int(asInt64)
				solver.Assert(totalPresses.LT(I(minumum)))
			}
		} else {
			return minumum
		}
	}
}

func main() {
	start := time.Now()
	input := GetFileContents("2025/10/input.txt", "\n")

	machines := make([]Machine, 0)
	for _, line := range input {
		m := Machine{}
		fields := strings.Split(line, " ")
		// first field is target, treat as binary
		m.setTarget(fields[0])

		// last field is joltage, treat as []int
		m.setJoltage(fields[len(fields)-1])

		// middle fields are buttons, treat as binaries
		for _, field := range fields[1 : len(fields)-1] {
			m.addButton(field)
		}

		machines = append(machines, m)
	}

	part1 := 0
	for _, machine := range machines {
		part1 += machine.p1()
	}

	fmt.Println("Part 1:", part1, "in", time.Since(start))

	part2 := 0
	for _, machine := range machines {
		ashamed := machine.ashamed(machine.joltage)
		part2 += ashamed
	}
	fmt.Println("Part 2:", part2, "in", time.Since(start))

}
