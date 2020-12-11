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

	//// add floor around the outer edges
	//seats := make([]string, 0)
	//seats = append(seats, "...")
	//for _, s := range input {
	//	row := "." + s + "."
	//	seats = append(seats, row)
	//}
	//seats = append(seats, "...")

	steps := 0
	for {
		steps++
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

	// 125 is too low
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
				if surrounding >= 4 {
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
	for i := y - 1; i <= y+1; i++ {
		if i < 0 {
			continue
		}
		if i == len(seats) {
			continue
		}
		for j := x - 1; j <= x+1; j++ {
			if j < 0 {
				continue
			}
			if j == len(seats[i]) {
				continue
			}
			if i == y && j == x {
				continue
			}
			if seats[i][j] == '#' {
				count++
			}
		}
	}

	return count
}
