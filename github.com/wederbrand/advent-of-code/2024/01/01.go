package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"math"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2016/01/input.txt", ", ")

	c := chart.Coord{0, 0}
	d := chart.N

	seen := make(map[chart.Coord]bool)
	seen[c] = true
	printed := false
	for _, line := range inFile {
		rotation := line[0]
		steps := util.Atoi(line[1:])
		if rotation == 'R' {
			d = d.Right()
		} else {
			d = d.Left()
		}
		for i := 0; i < steps; i++ {
			c = c.Move(d, 1)
			if !printed && seen[c] {
				fmt.Println("Part 2:", math.Abs(float64(c.X))+math.Abs(float64(c.Y)), "in", time.Since(start))
				printed = true
			}
			seen[c] = true
		}
	}

	fmt.Println("Part 1:", math.Abs(float64(c.X))+math.Abs(float64(c.Y)), "in", time.Since(start))
}
