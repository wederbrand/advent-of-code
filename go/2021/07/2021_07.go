package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.ReadFile("2021/07/2021_07.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), ",")

	crabs := make([]int, len(inFile))
	min := math.MaxInt
	max := math.MinInt
	for i, text := range inFile {
		atoi, _ := strconv.Atoi(text)
		crabs[i] = atoi
		if atoi < min {
			min = atoi
		}
		if atoi > max {
			max = atoi
		}
	}

	minPrice := math.MaxInt
	for i := min; i < max; i++ {
		price := 0
		for _, crab := range crabs {
			distToMove := crab - i
			if distToMove < 0 {
				distToMove *= -1
			}

			crabPrice := 0
			for j := 1; j <= distToMove; j++ {
				crabPrice += j
			}
			price += crabPrice
		}

		if price < minPrice {
			minPrice = price
		}
	}

	fmt.Println(minPrice)

	os.Exit(0)
}
