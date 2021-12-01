package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("2021/2021_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := make([]int, 0)
	for scanner.Scan() {
		sValue := scanner.Text()
		iValue, _ := strconv.Atoi(sValue)
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
