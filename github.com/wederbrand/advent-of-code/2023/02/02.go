package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2023/02/input.txt", "\n")

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
		for _, game := range strings.Split(games, ";") {
			for _, pick := range strings.Split(game, ",") {
				pick = strings.TrimSpace(pick)
				cnt := 0
				color := ""
				fmt.Sscanf(pick, "%d %s", &cnt, &color)
				if color == "red" {
					minRed = util.MaxOf(minRed, cnt)
					if cnt > 12 {
						possible = false
					}
				}
				if color == "green" {
					minGreen = util.MaxOf(minGreen, cnt)
					if cnt > 13 {
						possible = false
					}
				}
				if color == "blue" {
					minBlue = util.MaxOf(minBlue, cnt)
					if cnt > 14 {
						possible = false
					}
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
