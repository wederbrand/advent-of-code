package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
)

func main() {
	inFile := util.GetFileContents("2015/05/2015_05.txt", "\n")
	count := 0
	for _, word := range inFile {
		if testIt(word) {
			count++
		}

	}

	fmt.Println("part 2", count)
}

func testIt(word string) bool {
	pair := false
	repeat := false
	pairs := make(map[string]int)
	for i := 0; i < len(word)-1; i++ {
		i2, found := pairs[word[i:i+2]]
		if found {
			if i2 <= i-2 {
				pair = true
			}
		} else {
			pairs[word[i:i+2]] = i
		}

		if i < len(word)-2 && word[i] == word[i+2] {
			repeat = true
		}
	}
	if pair && repeat {
		return true
	}
	return false
}
