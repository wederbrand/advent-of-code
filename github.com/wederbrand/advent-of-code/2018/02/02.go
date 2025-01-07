package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"slices"
	"time"
)

type Box []rune

func (b Box) count() (int, int) {
	testValue := slices.Clone(b)
	slices.Sort(testValue)

	var lastSeen rune
	count := 1
	two := 0
	three := 0
	for _, c := range testValue {
		if c == lastSeen {
			count++
		} else {
			lastSeen = c
			if count == 2 {
				two++
			} else if count == 3 {
				three++
			}
			count = 1
		}
	}
	if count == 2 {
		two++
	} else if count == 3 {
		three++
	}

	return min(two, 1), min(three, 1)
}

func main() {
	start := time.Now()
	inFile := GetFileContents("2018/02/input.txt", "\n")

	input := make([]Box, 0)
	for _, s := range inFile {
		box := Box(s)
		input = append(input, box)
	}

	two := 0
	three := 0
	for _, box := range input {
		a, b := box.count()
		two += a
		three += b
	}

	fmt.Println("Part 1: ", two*three, "in", time.Since(start))

	for indexA, boxA := range input {
		for _, boxB := range input[indexA:] {
			diff := 0
			for i := 0; i < len(boxA) && diff <= 1; i++ {
				if boxA[i] != boxB[i] {
					diff++
				}
			}
			if diff == 1 {
				common := ""
				for i := 0; i < len(boxA); i++ {
					if boxA[i] == boxB[i] {
						common += string(boxA[i])
					}
				}
				fmt.Println("Part 2: ", common, "in", time.Since(start))
			}
		}
	}
}
