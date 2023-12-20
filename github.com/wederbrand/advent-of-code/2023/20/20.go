package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strings"
	"time"
)

type Module struct {
	name      string
	operation string
	state     bool // % true == on, & true == high
	in        map[*Module]bool
	outString []string
	out       []*Module
}

type Pulse struct {
	headingFor *Module
	comingFrom *Module
	state      bool // true == high
}

func newModule(s string) *Module {
	m := new(Module)
	m.in = make(map[*Module]bool)
	split := strings.Split(s, " -> ")
	if split[0][0] == '%' {
		m.operation = "%"
		m.name = split[0][1:]
	} else if split[0][0] == '&' {
		m.operation = "&"
		m.name = split[0][1:]
	} else {
		m.name = split[0]
	}

	if len(split[1]) != 0 {
		split2 := strings.Split(split[1], ", ")
		for _, out := range split2 {
			m.outString = append(m.outString, out)
		}
	}
	return m
}

func main() {
	startTimer := time.Now()
	inFile := util.GetFileContents("2023/20/input.txt", "\n")

	modules := make(map[string]*Module)
	keys := make([]string, 0)
	// first pass, create all modules and save output as outString
	for _, s := range inFile {
		m := newModule(s)
		modules[m.name] = m
		keys = append(keys, m.name)
	}

	// second pass, match outString to in and out
	for _, m := range modules {
		for _, o := range m.outString {
			out, found := modules[o]
			if !found {
				// create a dead output
				out = newModule(o + " -> ")
				modules[out.name] = out
				keys = append(keys, out.name)
			}
			m.out = append(m.out, out)
			out.in[m] = false
		}
	}

	fmt.Println("parsing in", time.Since(startTimer))
	startTimer = time.Now()

	part1Low := 0
	part1High := 0
	seenFactors := make(map[string]int)
outer:
	for i := 0; ; i++ {
		q := make([]Pulse, 0)
		startPulse := Pulse{modules["broadcaster"], nil, false}
		q = append(q, startPulse)
		for len(q) > 0 {
			p := q[0]
			q = q[1:]

			if p.state {
				part1High++
			} else {
				part1Low++
			}

			m := p.headingFor
			if m.operation == "%" {
				// Flip-flop modules (prefix %) are either on or off; they are initially off. If
				// a flip-flop module receives a high pulse, it is ignored and nothing happens.
				// However, if a flip-flop module receives a low pulse, it flips between on and
				// off. If it was off, it turns on and sends a high pulse. If it was on, it turns
				// off and sends a low pulse.
				if p.state {
					// high, do nothing
				} else {
					// low flip and resend
					m.state = !m.state
					for _, nextModule := range m.out {
						pulse := Pulse{nextModule, m, m.state}
						q = append(q, pulse)
					}
				}
			} else if m.operation == "&" {
				// Conjunction modules (prefix &) remember the type of the most recent pulse
				// received from each of their connected input modules; they initially default to
				// remembering a low pulse for each input. When a pulse is received, the
				// conjunction module first updates its memory for that input. Then, if it
				// remembers high pulses for all inputs, it sends a low pulse; otherwise, it
				// sends a high pulse.
				m.in[p.comingFrom] = p.state
				allHigh := true
				for _, b := range m.in {
					allHigh = allHigh && b
				}

				if allHigh {
					// send low
					for _, nextModule := range m.out {
						pulse := Pulse{nextModule, m, false}
						q = append(q, pulse)
					}
				} else {
					// send high
					for _, nextModule := range m.out {
						pulse := Pulse{nextModule, m, true}
						q = append(q, pulse)
					}
				}
			} else {
				// must be broadcaster
				for _, nextModule := range m.out {
					pulse := Pulse{nextModule, m, p.state}
					q = append(q, pulse)
				}
			}

			// for part2, check rx
			target := modules["rx"]
			// reading the input shows that rx has a single input, mf
			for module := range target.in {
				for m2, b := range module.in {
					if b {
						_, found := seenFactors[m2.name]
						if !found {
							seenFactors[m2.name] = i + 1
							if len(seenFactors) == len(module.in) {
								part2 := 1
								for _, factor := range seenFactors {
									part2 = util.Lcd([]int{part2, factor})
								}
								fmt.Println("part2: ", part2, "in", time.Since(startTimer))
								break outer
							}
						}
					}
				}
			}
		}
		if i+1 == 1000 {
			part1 := part1Low * part1High
			fmt.Println("part1: ", part1, "in", time.Since(startTimer))
			startTimer = time.Now()
		}
	}
}
