package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"time"
)

type Point struct {
	x, y, z, t int
}

func main() {
	start := time.Now()
	inFile := GetFileContents("2018/25/input.txt", "\n")

	constellations := make([]*[]Point, 0)
	for _, line := range inFile {
		p := Point{}
		fmt.Sscanf(line, "%d,%d,%d,%d", &p.x, &p.y, &p.z, &p.t)

		possibleConstellations := make([]*[]Point, 0)
		for _, c := range constellations {
			for _, point := range *c {
				if IntAbs(point.x-p.x)+IntAbs(point.y-p.y)+IntAbs(point.z-p.z)+IntAbs(point.t-p.t) <= 3 {
					possibleConstellations = append(possibleConstellations, c)
					break
				}
			}
		}
		if len(possibleConstellations) > 1 {
			// join those constellations and join it
			newConstellation := make([]Point, 0)
			for _, c := range possibleConstellations {
				newConstellation = append(newConstellation, *c...)
				constellations = removeConstellation(constellations, c)
			}
			newConstellation = append(newConstellation, p)
			constellations = append(constellations, &newConstellation)
		} else if len(possibleConstellations) == 1 {
			// join that constellation
			*possibleConstellations[0] = append(*possibleConstellations[0], p)

		} else {
			// create new constellation
			constellations = append(constellations, &[]Point{p})
		}
	}

	fmt.Println("Part 1:", len(constellations), "in", time.Since(start))
}

func removeConstellation(constellations []*[]Point, toRemove *[]Point) []*[]Point {
	for i, c := range constellations {
		if c == toRemove {
			constellations = append(constellations[:i], constellations[i+1:]...)
			break
		}
	}
	return constellations
}
