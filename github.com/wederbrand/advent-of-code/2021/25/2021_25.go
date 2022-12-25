package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func (p *point) key() string {
	return strconv.Itoa(p.x) + "," + strconv.Itoa(p.y)
}

func main() {
	readFile, err := os.ReadFile("2021/25/2021_25.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	sea := make(map[point]rune)
	dx, dy := 0, 0
	for y, str := range inFile {
		dy++
		for x, r := range str {
			dx = len(str)
			p := point{x: x, y: y}
			if r == '>' || r == 'v' {
				sea[p] = r
			}
		}
	}

	printIt(sea, dx, dy)
	fmt.Println("part 1", part1(sea, dx, dy))
	//part2(cubes)
}

func printIt(sea map[point]rune, dx int, dy int) {
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			r, found := sea[point{x, y}]
			if found {
				fmt.Print(string(r))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func part2() {
}

func part1(sea map[point]rune, dx int, dy int) int {
	moves := 0
	for {
		moves++
		moved := false

		// east
		newSea := make(map[point]rune)
		for p, r := range sea {
			newP := p
			if r == '>' {
				newP.x = (newP.x + 1) % dx
			}
			_, found := sea[newP]
			if found {
				// unmoved
				newSea[p] = r
			} else {
				// moved
				newSea[newP] = r
				moved = true
			}
		}

		// south
		sea = newSea
		newSea = make(map[point]rune)

		for p, r := range sea {
			newP := p
			if r == 'v' {
				newP.y = (newP.y + 1) % dy
			}
			_, found := sea[newP]
			if found {
				// unmoved
				newSea[p] = r
			} else {
				// moved
				newSea[newP] = r
				moved = true
			}
		}
		sea = newSea
		// printIt(sea, dx, dy)
		if !moved {
			return moves
		}
	}

	return 0
}
