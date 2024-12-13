package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"math"
	"strings"
	"time"
)

type Play struct {
	a     Coord
	b     Coord
	prize Coord
}

func main() {
	start := time.Now()
	inFile := GetFileContents("2024/13/input.txt", "\n")

	plays := make([]Play, 0)
	current := Play{}
	for _, line := range inFile {
		if strings.Contains(line, "Button A") {
			fmt.Sscanf(line, "Button A: X+%d, Y+%d", &current.a.X, &current.a.Y)
		}
		if strings.Contains(line, "Button B") {
			fmt.Sscanf(line, "Button B: X+%d, Y+%d", &current.b.X, &current.b.Y)
		}
		if strings.Contains(line, "Prize") {
			fmt.Sscanf(line, "Prize: X=%d, Y=%d", &current.prize.X, &current.prize.Y)
			plays = append(plays, current)
			current = Play{}
		}
	}

	fmt.Println("Parsing in", time.Since(start))
	start = time.Now()

	p1 := 0
	for _, play := range plays {
		score := doIt(play)

		if score < math.MaxInt {
			p1 += score
		}
	}
	fmt.Println("Part 1:", p1, "in", time.Since(start))
	start = time.Now()

	p2 := 0
	for _, play := range plays {
		play.prize.X += 10000000000000
		play.prize.Y += 10000000000000
		score := doIt(play)
		if score < math.MaxInt {
			p2 += score
		}
	}

	fmt.Println("Part 2:", p2, "in", time.Since(start))
}

// Return the lowest score to reach the prize, or MaxInt if it's impossible
func doIt(play Play) int {
	one := Equation{A: play.a.X, B: play.b.X, C: play.prize.X}
	two := Equation{A: play.a.Y, B: play.b.Y, C: play.prize.Y}
	x, y, err := WikiCramer(one, two)

	buttonAPresses := x
	buttonBPresses := y
	if err != nil {
		return math.MaxInt
	}
	return buttonAPresses*3 + buttonBPresses
}
