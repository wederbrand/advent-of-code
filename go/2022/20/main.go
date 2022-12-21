package main

import (
	"container/ring"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	readFile, err := os.ReadFile("20/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	// build ring
	var r *ring.Ring
	lookup := make(map[int]*ring.Ring)
	var zero *ring.Ring
	for i, s := range inFile {
		atoi, _ := strconv.Atoi(s)
		r2 := ring.New(1)
		r2.Value = atoi * 811589153
		lookup[i] = r2
		if atoi == 0 {
			zero = r2
		}
		if r != nil {
			r.Link(r2)
		}
		r = r2
	}

	// rotate
	for times := 0; times < 10; times++ {
		for i := 0; i < len(lookup); i++ {
			src := lookup[i]
			if src.Value.(int) == 0 {
				continue
			}
			prev := src.Prev()
			unlink := prev.Unlink(1)
			n := src.Value.(int)
			n %= len(lookup) - 1
			dst := prev.Move(n)
			dst.Link(unlink)
		}
	}

	part21 := zero.Move(1000).Value.(int)
	part22 := zero.Move(2000).Value.(int)
	part23 := zero.Move(3000).Value.(int)

	part2 := part21 + part22 + part23

	fmt.Println("part2:", part2, "in", time.Since(start))
}
