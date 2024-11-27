package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2016/07/input.txt", "\n")
	part1, part2 := doIt(inFile)
	fmt.Println("Part 1: ", part1, "in", time.Since(start))
	fmt.Println("Part 2: ", part2, "in", time.Since(start))
}

func doIt(inFile []string) (int, int) {
	part1 := 0
	part2 := 0

	for _, line := range inFile {
		if supportsTLS(line) {
			part1++
		}
		if supportsSSL(line) {
			part2++
		}
	}

	return part1, part2
}

func supportsTLS(line string) bool {
	outside, inside := splitWithBrackets(line)

	potential := false
	for _, segment := range outside {
		if hasAbba(segment) {
			potential = true
			break
		}
	}
	if potential {
		for _, segment := range inside {
			if hasAbba(segment) {
				potential = false
				break
			}
		}
	}

	return potential
}

func supportsSSL(line string) bool {
	outside, inside := splitWithBrackets(line)

	abas := make([]string, 0)
	for _, segment := range outside {
		abas = append(abas, getAbas(segment)...)
	}

	for _, aba := range abas {
		bab := string(aba[1]) + string(aba[0]) + string(aba[1])
		for _, segment := range inside {
			if strings.Contains(segment, bab) {
				return true
			}
		}
	}
	return false
}

func hasAbba(segment string) bool {
	if len(segment) < 4 {
		return false
	}

	for i := 3; i < len(segment); i++ {
		if segment[i-3] == segment[i] && segment[i-2] == segment[i-1] && segment[i-3] != segment[i-2] {
			return true
		}
	}

	return false
}

func getAbas(segment string) []string {
	result := make([]string, 0)
	if len(segment) < 3 {
		return result
	}

	for i := 2; i < len(segment); i++ {
		if segment[i-2] == segment[i] && segment[i-2] != segment[i-1] {
			result = append(result, segment[i-2:i+1])
		}
	}

	return result
}

func splitWithBrackets(line string) (outside []string, inside []string) {
	current := ""
	for _, c := range line {
		if c == '[' {
			outside = append(outside, current)
			current = ""
		} else if c == ']' {
			inside = append(inside, current)
			current = ""
		} else {
			current += string(c)
		}
	}
	outside = append(outside, current)
	return
}
