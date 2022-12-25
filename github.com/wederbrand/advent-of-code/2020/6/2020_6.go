package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	readFile, err := ioutil.ReadFile("6/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	allGroups := strings.Split(strings.TrimSpace(string(readFile)), "\n\n")

	total := 0
	for _, group := range allGroups {
		answers := make(map[rune]int, 0)
		allPeople := strings.Split(group, "\n")
		for _, person := range allPeople {
			for _, questionWithYes := range person {
				answers[questionWithYes]++
			}
		}

		for _, i := range answers {
			if i == len(allPeople) {
				total++
			}
		}
	}

	fmt.Println(total)

}
