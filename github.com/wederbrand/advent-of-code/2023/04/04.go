package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"slices"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2023/04/input.txt", "\n")

	part1 := 0
	part2 := 0
	cards := make([]int, len(inFile)+1)
	for rowNbr, s := range inFile {
		cardNbr := rowNbr + 1
		cards[cardNbr] += 1
		part2 += cards[cardNbr]

		splits := util.MatchingNumbersAfterSplitOnAny(s, ":|", " ")
		winningNumbers := splits[1]
		iHave := splits[2]

		cnt := 0
		nbrOfWins := 0
		for _, i := range iHave {
			if slices.Contains(winningNumbers, i) {
				nbrOfWins++
				if cnt == 0 {
					cnt = 1
				} else {
					cnt *= 2
				}
			}
		}
		for i := 1; i <= nbrOfWins; i++ {
			cards[cardNbr+i] += cards[cardNbr]
		}
		part1 += cnt
	}

	fmt.Println("part1: ", part1, "in", time.Since(start))
	fmt.Println("part2: ", part2, "in", time.Since(start))
}
