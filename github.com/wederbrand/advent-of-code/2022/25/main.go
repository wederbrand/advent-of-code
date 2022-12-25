package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"time"
)

var stoi = map[string]int{
	"=": -2,
	"-": -1,
	"0": 0,
	"1": 1,
	"2": 2,
}

var itos = map[int]string{
	-2: "=",
	-1: "-",
	0:  "0",
	1:  "1",
	2:  "2",
}

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2022/25/input.txt", "\n")

	sum := 0
	for _, i := range inFile {
		sum += fromSnafu(i)
	}

	fmt.Println("part1:", toSnafu(sum), "in", time.Since(start))
}

func toSnafu(in int) string {
	result := ""
	for in > 0 {
		modulo := in % 5
		if modulo > 2 {
			modulo -= 5
		}
		s := itos[modulo]
		result = s + result
		in -= modulo
		in /= 5
	}
	return result
}

func fromSnafu(in string) int {
	result := 0

	for _, i := range in {
		result *= 5
		result += stoi[string(i)]
	}

	return result
}
