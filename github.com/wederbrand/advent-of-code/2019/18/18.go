package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/priorityqueue"
	"golang.org/x/exp/maps"
	"math"
	"slices"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2019/18/input.txt", "\n")

	m := MakeChart(inFile, "")

	allKeys := make(map[string]Coord)
	allDoors := make(map[string]Coord)

	currentKey := "a"
	for {
		c, err := m.FindLetter(currentKey)
		if err != nil {
			break
		}
		allKeys[currentKey] = c
		door := strings.ToUpper(currentKey)
		c, err = m.FindLetter(door)
		if err != nil {
			allDoors[door] = c
		}
		currentKey = string(currentKey[0] + 1)
	}

	startCoord, _ := m.FindLetter("@")
	starts := []Coord{startCoord}

	part1 := doIt(m, starts, map[string]bool{}, len(allKeys))
	fmt.Println("part1: ", part1, "in", time.Since(start))

	m[startCoord.Move(UP)] = "#"
	m[startCoord.Move(DOWN)] = "#"
	m[startCoord.Move(LEFT)] = "#"
	m[startCoord.Move(RIGHT)] = "#"
	m[startCoord] = "#"
	starts = []Coord{startCoord.Move(UPLEFT), startCoord.Move(UPRIGHT), startCoord.Move(DOWNLEFT), startCoord.Move(DOWNRIGHT)}

	part2 := doIt(m, starts, map[string]bool{}, len(allKeys))
	fmt.Println("part2: ", part2, "in", time.Since(start))
}

var cache = make(map[string]int)

func doIt(m Chart, currentCoords []Coord, keys map[string]bool, numKeys int) int {
	mapKeys := maps.Keys(keys)
	slices.Sort(mapKeys)
	slices.SortFunc(currentCoords, func(i, j Coord) int {
		if i.Y == j.Y {
			return i.X - j.X
		}
		return i.Y - j.Y
	})
	cacheKey := ""
	for _, c := range currentCoords {
		cacheKey += fmt.Sprintf("%d,%d,", c.X, c.Y)
	}
	cacheKey = fmt.Sprintf("%s%s", cacheKey, strings.Join(mapKeys, ""))

	if val, found := cache[cacheKey]; found {
		return val
	}

	if len(keys) == numKeys {
		return 0
	}
	allPaths := make(map[string][]Coord)
	for _, c := range currentCoords {
		newPaths := getAllPaths(m, c, keys)
		for k, v := range newPaths {
			allPaths[k] = v
		}
	}
	minPath := math.MaxInt
	for s, path := range allPaths {
		keys[s] = true
		newCurrentCoords := slices.Clone(currentCoords)
		for i, c := range newCurrentCoords {
			if c == path[0] {
				newCurrentCoords[i] = path[len(path)-1]
				break
			}
		}

		val := doIt(m, newCurrentCoords, keys, numKeys)
		if val+len(path)-1 < minPath {
			minPath = val + len(path) - 1
		}
		delete(keys, s)
	}

	cache[cacheKey] = minPath
	return minPath
}

type pathStateWithFoundKeys struct {
	current Coord
	path    []Coord
}

type PathCacheKey struct {
	start Coord
	keys  string
}

var pathCache = make(map[PathCacheKey]map[string][]Coord)

func getAllPaths(m Chart, start Coord, keys map[string]bool) map[string][]Coord {
	mapsKeys := maps.Keys(keys)
	slices.Sort(mapsKeys)
	cacheKey := PathCacheKey{start, strings.Join(mapsKeys, "")}
	if val, found := pathCache[cacheKey]; found {
		return val
	}

	allPaths := make(map[string][]Coord)

	q := priorityqueue.NewQueue()
	q.Add(&priorityqueue.State{Data: pathStateWithFoundKeys{current: start, path: []Coord{start}}, Priority: 0})

	seen := make(map[Coord]int)

	for q.HasNext() {
		s := q.Next()
		ps := s.Data.(pathStateWithFoundKeys)
		c := ps.current

		if m[c] >= "a" && m[c] <= "z" {
			// found a key
			if _, found := keys[m[c]]; !found {
				allPaths[m[c]] = ps.path
			}
		}

		for _, dir := range ALL {
			next := c.Move(dir)

			if slices.Contains(ps.path, next) {
				// we have been here before
				continue
			}

			if m[next] != "." {
				if m[next] == "#" || m[next] == "" {
					continue
				}
				if m[next] == strings.ToUpper(m[next]) && m[next] != "@" {
					// found a door
					if _, found := keys[strings.ToLower(m[next])]; !found {
						continue
					}
				}
			}

			newPath := make([]Coord, len(ps.path))
			copy(newPath, ps.path)
			newPath = append(newPath, next)

			nextState := priorityqueue.State{Data: pathStateWithFoundKeys{current: next, path: newPath}, Priority: s.Priority + 1}

			oldValue, found := seen[next]
			if found && oldValue <= s.Priority {
				// we have been here before and it was a shorter path
				continue
			}

			seen[c] = s.Priority

			q.Add(&nextState)
		}
	}

	pathCache[cacheKey] = allPaths
	return allPaths
}
