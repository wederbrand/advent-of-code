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

func (p point) key() string {
	return strconv.Itoa(p.x) + "," + strconv.Itoa(p.y)
}

type fold struct {
	updown string
	xy     int
}

func main() {
	readFile, err := os.ReadFile("2021/13/2021_13.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")
	foldRE := regexp.MustCompile("fold along (.)=(\\d+)")

	folds := make([]fold, 0)
	paper := make(map[string]*point)
	for _, s := range inFile {
		if len(s) == 0 {
			continue
		} else if foldRE.MatchString(s) {
			// fold instruction
			submatch := foldRE.FindStringSubmatch(s)
			atoi, _ := strconv.Atoi(submatch[2])
			folds = append(folds, fold{
				updown: submatch[1],
				xy:     atoi,
			})
		} else {
			split := strings.Split(strings.TrimSpace(s), ",")
			p := new(point)
			p.x, _ = strconv.Atoi(split[0])
			p.y, _ = strconv.Atoi(split[1])
			paper[p.key()] = p
		}
	}

	for i := 0; i < len(folds); i++ {
		foldOne(paper, folds[i])
		if i == 0 {
			fmt.Println("part 1", len(paper))
		}
	}

	fmt.Println("part2")
	printIt(paper)

}

func printIt(paper map[string]*point) {
	maxX := 0
	maxY := 0
	for _, p := range paper {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}

	}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			_, found := paper[point{x, y}.key()]
			if found {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func foldOne(paper map[string]*point, f fold) {
	maxX := 0
	maxY := 0
	for _, p := range paper {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}

	}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			p, found := paper[point{x, y}.key()]
			if !found {
				continue
			}

			delete(paper, p.key())
			if f.updown == "y" && p.y > f.xy {
				// fold it up
				p.y = f.xy - (p.y - f.xy)
			} else if f.updown == "x" && p.x > f.xy {
				// fold it left
				p.x = f.xy - (p.x - f.xy)
			}
			paper[p.key()] = p

		}
	}
}
