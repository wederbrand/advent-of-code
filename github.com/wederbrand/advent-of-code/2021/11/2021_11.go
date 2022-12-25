package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type point struct {
	x     int
	y     int
	value int
}

type key [2]int

func main() {
	readFile, err := os.ReadFile("2021/11/2021_11.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	octo := make(map[key]*point)
	for y := 0; y < 10; y++ {
		line := []byte(inFile[y])
		for x := 0; x < 10; x++ {
			p := point{x, y, int(line[x]) - 48}
			octo[key{x, y}] = &p
		}
	}

	flashes := 0
	for i := 0; ; i++ {
		increase(octo)
		explode(octo)
		newFlashes := cooldown(octo)
		flashes += newFlashes
		if i == 99 {
			fmt.Println("part 1", flashes)
		}
		if newFlashes == 100 {
			fmt.Println("part 2", i+1)
			os.Exit(0)
		}
	}

}

func cooldown(octo map[key]*point) int {
	f := 0
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			key := key{x, y}
			if octo[key].value == 0 {
				f++
			}
		}
	}
	return f
}

func printMap(octo map[key]*point) {
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			key := key{x, y}
			fmt.Print(octo[key].value)
		}
		fmt.Println()
	}
	fmt.Println()
}

func explode(octo map[key]*point) {
	c := make(chan key, 100)
	defer close(c)

	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			key := key{x, y}
			if octo[key].value > 9 {
				c <- key
				octo[key].value = 0
			}
		}
	}

	for len(c) > 0 {
		//fmt.Println("a que now", len(c))
		k := <-c
		//fmt.Println("b que now", len(c))
		// explode surrounding squids
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				surroundingKey := key{k[0] + dx, k[1] + dy}
				p, found := octo[surroundingKey]
				if found && p.value != 0 {
					p.value++
					if p.value > 9 {
						//fmt.Println("c que now", len(c))
						c <- surroundingKey
						octo[surroundingKey].value = 0
						//fmt.Println("d que now", len(c))
					}
				}
			}
		}
	}
}

func increase(octo map[key]*point) {
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			key := key{x, y}
			octo[key].value++
		}
	}
}
