package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strings"
	"unicode"
)

func main() {
	points := make(map[rune]int)
	p := 1
	for r := 'a'; r <= 'z'; r++ {
		points[r] = p
		points[unicode.ToUpper(r)] = p + 26
		p++
	}

	inFile := util.GetFileContents("2022/03/input.txt", "\n")

	part1(inFile, points)
	part2(inFile, points)
}

func part2(inFile []string, points map[rune]int) {
	score := 0
	for i := 0; i < len(inFile); i += 3 {
		s1 := inFile[i]
		s2 := inFile[i+1]
		s3 := inFile[i+2]

		for _, r := range s1 {
			if strings.ContainsRune(s2, r) && strings.ContainsRune(s3, r) {
				score += points[r]
				break
			}
		}
	}

	fmt.Println(score)
}
func part1(inFile []string, points map[rune]int) {
	score := 0
	for _, s := range inFile {
		s1 := s[0 : len(s)/2]
		s2 := s[len(s)/2:]

		for _, r1 := range s1 {
			if strings.ContainsRune(s2, r1) {
				score += points[r1]
				break
			}
		}
	}

	fmt.Println(score)
}
