package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"math"
	"slices"
	"time"
)

type Pos3D struct {
	x float64
	y float64
	z float64
}

type Hail struct {
	p Pos3D
	s Pos3D

	// standard for ax + by + c = 0
	a float64
	b float64
	c float64
}

const minValue = 200000000000000
const maxValue = 400000000000000

func main() {
	startTimer := time.Now()
	inFile := util.GetFileContents("2023/24/input.txt", "\n")

	hails := make([]Hail, 0)
	for _, s := range inFile {
		h := Hail{}
		fmt.Sscanf(s, "%f, %f, %f @ %f, %f, %f", &h.p.x, &h.p.y, &h.p.z, &h.s.x, &h.s.y, &h.s.z)
		hails = append(hails, h)
	}

	// rewrite hails as general form
	for i := range hails {
		toNormalForm(&hails[i])
	}

	fmt.Println("parsing:", time.Since(startTimer))
	startTimer = time.Now()

	// check each pair of hails
	// count if hitpoint is within given values

	part1 := 0
	for i := 0; i < len(hails)-1; i++ {
		a := hails[i]
		for j := i + 1; j < len(hails); j++ {
			b := hails[j]
			if check(a, b) {
				part1++
			}
		}
	}
	fmt.Println("part1: ", part1, "in", time.Since(startTimer))
	startTimer = time.Now()

	possibleDX, possibleDY, possibleDZ := []int{}, []int{}, []int{}
	for i := 0; i < len(hails)-1; i++ {
		for j := i + 1; j < len(hails); j++ {
			a, b := hails[i], hails[j]
			if a.s.x == b.s.x && len(possibleDX) != 1 {
				possible := findPossibleSpeeds(int(b.p.x-a.p.x), int(a.s.x))
				if len(possibleDX) == 0 {
					possibleDX = possible
				} else {
					possibleDX = getIntersect(possibleDX, possible)
				}
			}
			if a.s.y == b.s.y && len(possibleDY) != 1 {
				possible := findPossibleSpeeds(int(b.p.y-a.p.y), int(a.s.y))
				if len(possibleDY) == 0 {
					possibleDY = possible
				} else {
					possibleDY = getIntersect(possibleDY, possible)
				}
			}
			if a.s.z == b.s.z && len(possibleDZ) != 1 {
				possible := findPossibleSpeeds(int(b.p.z-a.p.z), int(a.s.z))
				if len(possibleDZ) == 0 {
					possibleDZ = possible
				} else {
					possibleDZ = getIntersect(possibleDZ, possible)
				}
			}
		}
	}

	// 	slope := float64(hail.s.y) / float64(hail.s.x)
	//
	//	hail.a = -slope
	//	hail.b = 1 // this is always 1 the way I solve it
	//	hail.c = float64(-hail.p.y) - (slope * float64(-hail.p.x))

	rockSpeed := Pos3D{float64(possibleDX[0]), float64(possibleDY[0]), float64(possibleDZ[0])}
	a, b := hails[0], hails[1]
	mA := (a.s.y - rockSpeed.y) / (a.s.x - rockSpeed.x)
	mB := (b.s.y - rockSpeed.y) / (b.s.x - rockSpeed.x)
	cA := a.p.y - (mA * a.p.x)
	cB := b.p.y - (mB * b.p.x)
	xPos := (cB - cA) / (mA - mB)
	yPos := mA*xPos + cA
	impactTime := (xPos - a.p.x) / (a.s.x - rockSpeed.x)
	zPos := a.p.z + (a.s.z-rockSpeed.z)*impactTime
	part2 := int(math.Round(xPos + yPos + zPos))

	fmt.Println("part2: ", part2, "in", time.Since(startTimer))
}

func findPossibleSpeeds(distance int, hailSpeed int) []int {
	possible := make([]int, 0)
	// assume rock speed is never higher than 1000
	for v := -1000; v < 1000; v++ {
		catchUpSpeed := v - hailSpeed
		if catchUpSpeed != 0 && distance%catchUpSpeed == 0 {
			possible = append(possible, v)
		}
	}
	return possible
}

func getIntersect(a, b []int) []int {
	result := make([]int, 0)
	for _, val := range a {
		if slices.Contains(b, val) {
			result = append(result, val)
		}
	}
	return result
}

func toNormalForm(hail *Hail) {
	slope := hail.s.y / hail.s.x

	hail.a = -slope
	hail.b = 1 // this is always 1 the way I solve it
	hail.c = -hail.p.y - (slope * -hail.p.x)
}

func check(a Hail, b Hail) bool {
	x, y := intersect(a, b)
	ta := (x - a.p.x) / a.s.x
	tb := (x - b.p.x) / b.s.x

	if ta == math.Inf(1) || ta == math.Inf(-1) {
		return false
	}

	if ta < 0 || tb < 0 {
		// negative time
		return false
	}
	return x >= minValue && x <= maxValue && y >= minValue && y <= maxValue
}

func intersect(a Hail, b Hail) (float64, float64) {
	x := (a.b*b.c - b.b*a.c) / (a.a*b.b - b.a*a.b)
	y := (a.c*b.a - b.c*a.a) / (a.a*b.b - b.a*a.b)

	return x, y
}
