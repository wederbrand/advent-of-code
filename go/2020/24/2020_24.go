package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type pos struct {
	x int
	y int
}

var minX int
var maxX int
var minY int
var maxY int

func (p *pos) move(dir string) {
	switch dir {
	case "ne":
		p.y--
	case "nw":
		p.y--
		p.x--
	case "se":
		p.y++
		p.x++
	case "sw":
		p.y++
	case "w":
		p.x--
	case "e":
		p.x++
	}
}

func main() {
	readFile, err := ioutil.ReadFile("24/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	floor := make(map[pos]bool)
	input := strings.Split(strings.TrimSpace(string(readFile)), "\n")
	for _, inputTile := range input {
		p := new(pos)
		for i := 0; i < len(inputTile); i++ {
			r := inputTile[i]
			switch r {
			case 'n':
				i++
				r2 := inputTile[i]
				switch r2 {
				case 'e':
					p.move("ne")
				case 'w':
					p.move("nw")
				}
			case 's':
				i++
				r2 := inputTile[i]
				switch r2 {
				case 'e':
					p.move("se")
				case 'w':
					p.move("sw")
				}
			case 'e':
				p.move("e")
			case 'w':
				p.move("w")
			}
		}
		flip(p, floor)
		fmt.Println(p, floor[*p])
	}

	cnt := countBlackOnes(floor)
	fmt.Println("part 1:", cnt)

	for i := 0; i < 100; i++ {
		newFloor := make(map[pos]bool)
		for y := minY - 2; y < maxY+2; y++ {
			for x := minX - 2; x < maxX+2; x++ {
				p := pos{x, y}
				cnt := countNeighbours(p, floor)
				if floor[p] {
					// black tile
					if cnt == 0 || cnt > 2 {
						// black -> white
					} else {
						// black -> black
						flip(&p, newFloor)
					}

				} else {
					// white tile
					if cnt == 2 {
						// white -> black
						flip(&p, newFloor)
					} else {
						// white -> white
					}
				}
			}
		}
		floor = newFloor
		blackOnes := countBlackOnes(floor)
		fmt.Println(i+1, blackOnes)
	}
}

func countNeighbours(p pos, floor map[pos]bool) int {
	cnt := 0
	if floor[pos{p.x, p.y - 1}] {
		cnt++
	}
	if floor[pos{p.x + 1, p.y}] {
		cnt++
	}
	if floor[pos{p.x + 1, p.y + 1}] {
		cnt++
	}
	if floor[pos{p.x, p.y + 1}] {
		cnt++
	}
	if floor[pos{p.x - 1, p.y}] {
		cnt++
	}
	if floor[pos{p.x - 1, p.y - 1}] {
		cnt++
	}

	return cnt
}

func flip(p *pos, floor map[pos]bool) {
	if p.x < minX {
		minX = p.x
	}
	if p.x > maxX {
		maxX = p.x
	}
	if p.y < minY {
		minY = p.y
	}
	if p.y > maxY {
		maxY = p.y
	}
	floor[*p] = !floor[*p]
}

func countBlackOnes(floor map[pos]bool) int {
	cnt := 0
	for _, black := range floor {
		if black {
			cnt++
		}
	}
	return cnt
}
