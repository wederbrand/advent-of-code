package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strings"
	"time"
	"unicode"
)

type polymer []rune

func maxInt(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func caseDiffers(a, b rune) bool {
	if a-b == 32 {
		return true
	}
	if a-b == -32 {
		return true
	}
	return false
}

func (p polymer) reduce() polymer {
	for i := 0; i+1 < len(p); i++ {
		if caseDiffers(p[i], p[i+1]) {
			p = append(p[:i], p[i+2:]...)
			i = maxInt(-1, i-2)
		}
	}

	return p
}

func main() {
	start := time.Now()
	inFile := GetFileContents("2018/05/input.txt", "\n")
	input := inFile[0]

	fmt.Println("Part 1:", len(polymer(input).reduce()), "in", time.Since(start))

	lowest := len(input)
	for i := 'a'; i <= 'z'; i++ {
		test := input
		test = strings.Replace(test, string(i), "", -1)
		test = strings.Replace(test, string(unicode.ToUpper(i)), "", -1)
		strippedSize := len(polymer(test).reduce())
		if strippedSize < lowest {
			lowest = strippedSize
		}
	}

	fmt.Println("Part 2:", lowest, "in", time.Since(start))
}
