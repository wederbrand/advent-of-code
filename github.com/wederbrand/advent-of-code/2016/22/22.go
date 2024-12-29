package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"time"
)

type Node struct {
	size  int
	used  int
	avail int
}

func main() {
	start := time.Now()
	inFile := GetFileContents("2016/22/input.txt", "\n")

	maxX := 0
	maxY := 0
	nodes := make(map[Coord]Node)
	for _, line := range inFile {
		if line[0] != '/' {
			continue
		}

		c := Coord{}
		n := Node{}
		fmt.Sscanf(line, "/dev/grid/node-x%d-y%d %dT %dT %dT", &c.X, &c.Y, &n.size, &n.used, &n.avail)
		nodes[c] = n
		if c.X > maxX {
			maxX = c.X
		}
		if c.Y > maxY {
			maxY = c.Y
		}
	}

	part1 := 0
	for _, n1 := range nodes {
		for _, n2 := range nodes {
			if n1 == n2 {
				continue
			}
			if n1.used > 0 && n1.used <= n2.avail {
				part1++
			}
		}
	}
	fmt.Println("Part 1: ", part1, "in", time.Since(start))

	dataC := Coord{}
	freeC := Coord{}

	for c, n := range nodes {
		if n.used == 0 {
			freeC = c
		}
		if c.Y == 0 && c.X == maxX {
			dataC = c
		}
	}

	printIt(nodes, maxX, maxY, dataC, freeC)

	// from looking at the print we see that I need to move
	moves := 0
	// freeC all the way to the left (X=0)
	for freeC.X != 0 {
		freeC.X--
		moves++
	}

	// then all the way to the top (Y=0)
	for freeC.Y != 0 {
		freeC.Y--
		moves++
	}

	// then right stopping just before dataC
	for freeC.X != dataC.X-1 {
		freeC.X++
		moves++
	}

	// then a swap
	freeC, dataC = dataC, freeC
	moves++

	// from there I can move the dataC one step by moving the freeC down, left x3, up and right
	for dataC.X != 0 {
		dataC.X--
		moves += 5
	}

	fmt.Println("Part 2: ", moves, "in", time.Since(start))
}

func printIt(nodes map[Coord]Node, maxX int, maxY int, dataC Coord, freeC Coord) {
	freeN := nodes[freeC]

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			n := nodes[Coord{X: x, Y: y}]
			if n.used == 0 {
				fmt.Print("_")
			} else if dataC.X == x && dataC.Y == y {
				fmt.Print("G")
			} else if n.used > freeN.size {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
