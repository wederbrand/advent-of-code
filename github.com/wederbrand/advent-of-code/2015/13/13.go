package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"math"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2015/13/input.txt", "\n")

	eaters := make(map[string]bool)
	gains := make(map[string]int)
	for _, s := range inFile {
		var a string
		var b string
		var gainLose string
		var units int
		fmt.Sscanf(s, "%s would %s %d happiness units by sitting next to %s", &a, &gainLose, &units, &b)
		b = strings.Trim(b, ".")
		eaters[a] = true
		if gainLose == "lose" {
			units *= -1
		}
		gains[util.Key(a, b)] = units
	}

	alternatives := util.Permutations(util.Keys(eaters))

	part1 := math.MinInt
	for _, alternative := range alternatives {
		happiness := 0
		for i := range alternative {
			j := i + 1
			if j == len(alternative) {
				j = 0
			}
			happiness += gains[util.Key(alternative[i], alternative[j])]
			happiness += gains[util.Key(alternative[j], alternative[i])]
		}
		part1 = max(part1, happiness)
	}

	fmt.Println("part1", part1, "in", time.Since(start))

	// add me
	for s := range eaters {
		gains[util.Key("me", s)] = 0
		gains[util.Key(s, "me")] = 0
	}
	eaters["me"] = true
	alternatives = util.Permutations(util.Keys(eaters))

	part2 := math.MinInt
	for _, alternative := range alternatives {
		happiness := 0
		for i := range alternative {
			j := i + 1
			if j == len(alternative) {
				j = 0
			}
			happiness += gains[util.Key(alternative[i], alternative[j])]
			happiness += gains[util.Key(alternative[j], alternative[i])]
		}
		part2 = max(part2, happiness)
	}

	fmt.Println("part2", part2, "in", time.Since(start))

}
