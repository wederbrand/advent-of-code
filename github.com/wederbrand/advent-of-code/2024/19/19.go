package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2024/19/input.txt", "\n")

	towels := make(map[string][]string)
	for _, s := range strings.Split(inFile[0], ",") {
		s = strings.Trim(s, " ")
		firstLetter := string(s[0])
		towels[firstLetter] = append(towels[firstLetter], s)
	}

	part1 := 0
	part2 := 0
	memory := make(map[string]int)
	for _, target := range inFile[2:] {
		ways := possible(towels, target, memory)
		part2 += ways
		if ways > 0 {
			part1++
		}
	}

	fmt.Println("Part 1:", part1, "in", time.Since(start))
	fmt.Println("Part 2:", part2, "in", time.Since(start))
}

func possible(towels map[string][]string, target string, memory map[string]int) int {
	if target == "" {
		return 1
	}

	mem, ok := memory[target]
	if ok {
		return mem
	}

	total := 0
	firstLetter := string(target[0])
	for _, towel := range towels[firstLetter] {
		if strings.HasPrefix(target, towel) {
			total += possible(towels, target[len(towel):], memory)
		}
	}
	memory[target] = total
	return total
}
