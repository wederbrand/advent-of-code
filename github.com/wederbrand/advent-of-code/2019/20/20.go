package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/priorityqueue"
	"time"
	"unicode"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2019/20/input.txt", "\n")

	m := MakeChart(inFile, "")

	minC, maxC := GetChartMaxes(m)

	teleporters := make(map[string][]Coord)

	for y := minC.Y; y <= maxC.Y; y++ {
		for x := minC.X; x <= maxC.X; x++ {
			c := Coord{X: x, Y: y}

			s, found := m[c]
			if !found {
				continue
			}
			r := rune(s[0])
			if !unicode.IsLetter(r) {
				continue
			}

			if m[c.Move(UP)] != "." && m[c.Move(LEFT)] != "." && m[c.Move(RIGHT)] != "." && m[c.Move(DOWN)] != "." {
				continue
			}

			// we know we're the "middle letter" in a portal

			var up, left, down, right rune
			upS := m[c.Move(UP)]
			if upS != "" {
				up = rune(upS[0])
			}
			leftS := m[c.Move(LEFT)]
			if leftS != "" {
				left = rune(leftS[0])
			}
			downS := m[c.Move(DOWN)]
			if downS != "" {
				down = rune(downS[0])
			}
			rightS := m[c.Move(RIGHT)]
			if rightS != "" {
				right = rune(rightS[0])
			}

			// top row, the next below is a dot and the one above is a letter
			if down == '.' && unicode.IsLetter(up) {
				teleporters[string(up)+string(r)] = append(teleporters[string(up)+string(r)], c.Move(DOWN))
			}

			// bottom row, the next above is a dot and the one below is a letter
			if up == '.' && unicode.IsLetter(down) {
				teleporters[string(r)+string(down)] = append(teleporters[string(r)+string(down)], c.Move(UP))
			}

			// left column, the next to the right is a letter and the one to the right is a dot
			if right == '.' && unicode.IsLetter(left) {
				teleporters[string(left)+string(r)] = append(teleporters[string(left)+string(r)], c.Move(RIGHT))
			}

			// right column, the next to the left is a letter and the one to the left is a dot
			if left == '.' && unicode.IsLetter(right) {
				teleporters[string(r)+string(right)] = append(teleporters[string(r)+string(right)], c.Move(LEFT))
			}
		}
	}

	inPortals := make(map[Coord]Coord)
	outPortals := make(map[Coord]Coord)
	minC, maxC = GetChartMaxes(m)
	startC := Coord{}
	endC := Coord{}
	for name, coords := range teleporters {
		if name == "AA" {
			startC = coords[0]
		} else if name == "ZZ" {
			endC = coords[0]
		} else {
			if coords[0].X == minC.X+2 || coords[0].Y == minC.Y+2 || coords[0].X == maxC.X-2 || coords[0].Y == maxC.Y-2 {
				// this is an outer portal
				outPortals[coords[0]] = coords[1]
				inPortals[coords[1]] = coords[0]
			} else {
				// this is an inner portal
				inPortals[coords[0]] = coords[1]
				outPortals[coords[1]] = coords[0]
			}
		}
	}

	part1 := getShortest(m, startC, endC, inPortals, outPortals, false)
	fmt.Println("part1: ", part1, "in", time.Since(start))

	part2 := getShortest(m, startC, endC, inPortals, outPortals, true)
	fmt.Println("part2: ", part2, "in", time.Since(start))
}

type RecurseState struct {
	c     Coord
	depth int
}

func getShortest(m Chart, start Coord, end Coord, inPortals map[Coord]Coord, outPortals map[Coord]Coord, recursive bool) int {
	q := priorityqueue.NewQueue()
	q.Add(&priorityqueue.State{Data: RecurseState{c: start, depth: 0}, Priority: 0})

	seen := make(map[RecurseState]int)

	for q.HasNext() {
		s := q.Next()
		rs := s.Data.(RecurseState)
		c := rs.c

		if c == end {
			if recursive && rs.depth != 0 {
				continue
			}
			return s.Priority
		}

		seen[rs] = s.Priority

		// teleport
		if next, found := inPortals[c]; found {
			depth := rs.depth
			if recursive {
				depth++
			}
			nextRS := RecurseState{next, depth}
			if oldValue, seenFound := seen[nextRS]; !seenFound || oldValue >= s.Priority {
				nextState := priorityqueue.State{Data: nextRS, Priority: s.Priority + 1}
				q.Add(&nextState)
			}
		}

		if next, found := outPortals[c]; found {
			depth := rs.depth
			if recursive {
				depth--
			}
			nextRS := RecurseState{next, depth}
			if oldValue, seenFound := seen[nextRS]; !seenFound || oldValue >= s.Priority {
				nextState := priorityqueue.State{Data: nextRS, Priority: s.Priority + 1}
				q.Add(&nextState)
			}
		}

		for _, dir := range ALL {
			next := c.Move(dir)

			if m[next] != "." {
				continue
			}

			if recursive && rs.depth != 0 && (next == start || next == end) {
				continue
			}

			_, found := outPortals[next]
			if recursive && rs.depth == 0 && found {
				continue
			}

			oldValue, found := seen[RecurseState{next, rs.depth}]
			if found && oldValue < s.Priority {
				// we have been here before and it was a shorter path
				continue
			}

			nextState := priorityqueue.State{Data: RecurseState{next, rs.depth}, Priority: s.Priority + 1}
			q.Add(&nextState)
		}
	}

	return 0
}
