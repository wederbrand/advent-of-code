package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/2019/computer"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"math"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2019/07/input.txt", "\n")

	phases := []string{"0", "1", "2", "3", "4"}
	permutations := util.Permutations(phases)
	part1 := doIt(permutations, inFile, false)
	fmt.Println("part1: ", part1, "in", time.Since(start))

	phases = []string{"5", "6", "7", "8", "9"}
	permutations = util.Permutations(phases)
	part2 := doIt(permutations, inFile, true)
	fmt.Println("part2: ", part2, "in", time.Since(start))
}

func doIt(permutations [][]string, inFile []string, loopManyTimes bool) int {
	result := math.MinInt
	for _, phase := range permutations {
		computers := make([]*Computer, 0)
		for i := 0; i < 5; i++ {
			computer := NewComputer(inFile)
			computers = append(computers, &computer)
			if i > 0 {
				computers[i].SetInput(computers[i-1].GetOutput())
			}
		}
		first := computers[0]
		last := computers[len(computers)-1]

		for i, computer := range computers {
			go computer.Run()
			computer.GetInput() <- util.Atoi(phase[i])
		}

		first.GetInput() <- 0

		out := 0
		for out = range last.GetOutput() {
			if loopManyTimes && first.IsRunning() {
				first.GetInput() <- out
			}
		}
		result = max(result, out)
	}
	return result
}
