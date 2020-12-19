package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	readFile, err := ioutil.ReadFile("18/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.TrimSpace(string(readFile)), "\n")
	sum := 0
	for _, s := range input {
		sum += solveOne(s)
	}

	fmt.Println(sum)
}

func solveOne(input string) int {
	result := 0

	// recursively replace all expressions in parentheses with their calculated value
	pFinder := regexp.MustCompile("\\(([0-9 +*])+\\)")
	for pFinder.MatchString(input) {
		input = pFinder.ReplaceAllStringFunc(input, func(s string) string {
			innerValue := solveOne(s[1 : len(s)-1])
			itoa := strconv.Itoa(innerValue)
			return itoa
		})
	}

	// recursively replace all expressions surrounding + with their calculated value
	pFinder = regexp.MustCompile("(\\d+) \\+ (\\d+)")
	for pFinder.MatchString(input) {
		input = pFinder.ReplaceAllStringFunc(input, func(s string) string {
			allString := pFinder.FindStringSubmatch(s)
			v1, _ := strconv.Atoi(allString[1])
			v2, _ := strconv.Atoi(allString[2])
			value := v1 + v2
			itoa := strconv.Itoa(value)
			return itoa
		})
	}

	// all remaining is * but the logic is unchanged from part 1
	split := strings.Split(input, " ")
	nextIsAddition := true
	for _, s := range split {
		if s == "+" {
			nextIsAddition = true
		} else if s == "*" {
			nextIsAddition = false
		} else {
			atoi, _ := strconv.Atoi(s)
			if nextIsAddition {
				result += atoi
			} else {
				result *= atoi
			}
		}
	}

	return result
}
