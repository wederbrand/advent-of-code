package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	result := "3113322113"
	for i := 0; i < 50; i++ {
		result = lookAndSay(result)
		fmt.Println("part2", i, " ", len(result), "in", time.Since(start))
		start = time.Now()
	}

	fmt.Println("part2", len(result), "in", time.Since(start))
}

func lookAndSay(in string) (out string) {
	currentNbr := 0
	currentRune := ' '
	for _, r := range []rune(in) {
		if r == currentRune {
			currentNbr++
		} else {
			if currentNbr > 0 {
				out += fmt.Sprintf("%d%d", currentNbr, currentRune-'0')
			}
			currentNbr = 1
			currentRune = r
		}
	}
	if currentNbr > 0 {
		out += fmt.Sprintf("%d%d", currentNbr, currentRune-'0')
	}

	return
}
