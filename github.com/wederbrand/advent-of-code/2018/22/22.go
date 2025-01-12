package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/priorityqueue"
	"time"
)

type cave struct {
	geoIndex int
	erosion  int
	region   string
}

func main() {
	start := time.Now()
	inFile := GetFileContents("2018/22/input.txt", "\n")

	depth := 0
	fmt.Sscanf(inFile[0], "depth: %d", &depth)
	target := Coord{0, 0}
	fmt.Sscanf(inFile[1], "target: %d,%d", &target.X, &target.Y)

	m := map[Coord]cave{}
	entrance := Coord{0, 0}

	part1 := 0
	for y := 0; y <= target.Y; y++ {
		for x := 0; x <= target.X; x++ {
			c := Coord{x, y}
			caveRoom := getRoom(c, m, depth, entrance, target)
			m[c] = caveRoom
			switch caveRoom.region {
			case ".":
				part1 += 0
			case "=":
				part1 += 1
			case "|":
				part1 += 2
			}
		}
	}

	fmt.Println("Part 1:", part1, "in", time.Since(start))

	part2 := findTarget(m, depth, entrance, target)
	fmt.Println("Part 2:", part2, "in", time.Since(start))
}

type CaveState struct {
	coord Coord
	tool  string
}

func findTarget(m map[Coord]cave, depth int, entrance Coord, target Coord) int {
	q := priorityqueue.NewQueue()
	q.Add(&priorityqueue.State{Data: CaveState{entrance, "torch"}, Priority: 0})

	seen := map[CaveState]int{}
	for q.HasNext() {
		s := q.Next()
		cs := s.Data.(CaveState)
		c := cs.coord
		currentRoom := m[c]
		t := cs.tool

		if c == target && t == "torch" {
			return s.Priority
		}

		if val, ok := seen[cs]; ok && val <= s.Priority {
			continue
		}
		seen[cs] = s.Priority

		// for all directions, and change tool if needed
		for _, d := range ALL {
			next := c.Move(d)
			if next.X < 0 || next.Y < 0 {
				continue
			}

			if _, ok := m[next]; !ok {
				m[next] = getRoom(next, m, depth, entrance, target)
			}

			nextRoom := m[next]

			if nextRoom.region == "." {
				// needs climbing or torch
				if t == "torch" || t == "climbing" {
					q.Add(&priorityqueue.State{Data: CaveState{next, t}, Priority: s.Priority + 1})
				}
			} else if nextRoom.region == "=" {
				// needs climbing or neither
				if t == "climbing" || t == "neither" {
					q.Add(&priorityqueue.State{Data: CaveState{next, t}, Priority: s.Priority + 1})
				}
			} else if nextRoom.region == "|" {
				// needs torch or neither
				if t == "torch" || t == "neither" {
					q.Add(&priorityqueue.State{Data: CaveState{next, t}, Priority: s.Priority + 1})
				}
			}
		}

		// stay and change tool
		if currentRoom.region == "." {
			if t == "torch" {
				q.Add(&priorityqueue.State{Data: CaveState{c, "climbing"}, Priority: s.Priority + 7})
			} else if t == "climbing" {
				q.Add(&priorityqueue.State{Data: CaveState{c, "torch"}, Priority: s.Priority + 7})
			}
		} else if currentRoom.region == "=" {
			if t == "climbing" {
				q.Add(&priorityqueue.State{Data: CaveState{c, "neither"}, Priority: s.Priority + 7})
			} else if t == "neither" {
				q.Add(&priorityqueue.State{Data: CaveState{c, "climbing"}, Priority: s.Priority + 7})
			}
		} else if currentRoom.region == "|" {
			if t == "torch" {
				q.Add(&priorityqueue.State{Data: CaveState{c, "neither"}, Priority: s.Priority + 7})
			} else if t == "neither" {
				q.Add(&priorityqueue.State{Data: CaveState{c, "torch"}, Priority: s.Priority + 7})
			}
		}
	}
	panic("No path found")
}

func getRoom(c Coord, m map[Coord]cave, depth int, entrance Coord, target Coord) cave {
	geoIndex := 0
	if c == entrance || c == target {
		geoIndex = 0
	} else if c.Y == 0 {
		geoIndex = c.X * 16807
	} else if c.X == 0 {
		geoIndex = c.Y * 48271
	} else {
		if _, ok := m[c.Move(LEFT)]; !ok {
			m[c.Move(LEFT)] = getRoom(c.Move(LEFT), m, depth, entrance, target)
		}
		if _, ok := m[c.Move(UP)]; !ok {
			m[c.Move(UP)] = getRoom(c.Move(UP), m, depth, entrance, target)
		}
		geoIndex = m[c.Move(LEFT)].erosion * m[c.Move(UP)].erosion
	}
	erosion := (geoIndex + depth) % 20183
	switch erosion % 3 {
	case 0:
		return cave{geoIndex, erosion, "."}
	case 1:
		return cave{geoIndex, erosion, "="}
	case 2:
		return cave{geoIndex, erosion, "|"}
	}
	panic("Unknown region")
}
