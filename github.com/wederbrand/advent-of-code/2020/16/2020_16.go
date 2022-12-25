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
	name     string
	min1     int
	max1     int
	min2     int
	max2     int
	position int
}

func (r rule) valid(atoi int) bool {
	return atoi >= r.min1 && atoi <= r.max1 || atoi >= r.min2 && atoi <= r.max2
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

		r := rule{submatch[1], min1, max1, min2, max2, 0}
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
	var valids []string
SUPEROUTER:
	for ; i < len(input); i++ {
		s := input[i]
		split := strings.Split(s, ",")
	OUTER:
		for _, val := range split {
			atoi, _ := strconv.Atoi(val)
			// check val against all valids rules
			for _, r := range rules {
				if r.valid(atoi) {
					// ok
					continue OUTER
				}
			}
			invalid = append(invalid, atoi)
			continue SUPEROUTER
		}
		valids = append(valids, s)
	}

	// analyze all the valid ones
	product := 1
	completedFields := make([]int, 0)
OHNO:
	for fieldIndex := 0; fieldIndex < len(myTicket); fieldIndex++ {
		for _, field := range completedFields {
			if field == fieldIndex {
				fmt.Println("already found")
				continue OHNO
			}
		}

		candidates := append([]rule{}, rules...)
		for _, valid := range valids {
			value := strings.Split(valid, ",")[fieldIndex]
			atoi, _ := strconv.Atoi(value)
			nextCandidates := candidates[:0]
			for _, candidate := range candidates {
				if candidate.valid(atoi) {
					nextCandidates = append(nextCandidates, candidate)
				}

			}
			candidates = nextCandidates
		}
		// should be exactly one candidate
		if len(candidates) > 1 {
			fmt.Println("restarting")
			continue OHNO
		}

		completedFields = append(completedFields, fieldIndex)

		newRules := rules[:0]
		for _, r := range rules {
			if r != candidates[0] {
				newRules = append(newRules, r)
			}
		}
		rules = newRules

		if strings.HasPrefix(candidates[0].name, "departure") {
			product *= myTicket[fieldIndex]
		}
		fieldIndex = -1
	}

	fmt.Println(product)
}
