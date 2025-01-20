package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"time"
)

func main() {
	start := time.Now()

	inFile := GetFileContents("2019/24/input.txt", "\n")

	m := MakeChart(inFile, "")

	seen := make(map[int]bool)
	for {
		if seen[hash(m)] {
			break
		}
		seen[hash(m)] = true

		m = doIt(m)
	}

	part1 := 0
	score := 1
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if m[Coord{X: x, Y: y}] == "#" {
				part1 += score
			}

			score *= 2
		}
	}
	fmt.Println("part1: ", part1, "in", time.Since(start))

	m = MakeChart(inFile, "")

	layers := make([]Chart, 1+200*2+2)

	for i := 0; i < len(layers); i++ {
		layers[i] = Chart{}
		for y := 0; y < 5; y++ {
			for x := 0; x < 5; x++ {
				if x == 2 && y == 2 {
					layers[i][Coord{X: x, Y: y}] = "?"
				} else {
					layers[i][Coord{X: x, Y: y}] = "."
				}
			}
		}
	}

	layers[201] = m
	for i := 0; i < 200; i++ {
		layers = doIt2(layers)
	}

	part2 := 0
	for i := 0; i < len(layers); i++ {
		for y := 0; y < 5; y++ {
			for x := 0; x < 5; x++ {
				if layers[i][Coord{X: x, Y: y}] == "#" {
					part2++
				}
			}
		}
	}

	fmt.Println("part2: ", part2, "in", time.Since(start))
}

func hash(m Chart) int {
	out := 0
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			out <<= 1
			if m[Coord{X: x, Y: y}] == "#" {
				out += 1
			}
		}
	}
	return out
}

func doIt(m Chart) Chart {
	nextM := Chart{}
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			c := Coord{X: x, Y: y}
			adjacentBugs := 0
			for _, dir := range ALL {
				if m[c.Move(dir)] == "#" {
					adjacentBugs++
				}
			}
			if m[c] == "#" && adjacentBugs == 1 {
				nextM[c] = "#"
			} else if m[c] == "." && (adjacentBugs == 1 || adjacentBugs == 2) {
				nextM[c] = "#"
			} else {
				nextM[c] = "."
			}
		}
	}

	return nextM
}

func doIt2(l []Chart) []Chart {
	nextL := make([]Chart, len(l))
	nextL[0] = Chart{}
	nextL[len(l)-1] = Chart{}

	for i := 1; i < len(l)-1; i++ {
		nextL[i] = Chart{}
		for y := 0; y < 5; y++ {
			for x := 0; x < 5; x++ {
				if x == 2 && y == 2 {
					nextL[i][Coord{X: x, Y: y}] = "?"
					continue
				}
				adjacentBugs := 0
				c := Coord{X: x, Y: y}
				for _, dir := range ALL {
					if c.Move(dir).X == -1 {
						// check level -1 square 1,2
						if l[i-1][Coord{X: 1, Y: 2}] == "#" {
							adjacentBugs++
						}
					} else if c.Move(dir).X == 5 {
						// check level -1 square 3,2
						if l[i-1][Coord{X: 3, Y: 2}] == "#" {
							adjacentBugs++
						}
					} else if c.Move(dir).Y == -1 {
						// check level -1 square 2,1
						if l[i-1][Coord{X: 2, Y: 1}] == "#" {
							adjacentBugs++
						}
					} else if c.Move(dir).Y == 5 {
						// check level -1 square 2,3
						if l[i-1][Coord{X: 2, Y: 3}] == "#" {
							adjacentBugs++
						}
					} else if c.Move(dir).X == 2 && c.Move(dir).Y == 2 {
						if dir == UP {
							// check level +1 row 5
							for x := 0; x < 5; x++ {
								if l[i+1][Coord{X: x, Y: 4}] == "#" {
									adjacentBugs++
								}
							}
						} else if dir == LEFT {
							// check level +1 column 5
							for y := 0; y < 5; y++ {
								if l[i+1][Coord{X: 4, Y: y}] == "#" {
									adjacentBugs++
								}
							}
						} else if dir == DOWN {
							// check level +1 row 0
							for x := 0; x < 5; x++ {
								if l[i+1][Coord{X: x, Y: 0}] == "#" {
									adjacentBugs++
								}
							}
						} else if dir == RIGHT {
							// check level +1 column 0
							for y := 0; y < 5; y++ {
								if l[i+1][Coord{X: 0, Y: y}] == "#" {
									adjacentBugs++
								}
							}
						} else {
							panic("inner fuckery")
						}
					} else {
						if l[i][c.Move(dir)] == "#" {
							adjacentBugs++
						}
					}
				}
				if l[i][Coord{X: x, Y: y}] == "#" && adjacentBugs == 1 {
					nextL[i][Coord{X: x, Y: y}] = "#"
				} else if l[i][Coord{X: x, Y: y}] == "." && (adjacentBugs == 1 || adjacentBugs == 2) {
					nextL[i][Coord{X: x, Y: y}] = "#"
				} else {
					nextL[i][Coord{X: x, Y: y}] = "."
				}
			}
		}
	}

	return nextL
}
