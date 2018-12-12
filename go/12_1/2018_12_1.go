package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"unicode/utf8"
)

var iterations = make([][]rune, 0)
var mapper = make(map[string]rune)
var runs = 200
var offset = runs * 2

func main() {
	file, err := os.Open("12_1/2018_12.input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	initialMatcher := regexp.MustCompile(`^initial state: (.+)$`)
	lineMatcher := regexp.MustCompile(`^(.+) => (.)$`)

	scanner.Scan()
	initialMatch := initialMatcher.FindStringSubmatch(scanner.Text())
	inputPots := initialMatch[1]
	pots := make([]rune, len(inputPots)+runs*4)
	for i := -runs * 2; i < len(inputPots)+runs*2; i++ {
		pots[i+offset] = '.'
	}

	for i, r := range inputPots {
		pots[i+offset] = r
	}

	iterations = append(iterations, pots)

	// skip one
	scanner.Scan()

	for scanner.Scan() {
		lineMatch := lineMatcher.FindStringSubmatch(scanner.Text())
		r, _ := utf8.DecodeRuneInString(lineMatch[2])
		mapper[lineMatch[1]] = r
	}

	lastSum := 0
	for i := 0; i < runs; i++ {
		pots := make([]rune, len(inputPots)+runs*4)
		previous := iterations[i]
		copy(pots, previous)
		start := math.MaxInt64
		stop := math.MinInt64
		for key, value := range previous {
			if value == '#' && key < start {
				start = key
			}
			if value == '#' && key > stop {
				stop = key
			}
		}
		sum := 0
		for j := start - 2; j < stop+2; j++ {
			runes := previous[j-2 : j+3]
			key := string(runes)
			r, ok := mapper[key]
			if ok {
				pots[j] = r
				if r == '#' {
					sum += j - offset
				}
			} else {
				pots[j] = '.'
			}
		}
		iterations = append(iterations, pots)
		fmt.Println(i, sum, sum-lastSum, string(pots))
		lastSum = sum
	}
	
	fmt.Println(lastSum)

	// not 8550000001881
	// not 8550000002052
}
