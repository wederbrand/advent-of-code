package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.ReadFile("01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	current := 0
	m := make([]int, 0)
	for _, s := range inFile {
		if s == "" {
			m = append(m, current)
			current = 0
		}
		atoi, _ := strconv.Atoi(s)
		current += atoi
	}
	m = append(m, current)

	sort.Sort(sort.Reverse(sort.IntSlice(m)))

	fmt.Println("part1: ", m[0])
	fmt.Println("part2: ", m[0]+m[1]+m[2])
}
