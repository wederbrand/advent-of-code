package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"slices"
	"strings"
	"time"
)

type lens struct {
	label string
	focal int
}

func main() {
	startTimer := time.Now()
	inFile := util.GetFileContents("2023/15/input.txt", "\n")

	part1 := 0
	var boxes [256][]lens
	for _, s := range strings.Split(inFile[0], ",") {
		part1 += hashOf(s)
		if strings.ContainsRune(s, '-') {
			split := strings.Split(s, "-")
			label := split[0]
			box := hashOf(label)
			boxes[box] = slices.DeleteFunc(boxes[box], func(l lens) bool {
				return l.label == label
			})
		} else {
			split := strings.Split(s, "=")
			label := split[0]
			focal := util.Atoi(split[1])
			box := hashOf(label)
			found := false
			for i, l := range boxes[box] {
				if l.label == label {
					boxes[box][i].focal = focal
					found = true
				}
			}
			if !found {
				boxes[box] = append(boxes[box], lens{label, focal})
			}
		}
	}

	part2 := 0
	for i, box := range boxes {
		for j, l := range box {
			part2 += (i + 1) * (j + 1) * l.focal
		}
	}

	fmt.Println("part1: ", part1, "in", time.Since(startTimer))
	fmt.Println("part2: ", part2, "in", time.Since(startTimer))
}

func hashOf(s string) int {
	hash := 0
	for _, r := range s {
		hash += int(r)
		hash *= 17
		hash %= 256
	}

	return hash
}
