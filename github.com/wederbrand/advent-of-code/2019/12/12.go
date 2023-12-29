package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"time"
)

type Pos3D [3]int

type Moon struct {
	p Pos3D
	v Pos3D
}

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2019/12/input.txt", "\n")

	moons := [4]Moon{}
	for i, s := range inFile {
		m := Moon{}
		fmt.Sscanf(s, "<x=%d, y=%d, z=%d>", &m.p[0], &m.p[1], &m.p[2])
		moons[i] = m
	}

	// apply gravity
	seen := [3]map[[8]int]bool{
		make(map[[8]int]bool),
		make(map[[8]int]bool),
		make(map[[8]int]bool),
	}
	cycle := []int{0, 0, 0}
	i := 0
	for {
		i++
		for a := 0; a < len(moons); a++ {
			for b := a + 1; b < len(moons); b++ {
				for j := 0; j < 3; j++ {
					if moons[a].p[j] > moons[b].p[j] {
						moons[a].v[j]--
						moons[b].v[j]++
					} else if moons[a].p[j] < moons[b].p[j] {
						moons[a].v[j]++
						moons[b].v[j]--
					}
				}
			}
		}

		// change positions
		for j := range moons {
			for k := 0; k < 3; k++ {
				moons[j].p[k] += moons[j].v[k]
			}
		}
		if i == 1000 {
			part1 := 0
			for _, m := range moons {
				pot := 0
				kin := 0
				for j := 0; j < 3; j++ {
					pot += util.IntAbs(m.p[j])
					kin += util.IntAbs(m.v[j])
				}
				part1 += pot * kin
			}
			fmt.Println("part1: ", part1, "in", time.Since(start))
		}

		for j := 0; j < 3; j++ {
			key := makeKey(moons, j)
			_, found := seen[j][key]
			if found && cycle[j] == 0 {
				cycle[j] = i - 1 // last iteration overlaps (ie is found in map) so -1 gives the full cycle
			} else {
				seen[j][key] = true
			}
		}
		if cycle[0] != 0 && cycle[1] != 0 && cycle[2] != 0 {
			part2 := util.Lcd(cycle)
			fmt.Println("part2: ", part2, "in", time.Since(start))
			break
		}
	}
}

func makeKey(moons [4]Moon, i int) [8]int {
	return [8]int{
		moons[0].p[i],
		moons[0].v[i],
		moons[1].p[i],
		moons[1].v[i],
		moons[2].p[i],
		moons[2].v[i],
		moons[3].p[i],
		moons[3].v[i],
	}
}
