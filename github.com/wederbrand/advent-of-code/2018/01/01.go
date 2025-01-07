package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2018/01/input.txt", "\n")

	input := make([]int, 0)
	for _, s := range inFile {
		input = append(input, Atoi(s))
	}

	seen := make(map[int]bool)
	freq := 0
	found := false
	first := true
	for !found {
		for _, iValue := range input {
			freq += iValue
			_, ok := seen[freq]
			if ok {
				fmt.Println("Part 2: ", freq, "in", time.Since(start))
				found = true
				break
			}
			seen[freq] = true
		}
		if first {
			fmt.Println("Part 1: ", freq, "in", time.Since(start))
		}
		first = false
	}
}
