package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("2021/2021_2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	matcher := regexp.MustCompile("(.+) (\\d+)")
	scanner := bufio.NewScanner(file)

	h := 0
	d := 0
	a := 0
	for scanner.Scan() {
		text := scanner.Text()
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
