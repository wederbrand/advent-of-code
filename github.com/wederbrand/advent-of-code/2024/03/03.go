package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"regexp"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2024/03/input.txt", "\n")

	p1, p2 := doIt(strings.Join(inFile, ""))

	fmt.Println("Part 1:", p1, "in", time.Since(start))
	fmt.Println("Part 2:", p2, "in", time.Since(start))
}

func doIt(inFile string) (int, int) {
	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	matches := re.FindAllString(inFile, -1)

	p1 := 0
	p2 := 0
	enabled := true
	for _, match := range matches {
		if match == "do()" {
			enabled = true
			continue
		}
		if match == "don't()" {
			enabled = false
			continue
		}
		var a, b int
		fmt.Sscanf(match, "mul(%d,%d)", &a, &b)

		p1 += a * b
		if enabled {
			p2 += a * b
		}
	}

	return p1, p2
}
