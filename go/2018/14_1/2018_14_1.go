package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	doIt(9)
	doIt(5)
	doIt(18)
	doIt(2018)
	doIt(157901)

	doIt2(51589)
	doIt2(01245)
	doIt2(92510)
	doIt2(59414)
	doIt2(157901)
}

func doIt(rounds int) {
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

	fmt.Print("part1 ")
	for _, value := range r[rounds : rounds+10] {
		fmt.Print(value)
	}
	fmt.Println()
}

func doIt2(target int) {
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
	if strings.Index(string(r), strconv.Itoa(target)) > -1 {
		fmt.Println("part2", strings.Index(string(r), strconv.Itoa(target)))
	} else {
		fmt.Println("part 2 need more")
	}

}
