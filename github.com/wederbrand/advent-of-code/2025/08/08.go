package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"time"

	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"golang.org/x/exp/maps"
)

type Box struct {
	x, y, z int
}

type Pair struct {
	a, b Box
}

func main() {
	start := time.Now()
	input := GetFileContents("2025/08/input.txt", "\n")

	boxes := make([]Box, 0)

	for _, s := range input {
		n := Box{}
		fmt.Sscanf(s, "%d,%d,%d", &n.x, &n.y, &n.z)
		boxes = append(boxes, n)
	}
	pairs := getSortedPairs(boxes)

	circuits := make([]*[]Box, 0)

	for i, pair := range pairs {
		// check if a is in a circuit
		foundA := findInCircuits(pair.a, circuits)
		// check if b is in a circuit
		foundB := findInCircuits(pair.b, circuits)

		if foundA != nil && foundB != nil && foundA != foundB {
			*foundA = append(*foundA, *foundB...)
			// remove foundB from circuits
			circuits = slices.Delete(circuits, slices.Index(circuits, foundB), slices.Index(circuits, foundB)+1)
		} else if foundA != nil && foundB != nil && foundA == foundB {
			// both found in the same circuit, do nothing
			continue
		} else if foundA != nil {
			*foundA = append(*foundA, pair.b)
		} else if foundB != nil {
			*foundB = append(*foundB, pair.a)
		} else {
			// neither found, will create new circuit below
			circuits = append(circuits, &[]Box{pair.a, pair.b})
		}

		if i == 1000-1 {
			part1 := doPart1(circuits)
			fmt.Println("Part 1:", part1, "in", time.Since(start))
		}
		if len(circuits) == 1 && len(*circuits[0]) == len(boxes) {
			fmt.Println("Part 2:", pair.a.x*pair.b.x, "in", time.Since(start))
			os.Exit(0)
		}
	}
}

func doPart1(circuits []*[]Box) int {
	slices.SortFunc(circuits, func(a, b *[]Box) int {
		{
			return len(*b) - len(*a)
		}
	})

	part1 := 1
	for i := 0; i < 3; i++ {
		part1 *= len(*circuits[i])
	}
	return part1
}

func findInCircuits(a Box, circuits []*[]Box) *[]Box {
	for _, circuit := range circuits {
		if slices.Contains(*circuit, a) {
			return circuit
		}
	}
	return nil
}

func getSortedPairs(boxes []Box) []Pair {
	var pairs []Pair

	all := make(map[float64]Pair)

	for i := 0; i < len(boxes); i++ {
		a := boxes[i]
		for j := i + 1; j < len(boxes); j++ {
			b := boxes[j]
			dist := getStraightLineDistance(a, b)
			all[dist] = Pair{a, b}
		}
	}

	keys := maps.Keys(all)
	slices.Sort(keys)

	for _, key := range keys {
		pairs = append(pairs, all[key])
	}

	return pairs
}

func getStraightLineDistance(a Box, b Box) float64 {
	dx := a.x - b.x
	dy := a.y - b.y
	dz := a.z - b.z
	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}
