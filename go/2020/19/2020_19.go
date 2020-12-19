package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	readFile, err := ioutil.ReadFile("19/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.TrimSpace(string(readFile)), "\n\n")
	inputRules := strings.Split(strings.TrimSpace(input[0]), "\n")
	inputImages := strings.Split(strings.TrimSpace(input[1]), "\n")

	// process input until rules[0] exists
	rules := make(map[int][]string)
	for len(rules[0]) == 0 {
	INPUT_LOOP:
		for _, s := range inputRules {
			split := strings.Split(s, " ")
			ruleStr := strings.TrimRight(split[0], ":")
			i, _ := strconv.Atoi(ruleStr)

			if len(rules[i]) > 0 {
				continue
			}

			if split[1] == "\"a\"" {
				rules[i] = append(rules[i], "a")
				continue
			}
			if split[1] == "\"b\"" {
				rules[i] = append(rules[i], "b")
				continue
			}

			newRules := make([]string, 0)
			current := make([]string, 1)
			for _, anotherRule := range split[1:] {
				if anotherRule == "|" {
					newRules = append(newRules, current...)
					current = make([]string, 1)
					continue
				}
				atoi, _ := strconv.Atoi(anotherRule)
				if len(rules[atoi]) == 0 {
					// rule doesn't exist yet, break and come back later
					continue INPUT_LOOP
				}

				newCurrent := make([]string, 0)
				for _, rule := range rules[atoi] {
					for _, oldCurrent := range current {
						newCurrent = append(newCurrent, oldCurrent+rule)
					}
				}
				current = newCurrent
			}
			newRules = append(newRules, current...)
			rules[i] = newRules
		}
	}

	// now let's check the rest of the input

	cnt := 0
	for _, check := range inputImages {
		for _, rule := range rules[0] {
			if rule == check {
				cnt++
				break
			}
		}
	}
	fmt.Println(cnt)
}
