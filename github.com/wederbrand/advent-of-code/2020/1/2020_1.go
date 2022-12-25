package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
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

	for i := 0; i < len(input)-2; i++ {
		for j := i + 1; j < len(input)-1; j++ {
			for k := j + 1; k < len(input); k++ {
				if input[i]+input[j]+input[k] == 2020 {
					fmt.Println(input[i] * input[j] * input[k])
					os.Exit(0)
				}
			}
		}
	}

	os.Exit(1)
}
