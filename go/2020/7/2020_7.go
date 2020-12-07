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
	readFile, err := ioutil.ReadFile("7/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	rulesInput := strings.Split(strings.TrimSpace(string(readFile)), "\n")
	ruleMatcher := regexp.MustCompile("(.*) bags contain (.*)")
	bagMatcher := regexp.MustCompile("(\\d+) (.+?) bag")

	rules := make(map[string]map[string]int)
	for _, rule := range rulesInput {
		submatch := ruleMatcher.FindStringSubmatch(rule)
		color := submatch[1]
		rules[color] = make(map[string]int)
		colorRules := bagMatcher.FindAllStringSubmatch(submatch[2], -1)
		for _, colorRule := range colorRules {
			rules[color][colorRule[2]], _ = strconv.Atoi(colorRule[1])
		}
	}

	count := 0
	for color, _ := range rules {
		if color == "shiny gold" {
			continue
		}
		if leadsToGold(color, rules) {
			count++
		}
	}

	fmt.Println("part 1: ", count)

	totalBags := counter("shiny gold", rules) - 1
	fmt.Println("part 2: ", totalBags)
}

func counter(color string, rules map[string]map[string]int) int {
	total := 1 // my self

	for innerColor, count := range rules[color] {
		total += count * counter(innerColor, rules)
	}

	return total
}

func leadsToGold(color string, rules map[string]map[string]int) bool {
	if color == "shiny gold" {
		return true
	}

	for s, _ := range rules[color] {
		if leadsToGold(s, rules) {
			return true
		}
	}

	return false
}
