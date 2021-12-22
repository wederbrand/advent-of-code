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
	z int
}

type cube struct {
	on bool
	p1 point
	p2 point
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func (c cube) valid() bool {
	return c.p1.x <= c.p2.x && c.p1.y <= c.p2.y && c.p1.z <= c.p2.z
}

func (c cube) sum() int {
	return (c.p2.x - c.p1.x + 1) * (c.p2.y - c.p1.y + 1) * (c.p2.z - c.p1.z + 1)
}

func (c cube) key() [2]point {
	return [2]point{c.p1, c.p2}
}

// this will at time produce invalid cubes but that can be tested for later
// also, the on/off is always false when returned
func (c cube) getOverlappingCube(oldCube cube) cube {
	return cube{
		p1: point{max(c.p1.x, oldCube.p1.x), max(c.p1.y, oldCube.p1.y), max(c.p1.z, oldCube.p1.z)},
		p2: point{min(c.p2.x, oldCube.p2.x), min(c.p2.y, oldCube.p2.y), min(c.p2.z, oldCube.p2.z)},
	}
}

func main() {
	readFile, err := os.ReadFile("2021/22/2021_22.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")
	re := regexp.MustCompile("(.+) x=(-?\\d+)..(-?\\d+),y=(-?\\d+)..(-?\\d+),z=(-?\\d+)..(-?\\d+)")

	cubes := make([]cube, 0)
	for _, str := range inFile {
		submatch := re.FindStringSubmatch(str)
		onOff := submatch[1]

		x1, _ := strconv.Atoi(submatch[2])
		x2, _ := strconv.Atoi(submatch[3])

		y1, _ := strconv.Atoi(submatch[4])
		y2, _ := strconv.Atoi(submatch[5])

		z1, _ := strconv.Atoi(submatch[6])
		z2, _ := strconv.Atoi(submatch[7])

		cubes = append(cubes, cube{
			on: onOff == "on",
			p1: point{x1, y1, z1},
			p2: point{x2, y2, z2},
		})
	}

	part1(cubes)
	part2(cubes)
}

func part2(cubes []cube) {
	onCubes := make([]cube, 0)

	for _, c := range cubes {
		temp := make([]cube, 0)
		if c.on {
			// if "turn on" add to list of onCubes
			temp = append(temp, c)
		}

		// we need to compensate all overlapping cubes since we are not removing from the original cube
		for _, oldCube := range onCubes {
			compensationCube := c.getOverlappingCube(oldCube)
			if compensationCube.valid() {
				compensationCube.on = !oldCube.on
				temp = append(temp, compensationCube)
			}
		}

		onCubes = append(onCubes, temp...)
	}

	sum := 0
	for _, onCube := range onCubes {
		if onCube.on {
			sum += onCube.sum()
		} else {
			sum -= onCube.sum()
		}
	}

	fmt.Println("part 2", sum)
}

func part1(cubes []cube) {
	reactor := make(map[point]bool)

	for _, c := range cubes {
		for x := c.p1.x; x <= c.p2.x && x >= -50 && x <= 50; x++ {
			for y := c.p1.y; y <= c.p2.y && y >= -50 && y <= 50; y++ {
				for z := c.p1.z; z <= c.p2.z && z >= -50 && z <= 50; z++ {
					p := point{x, y, z}
					if c.on {
						reactor[p] = true
					} else {
						delete(reactor, p)
					}
				}
			}
		}
	}

	fmt.Println("part 1", len(reactor))
}
