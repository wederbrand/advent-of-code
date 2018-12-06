package main

import (
	"unicode"
	"strings"
	"fmt"
	"os"
	"bufio"
)

type polymer []rune

func max(a int, b int) int {
	if (a > b) {
		return a
	} else {
		return b
	}
}

func caseDiffers(a, b rune) bool {
	if a - b == 32 {
		return true
	}
	if a - b == -32 {
		return true
	}
	return false
}

func (p polymer) reduce() polymer {
	for i := 0; i+1 < len(p) ; i++ {
		if caseDiffers(p[i], p[i+1]) {
			p = append(p[:i], p[i+2:]...)
			i = max(-1, i-2)
		}
	}

	return p
}

func main () {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	input := scanner.Text()
	fmt.Println("part 1", len(polymer(input).reduce()))

	min := len(input)
	for i:='a';i<='z';i++ {
		test := input
		test = strings.Replace(string(test), string(i), "", -1)
		test = strings.Replace(string(test), string(unicode.ToUpper(i)), "", -1)
		strippedSize := len(polymer(test).reduce())
		if (strippedSize < min) {
			min = strippedSize
		}
	}

	fmt.Println("part 2", min)

}