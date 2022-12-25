package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"sort"
	"strconv"
)

func main() {
	inFile := util.GetFileContents("2022/01/input.txt", "\n")

	current := 0
	m := make([]int, 0)
	for _, s := range inFile {
		if s == "" {
			m = append(m, current)
			current = 0
		}
		atoi, _ := strconv.Atoi(s)
		current += atoi
	}
	m = append(m, current)

	sort.Sort(sort.Reverse(sort.IntSlice(m)))

	fmt.Println("part1: ", m[0])
	fmt.Println("part2: ", m[0]+m[1]+m[2])
}
