package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"regexp"
	"time"
)

type fabric struct {
	id int
	x  int
	y  int
	w  int
	h  int
}

func main() {
	start := time.Now()
	inFile := GetFileContents("2018/03/input.txt", "\n")

	matcher := regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)

	fabrics := make([]fabric, 0)
	for _, s := range inFile {
		match := matcher.FindStringSubmatch(s)
		id := Atoi(match[1])
		x := Atoi(match[2])
		y := Atoi(match[3])
		w := Atoi(match[4])
		h := Atoi(match[5])
		f := fabric{id, x, y, w, h}
		fabrics = append(fabrics, f)
	}

	var overlaps [1001][1001]int

	part1 := 0
	for _, f := range fabrics {
		for i := f.x; i < f.w+f.x; i++ {
			for j := f.y; j < f.h+f.y; j++ {
				overlaps[i][j]++
				if overlaps[i][j] == 2 {
					part1++
				}
			}
		}
	}

	for _, f := range fabrics {
		found := false
		for i := f.x; i < f.w+f.x; i++ {
			for j := f.y; j < f.h+f.y; j++ {
				if overlaps[i][j] > 1 {
					found = true
				}
			}
		}
		if !found {
			fmt.Println("Part 2: ", f.id, "in", time.Since(start))
		}
	}

	fmt.Println("Part 1: ", part1, "in", time.Since(start))
}
