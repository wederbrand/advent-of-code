package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2015/16/input.txt", "\n")

	facit := make(map[string]int)
	facit["children"] = 3
	facit["cats"] = 7
	facit["samoyeds"] = 2
	facit["pomeranians"] = 3
	facit["akitas"] = 0
	facit["vizslas"] = 0
	facit["goldfish"] = 5
	facit["trees"] = 3
	facit["cars"] = 2
	facit["perfumes"] = 1

	part1 := p1(inFile, facit)
	fmt.Println("Part 1: ", part1, "in", time.Since(start))

	part2 := p2(inFile, facit)
	fmt.Println("Part 2: ", part2, "in", time.Since(start))
}

func p1(inFile []string, facit map[string]int) int {
	for _, s := range inFile {
		s = strings.TrimSpace(s)
		s = strings.ReplaceAll(s, ":", "")
		s = strings.ReplaceAll(s, ",", "")
		split := strings.Split(s, " ")
		// split 0 is the name, 1 is the aunt number
		// 2,3 then 4,5 and so on are the properties

		found := true
		for i := 2; i < len(split); i += 2 {
			if Atoi(split[i+1]) != facit[split[i]] {
				found = false
				break
			}
		}
		if found {
			return Atoi(split[1])
		}
	}
	return 0
}

func p2(inFile []string, facit map[string]int) int {
	for _, s := range inFile {
		s = strings.TrimSpace(s)
		s = strings.ReplaceAll(s, ":", "")
		s = strings.ReplaceAll(s, ",", "")
		split := strings.Split(s, " ")
		// split 0 is the name, 1 is the aunt number
		// 2,3 then 4,5 and so on are the properties

		found := true
		for i := 2; i < len(split); i += 2 {
			if split[i] == "cats" || split[i] == "trees" {
				if Atoi(split[i+1]) <= facit[split[i]] {
					found = false
					break
				}
				continue
			} else if split[i] == "pomeranians" || split[i] == "goldfish" {
				if Atoi(split[i+1]) >= facit[split[i]] {
					found = false
					break
				}
				continue
			} else if Atoi(split[i+1]) != facit[split[i]] {
				found = false
				break
			}
		}
		if found {
			return Atoi(split[1])
		}
	}
	return 0
}
