package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"regexp"
	"time"
	"unicode"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2023/03/input.txt", "\n")

	m := make(map[string]rune)
	for y, s := range inFile {
		for x, r := range s {
			if r != '.' && !unicode.IsNumber(r) {
				m[util.IntKey(x, y)] = r
			}
		}
	}

	// find each number
	part1 := 0
	gears := make(map[string][]int)
	re := regexp.MustCompile(`[0-9]+`)
	for y, s := range inFile {
		matches := re.FindAllStringSubmatch(s, -1)
		indexes := re.FindAllStringSubmatchIndex(s, -1)
		for i, match := range matches {
			index := indexes[i]
			for x := index[0] - 1; x <= index[1]; x++ {
				for yOffset := -1; yOffset <= 1; yOffset++ {
					key := util.IntKey(x, y+yOffset)
					r, found := m[key]
					if !found {
						continue
					}

					atoi := util.Atoi(match[0])
					if !unicode.IsNumber(r) {
						part1 += atoi
					}

					if r == '*' {
						gears[key] = append(gears[key], atoi)
					}
				}
			}
		}
	}
	fmt.Println("part1: ", part1, "in", time.Since(start))

	part2 := 0
	for _, ints := range gears {
		if len(ints) == 2 {
			part2 += ints[0] * ints[1]
		}
	}
	fmt.Println("part2: ", part2, "in", time.Since(start))
}
