package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type food struct {
	in []string
	al []string
}

func (f *food) removeIn(ingredient string) {
	newIn := make([]string, 0)
	for _, s := range f.in {
		if s != ingredient {
			newIn = append(newIn, s)
		}
	}
	f.in = newIn
}

func (f *food) removeAl(allergen string) {
	newAl := make([]string, 0)
	for _, s := range f.al {
		if s != allergen {
			newAl = append(newAl, s)
		}
	}
	f.al = newAl
}

func main() {
	readFile, err := ioutil.ReadFile("21/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.TrimSpace(string(readFile)), "\n")
	allFood := make([]*food, 0)
	for _, s := range input {
		f := new(food)
		allFood = append(allFood, f)
		s = strings.TrimSuffix(s, ")")
		split := strings.Split(s, "(contains")
		f.in = strings.Split(strings.TrimSpace(split[0]), " ")
		f.al = strings.Split(strings.TrimSpace(split[1]), ", ")
	}

	// iterate over all foods forever
RESTART:
	for i := 0; i < len(allFood); i++ {
		f1 := allFood[i]
		for _, f2 := range allFood {
			if f1 == f2 {
				continue
			}

			// two different food
			// find exactly one allergen that is the same (two of the same will have to wait)
			commonAllergen := findCommon(f1.al, f2.al)
			commonIngredient := findCommon(f1.in, f2.in)

			if commonAllergen != "" && commonIngredient != "" {
				for _, f := range allFood {
					f.removeIn(commonIngredient)
					f.removeAl(commonAllergen)
				}
				continue RESTART
			}
		}
	}

	// if we get this far we didn't star over and were done
	allIn := make(map[string]bool)
	for _, f := range allFood {
		for _, s := range f.in {
			allIn[s] = true
		}
	}

	// 200 is too low
	fmt.Println(len(allIn))
}

func findCommon(f1 []string, f2 []string) string {
	result := ""
	count := 0
	for _, s1 := range f1 {
		for _, s2 := range f2 {
			if s1 == s2 {
				count++
				result = s1
			}
		}
	}

	if count == 1 {
		return result
	} else {
		return ""
	}

}
