package main

import (
	"sort"
	"fmt"
	"strconv"
	"regexp"
	"os"
	"bufio"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lineMatcher := regexp.MustCompile(`^\[....-..-.. ..:(\d+)\] (.*)$`)

	lines := make(map[int]string, 0)
	minutes := make([]int, 0)
	for scanner.Scan() {
		sValue := scanner.Text()
		lineMatch := lineMatcher.FindStringSubmatch(sValue)
		minute, _ := strconv.Atoi(lineMatch[1])
		lines[minute] = lineMatch[2]
		minutes = append(minutes, minute)
	}

	sort.Ints(minutes)

	guardMatcher := regexp.MustCompile(`^Guard #(\d+) begins shift$`)
	sleep := make(map[int]int, 0)
	sleepSince := 0
	guard := 0

	for _, m := range minutes {
		fmt.Println("minute:", m)
		fmt.Println("line:", lines[m])
		guardMatch := guardMatcher.FindStringSubmatch(lines[m])
		if (guardMatch != nil) {
			// no new guard
			if (lines[m][0] == 'f') {
				// falls asleep
				sleepSince = m
			} else {
				// wakes up
				sleep[guard] += m - sleepSince
			}
		} else {
			// new guard
			guard, _ = strconv.Atoi(guardMatch[1])
		}	
	}

	fmt.Println(sleep)
}

