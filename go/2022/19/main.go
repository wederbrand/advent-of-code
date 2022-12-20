package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const ORE = 0
const CLAY = 1
const OBSIDIAN = 2

type blueprint struct {
	id          int
	ore         [3]int
	clay        [3]int
	obsidian    [3]int
	geode       [3]int
	maxOre      int
	maxClay     int
	maxObsidian int
}

func newBlueprint(s string) *blueprint {
	b := &blueprint{}
	fmt.Sscanf(s, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.\n", &b.id, &b.ore[ORE], &b.clay[ORE], &b.obsidian[ORE], &b.obsidian[CLAY], &b.geode[ORE], &b.geode[OBSIDIAN])
	b.figureOutNeeds()

	return b
}

func (b *blueprint) figureOutNeeds() {
	b.maxOre = maxInt(b.ore[ORE], b.clay[ORE], b.obsidian[ORE], b.geode[ORE])
	b.maxClay = maxInt(b.ore[CLAY], b.clay[CLAY], b.obsidian[CLAY], b.geode[CLAY])
	b.maxObsidian = maxInt(b.ore[OBSIDIAN], b.clay[OBSIDIAN], b.obsidian[OBSIDIAN], b.geode[OBSIDIAN])
}

func maxInt(a, b, c, d int) int {
	max := 0
	if a > max {
		max = a
	}
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	if d > max {
		max = d
	}

	return max
}

func (b *blueprint) needOre(s *state) bool {
	return s.oreRobot < b.maxOre
}

func (b *blueprint) needClay(s *state) bool {
	return s.clayRobot < b.maxClay
}
func (b *blueprint) needObsidian(s *state) bool {
	return s.obsidianRobot < b.maxObsidian
}

type state struct {
	timeLeft      int
	ore           int
	clay          int
	obsidian      int
	geode         int
	oreRobot      int
	clayRobot     int
	obsidianRobot int
	geodeRobot    int
}

func (s *state) getStuff() {
	s.ore += s.oreRobot
	s.clay += s.clayRobot
	s.obsidian += s.obsidianRobot
	s.geode += s.geodeRobot
}

func (s *state) clone() *state {
	return &state{
		timeLeft:      s.timeLeft,
		ore:           s.ore,
		clay:          s.clay,
		obsidian:      s.obsidian,
		geode:         s.geode,
		oreRobot:      s.oreRobot,
		clayRobot:     s.clayRobot,
		obsidianRobot: s.obsidianRobot,
		geodeRobot:    s.geodeRobot,
	}
}

func (s *state) maxPossible() int {
	// naive approach, assume one new geode bot is created every minute that remains
	geodes := s.geode

	nbrOfBots := s.geodeRobot
	for i := 0; i < s.timeLeft; i++ {
		geodes += nbrOfBots
		nbrOfBots++
	}

	return geodes
}

type queue struct {
	states []*state
}

func newQueue() *queue {
	q := new(queue)
	q.states = make([]*state, 0)

	return q
}

func (q *queue) add(s *state) {
	// add state to queue, sort by least time left
	if len(q.states) == 0 {
		q.states = append(q.states, s)
	} else {
		for i, s2 := range q.states {
			if s.timeLeft < s2.timeLeft {
				// insert here
				q.states = append(q.states, nil)   // make room for copy
				copy(q.states[i+1:], q.states[i:]) // make room at index i, overwriting the nil
				q.states[i] = s                    // insert s
				return
			}
		}
		q.states = append(q.states, s)
	}
}

func (q *queue) dequeue() *state {
	p := q.states[0]
	q.states = q.states[1:len(q.states)]
	return p
}

func (b *blueprint) calculate() int {
	q := newQueue()

	initialState := state{
		timeLeft: 24,
		oreRobot: 1,
	}

	q.add(&initialState)

	max := 0
	seen := make(map[state]bool)
	for len(q.states) > 0 {
		//		fmt.Println(len(q.states), len(seen), max)
		s := q.dequeue()
		_, found := seen[*s]
		if found {
			continue
		}
		seen[*s] = true
		if s.maxPossible() <= max {
			continue
		}
		if s.timeLeft == 0 {
			if s.geode > max {
				max = s.geode
			}
			continue
		}
		s.timeLeft -= 1
		if s.ore >= b.geode[ORE] && s.obsidian >= b.geode[OBSIDIAN] {
			s2 := s.clone()
			s2.ore -= b.geode[ORE]
			s2.obsidian -= b.geode[OBSIDIAN]
			s2.getStuff()
			s2.geodeRobot += 1
			q.add(s2)
		}
		if s.ore >= b.obsidian[ORE] && s.clay >= b.obsidian[CLAY] && b.needObsidian(s) {
			s2 := s.clone()
			s2.ore -= b.obsidian[ORE]
			s2.clay -= b.obsidian[CLAY]
			s2.getStuff()
			s2.obsidianRobot += 1
			q.add(s2)
		}
		if s.ore >= b.clay[ORE] && b.needClay(s) {
			s2 := s.clone()
			s2.ore -= b.clay[ORE]
			s2.getStuff()
			s2.clayRobot += 1
			q.add(s2)
		}
		if s.ore >= b.ore[ORE] && b.needOre(s) {
			s2 := s.clone()
			s2.ore -= b.ore[ORE]
			s2.getStuff()
			s2.oreRobot += 1
			q.add(s2)
		}
		s2 := s.clone()
		s2.getStuff()
		q.add(s2)
	}

	return b.id * max
}

func main() {
	start := time.Now()
	readFile, err := os.ReadFile("19/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	blueprints := make([]*blueprint, 0)

	part1 := 0
	for _, s := range inFile {
		b := newBlueprint(s)
		blueprints = append(blueprints, b)

		part1 += b.calculate()
	}

	fmt.Println("part1:", part1, "in", time.Since(start))
}
