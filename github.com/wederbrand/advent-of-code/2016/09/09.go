package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2016/09/input.txt", "\n")

	fmt.Println("Part 1: ", part1(inFile[0]), "in", time.Since(start))
	fmt.Println("Part 2: ", part2(inFile[0]), "in", time.Since(start))
}

func part1(input string) int {
	result := ""
	for i, r := range input {
		if r == '(' {
			end := strings.Index(input[i:], ")")
			if end == -1 {
				panic("no end")
			}
			marker := input[i+1 : i+end]
			parts := strings.Split(marker, "x")
			length := util.Atoi(parts[0])
			times := util.Atoi(parts[1])
			next := input[i+end+1 : i+end+1+length]
			return len(result) + times*len(next) + part1(input[i+end+1+length:])
		} else {
			result += string(r)
		}
	}

	return len(result)
}

func part2(input string) int {
	result := ""
	for i, r := range input {
		if r == '(' {
			end := strings.Index(input[i:], ")")
			if end == -1 {
				panic("no end")
			}
			marker := input[i+1 : i+end]
			parts := strings.Split(marker, "x")
			length := util.Atoi(parts[0])
			times := util.Atoi(parts[1])
			next := input[i+end+1 : i+end+1+length]
			return len(result) + times*part2(next) + part2(input[i+end+1+length:])
		} else {
			result += string(r)
		}
	}

	return len(result)
}
