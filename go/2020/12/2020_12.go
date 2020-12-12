package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	readFile, err := ioutil.ReadFile("12/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	y := 0
	x := 0
	dy := -1
	dx := 10

	for _, s := range input {
		command := s[0]
		atoi, _ := strconv.Atoi(s[1:])
		if command == 'F' {
			y += dy * atoi
			x += dx * atoi
		}
		switch command {
		case 'L':
			dy, dx = rotate(-atoi, dy, dx)
		case 'R':
			dy, dx = rotate(atoi, dy, dx)
		case 'N':
			dy -= atoi
		case 'S':
			dy += atoi
		case 'E':
			dx += atoi
		case 'W':
			dx -= atoi
		}
	}

	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func rotate(rotation int, dy int, dx int) (int, int) {
	for rotation < 0 {
		rotation += 360
	}
	rotation %= 360
	switch rotation {
	case 0:
		return dy, dx
	case 90:
		// east -> south
		// south -> west
		return dx, -dy
	case 180:
		// east -> west
		// south -> north
		return -dy, -dx
	case 270:
		// east -> north
		// south -> east
		return -dx, dy
	}
	panic("at the disco")
}
