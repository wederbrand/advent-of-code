package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2024/08/input.txt", "\n")
	m := MakeChart(inFile, ".")
	maxC := Coord{len(inFile[0]) - 1, len(inFile) - 1}

	p1, p2 := doIt(m, maxC.X, maxC.Y)

	fmt.Println("Part 1:", p1, "in", time.Since(start))
	fmt.Println("Part 2:", p2, "in", time.Since(start))
}

func doIt(m Chart, maxX int, maxY int) (int, int) {
	part1 := make(Chart)
	part2 := make(Chart)
	for a := range m {
		if m[a] == "." {
			continue
		}
		for b := range m {
			if m[b] == "." {
				continue
			}
			if a == b {
				continue
			}
			if m[a] != m[b] {
				continue
			}

			// we only need to get the vector between these points and extrapolate it once.
			// the other direction will be covered by finding the same two points, in the other order.
			dX := b.X - a.X
			dY := b.Y - a.Y

			gcd := util.Gcd(dX, dY)
			gcd = util.IntAbs(gcd)

			dX /= gcd
			dY /= gcd

			antiNode := Coord{a.X, a.Y}
			for antiNode.X >= 0 && antiNode.Y >= 0 && antiNode.X <= maxX && antiNode.Y <= maxY {
				if util.IntAbs(antiNode.X-a.X) == 2*util.IntAbs(antiNode.X-b.X) {
					part1[antiNode] = "#"
				}
				part2[antiNode] = "#"
				antiNode.X += dX
				antiNode.Y += dY
			}
		}
	}
	return len(part1), len(part2)
}
