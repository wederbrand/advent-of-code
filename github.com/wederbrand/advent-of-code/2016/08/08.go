package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2016/08/input.txt", "\n")
	screen := doIt(inFile)
	fmt.Println("Part 1: ", len(screen), "in", time.Since(start))

	chart.PrintChart(screen)
}

func doIt(inFile []string) chart.Chart {
	screen := chart.Chart{}

	for _, line := range inFile {
		if strings.HasPrefix(line, "rect") {
			var width, height int
			fmt.Sscanf(line, "rect %dx%d", &width, &height)
			// turn on pixels
			for x := 0; x < width; x++ {
				for y := 0; y < height; y++ {
					screen[chart.Coord{x, y}] = "#"
				}
			}
		} else if strings.HasPrefix(line, "rotate row") {
			var row, amount int
			fmt.Sscanf(line, "rotate row y=%d by %d", &row, &amount)
			// rotate row, go backwards
			pixelsToShift := make([]chart.Coord, 0)
			for x := 0; x <= 50; x++ {
				pixel := chart.Coord{x, row}
				_, found := screen[pixel]
				if found {
					delete(screen, pixel)
					pixel.X = (pixel.X + amount) % 50
					pixelsToShift = append(pixelsToShift, pixel)
				}
			}
			for _, pixel := range pixelsToShift {
				screen[pixel] = "#"
			}
		} else if strings.HasPrefix(line, "rotate column") {
			var col, amount int
			fmt.Sscanf(line, "rotate column x=%d by %d", &col, &amount)
			// rotate row, go backwards
			pixelsToShift := make([]chart.Coord, 0)
			for y := 0; y <= 6; y++ {
				pixel := chart.Coord{col, y}
				_, found := screen[pixel]
				if found {
					delete(screen, pixel)
					pixel.Y = (pixel.Y + amount) % 6
					pixelsToShift = append(pixelsToShift, pixel)
				}
			}
			for _, pixel := range pixelsToShift {
				screen[pixel] = "#"
			}
		}
	}

	return screen
}
