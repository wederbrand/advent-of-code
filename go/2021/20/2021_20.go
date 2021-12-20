package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type image map[string]int

func (i image) getValue(x int, y int) (out int) {
	out += i.getSingleValue(x+1, y+1) * 1
	out += i.getSingleValue(x+0, y+1) * 2
	out += i.getSingleValue(x-1, y+1) * 4
	out += i.getSingleValue(x+1, y+0) * 8
	out += i.getSingleValue(x+0, y+0) * 16
	out += i.getSingleValue(x-1, y+0) * 32
	out += i.getSingleValue(x+1, y-1) * 64
	out += i.getSingleValue(x+0, y-1) * 128
	out += i.getSingleValue(x-1, y-1) * 256

	return out
}

func (i image) getSingleValue(x int, y int) int {
	_, f := i[key(x, y)]
	if f {
		return 1
	}
	return 0
}

func main() {
	readFile, err := os.ReadFile("2021/20/2021_20.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")
	algo := inFile[0]

	y := 2
	img := make(image)
	for {
		row := inFile[y]
		split := strings.Split(row, "")

		for x, s := range split {
			if s == "#" {
				img[key(x, y-2)] = 1
			}
		}

		y++
		if y == len(inFile) {
			break
		}
	}

	outside := 0

	for i := 0; i < 50; i++ {
		pass1 := make(image)

		minX, maxX, minY, maxY := getBorders(img)

		if outside == 1 {
			// outside is black, color the borders
			for x := minX - 1; x <= maxX+1; x++ {
				img[key(x, minY-1)] = 1
				img[key(x, minY-2)] = 1
				img[key(x, maxY+1)] = 1
				img[key(x, maxY+2)] = 1
			}
			for y := minY - 1; y <= maxY+1; y++ {
				img[key(minX-1, y)] = 1
				img[key(minX-2, y)] = 1
				img[key(maxX+1, y)] = 1
				img[key(maxX+2, y)] = 1
			}
			//  corners
			img[key(minX-2, minY-2)] = 1
			img[key(minX-2, maxY+2)] = 1
			img[key(maxX+2, maxY+2)] = 1
			img[key(maxX+2, minY-2)] = 1
		}

		for y := minY - 1; y <= maxY+1; y++ {
			for x := minX - 1; x <= maxX+1; x++ {
				pass1[key(x, y)] = img.getValue(x, y)
			}
		}

		img = make(image)
		// second pass
		for k, v := range pass1 {
			// match all values using the algo to overwrite the first image with light
			if algo[v] == '#' {
				img[k] = 1
			}
		}

		// flip all that is on the outside
		if algo[0] == '#' {
			outside = (outside + 1) % 2
		}

		if i == 1 {
			fmt.Println("part 1", len(img))
		}
		if i == 49 {
			fmt.Println("part 2", len(img))
		}
	}

	// 5005 - too low
	// 5179 - RIGHT
	// 5213 - wrong
	// 5216 - wrong
	// 5227 - wrong
	// 5355 - too high
}

func getBorders(img image) (int, int, int, int) {
	minX := math.MaxInt
	maxX := math.MinInt
	minY := math.MaxInt
	maxY := math.MinInt

	for k, _ := range img {
		x, y := dekey(k)
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}
	return minX, maxX, minY, maxY
}

func key(x int, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}

func dekey(in string) (x int, y int) {
	x, _ = strconv.Atoi(strings.Split(in, ",")[0])
	y, _ = strconv.Atoi(strings.Split(in, ",")[1])

	return
}
