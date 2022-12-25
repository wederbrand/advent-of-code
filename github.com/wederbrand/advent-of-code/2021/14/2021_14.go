package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"unicode/utf8"
)

type poly map[string]int

func main() {
	readFile, err := os.ReadFile("2021/14/2021_14.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	var last rune
	data := make(poly)
	for _, r := range inFile[0] {
		if last == 0 {
			last = r
			continue
		}

		data[string(last)+string(r)]++
		last = r
	}

	rules := make(map[string][2]string)
	for i := 2; i < len(inFile); i++ {
		split := strings.Split(inFile[i], " -> ")
		a := string(split[0][0])
		b := string(split[0][1])
		c := split[1]
		rules[a+b] = [2]string{a + c, c + b}
	}

	for i := 0; i < 10; i++ {
		data = insert(data, rules)
	}
	fmt.Print("part 1 ")
	printIt(data, last)

	for i := 0; i < 30; i++ {
		data = insert(data, rules)
	}
	fmt.Print("part 2 ")
	printIt(data, last)
}

func printIt(data poly, last rune) {
	m := make(map[rune]int)
	for s, i := range data {
		r, _ := utf8.DecodeRuneInString(s)
		m[r] += i
	}

	least := math.MaxInt
	most := math.MinInt

	for r, i := range m {
		if r == last {
			i++
		}
		if i > most {
			most = i
		}
		if i < least {
			least = i
		}
	}

	fmt.Println(most - least)
}

func insert(data poly, rules map[string][2]string) poly {
	nextPoly := make(poly)
	for pair, cnt := range data {
		nextPoly[rules[pair][0]] += cnt
		nextPoly[rules[pair][1]] += cnt
	}

	return nextPoly
}
