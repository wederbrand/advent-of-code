package chart

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/priorityqueue"
	"math"
	"slices"
	"strings"
)

type Dir [2]int

var N = Dir{0, -1}
var S = Dir{0, +1}
var E = Dir{+1, 0}
var W = Dir{-1, 0}

var NW = Dir{-1, -1}
var NE = Dir{+1, -1}
var SW = Dir{-1, +1}
var SE = Dir{+1, +1}

var UP = Dir{0, -1}
var DOWN = Dir{0, +1}
var RIGHT = Dir{+1, 0}
var LEFT = Dir{-1, 0}

var UPLEFT = Dir{-1, -1}
var UPRIGHT = Dir{+1, -1}
var DOWNLEFT = Dir{-1, +1}
var DOWNRIGHT = Dir{+1, +1}

var ALL = [4]Dir{UP, LEFT, RIGHT, DOWN}
var ALL_AND_DIAG = [8]Dir{UP, UPRIGHT, RIGHT, DOWNRIGHT, DOWN, DOWNLEFT, LEFT, UPLEFT}

func FromCompassChar(r uint8) Dir {
	switch r {
	case 'N':
		return N
	case 'S':
		return S
	case 'E':
		return E
	case 'W':
		return W
	default:
		panic("hoho")
	}
}

func (d Dir) Left() Dir {
	switch d {
	case UP:
		return LEFT
	case LEFT:
		return DOWN
	case DOWN:
		return RIGHT
	case RIGHT:
		return UP
	default:
		panic("hoho")
	}
}

func (d Dir) Right() Dir {
	switch d {
	case UP:
		return RIGHT
	case LEFT:
		return UP
	case DOWN:
		return LEFT
	case RIGHT:
		return DOWN
	default:
		panic("hoho")
	}
}

func (d Dir) ToArrowString() string {
	switch d {
	case UP:
		return "^"
	case LEFT:
		return "<"
	case DOWN:
		return "v"
	case RIGHT:
		return ">"
	default:
		panic("hoho")
	}
}

type Coord struct {
	X int
	Y int
}

func (c Coord) Move(dir Dir, length ...int) Coord {
	l := 1
	if len(length) == 1 {
		l = length[0]
	}
	return Coord{c.X + dir[0]*l, c.Y + dir[1]*l}
}

func (c Coord) AllBut(last Coord, m Chart, but string) ([]Coord, []string) {
	out := make([]Coord, 0)
	outString := make([]string, 0)
	for _, dir := range ALL {
		next := c.Move(dir)
		if next == last {
			continue
		}
		s2 := m[next]

		if !strings.Contains(but, s2) {
			out = append(out, next)
			outString = append(outString, s2)
		}
	}

	return out, outString
}

type Chart map[Coord]string

func (m Chart) FindLetter(letter string) (Coord, error) {
	for coord, s := range m {
		if s == letter {
			return coord, nil
		}
	}

	return Coord{}, fmt.Errorf("Letter %s not found", letter)
}

type PathState struct {
	Current Coord
	Path    []Coord
}

// GetPathLength returns the length of the shortest path from start to end
// note that the empty squares can't be skipped
// # and Coords outside the map is treated as walls
func (m Chart) GetPathLength(start Coord, end Coord) int {
	q := priorityqueue.NewQueue()
	q.Add(&priorityqueue.State{Data: start, Priority: 0})

	seen := make(map[Coord]bool)
	for q.HasNext() {
		s := q.Next()
		c := s.Data.(Coord)
		if c == end {
			return s.Priority
		}

		if seen[c] {
			continue
		}
		seen[c] = true

		for _, dir := range ALL {
			next := c.Move(dir)
			val, found := m[next]
			if !found || val == "#" {
				continue
			}
			q.Add(&priorityqueue.State{Data: next, Priority: s.Priority + 1})
		}
	}

	return -1
}

// GetPath returns the first path leading from the start to the end
// The path consists of a slice of Coords
// If there are multiple paths use GetAllPaths instead
func (m Chart) GetPath(start Coord, end Coord) []Coord {
	ch := make(chan []Coord)
	go m.GetAllPaths(ch, start, end, false)
	p := <-ch
	return p
}

// GetAllPaths returns all paths leading from the start to the end on the given channel
// The paths consist of a slice of Coords
// note that the empty squares can't be skipped
// # and Coords outside the map is treated as walls
// setting exhaustive to true will make the search continue even if a path has been found
// If there is exactly one path it's faster to set exhaustive to false, or use GetPath
func (m Chart) GetAllPaths(outChan chan []Coord, start Coord, end Coord, exhaustive bool) {
	q := priorityqueue.NewQueue()
	q.Add(&priorityqueue.State{Data: PathState{Current: start, Path: []Coord{start}}})

	seen := make(map[Coord]int)

	for q.HasNext() {
		s := q.Next()
		ps := s.Data.(PathState)
		c := ps.Current

		if c == end {
			outChan <- ps.Path
			continue
		}

		for _, dir := range ALL {
			next := c.Move(dir)

			if slices.Contains(ps.Path, next) {
				// we have been here before
				continue
			}

			if m[next] == "#" || m[next] == "" { // blocked by wall or outside the map
				continue
			}

			newPath := make([]Coord, len(ps.Path))
			copy(newPath, ps.Path)
			newPath = append(newPath, next)

			nextState := priorityqueue.State{Data: PathState{Current: next, Path: newPath}, Priority: s.Priority + 1}

			oldValue, found := seen[next]
			if !exhaustive && found && oldValue < s.Priority {
				// we have been here before and it was a shorter path
				continue
			}

			seen[c] = s.Priority

			q.Add(&nextState)
		}
	}

	close(outChan)
}

func MakeChart(in []string, ignored string) Chart {
	m := make(Chart)

	for y, s := range in {
		for x, r := range s {
			if strings.ContainsRune(ignored, r) {
				continue
			}
			m[Coord{x, y}] = string(r)
		}
	}

	return m
}

func Manhattan(a Coord, b Coord) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}

func RotateClockWise(in Chart) Chart {
	// For my reversed Y clockwise is the same a normal counterclockwise
	// 90° counterclockwise rotation: (𝑥,𝑦) becomes (−𝑦,𝑥)

	out := make(Chart)

	for key, value := range in {
		out[Coord{-key.Y, key.X}] = value
	}

	return out
}

func RotateCounterClockWise(in Chart) Chart {
	// For my reversed Y counterclockwise is the same a normal clockwise
	// 90° clockwise rotation: (𝑥,𝑦) becomes (𝑦,-𝑥)

	out := make(Chart)

	for key, value := range in {
		out[Coord{key.Y, -key.X}] = value
	}

	return out
}

func GetChartMaxes(m Chart) (minC Coord, maxC Coord) {
	minX := math.MaxInt
	minY := math.MaxInt
	maxX := math.MinInt
	maxY := math.MinInt
	for k := range m {
		minX = min(minX, k.X)
		minY = min(minY, k.Y)
		maxX = max(maxX, k.X)
		maxY = max(maxY, k.Y)
	}
	return Coord{minX, minY}, Coord{maxX, maxY}
}

func Find(m Chart, target string) (bool, Coord) {
	for k, v := range m {
		if v == target {
			return true, k
		}
	}
	return false, Coord{0, 0}
}

func PrintChart(m Chart) {
	minC, maxC := GetChartMaxes(m)
	for y := minC.Y; y <= maxC.Y; y++ {
		for x := minC.X; x <= maxC.X; x++ {
			s, found := m[Coord{x, y}]
			if !found {
				fmt.Print(".")
			} else {
				fmt.Print(s)
			}
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println()
}

func CopyChart(m Chart) Chart {
	out := make(Chart)
	for k, v := range m {
		out[k] = v
	}
	return out
}

func AsString(m Chart) string {
	out := ""
	minC, maxC := GetChartMaxes(m)
	for y := minC.Y; y <= maxC.Y; y++ {
		for x := minC.X; x <= maxC.X; x++ {
			s, found := m[Coord{x, y}]
			if !found {
				out += "."
			} else {
				out += s
			}
		}
		out += "|"
	}

	return out
}

func FromString(in string) Chart {
	c := make(Chart)

	lines := strings.Split(in, "|")
	for y, s := range lines {
		for x, r := range s {
			if r != '.' {
				c[Coord{x, y}] = string(r)
			}
		}
	}

	return c
}
