package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type cube struct {
	x int
	y int
	z int
}

func main() {
	readFile, err := os.ReadFile("2021/22/2021_22.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")
	/*
		on x=11..13,y=11..13,z=11..13
		off x=9..11,y=9..11,z=9..11
	*/
	re := regexp.MustCompile("(.+) x=(-?\\d+)..(-?\\d+),y=(-?\\d+)..(-?\\d+),z=(-?\\d+)..(-?\\d+)")

	reactor := make(map[cube]bool)
	for _, str := range inFile {
		submatch := re.FindStringSubmatch(str)
		onOff := submatch[1]

		x1, _ := strconv.Atoi(submatch[2])
		x2, _ := strconv.Atoi(submatch[3])

		y1, _ := strconv.Atoi(submatch[4])
		y2, _ := strconv.Atoi(submatch[5])

		z1, _ := strconv.Atoi(submatch[6])
		z2, _ := strconv.Atoi(submatch[7])

		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				for z := z1; z <= z2; z++ {
					c := cube{x, y, z}
					if onOff == "on" {
						reactor[c] = true
					} else {
						delete(reactor, c)
					}
				}
			}
		}
	}

	fmt.Println(len(reactor))

}
