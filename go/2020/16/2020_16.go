package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type rule struct {
	name string
	min1 int
	max1 int
	min2 int
	max2 int
}

func main() {
	readFile, err := ioutil.ReadFile("16/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	// rules
	rules := make([]rule, 0)
	ruleRE := regexp.MustCompile("(.+): (.+)-(.+) or (.+)-(.+)")
	i := 0
	for ; i < len(input); i++ {
		s := input[i]

		if s == "" {
			i += 2
			break
		}

		submatch := ruleRE.FindStringSubmatch(s)
		min1, _ := strconv.Atoi(submatch[2])
		max1, _ := strconv.Atoi(submatch[3])
		min2, _ := strconv.Atoi(submatch[4])
		max2, _ := strconv.Atoi(submatch[5])

		r := rule{submatch[1], min1, max1, min2, max2}
		rules = append(rules, r)
	}

	// my ticket
	var myTicket []int
	s := input[i]
	split := strings.Split(s, ",")
	for _, val := range split {
		atoi, _ := strconv.Atoi(val)
		myTicket = append(myTicket, atoi)
	}
	i += 3

	// other tickets
	var invalid []int
	for ; i < len(input); i++ {
		s := input[i]
		split := strings.Split(s, ",")
	OUTER:
		for _, val := range split {
			atoi, _ := strconv.Atoi(val)
			// check val against all valid rules
			for _, r := range rules {
				if atoi >= r.min1 && atoi <= r.max1 || atoi >= r.min2 && atoi <= r.max2 {
					// ok
					continue OUTER
				}
			}
			invalid = append(invalid, atoi)
		}
	}

	sum := 0
	for _, i := range invalid {
		sum += i
	}

	fmt.Println(sum)
}
