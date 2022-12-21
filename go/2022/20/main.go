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
		r2.Value = atoi
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
	for i := 0; i < len(lookup); i++ {
		src := lookup[i]
		if src.Value.(int) == 0 {
			continue
		}
		prev := src.Prev()
		unlink := prev.Unlink(1)
		n := src.Value.(int)
		dst := prev.Move(n)
		dst.Link(unlink)

	}

	part1 := zero.Move(1000).Value.(int)
	part1 += zero.Move(2000).Value.(int)
	part1 += zero.Move(3000).Value.(int)

	fmt.Println("part1:", part1, "in", time.Since(start))
}
