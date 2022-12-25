package main

import (
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	readFile, err := os.ReadFile("2015/07/2015_07.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")
	providerRE := regexp.MustCompile("(\\d+) -> (.+)")
	andRE := regexp.MustCompile("(.+) AND (.+) -> (.+)")
	orRE := regexp.MustCompile("(.+) OR (.+) -> (.+)")
	leftShiftRE := regexp.MustCompile("(.+) LSHIFT (.+) -> (.+)")
	rightShiftRE := regexp.MustCompile("(.+) RSHIFT (.+) -> (.+)")
	notRE := regexp.MustCompile("NOT (.+) -> (.+)")
	for _, instruction := range inFile {
		if providerRE.MatchString(instruction) {
			
		}

	}
}
