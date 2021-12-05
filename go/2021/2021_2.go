package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.ReadFile("2021/2021_2.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	matcher := regexp.MustCompile("(.+) (\\d+)")

	h := 0
	d := 0
	a := 0
	for _, text := range inFile {
		submatch := matcher.FindStringSubmatch(text)
		command := submatch[1]
		value, _ := strconv.Atoi(submatch[2])

		if command == "forward" {
			h += value
			d += a * value
		} else if command == "down" {
			a += value
		} else if command == "up" {
			a -= value
		}
	}

	fmt.Println(h * d)

	os.Exit(1)
}
