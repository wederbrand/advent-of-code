package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"math"
	"strings"
	"time"
)

type Range struct {
	start int
	stop  int
}

func newRange(start int, stop int) *Range {
	r := new(Range)
	r.start = start
	r.stop = stop

	return r
}

type Translation struct {
	r      *Range
	change int
}

func newTranslation(dst int, src int, len int) *Translation {
	t := new(Translation)
	t.r = newRange(src, src+len-1)
	t.change = dst - src

	return t
}

type Mapping struct {
	translations []*Translation
}

func (t Mapping) explode(in []*Range) (out []*Range) {
	for len(in) > 0 {
		var r1 *Range
		r1, in = in[0], in[1:]
		matches := false
		for _, tr := range t.translations {
			r2 := tr.r
			if r1.start >= r2.start && r1.stop <= r2.stop {
				// covering start and stop, keep all with diff
				out = append(out, newRange(r1.start+tr.change, r1.stop+tr.change))
				matches = true
				break
			} else if r1.start >= r2.start && r1.start < r2.stop && r1.stop >= r2.stop {
				// covering start point, split on r2.stop
				out = append(out, newRange(r1.start+tr.change, r2.stop+tr.change))
				in = append(in, newRange(r2.stop+1, r1.stop))
				matches = true
				break
			} else if r1.start <= r2.start && r1.stop > r2.start && r1.stop <= r2.stop {
				// covering end point, split on r2.start
				in = append(in, newRange(r1.start, r2.start-1))
				out = append(out, newRange(r2.start+tr.change, r1.stop+tr.change))
				matches = true
				break
			} else if r1.start < r2.start && r1.stop > r2.stop {
				// within, split in tree
				in = append(in, newRange(r1.start, r2.start))
				out = append(out, newRange(r2.start+tr.change, r2.stop+tr.change))
				in = append(in, newRange(r2.stop, r2.stop))
				matches = true
				break
			}
		}

		if !matches {
			// no translation matches, append as is
			out = append(out, r1)
		}
	}
	return
}

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2023/05/input.txt", "\n")

	maps := make([]*Mapping, 0)
	var currentMap *Mapping
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
			currentMap = new(Mapping)
			maps = append(maps, currentMap)
			continue
		}

		// mappings to the current map
		mappings := util.MatchingNumbersAfterSplitOnAny(s, "")
		translation := newTranslation(mappings[0][0], mappings[0][1], mappings[0][2])
		currentMap.translations = append(currentMap.translations, translation)
	}

	ranges := make([]*Range, 0)
	for i := 0; i < len(seeds); i++ {
		ranges = append(ranges, newRange(seeds[i], seeds[i]))
	}
	for _, mapping := range maps {
		ranges = mapping.explode(ranges)
	}
	part1 := math.MaxInt
	for _, r := range ranges {
		part1 = min(part1, r.start)
	}
	fmt.Println("part1: ", part1, "in", time.Since(start))

	ranges = make([]*Range, 0)
	for i := 0; i < len(seeds); i += 2 {
		ranges = append(ranges, newRange(seeds[i], seeds[i]+seeds[i+1]-1))
	}
	for _, mapping := range maps {
		ranges = mapping.explode(ranges)
	}
	part2 := math.MaxInt
	for _, r := range ranges {
		part2 = min(part2, r.start)
	}
	fmt.Println("part2: ", part2, "in", time.Since(start))
}
