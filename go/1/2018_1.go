package main

import (
	"strconv"
	"fmt"
	"os"
	"bufio"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	input := make([]int, 0)
	for scanner.Scan() {
		sValue := scanner.Text();
		iValue, _ := strconv.Atoi(sValue)
		input = append(input, iValue)
	}

	fmt.Println(input)

	seen := make(map[int]struct{})
	freq := 0
	Loop:
		for {
			for _, iValue := range input {
				freq += iValue
				_, ok := seen[freq]
				if (ok) {
					break Loop
				}
				seen[freq] = struct{}{}
			}
		}

	fmt.Println(freq)
}