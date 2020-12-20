package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const size = 10

type pos struct {
	x int
	y int
}

type tile struct {
	id    int
	dot   map[pos]bool
	sides [4]*tile // 0 up, 1 right, 2 down, 3 left
}

func (t tile) getSide(side int, reverse bool) string {
	result := ""
	switch side {
	case 0: // up left to right
		for x := 0; x < size; x++ {
			if t.dot[pos{x, 0}] {
				result += "#"
			} else {
				result += "."
			}
		}
	case 1: // right top to bottom
		for y := 0; y < size; y++ {
			if t.dot[pos{size - 1, y}] {
				result += "#"
			} else {
				result += "."
			}
		}
	case 2: // down right to left
		for x := 0; x < size; x++ {
			if t.dot[pos{size - 1 - x, size - 1}] {
				result += "#"
			} else {
				result += "."
			}
		}
	case 3: // left down to up
		for y := 0; y < size; y++ {
			if t.dot[pos{0, size - 1 - y}] {
				result += "#"
			} else {
				result += "."
			}
		}
	}

	// possibly reverse it
	if reverse {
		runes := []rune(result)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		result = string(runes)
	}

	return result
}

var edge = tile{
	id: 0,
}

func newTile(input string) *tile {
	t := new(tile)
	split := strings.Split(strings.TrimSpace(input), "\n")

	t.dot = make(map[pos]bool)

	for i, s := range split {
		if i == 0 {
			idPattern := regexp.MustCompile("\\d+")
			idString := idPattern.FindString(s)
			atoi, _ := strconv.Atoi(idString)
			t.id = atoi
			continue
		}

		for j, r := range s {
			if r == '#' {
				t.dot[pos{j, i - 1}] = true
			}
		}
	}

	return t
}

func main() {
	readFile, err := ioutil.ReadFile("20/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.TrimSpace(string(readFile)), "\n\n")

	tiles := make([]*tile, 0)
	for _, s := range input {
		t := newTile(s)
		tiles = append(tiles, t)
	}

	for _, t1 := range tiles {
		for _, t2 := range tiles {
			if t1 != t2 {
				matchTiles(t1, t2)
			}
		}
	}

	prod := 1
	for _, t := range tiles {
		sides := 0
		for j := 0; j < 4; j++ {
			if t.sides[j] != nil {
				sides++
			}
		}
		if sides == 2 {
			prod *= t.id
		}
	}

	fmt.Println(prod)
}

func matchTiles(t1 *tile, t2 *tile) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			matchSides(t1, t2, i, j)
		}
	}
}

func matchSides(t1 *tile, t2 *tile, i int, j int) {
	side1 := t1.getSide(i, false)
	side2 := t2.getSide(j, false)
	side2flipped := t2.getSide(j, true)

	if side1 == side2 || side1 == side2flipped {
		t1.sides[i] = t2
		t2.sides[j] = t1
	}
}
