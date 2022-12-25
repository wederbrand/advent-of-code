package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.ReadFile("2021/03/2021_03.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	current := inFile
	index := 0
	for len(current) > 1 {
		count := count(current, '1', index)
		if 2*count >= len(current) {
			current = filterOn(current, '1', index)
		} else {
			current = filterOn(current, '0', index)
		}
		index++
	}

	oxygen, _ := strconv.ParseInt(current[0], 2, 64)

	current = inFile
	index = 0
	for len(current) > 1 {
		count := count(current, '0', index)
		if 2*count > len(current) {
			current = filterOn(current, '1', index)
		} else {
			current = filterOn(current, '0', index)
		}
		index++
	}

	co2, _ := strconv.ParseInt(current[0], 2, 64)

	fmt.Println(oxygen * co2)

	os.Exit(1)
}

func count(input []string, value uint8, i int) int {
	count := 0
	for _, row := range input {
		if row[i] == value {
			count++
		}
	}
	return count
}

func filterOn(input []string, value uint8, i int) []string {
	output := make([]string, 0)
	for _, row := range input {
		if row[i] == value {
			output = append(output, row)
		}
	}

	return output
}
