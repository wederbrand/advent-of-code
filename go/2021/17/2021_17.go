package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type possibility struct {
	value    int
	infinite bool
	max      int
	steps    int
}

func key(x int, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}

func main() {
	readFile, err := os.ReadFile("2021/17/2021_17.txt")
	if err != nil {
		log.Fatal(err)
	}

	// target area: x=20..30, y=-10..-5
	re := regexp.MustCompile("target area: x=(-*\\d+)..(-*\\d+), y=(-*\\d+)..(-*\\d+)")
	inFile := strings.TrimSpace(string(readFile))

	submatch := re.FindStringSubmatch(inFile)
	minX, _ := strconv.Atoi(submatch[1])
	maxX, _ := strconv.Atoi(submatch[2])
	minY, _ := strconv.Atoi(submatch[3])
	maxY, _ := strconv.Atoi(submatch[4])

	possibleX := getPossibleX(minX, maxX)
	possibleY := getPossibleY(minY, maxY)

	// match
	maxHeight := math.MinInt
	total := make(map[string]bool)
	for _, x := range possibleX {
		for _, y := range possibleY {
			if x.steps == y.steps || (x.infinite && y.steps >= x.steps) {
				total[key(x.value, y.value)] = true
				if y.max > maxHeight {
					maxHeight = y.max
				}
			}
		}
	}

	fmt.Println("part 1", maxHeight)
	fmt.Println("part 2", len(total))
}

func getPossibleX(minX int, maxX int) []possibility {
	out := make([]possibility, 0)
	for startdx := 1; startdx <= maxX; startdx++ {
		dx := startdx
		posX := 0
		for steps := 1; ; steps++ {
			posX += dx
			if posX >= minX && posX <= maxX {
				out = append(out, possibility{value: startdx, infinite: dx == 0, steps: steps})
			}
			if posX > maxX {
				break
			}

			if dx == 0 {
				break
			}

			dx--
		}
	}

	return out
}

func getPossibleY(minY int, maxY int) []possibility {
	out := make([]possibility, 0)
	for startdy := minY - 1; startdy <= -minY; startdy++ {
		dy := startdy
		posY := 0
		maxHeight := math.MinInt
		for steps := 1; ; steps++ {
			posY += dy
			if maxHeight < posY {
				maxHeight = posY
			}
			if posY <= maxY && posY >= minY {
				out = append(out, possibility{value: startdy, max: maxHeight, steps: steps})
			}
			if posY < minY {
				break
			}

			dy--
		}
	}

	return out
}
