package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	in := "hepxcrrq"
	runes := []rune(in)
	nextValid(runes)

	fmt.Println("part1", string(runes), "in", time.Since(start))

	nextValid(runes)
	fmt.Println("part2", string(runes), "in", time.Since(start))
}

func nextValid(in []rune) {
	for {
		increment(in)
		warpOutOfIllegal(in)

		if !threeFound(in) {
			continue
		}

		if !twoPairsFound(in) {
			continue
		}

		return
	}
}

func increment(in []rune) {
	for i := 7; i >= 0; i-- {
		if in[i]+1 <= 'z' {
			in[i]++
			return
		} else {
			in[i] = 'a'
		}
	}
}

func warpOutOfIllegal(in []rune) {
	// Passwords may not contain the letters i, o, or l, as these letters can be mistaken for other characters and are therefore confusing.
	nextAllAce := false
	for i := 0; i < 8; i++ {
		if nextAllAce {
			in[i] = 'a'
		} else if in[i] == 'i' || in[i] == 'o' || in[i] == 'l' {
			in[i]++
			nextAllAce = true
		}
	}
}

func threeFound(in []rune) bool {
	// Passwords must include one increasing straight of at least three letters, like abc, bcd, cde, and so on, up to xyz. They cannot skip letters; abd doesn't count.
	for i := 0; i < 6; i++ {
		if in[i]+1 == in[i+1] && in[i]+2 == in[i+2] {
			return true
		}
	}
	return false
}

func twoPairsFound(in []rune) bool {
	// Passwords must contain at least two different, non-overlapping pairs of letters, like aa, bb, or zz.
	pairs := make(map[rune]bool)
	for i := 0; i < 7; i++ {
		if in[i] == in[i+1] {
			pairs[in[i]] = true
		}
	}
	return len(pairs) >= 2
}
