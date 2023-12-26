package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"slices"
	"strings"
	"time"
)

type Orbit struct {
	name   string
	orbits *Orbit
}

func getOrMake(name string, all map[string]*Orbit) *Orbit {
	orbit, found := all[name]
	if found {
		return orbit
	} else {
		o := Orbit{name, nil}
		all[name] = &o
		return &o
	}
}

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2019/06/input.txt", "\n")

	all := make(map[string]*Orbit)
	for _, s := range inFile {
		split := strings.Split(s, ")")
		src := getOrMake(split[0], all)
		dst := getOrMake(split[1], all)
		dst.orbits = src
	}

	part1 := 0
	for _, orbit := range all {
		part1 += getOrbitCount(orbit)
	}
	fmt.Println("part1: ", part1, "in", time.Since(start))

	you := make([]*Orbit, 0)
	you = append(you, all["YOU"])
	for you[len(you)-1] != nil {
		you = append(you, you[len(you)-1].orbits)
	}

	san := make([]*Orbit, 0)
	san = append(san, all["SAN"])
	for san[len(san)-1] != nil {
		san = append(san, san[len(san)-1].orbits)
	}

	slices.Reverse(you)
	slices.Reverse(san)

	i := 0
	for you[i] == san[i] {
		i++
	}

	part2 := (len(you) - i - 1) + (len(san) - i - 1)
	fmt.Println("part2: ", part2, "in", time.Since(start))
}

func getOrbitCount(orbit *Orbit) int {
	if orbit.orbits == nil {
		return 0
	} else {
		return 1 + getOrbitCount(orbit.orbits)
	}
}
