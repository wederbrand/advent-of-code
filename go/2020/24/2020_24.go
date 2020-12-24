package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type pos struct {
	x int
	y int
}

func (p *pos) move(dir string) {
	switch dir {
	case "ne":
		p.y--
	case "nw":
		p.y--
		p.x--
	case "se":
		p.y++
		p.x++
	case "sw":
		p.y++
	case "w":
		p.x--
	case "e":
		p.x++
	}
}

func main() {
	readFile, err := ioutil.ReadFile("24/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	floor := make(map[pos]bool)
	input := strings.Split(strings.TrimSpace(string(readFile)), "\n")
	for _, inputTile := range input {
		p := new(pos)
		for i := 0; i < len(inputTile); i++ {
			// 			fmt.Println(i, p, inputTile[i:])
			r := inputTile[i]
			switch r {
			case 'n':
				i++
				r2 := inputTile[i]
				switch r2 {
				case 'e':
					p.move("ne")
				case 'w':
					p.move("nw")
				}
			case 's':
				i++
				r2 := inputTile[i]
				switch r2 {
				case 'e':
					p.move("se")
				case 'w':
					p.move("sw")
				}
			case 'e':
				p.move("e")
			case 'w':
				p.move("w")
			}
		}
		floor[*p] = !floor[*p]
		fmt.Println(p, floor[*p])
	}

	cnt := 0
	for _, black := range floor {
		if black {
			cnt++
		}
	}

	fmt.Println(cnt)
}
