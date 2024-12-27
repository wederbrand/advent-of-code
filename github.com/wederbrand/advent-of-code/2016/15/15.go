package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"time"
)

type Disc struct {
	positions int
	current   int
}

func (d *Disc) rotate(steps int) {
	newCurrent := (d.current + steps) % d.positions
	d.current = newCurrent
}

func main() {
	start := time.Now()
	inFile := GetFileContents("2016/15/input.txt", "\n")

	part1 := doIt(inFile)
	fmt.Println("Part 1: ", part1, "in", time.Since(start))

	part2 := doIt(append(inFile, "Disc #7 has 11 positions; at time=0, it is at position 0."))
	fmt.Println("Part 2: ", part2, "in", time.Since(start))
}

func doIt(inFile []string) int {
	discs := make([]*Disc, len(inFile))

	for _, line := range inFile {
		var d Disc
		var i int
		_, err := fmt.Sscanf(line, "Disc #%d has %d positions; at time=0, it is at position %d.", &i, &d.positions, &d.current)
		if err != nil {
			panic(err)
		}
		discs[i-1] = &d
	}

	// move disc start position according to time
	for i, d := range discs {
		d.rotate(i + 1)
	}

	// find the first time when all discs are at 0
	dropTime := 0
	found := false
	for !found {
		for _, d := range discs {
			d.rotate(1)
		}

		found = true
		for _, d := range discs {
			if d.current != 0 {
				found = false
				break
			}
		}

		dropTime++
	}
	return dropTime
}
