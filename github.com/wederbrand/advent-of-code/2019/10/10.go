package main

import (
	"cmp"
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"golang.org/x/exp/maps"
	"math"
	"slices"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2019/10/input.txt", "\n")

	m := MakeChart(inFile, ".")

	part1 := 0
	var selectedStation Coord
	for station := range m {
		seen := getSeen(m, station)
		if len(seen) > part1 {
			part1 = len(seen)
			selectedStation = station
		}
	}

	last := zap200(m, selectedStation)

	part2 := last.X*100 + last.Y
	fmt.Println("part1: ", part1, "in", time.Since(start))
	fmt.Println("part2: ", part2, "in", time.Since(start))
}

func zap200(m Chart, selectedStation Coord) Coord {
	killCount := 0
	for {
		seen := getSeen(m, selectedStation)
		asteroids := maps.Values(seen)
		slices.SortFunc(asteroids, func(a, b Coord) int {
			aa := angle(a, selectedStation)
			ba := angle(b, selectedStation)
			return cmp.Compare(aa, ba)
		})
		for _, asteroid := range asteroids {
			delete(m, asteroid)
			killCount++
			if killCount == 200 {
				return asteroid
			}

		}
	}
}

func angle(asteroid Coord, station Coord) float64 {
	dx := asteroid.X - station.X
	dy := asteroid.Y - station.Y

	rad := math.Atan2(float64(dy), float64(dx))
	deg := rad * (180 / math.Pi)
	deg += 90
	if deg > 360 {
		deg -= 360
	}
	if deg < 0 {
		deg += 360
	}
	return deg
}

func getSeen(m Chart, station Coord) map[Coord]Coord {
	asteroids := make([]Coord, 0)
	for asteroid := range m {
		if asteroid != station {
			asteroids = append(asteroids, asteroid)
		}
	}
	// reverse sort, furthers first
	slices.SortFunc(asteroids, func(a, b Coord) int {
		ma := Manhattan(station, a)
		mb := Manhattan(station, b)
		return cmp.Compare(mb, ma)
	})

	seen := make(map[Coord]Coord)
	for _, asteroid := range asteroids {
		asteroidDiff := getSmallestDX(asteroid, station)
		seen[asteroidDiff] = asteroid
	}
	return seen
}

func getSmallestDX(asteroid Coord, station Coord) Coord {
	dx := asteroid.X - station.X
	dy := asteroid.Y - station.Y
	smallestDX := dx / util.IntAbs(util.Gcd(dx, dy))
	smallestDY := dy / util.IntAbs(util.Gcd(dx, dy))

	return Coord{smallestDX, smallestDY}
}
