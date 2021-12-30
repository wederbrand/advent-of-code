package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func main() {
	readFile, err := os.ReadFile("2015/06/2015_06.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")
	re := regexp.MustCompile("(.*) (\\d+),(\\d+) through (\\d+),(\\d+)")
	lights := make(map[point]int)
	for _, instruction := range inFile {
		submatch := re.FindStringSubmatch(instruction)
		minX, _ := strconv.Atoi(submatch[2])
		minY, _ := strconv.Atoi(submatch[3])
		maxX, _ := strconv.Atoi(submatch[4])
		maxY, _ := strconv.Atoi(submatch[5])

		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {
				p := point{x, y}
				if submatch[1] == "turn on" {
					lights[p]++
				} else if submatch[1] == "turn off" {
					lights[p]--
					if lights[p] <= 0 {
						delete(lights, p)
					}
				} else {
					lights[p] += 2
				}
			}
		}
	}

	sum := 0
	for _, i := range lights {
		sum += i
	}

	fmt.Println("part 2", sum)
}
