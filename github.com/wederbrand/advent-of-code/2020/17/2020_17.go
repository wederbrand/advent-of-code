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
	z int
	w int
}

var offset = 7
var size = 20

func main() {
	readFile, err := ioutil.ReadFile("17/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	world := make(map[pos]bool)

	for y, s := range input {
		for x, r := range s {
			if r == '#' {
				world[pos{x + offset, y + offset, 0 + offset, 0 + offset}] = true
			}
		}
	}

	for i := 0; i < 6; i++ {
		world = expand(world)
	}

	cnt := 0
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			for z := 0; z < size; z++ {
				for w := 0; w < size; w++ {
					if world[pos{x, y, z, w}] {
						cnt++
					}
				}
			}
		}
	}

	fmt.Println(cnt)
}

func expand(world map[pos]bool) map[pos]bool {
	newWorld := make(map[pos]bool)
	for w := 0; w < size; w++ {
		for z := 0; z < size; z++ {
			for y := 0; y < size; y++ {
				for x := 0; x < size; x++ {
					p := pos{x, y, z, w}
					cnt := count(world, p)
					if world[p] {
						// rules for when it's active
						if cnt == 2 || cnt == 3 {
							newWorld[p] = true
						}
					} else {
						// rules for when it's inactive
						if cnt == 3 {
							newWorld[p] = true
						}
					}
				}
			}
		}
	}
	return newWorld
}

func count(world map[pos]bool, p pos) int {
	cnt := 0
	for x := p.x - 1; x <= p.x+1; x++ {
		for y := p.y - 1; y <= p.y+1; y++ {
			for z := p.z - 1; z <= p.z+1; z++ {
				for w := p.w - 1; w <= p.w+1; w++ {
					if x == p.x && y == p.y && z == p.z && w == p.w {
						continue
					}
					if world[pos{x, y, z, w}] {
						cnt++
					}
				}
			}
		}
	}

	return cnt
}
