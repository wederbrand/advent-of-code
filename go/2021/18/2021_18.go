package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	left  *pair
	right *pair
	value int
}

func (p *pair) leaf() bool {
	return p.left == nil
}

func (p *pair) print() string {
	out := ""
	if !p.leaf() {
		out += fmt.Sprint("[")
		out += p.left.print()
		out += fmt.Sprint(",")
		out += p.right.print()
		out += fmt.Sprint("]")
	} else {
		out += fmt.Sprint(p.value)
	}
	return out
}

func (p *pair) reduce() {
outer:
	for {
		// explode until none to explode
		siblings := p.getSiblings()
		if p.explode(0, siblings) {
			continue outer
		}
		// split until none left to split
		for _, sibling := range siblings {
			if sibling.value >= 10 {
				sibling.split()
				continue outer
			}
		}
		// return if all is done
		return
	}
}

func (p *pair) explode(d int, siblings []*pair) bool {
	if p.leaf() {
		return false
	}
	if p.left.leaf() && p.right.leaf() {
		i := 0 // index of left leaf
		var sibling *pair
		for i, sibling = range siblings {
			if sibling == p.left {
				break
			}
		}
		// two values under me, explode if deep enough
		if d >= 4 {
			lValue := p.left.value
			rValue := p.right.value
			p.left = nil
			p.right = nil
			p.value = 0 // but already 0
			if i-1 >= 0 {
				siblings[i-1].value += lValue
			}
			if i+2 < len(siblings) {
				siblings[i+2].value += rValue
			}
			return true
		}
		return false
	} else {
		exploded := p.left.explode(d+1, siblings)
		if !exploded {
			exploded = p.right.explode(d+1, siblings)
		}
		return exploded
	}
}

func (p *pair) split() {
	p.left = new(pair)
	p.right = new(pair)
	p.left.value = p.value / 2
	p.right.value = p.value - p.left.value
	p.value = 0
}

func (p *pair) getSiblings() []*pair {
	out := make([]*pair, 0)
	if !p.leaf() {
		out = append(out, p.left.getSiblings()...)
		out = append(out, p.right.getSiblings()...)
	} else {
		out = append(out, p)
	}
	return out

}

func (p *pair) magnitude() int {
	out := 0
	if !p.leaf() {
		out += 3 * p.left.magnitude()
		out += 2 * p.right.magnitude()
	} else {
		return p.value
	}
	return out
}

func newPair(in string) (*pair, string) {
	p := new(pair)
	if in[0] == '[' {
		var rest string
		p.left, rest = newPair(in[1:])
		p.right, rest = newPair(rest[1:]) // skip the comma
		return p, rest[1:]                // skip the ]
	} else {
		p.value, _ = strconv.Atoi(string(in[0]))
		return p, in[1:]
	}
}

func main() {
	readFile, err := os.ReadFile("2021/18/2021_18.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")
	var p *pair
	all := make([]string, 0)
	for _, str := range inFile {
		all = append(all, str)
		parsed, _ := newPair(str)
		if p == nil {
			p = parsed
		} else {
			// add it
			newPair := new(pair)
			newPair.left = p
			newPair.right = parsed
			p = newPair
			p.reduce()
		}
	}

	fmt.Println("part 1", p.magnitude())

	maxMag := math.MinInt
	for _, s1 := range all {
		for _, s2 := range all {
			if s1 == s2 {
				continue
			}

			p1, _ := newPair(s1)
			p2, _ := newPair(s2)

			newPair := new(pair)
			newPair.left = p1
			newPair.right = p2
			newPair.reduce()
			magnitude := newPair.magnitude()
			if magnitude > maxMag {
				maxMag = magnitude
			}
		}
	}

	fmt.Println("part 2", maxMag)

}
