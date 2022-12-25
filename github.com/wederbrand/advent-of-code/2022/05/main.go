package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strconv"
	"strings"
)

func main() {
	inFile := util.GetFileContents("2022/05/input.txt", "\n")

	var stacks []string

	nbrOfStacks := 0
	programStart := 0
	for i, s := range inFile {
		if strings.Contains(s, "[") {
			continue
		}
		// find the one with number
		splits := strings.Split(s, " ")
		nbrOfStacks, _ = strconv.Atoi(splits[len(splits)-1])
		stacks = make([]string, nbrOfStacks)
		for j := i - 1; j >= 0; j-- {
			for c := 0; c < nbrOfStacks; c++ {
				s2 := inFile[j]
				if c*4+1 >= len(s2) {
					continue
				}
				crate := string(s2[c*4+1])
				stacks[c] += crate
			}
		}
		programStart = i + 2
		break
	}

	for i := 0; i < nbrOfStacks; i++ {
		stacks[i] = strings.TrimSpace(stacks[i])
	}

	for i := programStart; i < len(inFile); i++ {
		nbr := 0
		from := 0
		to := 0
		fmt.Sscanf(inFile[i], "move %d from %d to %d", &nbr, &from, &to)
		if nbr == 0 {
			continue
		}
		from -= 1
		to -= 1
		stacks[to] += stacks[from][len(stacks[from])-nbr:]
		stacks[from] = stacks[from][0 : len(stacks[from])-nbr]
	}

	result := ""
	for _, stack := range stacks {
		result += string(stack[len(stack)-1])
	}
	fmt.Println(result)

}
