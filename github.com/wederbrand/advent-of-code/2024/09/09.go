package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2024/09/input.txt", "\n")

	p1 := doIt(inFile[0], false)
	fmt.Println("Part 1:", p1, "in", time.Since(start))

	p2 := doIt(inFile[0], true)
	fmt.Println("Part 2:", p2, "in", time.Since(start))
}

func doIt(inFile string, part2 bool) int {
	disk := make([]int, 0)
	index := 0
	file := true
	for _, i := range inFile {
		atoi := util.Atoi(string(i))
		for j := 0; j < atoi; j++ {
			if file {
				disk = append(disk, index)
			} else {
				disk = append(disk, -1)
			}
		}

		file = !file
		if file {
			index++
		}
	}

	// defrag
	if !part2 {
		part1Defrag(disk)
	} else {
		part2Defrag(disk)
	}

	checkSum := 0
	for i, s := range disk {
		if s == -1 {
			continue
		}
		checkSum += i * s
	}

	return checkSum
}

func part1Defrag(disk []int) {
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] == -1 {
			continue
		}
		free := findFree(disk, i, 1)
		if free == -1 {
			break
		} else {
			disk[free] = disk[i]
			disk[i] = -1
		}
	}
}

func part2Defrag(disk []int) {
	for end := len(disk) - 1; end >= 0; end-- {
		if disk[end] == -1 {
			continue
		}

		// find the size of the block
		start := end
		for j := end; j >= 0 && disk[j] == disk[end]; j-- {
			start = j
		}

		// find any free space of the same size, or larger
		size := end - start + 1
		free := findFree(disk, start, size)

		if free > -1 {
			// move the block there
			for j := 0; j < size; j++ {
				disk[free+j], disk[start+j] = disk[start+j], disk[free+j]
			}
		}

		// adjust end
		end = start
	}
}

func findFree(disk []int, bound int, minSize int) int {
	for i, s := range disk {
		if i >= bound {
			return -1
		}
		if s == -1 {
			gapSize := 0
			for j := i; j < len(disk) && disk[j] == -1; j++ {
				gapSize++
			}
			if gapSize >= minSize {
				return i
			}
		}
	}
	return -1
}
