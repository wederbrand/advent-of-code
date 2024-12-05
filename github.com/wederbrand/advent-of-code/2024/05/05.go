package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"slices"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2024/05/input.txt", "\n")

	p1, p2 := doIt(inFile)

	fmt.Println("Part 1:", p1, "in", time.Since(start))
	fmt.Println("Part 2:", p2, "in", time.Since(start))
}

func doIt(inFile []string) (int, int) {
	p1 := 0
	p2 := 0

	rulesMode := true
	afters := make(map[int][]int)
	for _, line := range inFile {
		if line == "" {
			rulesMode = false
			continue
		}

		if rulesMode {
			var a, b int
			fmt.Sscanf(line, "%d|%d", &a, &b)
			afters[a] = append(afters[a], b)
		} else {
			result := util.MatchingNumbersAfterSplitOnAny(line, " ", ",")[0]
			sorted := util.CloneSlice(result)
			slices.SortFunc(sorted, func(a, b int) int {
				if slices.Contains(afters[a], b) {
					// b is in the list of what's after a
					// so a is before b => negative
					return -1
				} else {
					// b is not in the list of what's after a
					// so a is after b => positive
					return 1
				}
			})

			if slices.Equal(result, sorted) {
				p1 += result[len(result)/2]
			} else {
				p2 += sorted[len(sorted)/2]
			}
		}
	}

	return p1, p2
}
