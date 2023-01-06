package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2015/08/input.txt", "\n")

	totalCode := 0
	totalString := 0
	totalEncoded := 0
	for _, code := range inFile {
		decoded := decode(code)
		encoded := strconv.QuoteToASCII(code)
		totalCode += len(code)
		totalString += len(decoded)
		totalEncoded += len(encoded)
	}

	fmt.Println("part1:", totalCode-totalString, "in", time.Since(start))
	fmt.Println("part2:", totalEncoded-totalCode, "in", time.Since(start))
}

func decode(code string) []rune {
	result := make([]rune, 0)
	runes := []rune(code)

	for i := 0; i < len(runes); i++ {
		if i == 0 || i == len(runes)-1 {
			continue
		}
		r := runes[i]
		if r == '\\' {
			r2 := runes[i+1]
			if r2 == '"' {
				i++
				r = r2
			} else if r2 == '\\' {
				i++
				r = r2
			} else if r2 == 'x' {
				ascii, _ := strconv.ParseInt(string(runes[i+2:i+4]), 16, 64)
				r = rune(ascii)
				i += 3
			}
		}

		result = append(result, r)
	}

	return result
}
