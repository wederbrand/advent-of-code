package main

import (
	"fmt"
	"strings"
	"time"

	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
)

func main() {
	start := time.Now()
	input := GetFileContents("2025/12/input.txt", "\n")

	bradgards := make([]int, 0)
	currentShapeStart := 0
	regions := []string{}
	for i, s := range input {
		if strings.Contains(s, "x") {
			// region
			regions = append(regions, s)
		} else if strings.Contains(s, ":") {
			// shape start
			currentShapeStart = i + 1
		} else if s == "" {
			// shape end
			chart := MakeChart(input[currentShapeStart:i], ".")
			bradgards = append(bradgards, len(chart))
		}
	}

	part1 := 0
	for _, region := range regions {
		splits := strings.Split(region, " ")

		var x, y int
		fmt.Sscanf(splits[0], "%dx%d:", &x, &y)

		totalShapes := 0
		totalBradgards := 0
		for i := 1; i < len(splits); i++ {
			atoi := Atoi(strings.TrimSpace(splits[i]))
			totalShapes += atoi
			totalBradgards += bradgards[i-1] * atoi
		}

		area := x * y
		if totalBradgards > area {
			// impossible, more # than fits
			continue
		}

		worstCase := x / 3 * y / 3
		if worstCase >= totalShapes {
			// even shapes with 3x3 # would fit
			part1++
			continue
		}

		panic("ho no")
	}
	fmt.Println("Part 1:", part1, "in", time.Since(start))
}
