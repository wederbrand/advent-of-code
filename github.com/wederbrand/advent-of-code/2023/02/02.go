package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"regexp"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2023/02/input.txt", "\n")

	colorRe := regexp.MustCompile(`([0-9]+) (red|green|blue)`)
	part1 := 0
	part2 := 0
	for _, s := range inFile {
		gameNbr := 0
		possible := true
		minRed := 0
		minGreen := 0
		minBlue := 0
		fmt.Sscanf(s, "Game %d:.*", &gameNbr)
		games := strings.Split(s, ":")[1]
		all := colorRe.FindAllStringSubmatch(games, -1)
		for _, matches := range all {
			color := matches[2]
			cnt := util.Atoi(matches[1])
			if color == "red" {
				minRed = max(minRed, cnt)
				if cnt > 12 {
					possible = false
				}
			}
			if color == "green" {
				minGreen = max(minGreen, cnt)
				if cnt > 13 {
					possible = false
				}
			}
			if color == "blue" {
				minBlue = max(minBlue, cnt)
				if cnt > 14 {
					possible = false
				}
			}
		}
		if possible {
			part1 += gameNbr
		}
		part2 += minRed * minGreen * minBlue
	}

	fmt.Println("part1: ", part1, "in", time.Since(start))
	fmt.Println("part2: ", part2, "in", time.Since(start))
}
