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

	dX := t.X - s.X
	dY := t.Y - s.Y

	var movements []Dir
	for dX < 0 {
		movements = append(movements, LEFT)
		dX++
	}
	for dX > 0 {
		movements = append(movements, RIGHT)
		dX--
	}
	for dY < 0 {
		movements = append(movements, UP)
		dY++
	}
	for dY > 0 {
		movements = append(movements, DOWN)
		dY--
	}

	permutations := Permutations(movements)

	// de-duplicate and stringify
	result := make(map[string]bool)
	for _, p := range permutations {
		movement := ""
		curr := s
		illegal := false
		for _, d := range p {
			movement += d.ToArrowString()
			curr = curr.Move(d)

			if curr == pad["X"] {
				illegal = true
			}
		}
		if !illegal {
			result[movement+"A"] = true
		}
	}

	pathCache[cacheKey] = Keys(result)
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

	pushes := 0
	current := pad["A"]
	for i := 0; i < len(path); i++ {
		next := pad[string(path[i])]

		paths := getPathsBetweenButtons(current, next)

		minPushes := math.MaxInt
		for _, s := range paths {
			candidate := minPresses(s, levelsLeft-1)
			if candidate < minPushes {
				minPushes = candidate
			}
		}
		current = next
		pushes += minPushes
	}

	keyPressCache[cacheKey] = pushes
	return pushes
}
