package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.ReadFile("2021/2021_1.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	input := make([]int, 0)
	for _, text := range inFile {
		iValue, _ := strconv.Atoi(text)
		input = append(input, iValue)
	}

	increases := 0
	for i := 3; i < len(input); i++ {
		if input[i] > input[i-3] {
			increases++
		}
	}

	fmt.Println(increases)
	os.Exit(1)
}
