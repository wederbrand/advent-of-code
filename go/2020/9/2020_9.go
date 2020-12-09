package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	readFile, err := ioutil.ReadFile("9/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.TrimSpace(string(readFile)), "\n")
	numbers := make([]int, 0)
	for _, s := range input {
		atoi, _ := strconv.Atoi(s)
		numbers = append(numbers, atoi)
	}
	preamble := 25

	i := preamble

	weakness := 0
	for {
		if !check(i, numbers, preamble) {
			weakness = numbers[i]
			fmt.Println("part 1:", numbers[i])
			break
		}
		i++
	}

	i = 0
	for {
		a, b, err := findSet(i, numbers, preamble, weakness)
		if err == nil {
			fmt.Println("part 2:", a+b)
			break
		}
		i++
	}

}

func findSet(i int, numbers []int, preamble int, weakness int) (int, int, error) {
	sum := 0
	min := math.MaxInt64
	max := math.MinInt64

	for {
		sum += numbers[i]
		if numbers[i] < min {
			min = numbers[i]
		}
		if numbers[i] > max {
			max = numbers[i]
		}
		if sum == weakness {
			return min, max, nil
		}
		if sum > weakness {
			return 0, 0, errors.New("too big")
		}
		i++
	}
}

func check(i int, numbers []int, preamble int) bool {
	for a := i - preamble; a < i-1; a++ {
		for b := a + 1; b < i; b++ {
			candidate := numbers[i]
			theFirst := numbers[a]
			theSecond := numbers[b]
			if theFirst+theSecond == candidate {
				return true
			}
		}
	}

	return false
}
