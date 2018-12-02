package main

import (
	"sort"
	"strings"
	"fmt"
	"os"
	"bufio"
)

type Box []rune

func (b Box) Len() int {
	return len(b)
}

func (box Box) Less(a int, b int) bool {
	return box[a] < box[b]
}

func (box Box) Swap(a int, b int) {
	box[a], box[b] = box[b], box[a]
}

func min(a int, b int) int {
	if (a < b) {
		return a
	} else {
		return b
	}
}

func (box Box) count() (int, int) {
	sort.Sort(box)
	
	var lastSeen rune
	count := 1
	two := 0
	three := 0
	for _, c := range box {
		if (c == lastSeen) {
			count++
		} else {
			lastSeen = c
			if (count == 2) {
				two++
			} else if (count == 3) {
				three++
			}
			count = 1
		}
	}
	if (count == 2) {
		two++
	} else if (count == 3) {
		three++
	}

	return min(two, 1), min(three, 1)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	input := make([]Box, 0)
	for scanner.Scan() {
		sValue := scanner.Text()
		box := Box(strings.TrimSpace(sValue))
		input = append(input, box)
	}

	two := 0
	three := 0
	for _, box := range input {
		a, b := box.count()
		two += a
		three += b
	}

	fmt.Println(two * three)
}