package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Floor map[string]int

func (f Floor) key(x int, y int) string {
	return fmt.Sprint(x, ",", y)
}

func (f Floor) getVal(x int, y int) int {
	depth, found := f[f.key(x, y)]
	if !found {
		return math.MaxInt
	}
	return depth
}

func (f Floor) countAndRaise(x int, y int) int {
	if f.inBasin(x, y) {
		f[f.key(x, y)] = 9
		return 1 + f.countAndRaise(x, y-1) + f.countAndRaise(x, y+1) + f.countAndRaise(x+1, y) + f.countAndRaise(x-1, y)
	}
	return 0
}

func (f Floor) inBasin(x int, y int) bool {
	depth, found := f[f.key(x, y)]
	if !found || depth == 9 {
		return false
	}
	return true
}

func main() {
	readFile, err := os.ReadFile("2021/09/2021_09.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	floor := make(Floor)
	maxX := 0
	maxY := 0
	for y, text := range inFile {
		split := strings.Split(text, "")
		maxX = len(split)
		maxY++
		for x, s := range split {
			atoi, _ := strconv.Atoi(s)
			floor[floor.key(x, y)] = atoi
		}
	}

	part1(maxY, maxX, floor)
	part2(maxY, maxX, floor)
	os.Exit(0)
}

func part1(maxY int, maxX int, floor Floor) {
	risk := 0
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if floor.getVal(x, y) < floor.getVal(x, y-1) && floor.getVal(x, y) < floor.getVal(x, y+1) && floor.getVal(x, y) < floor.getVal(x-1, y) && floor.getVal(x, y) < floor.getVal(x+1, y) {
				risk += floor[floor.key(x, y)] + 1
			}
		}
	}

	fmt.Println("part 1", risk)
}

func part2(maxY int, maxX int, floor Floor) {
	sizes := make([]int, 0)
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if floor.inBasin(x, y) {
				size := floor.countAndRaise(x, y)
				sizes = append(sizes, size)
			}
		}
	}

	sort.Ints(sizes)
	result := 1
	for i := 0; i < 3; i++ {
		result *= sizes[len(sizes)-i-1]
	}

	fmt.Println("part 2", result)
}
