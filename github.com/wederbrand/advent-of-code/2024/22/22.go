package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"golang.org/x/exp/maps"
	"time"
	// . "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2024/22/input.txt", "\n")

	part1 := 0
	all := make([]map[string]int, len(inFile))
	for i, line := range inFile {
		all[i] = make(map[string]int)
		secret := Atoi(line)
		digit := secret % 10
		d0, d1, d2, d3, d4 := 0, 0, 0, 0, digit
		for j := 0; j < 2000; j++ {
			secret = doSecret(secret)
			digit = secret % 10
			d0, d1, d2, d3, d4 = d1, d2, d3, d4, digit
			if j >= 3 {
				lastFourString := fmt.Sprintf("%d,%d,%d,%d", d1-d0, d2-d1, d3-d2, d4-d3)

				_, found := all[i][lastFourString]
				if !found {
					all[i][lastFourString] = digit
				}
			}
		}
		part1 += secret
	}
	fmt.Println("Part 1:", part1, "in", time.Since(start))

	part2 := doPart2(all)
	fmt.Println("Part 2:", part2, "in", time.Since(start))
}

func doPart2(all []map[string]int) int {
	allSeq := make(map[string]bool)
	for _, m := range all {
		keys := maps.Keys(m)
		for _, k := range keys {
			allSeq[k] = true
		}
	}

	part2 := 0
	for _, key := range maps.Keys(allSeq) {
		sum := 0
		for _, m := range all {
			v := m[key]
			sum += v
		}
		if sum > part2 {
			part2 = sum
		}
	}

	return part2
}

func doSecret(secret int) int {
	// Calculate the result of multiplying the secret number by 64. Then, mix this result into the secret number. Finally, prune the secret number.
	secret ^= secret * 64
	secret %= 16777216

	// Calculate the result of dividing the secret number by 32. Round the result down to the nearest integer. Then, mix this result into the secret number. Finally, prune the secret number.
	secret ^= secret / 32
	secret %= 16777216

	// Calculate the result of multiplying the secret number by 2048. Then, mix this result into the secret number. Finally, prune the secret number.
	secret ^= secret * 2048
	secret %= 16777216

	return secret
}
