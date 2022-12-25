package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type fabric struct {
	id int
	x int
	y int
	w int
	h int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	matcher := regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)

	fabrics := make([]fabric, 0)
	for scanner.Scan() {
		sValue := scanner.Text()
		fmt.Println(sValue)
		match := matcher.FindStringSubmatch(sValue)
		id, _ := strconv.Atoi(match[1])
		x, _ := strconv.Atoi(match[2])
		y, _ := strconv.Atoi(match[3])
		w, _ := strconv.Atoi(match[4])
		h, _ := strconv.Atoi(match[5])
		f := fabric{id, x, y, w, h}
		fmt.Println(f)
		fabrics = append(fabrics, f)
	}

	var overlaps [1001][1001]int

	for _, f := range fabrics {
		for i := f.x; i < f.w+f.x; i++ {
			for j := f.y; j < f.h+f.y; j++ {
				overlaps[i][j]++
			}
		}
	}

	for _, f := range fabrics {
		found := false
		for i := f.x; i < f.w+f.x; i++ {
			for j := f.y; j < f.h+f.y; j++ {
				if overlaps[i][j] > 1 {
					found = true
				}
			}
		}
		if (!found) {
			fmt.Println(f.id)
		}
	}
}
