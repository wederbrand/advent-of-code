package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"log"
	"sort"
)

func main() {
	inFile := util.GetFileContents("2022/06/input.txt", "\n")

	inData := inFile[0]
	fmt.Println(checkString(inData, 4))
	fmt.Println(checkString(inData, 14))
}

func checkString(inData string, length int) int {
	for i := range inData {
		if i < length-1 {
			continue
		}
		if different(inData[i-length+1 : i+1]) {
			return i + 1
		}
	}
	log.Fatal("didn't find key")
	return 0
}

func different(in string) bool {
	runes := []rune(in)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	for i := 0; i < len(runes)-1; i++ {
		if runes[i] == runes[i+1] {
			return false
		}
	}
	return true
}
