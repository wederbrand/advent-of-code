package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	row := 2981
	col := 3075

	code := getcode(row, col)

	// then move up to the target row
	fmt.Println("Part 1: ", code, "in", time.Since(start))
}

func getcode(row int, col int) int {
	code := 20151125
	currentRow := 1
	currentCol := 1

	for {
		currentRow = currentCol + 1
		currentCol = 1
		for {
			code = (code * 252533) % 33554393
			if currentRow == row && currentCol == col {
				return code
			}
			currentRow--
			if currentRow == 0 {
				break
			}
			currentCol++
		}
	}
	panic("ho ho ho")
}
