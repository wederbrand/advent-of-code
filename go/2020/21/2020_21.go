package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
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

	allAl := make(map[string][]string)
	for _, food := range allFood {
		for _, allergen := range food.al {
			if allAl[allergen] == nil {
				// add all
				allAl[allergen] = food.in
			} else {
				allAl[allergen] = union(allAl[allergen], food.in)
			}
		}
	}

	// clear out each ingredient that is uniq from all the others, until none is removed
OUTER:
	for {
		for al, in := range allAl {
			if len(in) == 1 {
				for innerAl := range allAl {
					if al != innerAl {
						minus := setMinus(allAl[innerAl], in[0])
						if len(minus) != len(allAl[innerAl]) {
							allAl[innerAl] = minus
							continue OUTER
						}
					}
				}
			}
		}

		// no changes, all done
		break
	}

	// clear out known allergens
	for _, food := range allFood {
		for _, ingredient := range allAl {
			food.in = setMinus(food.in, ingredient[0])
		}
	}

	cnt := 0
	for _, f := range allFood {
		cnt += len(f.in)
	}

	fmt.Println("part 1:", cnt)

	sortedAl := make([]string, 0)
	for al := range allAl {
		sortedAl = append(sortedAl, al)
	}

	sort.Strings(sortedAl)

	result := ""
	for _, al := range sortedAl {
		result += allAl[al][0] + ","
	}

	result = strings.TrimRight(result, ",")
	fmt.Println("part 2:", result)
}

func setMinus(list []string, entryToRemove string) []string {
	result := make([]string, 0)
	for _, s := range list {
		if s != entryToRemove {
			result = append(result, s)
		}
	}

	return result
}

func union(a []string, b []string) []string {
	result := make([]string, 0)
	for _, s1 := range a {
		for _, s2 := range b {
			if s1 == s2 {
				result = append(result, s1)
			}
		}
	}
	return result
}
