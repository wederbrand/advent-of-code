package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strings"
	"time"
)

type Instruction struct {
	ruleIndex int
	rule      string
	value     int
	target    string
}

type Part [4]int

type Range [4][2]int

type State struct {
	instructions     []*Instruction
	instructionIndex int
	xmas             Range
}

func main() {
	startTimer := time.Now()
	inFile := util.GetFileContents("2023/19/input.txt", "\n")

	workflows := make(map[string][]*Instruction)
	parts := make([]Part, 0)
	for _, s := range inFile {
		if s == "" {
			continue
		}
		if strings.HasPrefix(s, "{") {
			// parse Part
			split := strings.Split(s[1:len(s)-1], ",")
			p := Part{}
			for i, s2 := range split {
				value := util.Atoi(s2[2:])
				p[i] = value
			}
			parts = append(parts, p)
		} else {
			// parse Instruction
			name := strings.Split(s, "{")[0]
			rest := strings.Split(s, "{")[1]
			rest = rest[:len(rest)-1]
			for _, inst := range strings.Split(rest, ",") {
				i := new(Instruction)
				workflows[name] = append(workflows[name], i)

				if !strings.Contains(inst, ":") {
					i.ruleIndex = -1
					i.target = inst
				} else {
					split := strings.Split(inst, ":")
					i.target = split[1]
					switch string(split[0][0]) {
					case "x":
						i.ruleIndex = 0
					case "m":
						i.ruleIndex = 1
					case "a":
						i.ruleIndex = 2
					case "s":
						i.ruleIndex = 3
					}
					i.rule = string(split[0][1])
					i.value = util.Atoi(split[0][2:])
				}
			}
		}
	}
	fmt.Println("parsing in", time.Since(startTimer))
	startTimer = time.Now()

	part1 := 0
	for _, p := range parts {
		wf := workflows["in"]

		for i := 0; i < len(wf); i++ {
			instruction := wf[i]
			target := ""
			if instruction.ruleIndex == -1 {
				target = instruction.target
			} else {
				xmas := p[instruction.ruleIndex]
				if (instruction.rule == "<" && xmas < instruction.value) || (instruction.rule == ">" && xmas > instruction.value) {
					target = instruction.target
				}
			}

			if target != "" {
				if target == "A" {
					part1 += p[0] + p[1] + p[2] + p[3]
				}
				if target == "R" {
					break
				}

				wf = workflows[instruction.target]
				i = -1
			}
		}
	}

	fmt.Println("part1: ", part1, "in", time.Since(startTimer))
	startTimer = time.Now()

	part2 := 0

	q := make([]State, 0)
	startState := State{workflows["in"], 0, Range{{1, 4000}, {1, 4000}, {1, 4000}, {1, 4000}}}
	q = append(q, startState)

	for len(q) > 0 {
		s := q[0]
		q = q[1:]
		// split just one Instruction in each iteration
		instruction := s.instructions[s.instructionIndex]
		if instruction.ruleIndex == -1 {
			product, nextState := calculateOutcome(instruction, s.xmas, workflows)
			part2 += product
			if nextState != nil {
				q = append(q, *nextState)
			}

		} else {
			// handle split
			matchingXmas := s.xmas
			remainingXmas := s.xmas
			splitValue := s.xmas[instruction.ruleIndex]

			// split - assume <
			matching := [2]int{splitValue[0], instruction.value - 1}
			remaining := [2]int{instruction.value, splitValue[1]}

			if instruction.rule == ">" {
				matching = [2]int{instruction.value + 1, splitValue[1]}
				remaining = [2]int{splitValue[0], instruction.value}
			}

			matchingXmas[instruction.ruleIndex] = matching
			remainingXmas[instruction.ruleIndex] = remaining

			product, nextState := calculateOutcome(instruction, matchingXmas, workflows)
			part2 += product
			if nextState != nil {
				q = append(q, *nextState)
			}

			remainingState := State{s.instructions, s.instructionIndex + 1, remainingXmas}
			q = append(q, remainingState)
		}
	}

	fmt.Println("part2: ", part2, "in", time.Since(startTimer))
}

func calculateOutcome(inst *Instruction, s Range, workflows map[string][]*Instruction) (int, *State) {
	if inst.target == "A" {
		prod := 1
		for i := 0; i < 4; i++ {
			prod *= s[i][1] - s[i][0] + 1
		}
		return prod, nil
	} else if inst.target != "R" {
		nextState := State{workflows[inst.target], 0, s}
		return 0, &nextState
	}
	return 0, nil
}
