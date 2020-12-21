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
	id           int
	dot          map[pos]bool
	sides        [4]*tile // 0 up, 1 right, 2 down, 3 left
	flippedSides [4]bool
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

func (t tile) getMapPart(right int, flip bool, i int) string {
	result := ""
	// return the 8 characters with this dir1/dir2 and offset, and index (1 -> 8)
	switch right {
	case 0:
		if flip {
			for y := 1; y < size-1; y++ {
				if t.dot[pos{size - 1 - (i + 1), size - 1 - y}] {
					result += "#"
				} else {
					result += "."
				}
			}
		} else {
			for y := 1; y < size-1; y++ {
				if t.dot[pos{i + 1, size - 1 - y}] {
					result += "#"
				} else {
					result += "."
				}
			}
		}
	case 1:
		if flip {
			for x := 1; x < size-1; x++ {
				if t.dot[pos{x, size - 1 - (i + 1)}] {
					result += "#"
				} else {
					result += "."
				}
			}
		} else {
			for x := 1; x < size-1; x++ {
				if t.dot[pos{x, i + 1}] {
					result += "#"
				} else {
					result += "."
				}
			}
		}
	case 2:
		if flip {
			for y := 1; y < size-1; y++ {
				if t.dot[pos{i + 1, y}] {
					result += "#"
				} else {
					result += "."
				}
			}
		} else {
			for y := 1; y < size-1; y++ {
				if t.dot[pos{size - 1 + (i + 1), y}] {
					result += "#"
				} else {
					result += "."
				}
			}
		}
	case 3:
		if flip {
			for x := 1; x < size-1; x++ {
				if t.dot[pos{size - 1 - x, size - 1 - (i + 1)}] {
					result += "#"
				} else {
					result += "."
				}
			}
		} else {
			for x := 1; x < size-1; x++ {
				if t.dot[pos{size - 1 - x, i + 1}] {
					result += "#"
				} else {
					result += "."
				}
			}
		}
	}

	return result
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
	corners := make([]*tile, 0)
	for _, t := range tiles {
		sides := 0
		for j := 0; j < 4; j++ {
			if t.sides[j] != nil {
				sides++
			}
		}
		if sides == 2 {
			corners = append(corners, t)
			prod *= t.id
		}
	}

	fmt.Println("part1", prod)

	// PART 2
	// start on all the four corners, flipping each twice
	middleRE := regexp.MustCompile("^(.*)#....##....##....###(.*)$")
	for _, corner := range corners {
		for i := 0; i < 2; i++ {
			fmt.Println("map for corner: ", corner.id, "flip: ", i == 1)
			giantMap := makeGiantMap(corner, i == 1)
			for _, s := range giantMap {
				fmt.Println(s)
			}

			for i := 1; i < len(giantMap)-2; i++ {
				s := giantMap[i]
				// look for the middle line
				if middleRE.MatchString(s) {
					fmt.Println("found")
					// if found, check the -1 and +1 once
					// if found, replace all 3 with 0
					// count #
					// print
					// exit
				}
			}
		}
	}
	fmt.Println("result")
}

// should return 1 map, starting with this tile,
// in either it's first or second direction
func makeGiantMap(corner *tile, flip bool) []string {
	result := make([]string, 0)
	dir1 := -1
	dir2 := -1

	for dir := 0; dir < 3; dir++ {
		nextDir := dir + 1
		if corner.sides[dir] != nil && corner.sides[nextDir%4] != nil {
			// found the two sides
			dir1 = dir
			dir2 = nextDir
			break
		}
	}

	// calc what direction is RIGHT and what is DOWN. Note the flipping part.
	right := dir1
	down := dir2

	if flip {
		right, down = down, right
	}

	right %= 4
	down %= 4

	// call something recursive with tile and right direction.
	// then, go down, find tile underneath, turn it and call the recursive part with that one, and it's right direction!
	result = append(result, someThingRecursive(corner, right, flip, make([]string, 8))...)

	return result
}

// should return all lines for this tile and to the right
func someThingRecursive(t *tile, right int, flip bool, partial []string) []string {
	for i := range partial {
		partial[i] += " " + t.getMapPart(right, flip, i)
	}

	nextTile := t.sides[right]
	if nextTile == nil {
		// we're done
		return partial
	} else {
		// find the correct orientation
		mySide := t.getSide(right, flip)
		for i := 0; i < 4; i++ {
			for j := 0; j < 2; j++ {
				thatSide := nextTile.getSide(i, j == 0)
				if mySide == thatSide {
					return someThingRecursive(nextTile, (i+2)%4, j == 1, partial)
				}
			}
		}

		panic("at the disco")
	}
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

	if side1 == side2 {
		t1.sides[i] = t2
		t2.sides[j] = t1
	}

	if side1 == side2flipped {
		t1.sides[i] = t2
		t1.flippedSides[i] = true
		t2.sides[j] = t1
		t2.flippedSides[j] = true
	}
}
