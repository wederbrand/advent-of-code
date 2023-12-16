package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"strings"
	"time"
)

type Pipe struct {
	r rune

	x int
	y int

	n bool
	s bool
	e bool
	w bool

	touched bool
}

func (p Pipe) getNext(dir Dir) Dir {
	if p.n && dir != S {
		return N
	}
	if p.s && dir != N {
		return S
	}
	if p.e && dir != W {
		return E
	}
	if p.w && dir != E {
		return W
	}

	panic("ho ho")
}

func newPipe(r rune, x int, y int) *Pipe {
	p := new(Pipe)
	p.r = r
	p.x = x
	p.y = y

	p.n = strings.ContainsRune("|LJ", r)
	p.s = strings.ContainsRune("|7F", r)
	p.e = strings.ContainsRune("-LF", r)
	p.w = strings.ContainsRune("-7J", r)
	return p
}

func main() {
	startTimer := time.Now()
	inFile := util.GetFileContents("2023/10/input.txt", "\n")

	m := make(map[string]*Pipe)
	var start *Pipe
	for y, s := range inFile {
		for x, r := range s {
			p := newPipe(r, x, y)
			if p.r == 'S' {
				start = p
			}
			m[util.IntKey(x, y)] = p
		}
	}

	// find directions of start pipe. also find one initial direction
	var dir Dir
	n, foundN := m[util.IntKey(start.x, start.y-1)]
	if foundN && n.s {
		start.n = true
		dir = N
	}
	s, foundS := m[util.IntKey(start.x, start.y+1)]
	if foundS && s.n {
		start.s = true
		dir = S
	}
	e, foundE := m[util.IntKey(start.x+1, start.y)]
	if foundE && e.w {
		start.e = true
		dir = E
	}
	w, foundW := m[util.IntKey(start.x-1, start.y)]
	if foundW && w.e {
		start.w = true
		dir = W
	}

	sRune := 'S'
	if start.n && start.s {
		sRune = '|'
	}
	if start.e && start.w {
		sRune = '-'
	}
	if start.n && start.e {
		sRune = 'L'
	}
	if start.n && start.w {
		sRune = 'J'
	}
	if start.s && start.w {
		sRune = '7'
	}
	if start.s && start.e {
		sRune = 'F'
	}

	fmt.Println("parsing:", time.Since(startTimer))
	startTimer = time.Now()

	dist := 0
	curr := start
	for curr != start || dist == 0 {
		curr.touched = true
		dist++
		curr, _ = m[util.IntKey(curr.x+dir[0], curr.y+dir[1])]
		dir = curr.getNext(dir)
	}

	part1 := dist / 2
	fmt.Println("part1: ", part1, "in", time.Since(startTimer))
	startTimer = time.Now()

	part2 := 0
	for y, s := range inFile {
		in := false
		var lastCorner rune
		for x, r := range s {
			p := m[util.IntKey(x, y)]

			if r == 'S' {
				// the start rune needs special attention
				r = sRune
			}

			if !p.touched || r == '.' {
				// treat all pipes not part of the main loop as empty
				// ie, don't flip in/out and count if currently enclosed
				if in {
					part2++
				}
			} else if r == '|' {
				in = !in
			} else if r == 'F' || r == 'L' {
				lastCorner = r
			} else if r == 'J' && lastCorner == 'F' {
				in = !in
			} else if r == '7' && lastCorner == 'L' {
				in = !in
			}
		}
	}
	fmt.Println("part2: ", part2, "in", time.Since(startTimer))
}
