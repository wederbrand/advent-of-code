package main

import (
	"fmt"
	"strings"
	"time"

	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	// . "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	// . "github.com/wederbrand/advent-of-code/github.com/wederbrand/priorityqueue"
)

var seen = make(map[[2]string]int)

func main() {
	start := time.Now()
	input := GetFileContents("2025/11/input.txt", "\n")

	paths := make(map[string][]string)
	for _, line := range input {
		split := strings.Split(line, ":")
		key := strings.TrimSpace(split[0])
		values := strings.Split(strings.TrimSpace(split[1]), " ")
		paths[key] = values
	}

	part1 := pathIt("you", "out", paths)
	fmt.Println("Part 1:", part1, "in", time.Since(start))

	svrDac := pathIt("svr", "dac", paths)
	dacFft := pathIt("dac", "fft", paths)
	fftOut := pathIt("fft", "out", paths)
	svrFft := pathIt("svr", "fft", paths)
	fftDac := pathIt("fft", "dac", paths)
	dacOut := pathIt("dac", "out", paths)

	part2 := svrDac*dacFft*fftOut + svrFft*fftDac*dacOut
	fmt.Println("Part 2:", part2, "in", time.Since(start))
}

func pathIt(in, out string, paths map[string][]string) int {
	if in == out {
		return 1
	}

	key := [2]string{in, out}
	if val, ok := seen[key]; ok {
		return val
	}

	total := 0
	for _, next := range paths[in] {
		total += pathIt(next, out, paths)
	}

	seen[key] = total
	return total
}
