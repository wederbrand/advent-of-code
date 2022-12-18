package main

import (
	"fmt"
	"time"
)

type rock map[string]bool

func (r rock) clone(dx int, dy int) rock {
	n := make(rock)
	for s, b := range r {
		x, y := dekey(s)
		n[key(x+dx, y+dy)] = b
	}
	return n
}

func (r rock) hash(row int, height int) (result string) {
	for y := height; y < height+row; y++ {
		binary := ""
		for x := 0; x < 7; x++ {
			if r[key(x, y)] {
				binary += "1"
			} else {
				binary += "0"
			}
		}
		fmt.Sscanf(binary, "%b", &result)
	}

	return
}

func (r rock) outOfBounds() bool {
	for s := range r {
		x, y := dekey(s)
		if x < 0 || x >= 7 || y >= 0 {
			return true
		}
	}

	return false
}

func (r rock) collide(o rock) bool {
	for s := range r {
		_, found := o[s]
		if found {
			return true
		}
	}

	return false
}

func (r rock) draw(height int) {
	for y := height; y < 0; y++ {
		fmt.Print("|")
		for x := 0; x < 7; x++ {
			if r[key(x, y)] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("|")
	}
	fmt.Println()
}

type state struct {
	height int
	lastX  int
}

func key(x int, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

func dekey(key string) (x int, y int) {
	fmt.Sscanf(key, "%d:%d", &x, &y)
	return
}

const rowsToCompare = 10

func main() {
	start := time.Now()

	minusRock := make(rock)
	minusRock[key(2, 0)] = true
	minusRock[key(3, 0)] = true
	minusRock[key(4, 0)] = true
	minusRock[key(5, 0)] = true

	plusRock := make(rock)
	plusRock[key(3, -2)] = true
	plusRock[key(2, -1)] = true
	plusRock[key(3, -1)] = true
	plusRock[key(4, -1)] = true
	plusRock[key(3, 0)] = true

	lRock := make(rock)
	lRock[key(4, -2)] = true
	lRock[key(4, -1)] = true
	lRock[key(2, 0)] = true
	lRock[key(3, 0)] = true
	lRock[key(4, 0)] = true

	iRock := make(rock)
	iRock[key(2, -3)] = true
	iRock[key(2, -2)] = true
	iRock[key(2, -1)] = true
	iRock[key(2, 0)] = true

	boxRock := make(rock)
	boxRock[key(2, -1)] = true
	boxRock[key(3, -1)] = true
	boxRock[key(2, 0)] = true
	boxRock[key(3, 0)] = true

	rocks := [5]rock{
		minusRock,
		plusRock,
		lRock,
		iRock,
		boxRock,
	}

	input := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
	// input := ">><<<<>><>>>><><<>>>><<<>>>><<<<>>><>>>><<<<>>>><<><><<<><<><<<<><<<<>>>><<>>><<<>><<>>><<><<<>>>><>>><<><>><<>>><<<<><>><<>>><<<<><<<<><>><<<<>>><<<<>>>><<<>><><<<<>>><>><<<>><><<<>>><<<<>>><>>>><<<><<>>>><><<<<>>>><<<>>>><>>>><>>>><><<><<>>>><<<>>><<<>>>><><<<><<<<>><>>><<<<><<<><>><><><<><<><<<<>>>><<<<>>><>>><<<><<<<>>><<<<>>>><<<<>>><>>>><<<><<<<>>><<>><>><>><><>>>><<>>><>>>><<<<>><<<>>>><<>><<<>>>><><<<<><>><>>>><>><<<<>>>><<<>>>><>><>>><><<<<>><><<<><<<<>>><<><<<>>><<<<><<>>><<>><<<<><<>>>><<><>>><<><<<>>><<<<><<<<>><>>><<<>>><<<<>>>><<><<>><<<>><<>>>><<<>>>><<>><<>><<><<<<>><<<<>>><<><<<<>>><>><<>><<<<><<<><<<>>>><<<<>><>>>><<<<>><<>>><<>>>><<<<><<<><<>>><>>><<<>>>><>><<><><<<<>>>><<<<>>>><<<<><<<>>><<>><>>>><>>>><<<<>><<<<><<><>><<<<>><<<>>>><<>>>><<<<><<<>>><<<<>>><<>>><<<>>><<><<><<<>>><><<>>><<<>>>><>>>><<<><<<<><<>><><<>>>><><<<<>>><<><<<<><<<<><<<<>>>><<>><<<><<<>><<>>><<>><<<<><<><<<>>><<>><<<><>><<><<<<><<<<>><<>><<><>>><>>><<<>><<<>>>><>>><>><<><<>><><<<<><<<<>>><<<<>>><<<<><<<>>><<<>>>><<<>><<>><><<><<<<><>>>><>>>><<<<><>>><<<>><<<>>>><<<<>><<<<>><<<<>>>><>><<>>>><<<>>><<><<<>>>><<<><<><<<>>>><<<>>>><<>><<<<>>>><><<<>>><>>>><<<>>>><<>><<<<>>>><><<<>><<<>>><<>>><>>>><<>>><><<<<><>>>><><<<<>>>><<<<>><<<>>>><<<>><<<>><<<>>>><><<<<>>><<>>>><<>><><<<>>><<<>><<<>>>><>>>><<>>><<<<>>><<<>><><<<>>><<<<>>>><<>><<>>>><<<<>>><<<><>>><<<>>>><<<<>><<<<><<<<><<<<>>><<<>>><<<>>>><><>>>><>>><<>>><><>>><<>>><>><><<<>>>><<<>><<<<>>><<>>>><<>>><<<>><<<<>><>>><<<>>>><<<>>><<<>>><><><<<>>><<<>><<<>><<<<>><>>><<><>>><<>>><<>>><<<><>>><<<<>>><<<>>><<<>>><>>><<>><<>>>><><<<<>><>>>><<<<><<><<<<>>><>><>>>><<>>>><<>>>><<<>><>>><><<<<>>><><<>><<<>><<<<><<<><<<<><<<>><<>>>><<<<>>><>>><<<<><<<>><<<>>><<>>>><<<<>>>><<>>><><><<<<>>>><>>>><<<<>>><><<<><<<>>><>>><<<<><<<<><>>>><>>>><<><<<>>>><<>>><>>><<<<>>><<<<><<>>>><<<>>><<<<>>><>><<<><>>>><<>>><<<>><<<<>>>><<<>>><<<<><<<<><<>><>>><>><<<<><>>><>>>><<>><>>><<><<<><<>><<<>><>>><<<<>>>><>>>><<<<>>><<<><<<<>>>><<<<>>><<<<>>>><>>><<<>><<>>>><<<<>>>><<<<><<>>>><<<>><>>><<>><><<<<>><<<<>>><<<><<>>><>><<<><>><<>>><>>>><<>>><>>><>><<<><<>>>><<<<><<>><<<<><<<><<<<>>><<<<>>><<<>><<<<>>><<<<>>>><<><<>>><<<<><<>><<<<><<><<<<>>>><<<><<<<>><<<<><>>>><<><<<<>><>>>><><<<>>>><<<>><<<<>>>><<><<<<>><<>><<<<><<<><<>>><<>><>><<><<><>><<<<>>><<>>><>>><<<><<<<><<>>>><<>>><<>><>>>><><<<<><<>><>>><><<>>><<>>>><<<<>>><<<<>>>><<<>><<<<><<<>>><<<>><<<<><<<><><<>>>><<<>>>><<<>>>><<<>>><<>>>><<>>>><<<<>>><<>>><<<>><<>><<>>>><<<>>>><<<<><<>><>>><<<<>>>><<<><<>><>>>><>>>><>>>><<>><<<><<><>><<<>><<<<>><><<>><<>><<>>><><<<>><>><<><<<<><>>><>><<<>>>><<<>>><>>>><<>><><<>><<>>>><<<<><>>>><<>><<>>><<<>>>><<<<>><<>>><<<<>>>><>>><<<>><<<<>>>><<<<>>><<<>><<>>>><<<><<<<>><<<<>><<><<<>>><<>><<<><<<>>>><<<<>>>><<<<>><<<<>><><<>><<<<>>>><<<<>><<>>>><<>><>>>><<<><>><<><>><<>><><<><>>><<<>>>><>>><<>><>><>>>><<<<><<>><<<<><>>>><<>>><<>>><<><>>>><>><<<><<<><>>>><><<><<<<><<><<<<>>>><><<<<>>>><<<>><<>><<<<>>>><<<<>>>><<<>>>><>>><<<>>>><<<>><<<<>>>><<<>><<>>><<<<>>><<><<<<>>>><>><<<>><<>>>><<<<>>>><<>>><<<<>><><<<<>>><<<>>>><<>>>><><>>>><<>>><>>><<<>>><<<>><<<>><<<<><><<><<<<>><>>><>>><><>>><<<>>><<>><<><<<>>>><<<>>><>><><<>>>><<<><>><<<<>><<<><>>><<><<<<>>>><<>>>><<<<>><<<<>><<<<>>>><<<<>>><<<<><<<>><<>><<<<>><<><<<<>><<<<>>><<<<><<<<>>>><<><<<><<<<>>>><<>>><<<<>>>><<<<><<>>><>>><<<><<<<>><<<>>><>>>><><>>>><><>>><<<>>>><<<><>>>><<<>><<<><<>><<><><<<<><<>>><<>><<><<<>>>><<<>>><>>><>><<<<>>><<><<<<>><>>><<<>>>><<>><>>><<<>>><<<>><<<<>>>><>>><>>>><>>><<<>>><<<>>>><<<>><<<>><<<><><>>>><<<<>>>><<<<><>>><><<<<>>><<<<>>><<><><><<>><<><<<<>>><>>>><<>>><<>>>><<<>>><<<>><<<<><>>><>><<<<><<>><<<<><<<<>>>><<<>>><>>>><<<<>>>><<<<><<<<>>>><<<<>>>><<<>><<<<>>>><<>><<<>>>><<<>>>><<<<>><<<><>>>><<><>>>><<<<><>>><><<<><>>>><<>><<<<>><>><<<<>>>><<<<>><<>>><><<<><<<<>>>><>>>><<>><<<><<>>><<>><>><<>><<><<<<>><>><<<<>>>><<<<>><<<<>>><<><<>><>>><<<<>>>><<>>>><<><<>>>><<>><<>>>><<<>><<<<>>>><>>>><<<<><<<><><<<><<<<>><<<><<>><<<<>><<<<>>><<<><<<>><<>><<<>>><<><<>><<>>><<<>>>><>>><<<>>>><<>>><<<<>>>><><>>><<<<>>>><>>>><<<>>><<>>>><<<>><>><>><<<>>><<<>>>><<<>>><<<>>>><<>>>><<<><>><<>>><<<<>>>><><<<<><<>><>><>>><>>>><<>>><<<<>>><<<<><<<>>>><<<<>>>><<<<>>><>><>><<<<>>><<<<>>>><<>>>><<<<><>><<>>>><>>><<>><<<>>><<<>><<<><<<<>>>><<<<>><<>><<>>>><<<<>>>><>><<>>><<<<><<<><><<<>>><<>>><<<>>><<<>><<>>><>><<<>>>><<>>>><<<<>>>><<>><>>>><<<<><>>><<<<>>>><<>>>><<<<>>><<<<>>><<<<>><<>>><<>>>><<>>><<<><<<>>><<>>>><<<<><<>><<<<>>>><>>>><>><><<><<><<<<>><<<><<>>><>>><>>><<<<><<<<><<<<>>>><<>><>>><>>><<<><>>>><>>><>>><<<>><<<><<<>><<<<>>>><<<>>>><<<>><>><<>><<<>>>><<<<>>>><<<>>><><>><<><<<>><<<<>>><<>><<>><<<<>>><<<>>>><<>>>><<<<>>><<<>>><<<><>>>><<<<>><>><<<<>>>><<>>><<<>><><<<><<<<><>>>><>><><>><<<>>>><<<>>><<>>>><><<<<><<<>><<<<>>><<<><<><<<>><>>><><<<>>>><<<<>>>><<<>><>><<<>><<<>><<><<<<>>>><<>>><<<>>><>>>><<<>><<<<>>><<><>><<<<><<<<>><<<>><><>><<<><<<><><<<<>>>><<<><<<<><<<>>>><<>><<><><>>>><><>>><<>>><<<>><<<><<<<>>>><<<><<>><<<<>>><>><<><<>>>><<<>>><<<><<<>>><><<<<>>>><<><>><>>><<>>><<>>>><<>>><<>>>><>><><<<<>><>>><<<>>><<<>><<>>><<<<>><>>><<>>>><<<>>><>>><<<>><<<><<<>>><<><<<>>>><<>><<<<>>><<<>>><<<>>>><>><<<>>><<<<>>>><<><>>>><>>>><<<<>><<<>>>><>>>><<<>><<<<>><<><<>>>><<<<>>><<<<>>><<>><<<<>><<>>><<>>>><>><<>>><<>>>><>>>><>>><<<><<>><<<<>>>><<<>><<<<>>>><<<>><<<><<<>>><<<>><>>><<<<>>>><>>>><<<>><<<<>>><>>>><>><<><><><<<<>>>><<<<><>>><<>><<<>><<<>><<><<>><>><<<><<<>><<<>>><<<>><<>>><<<>>><<<<><>>><>>>><><<>>>><>>>><<<<>>><>>><<<<>>><>>>><>>>><<<<>>>><<>>>><<<>>><<>><<>>><<>>>><<>><>>>><<<>><<<><<><<<>><<<<>>><<>>>><<<<><>><<<>>>><<<>>>><<>><<>>>><<><>><>>><<<<>><<<>>>><<<>>>><<<<>>>><>>>><>>><<<<>>>><<<<><<>>>><<<<>>><<<<>><<>><>>><<<>><<<><<<<>>><><>><>>>><><<><<<<>>>><><<<>><<>>><<<<>>><<<<><>>><<>><>><<<><><>><<>>><<><><<><>><<>>>><<<<><<>><<<><<<><<>>><<<<><>><>>><<<><<<>><<><<<>><<<<>>>><<<>>><<>>>><><>><>>>><>>><>>>><>>><<<<>>>><<<>>><<<<>>>><>><<<<><<>><>><<>>><<<<>>>><<<>>><<<>><<<>><<>><<<<><>>>><>>>><<>>><<<><>><>><<<<>>><>>>><>><<>>>><<<<>><>>><<>>><<<<>>>><<><<<<>>>><<><<>><><>>>><>>><>>>><>>><<><>>><<>>><<<>>><<<<>>>><>>><<>><>><>>>><<>><<<>>><<<>>>><>><<>>><<>>><<><<>>>><<<<>>>><>><<<<>>>><<<<>>>><<>><<>>>><<<<>>>><<<<>><>>><>>>><<<>><>>>><<<><<>>>><<>>><<>>>><<><<<><<<<><<<>>><<>>><<<>>>><<><<<><><<>><<<<><<<<>>><<><<<<>><<<<>>>><<<>>>><<<>><<<<>>><>>><<<<>>><<<>>><<<<>><<<<>>><<>>><<<>><<<<>>><>>><<<>><<<<><<<<>>><<><<<<>>>><<<>>>><<>>><<><<<>>>><>>>><<>>>><>><<>><<>>>><<<<>>><<<>>>><<><<<><>><<<<>>><><>><<>>><<<>><<<>>><<<<><<<>>><<<<>>><<<<>><<><<<<>>>><<<><>><<>><>>>><>><<>><><<<>><<<>><<<<>>>><<<>>><<><<>>>><<>>>><<>>><<>><<<<>>>><<<<>><<>>>><<<>>>><<<<>>><<><<<<>>>><<<><>>><<<>><><>>>><><<<>>>><<<<>><<>><<<>><<><<><<<<>><><<<<>>><<><<>><<<>>>><<<>>><<>>><<<<>><<><<<<>>><<<<><><<<>>><<<><<<<>>><<<<>>><<<<>>><<<<>>><<>>><<<<>>><>><<<>><<<<>>>><><>>>><<><<><<<<>><><<<>>><><<<><<<>>>><<>><<<<>>><<<<><<<>>>><<>>>><<<<>><<<>>>><>>><<<>>>><<<>>>><<<>>><>>>><<<<>><>><<>><><<<<><<<>>>><<<>>><>>>><>>>><<<><<<>>><<>>>><>>><><<<<>><<>>><<<>><<>>><<<<><<<<>>>><<><<<><>>>><<<>>><<>>>><<<<>>>><<<<>>>><>><<<<>>><>>>><<>>>><<<><<>><<<>>><<<><<>>><<<<>>><<<<><<>>>><<>><<<>>><<<<>><<><<<<>>>><<><<>>><<<>>>><<<><<>><<<>>>><<><<>>>><>><<<<>>>><<>><>><<>>><<<><<<<><<<<>>><>><><>>><<<>>><<>><><<>><<<>>><<<>><<<<>>><<<<><<<><>><>>><<<<>>>><<<<>><<<<>><>>>><>>>><>>>><<<<>><>>>><<>><<<>>><<>><><>>>><<><<<<>><<>>><<>>><<>>><<<<>>>><<<<>><<<>>><<><<<<>><<<<>><<><>>>><<<>>>><<<><<>><<<<>>><<<<><>>>><<<<>>>><<>>>><<>><>>><>>><<><<<<>><<<<>><><<>><>><>>><<><<<>>>><<<>>><<<<>>>><<><<><>>><<<<><><<>>>><<<<>><<<<>><<>><<>>>><<>>>><<>>><<<>>>><<>>>><<<<>><>><>><<>>>><<<>>>><<<<>><<<>>><<>>><<<<>>>><>>>><<>>><<>>>><<<<><<>>><<<<>><<<>>>><><<<>>>><<<>>><<<<>>><<<<><<><<><<<<>>><<<>><>>><>><<<<>>><<<<>>><<><<<<>>><>>><<<>>><<>>>><<<>><<>>>><<><<<>><<<>>><<<<>>>><<<<>>>><>>><<>>><<<><<<>>>><<>>>><<><<<><>>><<>><<><<<<>>><>><>>><<<><><<><<<>>><<<><><<>>><<><<>><>>>><>><<<>>><<<>>>><<<>><<<<>>><>>>><<<<>><<<<><<<<><<<><>><<<<><<>>><<<<>>>><<<>>><<<>>>><>>>><<<<><<<<><<<<><<>><><<<<>>>><<<<>><<<>>>><<<>><<<>>>><<<>>>><<<><<>>><>><<<>>><<>><<<>>><>><<<>>><<<<>><>>>><>>>><><>>>><<<>>>><<<<><<><<><>><>>>><><<>><<<>>><<<<><>><<>><>>>><><<<>>>><<<>>>><<>>>><>>><>>>><<<>>><>><<<<><<<<>>>><<<<>><<<>>><<>>><<<<>>><<<>><>>><<><<<><<>>><<>><<>>><<<<>>><<<>>>><<<<><<<<>>>><<<>>>><<<<>>>><<><>><<<><><<>>>><><<<>>><<>>><<>>>><<<<>>><<<>>><<<<>>>><<>>><<<>><>>>><<>><<<><<><<<>><<<<>>><><>>>><><<><><<>>>><<>>>><>><><>>>><<<><<<<>>>><<<>><<<>><<<>>><<<>><<<>>><<<><<<>>>><<<<><<>><<<>>><<<<>>>><<<<>><<<<>>>><>>><<><<>>><><<<>><<><<<<>>><<<<>>><<<<>><<<>>><<<<><>><<>>><<<<>>><<<>><<<>><<<>>><><<<>><<>>><>><>><<>>><<<>>>><<<>>>><<>>><>><>><<>>><<>>>><<>>><><<<>><<>>><<>><<>><<<<>>>><>>>><<>><<<><<>><<<<>>>><<<<>><<>><>><<<>><<<<><<><<>>>><<>>>><<<>><<<<>>><<<<><<<>><<<>><>>>><>>>><<><>>>><<<>>><<<<>>><<<>><<<<>>><<<><<<>>><>>><<<<>><<>>>><<>><<<>><<<>>>><<<<>><<>>><<<<>><<>>>><<><<<>>><<<<>><>>>><<<<>>>><>>>><>>><<<>>><<<>>><>><>>>><<<<>>><<<<>>><<>>><>>><><<<>><<<<>>>><>>>><<<>><>>>><<>>>><<<><<>>><<<<>>>><<><>>><<<<>><<<>>><>>><<>>><<<><>>><<><<<>>>><<<<>><>>>><<><<>><>>><<>><>>>><<<<>>>><<<<>>><<>><<><><<<<>>>><<>>><<<>>>><<<>>><<<<>><<<<>><<>>>><<<<><<>>><<<><<<>><<<<>>><<<<>><<<>>><<><<>><<<<>>><<<<>>><<><>>><<<><<<<>>>><<<<>>>><>><>>><<<>>>><<<><<>>><<>><<<<>>><<<<><<<><<<>>><<>><>><>>>><>>>><><<<><><<<<><>>>><<>>><>>>><<>>><<<<>>>><<>>><>>>><<<>><><<>>>><<>>><<<>>><<<><<<>><<<>>><<<>>><<><<>><<<>>><<<<>>><<>>><>>><<<<>>><<>>>><<<><<<<>>><>><<<<>><>>><<>><>>>><<><<<<>>><<<<>><<>>><<<>>>><<<<>>><<<>><>>>><<<<>>><>>>><<>><>>><><>>><>>>><<<<>><<<><><<>><>><<<><<>>><<>>><<>><<<>>><<<<>><>><<<>><<<<>>>><<>>><>><<<<>>><<<><<>><>>>><<><<>>><<<>><><<<>>>><<<<><<<<>>>><<<>>><<<><>>><<<>>><<<<>><<<><<>>>><<<<>>>><<>><><<>>><<<<>>><>><<<<>>><>>>><>>><<<<>>><<><<<>>>><<<><<<>>><<<<>>><<<>>>><<<>><>><<<<>>>><<<<>><<<><>><<<><<>>>><<<>><<<<>>><<<<>>>><>><<<>>><<>>><<><<>><<<<>><<<<>><>><<<>>>><<<>>><<<>>><<>>>><<>>>><<>><<>><<>><<>><<>><<><>><>>><<<>>><<><<<<>>>><><<<>><<<<>>>><>><>>>><<><<<><>>"

	wind := make([]int, len(input))
	for i, r := range input {
		if r == '<' {
			wind[i] = -1
		} else if r == '>' {
			wind[i] = 1
		}
	}

	rockIndex := 0
	windIndex := 0

	world := rock{}
	height := 0
	current := rocks[rockIndex%len(rocks)].clone(0, height-4)
	rockIndex++

	limit := 1000000000000

	stateLog := make([]state, 0)

	for rockIndex <= limit {
		// blow
		next := current.clone(wind[windIndex%len(wind)], 0)
		windIndex++
		if !next.outOfBounds() && !next.collide(world) {
			current = next
		}

		// fall
		next = current.clone(0, 1)
		if next.outOfBounds() || next.collide(world) {
			// settle
			for s, b := range current {
				world[s] = b
				_, y := dekey(s)
				if y < height {
					height = y
				}

			}
			state{
				height: height,
				lastX:  world.hash(rowsToCompare, height),
			}
			current = rocks[rockIndex%len(rocks)].clone(0, height-4)
			rockIndex++
		} else {
			current = next
		}
	}

	fmt.Println("part1:", -height, "in", time.Since(start))

}
