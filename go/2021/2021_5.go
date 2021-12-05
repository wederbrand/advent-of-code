package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

func (l Line) feed(c chan Point) {
	// feed all the points

	dx := 0
	if l.start.x < l.end.x {
		dx = 1
	} else if l.start.x > l.end.x {
		dx = -1
	}
	dy := 0
	if l.start.y < l.end.y {
		dy = 1
	} else if l.start.y > l.end.y {
		dy = -1
	}

	p := l.start
	c <- p

	for ok := true; ok; ok = p != l.end {
		p.x += dx
		p.y += dy
		c <- p
	}

	close(c)
}

func main() {
	file, err := os.Open("2021/2021_5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	matcher := regexp.MustCompile("(\\d+),(\\d+) -> (\\d+),(\\d+)")
	scanner := bufio.NewScanner(file)
	input := make([]Line, 0)
	for scanner.Scan() {
		text := scanner.Text()
		submatch := matcher.FindStringSubmatch(text)
		var l Line

		l.start.x, _ = strconv.Atoi(submatch[1])
		l.start.y, _ = strconv.Atoi(submatch[2])

		l.end.x, _ = strconv.Atoi(submatch[3])
		l.end.y, _ = strconv.Atoi(submatch[4])

		input = append(input, l)
	}

	floor := make(map[string]int)

	for _, l := range input {
		c := make(chan Point)
		go l.feed(c)
		for p := range c {
			floor[p.key()]++
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
