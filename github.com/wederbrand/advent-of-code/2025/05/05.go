package main

import (
	"fmt"
	"maps"
	"sort"
	"strings"
	"time"

	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
)

func main() {
	start := time.Now()
	input := GetFileContents("2025/05/input.txt", "\n")

	ranges := [][]int{}
	ingredients := []int{}
	for _, s := range input {
		// range
		if strings.Contains(s, "-") {
			var a, b int
			fmt.Sscanf(s, "%d-%d", &a, &b)
			ranges = append(ranges, []int{a, b})
		} else if s != "" {
			// ingredient
			ingredients = append(ingredients, Atoi(s))
		}
	}

	part1 := 0
	for _, ingredient := range ingredients {
		spoiled := true
		for _, r := range ranges {
			if ingredient >= r[0] && ingredient <= r[1] {
				spoiled = false
				break
			}
		}
		if !spoiled {
			part1++
		}
	}
	fmt.Println("Part 1:", part1, "in", time.Since(start))

	var all = make(map[int]bool)
	for _, r := range ranges {
		all[r[0]] = true
		all[r[1]] = true
	}

	keys := []int{}
	for k := range maps.Keys(all) {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	part2 := 0
	in := false
	entry := 0
	for _, i := range keys {
		if in {
			// in - look for an exit
			found := false
			for _, r := range ranges {
				if i+1 >= r[0] && i+1 <= r[1] {
					found = true
					break
				}
			}
			// outside all ranges
			if !found {
				part2 += i - entry + 1
				in = false
			}
		} else {
			// out - look for an entry
			for _, r := range ranges {
				if i >= r[0] && i <= r[1] {
					in = true
					entry = i
					break
				}
			}
		}
	}

	// 353863745078671 is correct
	fmt.Println("Part 2:", part2, "in", time.Since(start))
}
