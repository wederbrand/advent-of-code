package main

import (
	"fmt"
	"log"
	"os"
	"strings"
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
	readFile, err := os.ReadFile("25/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fromSnafu("1=0")
	toSnafu(15)
	toSnafu(10)
	toSnafu(20)
	toSnafu(12345)

	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")
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
