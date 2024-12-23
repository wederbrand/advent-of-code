package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"slices"
	"strings"
	"time"
)

type Computer struct {
	name     string
	connects map[string]bool
}

func main() {
	start := time.Now()
	inFile := GetFileContents("2024/23/input.txt", "\n")

	computers := make(map[string]Computer)
	for _, line := range inFile {
		one := strings.Split(line, "-")[0]
		two := strings.Split(line, "-")[1]

		if _, ok := computers[one]; !ok {
			computers[one] = Computer{name: one, connects: make(map[string]bool)}
		}
		if _, ok := computers[two]; !ok {
			computers[two] = Computer{name: two, connects: make(map[string]bool)}
		}

		computers[one].connects[two] = true
		computers[two].connects[one] = true
	}

	rings := make(map[string]bool)
	for _, one := range computers {
		if one.name[0] != 't' {
			// we are doing all computers even if found from a different one
			// so we can skip the ones that don't start with t, they will be found from the others
			continue
		}
		for oneConnect := range one.connects {
			two := computers[oneConnect]
			if two.name == one.name {
				// no backtracking
				continue
			}
			for twoConnect := range two.connects {
				three := computers[twoConnect]
				if three.name == one.name || three.name == two.name {
					// no backtracking
					continue
				}

				if three.connects[one.name] {
					names := []string{one.name, two.name, three.name}
					slices.Sort(names)
					rings[strings.Join(names, " ")] = true
				}
			}
		}
	}

	part1 := len(rings)
	fmt.Println("Part 1:", part1, "in", time.Since(start))

	biggest := 0
	var biggestLan []string
	for _, one := range computers {
		// create a tiny LAN with one and any other computer it sees
		for two := range one.connects {
			lan := []string{one.name, two}
			// then add any computer connected to all the computers already in the LAN
			for three := range one.connects {
				if two == three {
					// no backtracking
					continue
				}

				valid := true
				for _, member := range lan {
					if !computers[three].connects[member] {
						valid = false
						break
					}
				}

				if valid {
					lan = append(lan, three)
				}
			}

			if len(lan) > biggest {
				biggest = len(lan)
				biggestLan = lan
			}
		}
	}

	slices.Sort(biggestLan)
	part2 := strings.Join(biggestLan, ",")
	fmt.Println("Part 2:", part2, "in", time.Since(start))
}
