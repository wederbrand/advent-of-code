package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	readFile, err := os.ReadFile("2021/2021_10.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	corruptScore := 0
	incompleteScores := make([]int, 0)
	for _, text := range inFile {
		i, j := process(text)
		if i > 0 {
			corruptScore += i
		} else {
			incompleteScores = append(incompleteScores, j)
		}
	}

	fmt.Println("part 1", corruptScore)

	sort.Ints(incompleteScores)
	fmt.Println("part 2", incompleteScores[len(incompleteScores)/2])
	os.Exit(0)
}

func process(text string) (int, int) {
	changed := true
	for changed {
		newText := text
		newText = strings.ReplaceAll(newText, "()", "")
		newText = strings.ReplaceAll(newText, "[]", "")
		newText = strings.ReplaceAll(newText, "<>", "")
		newText = strings.ReplaceAll(newText, "{}", "")

		if len(newText) != len(text) {
			changed = true
		} else {
			changed = false
		}

		text = newText
	}

	if strings.ContainsAny(text, ")]>}") {
		// corrupt
		for _, r := range text {
			if corruptValue(r) > 0 {
				return corruptValue(r), 0
			}
		}
	} else {
		// incomplete
		score := 0
		for _, r := range Reverse(text) {
			score *= 5
			score += incompleteValue(r)
		}
		return 0, score
	}
	return 0, 0
}

func corruptValue(r int32) int {
	switch r {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	}
	return 0
}

func incompleteValue(r int32) int {
	switch r {
	case '(':
		return 1
	case '[':
		return 2
	case '{':
		return 3
	case '<':
		return 4
	}
	return 0
}
