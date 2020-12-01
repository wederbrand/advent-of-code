package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(getPowerLevel(3, 5, 8))
	//fmt.Println(getPowerLevel(122, 79, 57))
	//fmt.Println(getPowerLevel(217, 196, 39))
	//fmt.Println(getPowerLevel(101, 153, 71))
	//
	//fmt.Println(checkGrid(getGrid(18), 3))
	//fmt.Println(checkGrid(getGrid(42), 3))
	//
	//fmt.Println("part1")
	//fmt.Println(checkGrid(getGrid(7672), 3))
	//
	//fmt.Println(getMaxGrid(18))
	//fmt.Println(getMaxGrid(42))
	fmt.Println("part 2")
	fmt.Println(getMaxGrid(7672))
}

func getMaxGrid(serial int) (int, int, int, int) {
	grid := getGrid(serial)

	max := 0
	maxSize := 1

	for i := 1; i <= 300; i++ {
		x, y, _, power := checkGrid(grid, i)
		fmt.Println("checking",  i, x, y, power)

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
