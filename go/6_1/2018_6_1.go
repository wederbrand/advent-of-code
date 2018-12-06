package main

import (
	"fmt"
	"strconv"
	"regexp"
	"os"
	"bufio"
)

var world [500][500]int
var world2 [500][500]int
var inf = make(map[int]bool)
var BLOCKED = 999

func grow(x int, y int) {
	marker := world[x][y]

	if (marker == 0) {
		return
	} else if (marker == BLOCKED) {
		return
	}

	mark(x-1, y, marker)
	mark(x+1, y, marker)
	mark(x, y-1, marker)
	mark(x, y+1, marker)
}

func mark(x int, y int, marker int) {
	if (x < 0) {
		inf[marker] = true
		return
	}
	if (x > len(world)-1) {
		inf[marker] = true
		return
	}
	if (y < 0) {
		inf[marker] = true
		return
	}
	if (y > len(world)-1) {
		inf[marker] = true
		return
	}
	
	// this point can be
	if (world2[x][y] > 0 && world2[x][y] != marker) {
		// - marked by another, mark with BLOCKED
		world2[x][y] = BLOCKED
	} else if (world2[x][y] == 0 && world[x][y] == 0) {
		// - taken by nothing, mark it	
		world2[x][y] = marker
	}
}

func resolve(x int, y int) {
	// this point can be
	if (world2[x][y] == BLOCKED) {
		// - marked with BLOCKED, BLOCK it in world
		world[x][y] = BLOCKED
	} else if (world2[x][y] > 0) {
		// - marked, transfer to world
		world[x][y] = world2[x][y]
	}	

	world2[x][y] = 0
}

func main() {
	file, _ := os.Open("2018_6.input")
	scanner := bufio.NewScanner(file)

	lineMatcher := regexp.MustCompile(`^(\d+), (\d+)$`)

	marker := 0
	for scanner.Scan() {
		marker += 1
		lineMatch := lineMatcher.FindStringSubmatch(scanner.Text())
		x, _ := strconv.Atoi(lineMatch[1])
		y, _ := strconv.Atoi(lineMatch[2])
		world[x][y] = marker
	}

	for {
		// phase 1 - write candidates to world 2
		for x, _ := range(world) {
			for y, _ := range(world[x]) {
				grow(x, y)
			}
		}

		// phase 1.5 - abort if none where marked
		found := false
		for x, _ := range(world) {
			for y, _ := range(world[x]) {
				if (world2[x][y] != 0) {
					found = true
					break
				}				
			}
		}

		if !found {
			break
		}
	
		// phase 2 - resolve candidates back to world
		for x, _ := range(world) {
			for y, _ := range(world[x]) {
				resolve(x, y)
			}
		}		
	}

	// phase 3 - count
	count := make(map[int]int)
	max := 0
	for x, _ := range(world) {
		for y, _ := range(world[x]) {
			marker := world[x][y]
			if !inf[marker] {
				count[marker]++
				if count[marker] > max {
					max = count[marker]
				}
			}
		}
	}		

	fmt.Println(inf)
	fmt.Println(max)
}