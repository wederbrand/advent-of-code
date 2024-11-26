package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2016/04/input.txt", "\n")
	part1 := part1(inFile)
	fmt.Println("Part 1: ", part1, "in", time.Since(start))
	fmt.Println("Part 2 found manually to be \"northpole object storage\" 984")
}

func part1(inFile []string) int {
	sum := 0
	for _, line := range inFile {
		split := strings.Split(line, "-")
		words := split[0 : len(split)-1]
		rest := split[len(split)-1]
		value := util.Atoi(strings.Split(rest, "[")[0])
		checksum := rest[len(rest)-6 : len(rest)-1]
		calculatedSum := getCheckSum(words)
		if checksum == calculatedSum {
			sum += value
			decryptAndPrint(words, value)
		}
	}

	return sum
}

func decryptAndPrint(words []string, value int) {
	for _, letters := range words {
		for _, l := range letters {
			if l == '-' {
				fmt.Print(" ")
			} else {
				shifted := rune((int(l)-'a'+value)%('z'-'a'+1) + 'a')
				fmt.Printf("%c", shifted)
			}
		}
		fmt.Print(" ")
	}
	fmt.Println(value)
}

func getCheckSum(words []string) string {
	result := ""

	counts := make(map[string]int)
	for _, letters := range words {
		for _, l := range letters {
			counts[string(l)]++
		}
	}

	for len(result) < 5 {
		// find the first occurrence of the max count
		maxFound := 0
		maxLetter := ""
		for letter, count := range counts {
			if count > maxFound || (count == maxFound && letter < maxLetter) {
				maxFound = count
				maxLetter = letter
			}
		}
		result += maxLetter
		counts[maxLetter] = 0 // hide the one we just used
	}

	return result
}
