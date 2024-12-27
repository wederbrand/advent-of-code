package main

import (
	"fmt"
	"slices"
	"time"
)

func main() {
	start := time.Now()
	input := "00101000101111010"

	part1 := doIt(input, 272)
	fmt.Println("Part 1: ", part1, "in", time.Since(start))

	part2 := doIt(input, 35651584)
	fmt.Println("Part 2: ", part2, "in", time.Since(start))
}

func doIt(input string, length int) string {
	a := []rune(input)

	for len(a) < length {
		b := slices.Clone(a)
		b = append(b, '0')
		slices.Reverse(a)
		for i, r := range a {
			if r == '0' {
				a[i] = '1'
			} else {
				a[i] = '0'
			}
		}
		b = append(b, a...)
		a = b
	}

	checksum := a[:length]
	for len(checksum)%2 == 0 {
		newChecksum := make([]rune, 0)
		for i := 0; i < len(checksum); i += 2 {
			if checksum[i] == checksum[i+1] {
				newChecksum = append(newChecksum, '1')
			} else {
				newChecksum = append(newChecksum, '0')
			}
		}
		checksum = newChecksum
	}
	return string(checksum)
}
