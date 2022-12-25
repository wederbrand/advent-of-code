package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func (p Point) key() string {
	return strconv.Itoa(p.x) + "," + strconv.Itoa(p.y)
}

type Line struct {
	start Point
	end   Point
}

func (line Line) feedPoints(c chan Point) {
	// feedPoints all the points

	dx := 0
	if line.start.x < line.end.x {
		dx = 1
	} else if line.start.x > line.end.x {
		dx = -1
	}
	dy := 0
	if line.start.y < line.end.y {
		dy = 1
	} else if line.start.y > line.end.y {
		dy = -1
	}

	p := line.start
	c <- p

	for ok := true; ok; ok = p != line.end {
		p.x += dx
		p.y += dy
		c <- p
	}

	close(c)
}

func main() {
	readFile, err := os.ReadFile("2021/05/2021_05.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	matcher := regexp.MustCompile("(\\d+),(\\d+) -> (\\d+),(\\d+)")
	input := make([]Line, 0)
	for _, text := range inFile {
		submatch := matcher.FindStringSubmatch(text)
		var l Line

		l.start.x, _ = strconv.Atoi(submatch[1])
		l.start.y, _ = strconv.Atoi(submatch[2])

		l.end.x, _ = strconv.Atoi(submatch[3])
		l.end.y, _ = strconv.Atoi(submatch[4])

		input = append(input, l)
	}

	floor := make(map[string]int)

	for _, line := range input {
		c := make(chan Point)
		go line.feedPoints(c)
		for point := range c {
			floor[point.key()]++
		}
	}

	twos := 0
	for _, v := range floor {
		if v >= 2 {
			twos++
		}
	}

	fmt.Println(twos)

	os.Exit(1)
}
