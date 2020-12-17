package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	readFile, err := ioutil.ReadFile("17/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	offset := 7
	var world [20][20][20]bool

	for y, s := range input {
		for x, r := range s {
			if r == '#' {
				world[x+offset][y+offset][offset] = true
			}
		}
	}

	printWorld(world)
	for i := 0; i < 6; i++ {
		world = expand(world)
		fmt.Println("one iteration", "")
		printWorld(world)
	}

	cnt := 0
	for x := 0; x < len(world); x++ {
		for y := 0; y < len(world); y++ {
			for z := 0; z < len(world); z++ {
				if world[x][y][z] {
					cnt++
				}
			}
		}
	}

	fmt.Println(cnt)
}

func printWorld(world [20][20][20]bool) {
	for x := 0; x < len(world); x++ {
		for y := 0; y < len(world); y++ {
			for z := 0; z < len(world); z++ {
				if world[x][y][z] {
					fmt.Println(x, y, z)
				}
			}
		}
	}
}

func expand(world [20][20][20]bool) [20][20][20]bool {
	var newWorld [20][20][20]bool
	for z := 0; z < len(newWorld); z++ {
		for y := 0; y < len(newWorld); y++ {
			for x := 0; x < len(newWorld); x++ {
				cnt := count(world, x, y, z)
				if world[x][y][z] {
					// rules for when it's active
					if cnt == 2 || cnt == 3 {
						newWorld[x][y][z] = true
					}
				} else {
					// rules for when it's inactive
					if cnt == 3 {
						newWorld[x][y][z] = true
					}
				}
			}
		}
	}

	return newWorld
}

func count(world [20][20][20]bool, x int, y int, z int) int {
	cnt := 0
	for xx := x - 1; xx <= x+1; xx++ {
		if xx < 0 || xx >= len(world) {
			continue
		}
		for yy := y - 1; yy <= y+1; yy++ {
			if yy < 0 || yy >= len(world) {
				continue
			}
			for zz := z - 1; zz <= z+1; zz++ {
				if zz < 0 || zz >= len(world) {
					continue
				}
				if xx == x && yy == y && zz == z {
					continue
				}
				if world[xx][yy][zz] {
					cnt++
				}
			}
		}
	}

	return cnt
}
