package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/priorityqueue"
	"golang.org/x/exp/maps"
	"math"
	"sort"
	"time"
)

type Unit struct {
	name string
	hp   int
	dam  int
}

func main() {
	start := time.Now()
	inFile := GetFileContents("2018/15/input.txt", "\n")

	part1 := doIt(inFile, 3, false)
	fmt.Println("Part 1:", part1, "in", time.Since(start))

	attack := 4
	part2 := 0
	for {
		part2 = doIt(inFile, attack, true)
		if part2 != 0 {
			break
		}
		attack++
	}
	fmt.Println("Part 2:", part2, "in", time.Since(start))
}

func doIt(inFile []string, elvesDamage int, abortOnElvesDeath bool) int {
	m := MakeChart(inFile, "")
	elves := make(map[Coord]*Unit)
	goblins := make(map[Coord]*Unit)

	for coord, s := range m {
		if s == "E" {
			elves[coord] = &Unit{name: "E", hp: 200, dam: elvesDamage}
		} else if s == "G" {
			goblins[coord] = &Unit{name: "G", hp: 200, dam: 3}
		}
	}

	turns := 0
	for {
		// sort units
		units := append(maps.Keys(elves), maps.Keys(goblins)...)
		sort.Slice(units, func(i, j int) bool {
			if units[i].Y != units[j].Y {
				return units[i].Y < units[j].Y
			}
			return units[i].X < units[j].X
		})

		for _, unit := range units {
			if m[unit] == "." {
				// unit was just killed
				continue
			}
			attacker := goblins[unit]
			// find targets
			enemies := elves
			if m[unit] == "E" {
				enemies = goblins
				attacker = elves[unit]
			}
			if len(enemies) == 0 {
				// game over
				totalHP := 0
				for _, e := range elves {
					totalHP += e.hp
				}
				for _, g := range goblins {
					totalHP += g.hp
				}
				return totalHP * turns
			}

			// move
			path := getBestPath(m, unit, enemies)
			if path != nil {
				if m[unit] == "E" {
					elves[path[1]] = elves[unit]
					delete(elves, unit)
				} else {
					goblins[path[1]] = goblins[unit]
					delete(goblins, unit)
				}
				m[unit], m[path[1]] = m[path[1]], m[unit]
				unit = path[1]
			}

			// attack
			bestTarget := Coord{X: -1}
			for _, dir := range ALL {
				target := unit.Move(dir)
				if enemies[target] != nil {
					if bestTarget.X == -1 {
						bestTarget = target
					} else if enemies[target].hp < enemies[bestTarget].hp {
						bestTarget = target
					} else if enemies[bestTarget].hp == enemies[target].hp {
						if target.Y < bestTarget.Y || (target.Y == bestTarget.Y && target.X < bestTarget.X) {
							bestTarget = target
						}
					}
				}
			}

			if bestTarget.X != -1 {
				enemies[bestTarget].hp -= attacker.dam
				if enemies[bestTarget].hp <= 0 {
					if m[bestTarget] == "E" && abortOnElvesDeath {
						return 0
					}
					delete(enemies, bestTarget)
					m[bestTarget] = "."
				}
			}
		}
		turns++
	}
}

func getBestPath(m Chart, start Coord, targets map[Coord]*Unit) []Coord {
	allEnds := make(map[Coord]bool)
	for target := range targets {
		manhattan := Manhattan(start, target)
		if manhattan == 1 {
			return nil // adjacent, don't move
		}

		if m[target.Move(UP)] == "." {
			allEnds[target.Move(UP)] = true
		}
		if m[target.Move(DOWN)] == "." {
			allEnds[target.Move(DOWN)] = true
		}
		if m[target.Move(LEFT)] == "." {
			allEnds[target.Move(LEFT)] = true
		}
		if m[target.Move(RIGHT)] == "." {
			allEnds[target.Move(RIGHT)] = true
		}
	}

	paths := getPaths(m, start, allEnds)

	if len(paths) == 0 {
		return nil // no paths found, don't move
	}

	sort.Slice(paths, func(i, j int) bool {
		if len(paths[i]) != len(paths[j]) {
			return len(paths[i]) < len(paths[j])
		}

		// return the one that ends first in reading order
		if paths[i][len(paths[i])-1].Y != paths[j][len(paths[j])-1].Y {
			return paths[i][len(paths[i])-1].Y < paths[j][len(paths[j])-1].Y
		}

		if paths[i][len(paths[i])-1].X != paths[j][len(paths[j])-1].X {
			return paths[i][len(paths[i])-1].X < paths[j][len(paths[j])-1].X
		}

		// return the one that first steps different in reading order
		for k := 0; k < len(paths[i]); k++ {
			if paths[i][k].Y != paths[j][k].Y {
				return paths[i][k].Y < paths[j][k].Y
			}
			if paths[i][k].X != paths[j][k].X {
				return paths[i][k].X < paths[j][k].X
			}
		}

		panic("ho ho no")
	})

	return paths[0]
}

func getPaths(m Chart, start Coord, allEnds map[Coord]bool) [][]Coord {
	result := make([][]Coord, 0)

	q := priorityqueue.NewQueue()
	q.Add(&priorityqueue.State{Data: PathState{Current: start, Path: []Coord{start}}})

	foundLength := math.MaxInt
	seen := make(map[Coord]bool)
	for q.HasNext() {
		s := q.Next()
		ps := s.Data.(PathState)
		c := ps.Current

		if allEnds[c] {
			if len(ps.Path) > foundLength {
				// we have found a longer path than the shortest one
				return result
			}
			result = append(result, ps.Path)
			foundLength = len(ps.Path)
			continue
		}

		for _, dir := range ALL {
			next := c.Move(dir)

			if seen[next] {
				continue
			}

			seen[next] = true

			if m[next] == "#" || m[next] == "E" || m[next] == "G" {
				continue
			}

			newPath := make([]Coord, len(ps.Path))
			copy(newPath, ps.Path)
			newPath = append(newPath, next)

			nextState := priorityqueue.State{Data: PathState{Current: next, Path: newPath}, Priority: s.Priority + 1}

			q.Add(&nextState)
		}
	}

	return result
}
