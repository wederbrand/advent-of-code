package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"math"
	"strings"
	"time"
)

type fromto struct {
	dst int
	src int
	len int
}

type trans struct {
	fromto []fromto
}

func (t trans) findOut(pos int) int {
	for _, f := range t.fromto {
		if pos >= f.src && pos <= f.src+f.len {
			return pos + f.dst - f.src
		}
	}
	return pos
}

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2023/05/input.txt", "\n")

	maps := make([]*trans, 0)
	var currentMap *trans
	var seeds []int
	for _, s := range inFile {
		if strings.HasPrefix(s, "seeds: ") {
			seeds = util.MatchingNumbersAfterSplitOnAny(s, ":")[1]
			continue
		}

		if s == "" {
			continue
		}

		if strings.HasSuffix(s, "map:") {
			// a new map, name is not important, order is
			currentMap = new(trans)
			maps = append(maps, currentMap)
			continue
		}

		// mappings to the current map
		mappings := util.MatchingNumbersAfterSplitOnAny(s, "")
		currentMap.fromto = append(currentMap.fromto, fromto{mappings[0][0], mappings[0][1], mappings[0][2]})
	}

	part1 := math.MaxInt
	for _, seed := range seeds {
		currentPos := seed
		for _, m := range maps {
			currentPos = m.findOut(currentPos)
		}
		part1 = util.MinOf(part1, currentPos)
	}

	fmt.Println("part1: ", part1, "in", time.Since(start))
	//fmt.Println("part2: ", part2, "in", time.Since(start))
}
