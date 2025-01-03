package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2015/19/input.txt", "\n\n")

	inMolecule := inFile[1]
	replacements := make(map[string][]string)
	reversed := make(map[string]string)
	for _, s := range strings.Split(inFile[0], "\n") {
		fields := strings.Fields(s)
		from := fields[0]
		to := fields[2]
		if _, ok := replacements[from]; !ok {
			replacements[from] = make([]string, 0)
		}

		replacements[from] = append(replacements[from], to)
		reversed[to] = from
	}

	part1 := replace(inMolecule, replacements)
	fmt.Println("Part 1: ", len(part1), "in", time.Since(start))

	part2 := 0
	for inMolecule != "e" {
		for to, from := range reversed {
			if strings.Contains(inMolecule, to) {
				inMolecule = strings.Replace(inMolecule, to, from, 1)
				part2++
			}
		}
	}

	fmt.Println("Part 2: ", part2, "in", time.Since(start))
}

func replace(molecule string, replacements map[string][]string) map[string]bool {
	results := make(map[string]bool)
	for from, tos := range replacements {
		for _, to := range tos {
			for i := 0; i < len(molecule); i++ {
				if i+len(from) <= len(molecule) && molecule[i:i+len(from)] == from {
					newMolecule := molecule[:i] + to + molecule[i+len(from):]
					results[newMolecule] = true
				}
			}
		}
	}

	return results
}
