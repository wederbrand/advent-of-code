package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	keysOrLocks := GetFileContents("2024/25/input.txt", "\n\n")

	locks := make([][]int, 0)
	keys := make([][]int, 0)

	for _, keyOrLock := range keysOrLocks {
		m := MakeChart(strings.Split(keyOrLock, "\n"), ".")
		if keyOrLock[0] == '#' {
			locks = append(locks, countHashes(m))
		} else {
			keys = append(keys, countHashes(m))
		}
	}

	height := len(strings.Split(keysOrLocks[0], "\n"))

	part1 := 0
	for _, lock := range locks {
		for _, key := range keys {
			overlap := false
			for x := 0; x < len(lock); x++ {
				if lock[x]+key[x] > height {
					overlap = true
				}
			}
			if !overlap {
				part1++
			}
		}
	}

	fmt.Println("Part 1:", part1, "in", time.Since(start))
}

func countHashes(m Chart) []int {
	_, maxC := GetChartMaxes(m)
	count := make([]int, maxC.X+1)
	for c := range m {
		count[c.X]++
	}
	return count
}
