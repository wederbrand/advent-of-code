package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	readFile, err := ioutil.ReadFile("11/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	seats := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	for {
		newSeats, changes := stepOne(seats)
		seats = newSeats
		if changes == 0 {
			break
		}
	}

	taken := 0
	for _, seat := range seats {
		taken += strings.Count(seat, "#")
	}

	fmt.Println(taken)
}

func stepOne(seats []string) ([]string, int) {
	newSeats := make([]string, 0)
	diff := 0
	for y, row := range seats {
		newSeats = append(newSeats, "")
		for x, seat := range row {
			surrounding := countOccupied(seats, y, x)
			switch seat {
			case '.':
				newSeats[y] += "."
			case 'L':
				if surrounding == 0 {
					newSeats[y] += "#"
					diff++
				} else {
					newSeats[y] += "L"
				}
			case '#':
				if surrounding >= 5 {
					newSeats[y] += "L"
					diff++
				} else {
					newSeats[y] += "#"
				}
			}
		}
	}

	return newSeats, diff
}

func countOccupied(seats []string, y int, x int) int {
	count := 0
	count += countLine(seats, y, x, -1, 0)  // N
	count += countLine(seats, y, x, -1, 1)  // NE
	count += countLine(seats, y, x, 0, 1)   // E
	count += countLine(seats, y, x, +1, 1)  // SE
	count += countLine(seats, y, x, +1, 0)  // S
	count += countLine(seats, y, x, +1, -1) // SW
	count += countLine(seats, y, x, 0, -1)  // W
	count += countLine(seats, y, x, -1, -1) // NW

	return count
}

func countLine(seats []string, y int, x int, dy int, dx int) int {
	for {
		y += dy
		x += dx
		if y < 0 || x < 0 {
			return 0
		}
		if y == len(seats) || x == len(seats[y]) {
			return 0
		}

		if seats[y][x] == 'L' {
			return 0
		}

		if seats[y][x] == '#' {
			return 1
		}
	}
}
