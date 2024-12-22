package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"math"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2024/22/input.txt", "\n")

	part1 := 0
	all := make(map[string]int)
	for _, line := range inFile {
		secret := Atoi(line)
		d0, d1, d2, d3 := 0, 0, 0, 0
		seen := make(map[DiffKey]bool)
		last := 0
		for j := 0; j < 2000; j++ {
			secret = doSecret(secret)
			digit := secret % 10
			d0, d1, d2, d3 = d1, d2, d3, digit-last
			if j >= 3 {
				lastFourString := fmt.Sprintf("%d,%d,%d,%d", d1-d0, d2-d1, d3-d2, d4-d3)
				if !seen[lastFourString] {
					all[lastFourString] += digit
					seen[lastFourString] = true
				}
			}
		}
		part1 += secret
	}
	fmt.Println("Part 1:", part1, "in", time.Since(start))

	part2 := math.MinInt
	for _, value := range all {
		if value > part2 {
			part2 = value
		}
	}
	fmt.Println("Part 2:", part2, "in", time.Since(start))
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
