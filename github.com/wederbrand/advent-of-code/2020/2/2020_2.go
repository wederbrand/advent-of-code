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
	file, err := os.Open("2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	matcher := regexp.MustCompile("(\\d+)-(\\d+) (.+): (.+)")

	var valid int
	var invalid int

	for scanner.Scan() {
		text := scanner.Text()
		submatch := matcher.FindStringSubmatch(text)
		min, _ := strconv.Atoi(submatch[1])
		max, _ := strconv.Atoi(submatch[2])
		rule := submatch[3]
		password := submatch[4]

		var count int
		a := password[min-1]
		b := password[max-1]
		if a == rule[0] {
			count++
		}
		if b == rule[0] {
			count++
		}
		if count == 1 {
			valid++
		} else {
			invalid++
		}
	}
	fmt.Println("valid ", valid)
}
