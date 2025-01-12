package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2018/20/input.txt", "\n")

	m := Chart{}

	c := Coord{0, 0}
	m[c] = "X"

	split := strings.Split(inFile[0], "")
	split = split[1 : len(split)-1]

	createMap(m, c, strings.Join(split, ""))
	fmt.Println("Map created in", time.Since(start))
	
	newM := CopyChart(m)
	part1 := floodFill(newM, c, 0)
	fmt.Println("Part 1:", part1, "in", time.Since(start))

	part2 := 0
	for _, v := range newM {
		atoi, err := strconv.Atoi(v)
		if err == nil && atoi >= 1000 {
			part2++
		}
	}

	fmt.Println("Part 2:", part2, "in", time.Since(start))

}

func floodFill(m Chart, c Coord, i int) int {
	// mark this as 0, don't mark the X
	// walk in all directions with doors recursively
	// if we hit a number stop walking

	if m[c] != "." && m[c] != "X" {
		return i - 1
	}

	if m[c] != "X" {
		m[c] = fmt.Sprintf("%d", i)
	}

	maxVal := 0
	for _, d := range ALL {
		if m[c.Move(d)] == "-" || m[c.Move(d)] == "|" {
			v := floodFill(m, c.Move(d, 2), i+1)
			if v > maxVal {
				maxVal = v
			}
		}
	}
	return maxVal

}

type cacheKey struct {
	c  Coord
	in string
}

var cache = make(map[cacheKey]bool)

func createMap(m Chart, c Coord, in string) {
	key := cacheKey{c, in}
	if _, ok := cache[key]; ok {
		return
	}
	cache[key] = true

	if len(in) == 0 {
		return
	}
	if in[0] == '(' {
		inputs := make([]string, 0)
		input := ""
		i := 0
		open := 1
		for i = 1; i < len(in); i++ {
			if in[i] == '|' {
				if open == 1 {
					inputs = append(inputs, input)
					input = ""
				} else {
					input += string(in[i])
				}
			} else if in[i] == ')' {
				open--
				if open == 0 {
					inputs = append(inputs, input)
					break
				} else {
					input += string(in[i])
				}
			} else if in[i] == '(' {
				open++
				input += string(in[i])
			} else {
				input += string(in[i])
			}
		}
		for _, j := range inputs {
			createMap(m, c, j+in[i+1:])
		}
	} else {
		d := FromCompassChar(in[0])
		if d == N || d == S {
			m[c.Move(d)] = "-"
		} else {
			m[c.Move(d)] = "|"
		}
		c = c.Move(d, 2)
		m[c] = "."
		createMap(m, c, in[1:])
	}
}
