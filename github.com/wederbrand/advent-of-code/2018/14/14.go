package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	part1 := doIt(157901)
	fmt.Println("Part 1:", part1, "in", time.Since(start))

	part2 := doIt2(157901)
	fmt.Println("Part 2:", part2, "in", time.Since(start))
}

func doIt(rounds int) string {
	r := make([]int, 0)
	r = append(r, 3, 7)
	elf1 := 0
	elf2 := 1

	for {
		// create new recipes
		newRecipe := r[elf1] + r[elf2]
		if newRecipe > 9 {
			r = append(r, newRecipe/10, newRecipe%10)
		} else {
			r = append(r, newRecipe)
		}

		// move elves
		elf1 = (elf1 + r[elf1] + 1) % len(r)
		elf2 = (elf2 + r[elf2] + 1) % len(r)
		if len(r) > rounds+10 {
			break
		}
	}

	result := ""
	for _, value := range r[rounds : rounds+10] {
		result += strconv.Itoa(value)
	}
	return result
}

func doIt2(target int) int {
	r := []byte{'3', '7'}
	elf1 := 0
	elf2 := 1

	for len(r) < 30000000 {
		// create new recipes
		newRecipe := []byte(strconv.Itoa(int(r[elf1] - '0' + r[elf2] - '0')))
		r = append(r, newRecipe...)

		// move elves
		elf1 = (elf1 + int(r[elf1]-'0') + 1) % len(r)
		elf2 = (elf2 + int(r[elf2]-'0') + 1) % len(r)

	}

	return strings.Index(string(r), strconv.Itoa(target))
}
