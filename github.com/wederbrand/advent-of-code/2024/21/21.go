package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"math"
	"time"
)

var pad map[string]Coord

type keyPressCacheKey struct {
	path       string
	levelsLeft int
}

var keyPressCache = make(map[keyPressCacheKey]int)

type pathCacheKey struct {
	s Coord
	t Coord
}

var pathCache = make(map[pathCacheKey][]string)

func main() {
	start := time.Now()
	inFile := GetFileContents("2024/21/input.txt", "\n")

	pad = make(map[string]Coord)

	pad["7"] = Coord{0, 0}
	pad["8"] = Coord{1, 0}
	pad["9"] = Coord{2, 0}

	pad["4"] = Coord{0, 1}
	pad["5"] = Coord{1, 1}
	pad["6"] = Coord{2, 1}

	pad["1"] = Coord{0, 2}
	pad["2"] = Coord{1, 2}
	pad["3"] = Coord{2, 2}

	// The bottom line of the numerical pad, and the top one of the arrow one are identical except for ^/0.
	// We can keep these in the same map.
	pad["X"] = Coord{0, 3}
	pad["0"] = Coord{1, 3}
	pad["^"] = Coord{1, 3}
	pad["A"] = Coord{2, 3}

	pad["<"] = Coord{0, 4}
	pad["v"] = Coord{1, 4}
	pad[">"] = Coord{2, 4}

	part1 := 0
	for _, line := range inFile {
		cheapest := minPresses(line, 2+1)
		part1 += cheapest * Atoi(line[0:len(line)-1])
	}
	fmt.Println("Part 1:", part1, "in", time.Since(start))

	part2 := 0
	for _, line := range inFile {
		cheapest := minPresses(line, 25+1)
		part2 += cheapest * Atoi(line[0:len(line)-1])
	}
	fmt.Println("Part 2:", part2, "in", time.Since(start))
}

func getPathsBetweenButtons(s Coord, t Coord) []string {
	cacheKey := pathCacheKey{s, t}
	path, ok := pathCache[cacheKey]
	if ok {
		return path
	}

	if s == t {
		return []string{"A"}
	}

func getPathsBetweenNumbers(s Coord, t Coord) []string {
	dX := t.X - s.X
	dY := t.Y - s.Y

	movements := []string{}
	for dX < 0 {
		movements = append(movements, "<")
		dX++
	}
	for dX > 0 {
		movements = append(movements, ">")
		dX--
	}
	for dY < 0 {
		movements = append(movements, "^")
		dY++
	}
	for dY > 0 {
		movements = append(movements, "v")
		dY--
	}

	permutations := Permutations(movements)

	// de-duplicate and stringify
	result := make(map[string]bool)
outer:
	for _, p := range permutations {
		movement := ""
		curr := s
		for _, r := range p {
			// TODO: check the avoid-box here
			movement += r

			if r == "<" {
				curr = curr.Move(LEFT)
			}
			if r == ">" {
				curr = curr.Move(RIGHT)
			}
			if r == "^" {
				curr = curr.Move(UP)
			}
			if r == "v" {
				curr = curr.Move(DOWN)
			}

			if curr == numbers["X"] {
				// illegal move
				continue outer
			}
		}
		result[movement+"A"] = true
	}
	if len(result) == 0 {
		return []string{"A"}
	}
	return Keys(result)
}

func minPresses(path string, levelsLeft int) int {
	cacheKey := keyPressCacheKey{path, levelsLeft}
	mem, ok := keyPressCache[cacheKey]
	if ok {
		return mem
	}

	if levelsLeft == 0 {
		return len(path)
	}

	slutSvaret := 0
	current := arrows["A"]
	for i := 0; i < len(path); i++ {
		next := arrows[string(path[i])]

		allPathsBetweenArrows := getPathsBetweenNumbers(current, next)

		minPushes := math.MaxInt
		for _, s := range paths {
			candidate := minPresses(s, levelsLeft-1)
			if candidate < minPushes {
				minPushes = candidate
			}
		}
		current = next
		slutSvaret += minsta
	}

	cache[key] = slutSvaret
	return slutSvaret
}
