package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2019/16/input.txt", "\n")

	input := make([]int, 0)
	for _, r := range inFile[0] {
		input = append(input, int(r-'0'))
	}

	for i := 0; i < 100; i++ {
		input = doIt(input)
	}

	part1 := ""
	for i := 0; i < 8; i++ {
		part1 += strconv.Itoa(input[i])
	}
	fmt.Println("part1: ", part1, "in", time.Since(start))

	input = make([]int, 0)
	for i := 0; i < 10000; i++ {
		for _, r := range inFile[0] {
			input = append(input, int(r-'0'))
		}
	}

	offset := Atoi(inFile[0][:7])
	input = input[offset:]
	for i := 0; i < 100; i++ {
		input = cheatIt(input)
	}

	part2 := ""
	for i := 0; i < 8; i++ {
		part2 += strconv.Itoa(input[i])
	}
	fmt.Println("part2: ", part2, "in", time.Since(start))
}

func cheatIt(input []int) []int {
	output := make([]int, len(input))
	sum := 0
	for i := len(input) - 1; i >= 0; i-- {
		sum += input[i]
		output[i] = sum % 10
	}
	return output
}

func doIt(input []int) []int {
	pattern := [4]int{0, 1, 0, -1}
	output := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		sum := 0
		for j := 0; j < len(input); j++ {
			patternIndex := (j + 1) / (i + 1)
			patternIndex %= 4
			sum += input[j] * pattern[patternIndex]
		}
		output[i] = IntAbs(sum) % 10
	}
	return output
}
