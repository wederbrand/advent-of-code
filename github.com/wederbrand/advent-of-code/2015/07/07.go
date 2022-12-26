package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"math"
	"time"
)

type gate struct {
	target      string
	operation   string
	a           string
	b           string
	v           int
	cachedValue int
}

func (g *gate) getValue() int {
	if g.cachedValue == math.MinInt {
		ga, founda := gates[g.a]
		gb, foundb := gates[g.b]
		switch g.operation {
		case "DIRECT":
			a := 0
			if founda {
				a = ga.getValue()
			} else {
				a = util.Atoi(g.a)
			}
			g.cachedValue = a
		case "AND":
			a := 0
			if founda {
				a = ga.getValue()
			} else {
				a = util.Atoi(g.a)
			}
			b := 0
			if foundb {
				b = gb.getValue()
			} else {
				b = util.Atoi(g.b)
			}
			g.cachedValue = a & b
		case "OR":
			a := 0
			if founda {
				a = ga.getValue()
			} else {
				a = util.Atoi(g.a)
			}
			b := 0
			if foundb {
				b = gb.getValue()
			} else {
				b = util.Atoi(g.b)
			}
			g.cachedValue = a | b
		case "LSHIFT":
			a := ga.getValue()
			if !founda {
				a = util.Atoi(g.a)
			}
			g.cachedValue = (a << g.v) % 65536
		case "RSHIFT":
			a := 0
			if founda {
				a = ga.getValue()
			} else {
				a = util.Atoi(g.a)
			}
			g.cachedValue = a >> g.v
		case "NOT":
			a := 0
			if founda {
				a = ga.getValue()
			} else {
				a = util.Atoi(g.a)
			}
			g.cachedValue = a ^ 65535
		}
	}
	return g.cachedValue
}

func parseGate(in string) *gate {
	g := gate{cachedValue: math.MinInt}

	_, err := fmt.Sscanf(in, "%s -> %s", &g.a, &g.target)
	if err == nil {
		g.operation = "DIRECT"
		return &g
	}

	_, err = fmt.Sscanf(in, "%s AND %s -> %s", &g.a, &g.b, &g.target)
	if err == nil {
		g.operation = "AND"
		return &g
	}

	_, err = fmt.Sscanf(in, "%s OR %s -> %s", &g.a, &g.b, &g.target)
	if err == nil {
		g.operation = "OR"
		return &g
	}

	_, err = fmt.Sscanf(in, "%s LSHIFT %d -> %s", &g.a, &g.v, &g.target)
	if err == nil {
		g.operation = "LSHIFT"
		return &g
	}

	_, err = fmt.Sscanf(in, "%s RSHIFT %d -> %s", &g.a, &g.v, &g.target)
	if err == nil {
		g.operation = "RSHIFT"
		return &g
	}

	_, err = fmt.Sscanf(in, "NOT %s -> %s", &g.a, &g.target)
	if err == nil {
		g.operation = "NOT"
		return &g
	}

	panic(in)
	return nil
}

var gates = make(map[string]*gate)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2015/07/input.txt", "\n")

	for _, instruction := range inFile {
		g := parseGate(instruction)
		gates[g.target] = g
	}

	// answer from part1
	gates["b"].cachedValue = 46065

	// fmt.Println("part1:", gates["a"].getValue(), "in", time.Since(start))
	// 1674 is too low
	fmt.Println("part2:", gates["a"].getValue(), "in", time.Since(start))
}
