package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"math"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2016/24/input.txt", "\n")

	m := MakeChart(inFile, "")

	numbers := make([]Coord, 0)
	for i := 1; ; i++ {
		letter, err := m.FindLetter(strconv.Itoa(i))
		if err != nil {
			break
		}
		numbers = append(numbers, letter)
	}

	zero, _ := m.FindLetter("0")

	part1 := doIt(numbers, m, zero, false)
	fmt.Println("Part 1: ", part1, "in", time.Since(start))

	part2 := doIt(numbers, m, zero, true)
	fmt.Println("Part 2: ", part2, "in", time.Since(start))

	fmt.Println("Cache hits:", cacheHits)
	fmt.Println("Cache misses:", cacheMisses)
}

func doIt(numbers []Coord, m Chart, zero Coord, returnHome bool) int {
	permutations := Permutations(numbers)
	lowest := math.MaxInt
	for _, permutation := range permutations {
		sum := 0
		for i := 0; i < len(permutation); i++ {
			if i == 0 {
				sum += walkItCached(m, zero, permutation[i])
			} else {
				sum += walkItCached(m, permutation[i-1], permutation[i])
			}
			if i == len(permutation)-1 && returnHome {
				sum += walkItCached(m, permutation[i], zero)
			}
		}

		if sum < lowest {
			lowest = sum
		}
	}
	return lowest
}

var cache = make(map[[2]Coord]int)
var cacheHits = 0
var cacheMisses = 0

func walkItCached(m Chart, a Coord, b Coord) int {
	if val, ok := cache[[2]Coord{a, b}]; ok {
		cacheHits++
		return val
	}
	cacheMisses++
	length := m.GetPathLength(a, b)
	cache[[2]Coord{a, b}] = length
	return length
}
