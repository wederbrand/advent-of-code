package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	x, y, size, _ := getMaxGrid(7672)
	fmt.Println("Part 2:", x, y, size, "in", time.Since(start))
}

func getMaxGrid(serial int) (int, int, int, int) {
	grid := getGrid(serial)

	max := 0
	maxSize := 1

	for i := 1; i <= 300; i++ {
		_, _, _, power := checkGrid(grid, i)

		if power > max {
			max = power
			maxSize = i
		}
	}

	return checkGrid(grid, maxSize)
}

func checkGrid(grid *[300][300]int, size int) (int, int, int, int) {
	max := math.MinInt64
	maxX := 0
	maxY := 0
	for x := 0; x < 300-size+1; x++ {
		for y := 0; y < 300-size+1; y++ {
			sum := 0
			for i := 0; i < size; i++ {
				for j := 0; j < size; j++ {
					sum += grid[x+i][y+j]
				}
			}

			if sum > max {
				max = sum
				maxX = x
				maxY = y
			}
		}
	}

	return maxX + 1, maxY + 1, size, max
}

func getGrid(serial int) *[300][300]int {
	var grid [300][300]int
	for x := 0; x < 300; x++ {
		for y := 0; y < 300; y++ {
			grid[x][y] = getPowerLevel(x+1, y+1, serial)
		}
	}
	return &grid
}

func getPowerLevel(x int, y int, serial int) int {
	result := x + 10
	result *= y
	result += serial
	result *= x + 10
	result /= 100
	result %= 10
	result -= 5

	return result
}
