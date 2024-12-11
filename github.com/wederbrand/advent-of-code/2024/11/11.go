package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strconv"
	"strings"
	"time"
)

type Pair struct {
	stone      int
	iterations int
}

var memory = make(map[Pair]int)

func main() {
	start := time.Now()
	inFile := GetFileContents("2024/11/input.txt", "\n")

	stones := make([]int, 0)
	split := strings.Split(inFile[0], " ")
	for _, s := range split {
		stones = append(stones, Atoi(s))
	}

	p1 := 0
	for _, s := range stones {
		p1 += doIt(s, 25)
	}
	fmt.Println("Part 1:", p1, "in", time.Since(start))

	p2 := 0
	for _, s := range stones {
		p2 += doIt(s, 75)
	}
	fmt.Println("Part 2:", p2, "in", time.Since(start))
}

func doIt(stone int, iterations int) int {
	if mem, found := memory[Pair{stone, iterations}]; found {
		return mem
	}

	if iterations == 0 {
		return 1
	}
	count := 0
	switch {
	case stone == 0:
		count = doIt(1, iterations-1)
	case len(strconv.Itoa(stone))%2 == 0:
		stringValue := strconv.Itoa(stone)
		half := Atoi(stringValue[:len(stringValue)/2])
		otherHalf := Atoi(stringValue[len(stringValue)/2:])
		count += doIt(half, iterations-1)
		count += doIt(otherHalf, iterations-1)
	default:
		count += doIt(stone*2024, iterations-1)
	}

	memory[Pair{stone, iterations}] = count
	return count
}
