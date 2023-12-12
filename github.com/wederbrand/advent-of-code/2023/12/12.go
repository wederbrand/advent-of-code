package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strings"
	"time"
)

type State struct {
	pattern string
	numbers string
}

var cache map[State]int

// put value in cache and return it
func cacheIt(s State, v int) int {
	cache[s] = v
	return v
}

func main() {
	startTimer := time.Now()
	inFile := util.GetFileContents("2023/12/input.txt", "\n")

	part1 := 0
	part2 := 0
	for _, s := range inFile {
		cache = make(map[State]int)
		split := strings.Split(s, " ")
		state := State{split[0], split[1]}

		solved := solve(state)
		part1 += solved

		for i := 0; i < 4; i++ {
			state.pattern = state.pattern + "?" + split[0]
			state.numbers = state.numbers + "," + split[1]
		}

		solved = solve(state)
		part2 += solved
	}

	fmt.Println("part1: ", part1, "in", time.Since(startTimer))
	fmt.Println("part2: ", part2, "in", time.Since(startTimer))
}

func solve(s State) int {
	cachedValue, found := cache[s]
	if found {
		return cachedValue
	}
	if len(s.pattern) == 0 {
		if len(s.numbers) == 0 {
			return cacheIt(s, 1)
		} else if len(s.numbers) > 0 {
			return cacheIt(s, 0)
		}
	}

	if len(s.numbers) == 0 {
		// no patterns to match
		// error if the next rune is '#'
		if s.pattern[0] == '#' {
			return cacheIt(s, 0)
		}
		return cacheIt(s, solve(State{s.pattern[1:], s.numbers}))
	}

	// if . remove it and send down
	if s.pattern[0] == '.' {
		return cacheIt(s, solve(State{s.pattern[1:], s.numbers}))
	}

	// if # pick all it needs and send the rest down
	if s.pattern[0] == '#' {
		indexOfRuneAfterPattern := util.Atoi(strings.Split(s.numbers, ",")[0])
		if len(s.pattern) < indexOfRuneAfterPattern {
			// pattern too short
			return cacheIt(s, 0)
		}

		// iterate until indexOfRuneAfterPattern
		for i := 0; i < indexOfRuneAfterPattern; i++ {
			// check if it's # or ? (assume #)
			if s.pattern[i] == '.' {
				// impossible pattern
				return cacheIt(s, 0)
			}
		}

		// after iteration, check next character
		if len(s.pattern) <= indexOfRuneAfterPattern {
			// eol is fine
			join := strings.Join(strings.Split(s.numbers, ",")[1:], ",")
			return cacheIt(s, solve(State{"", join}))
		}
		nextRune := s.pattern[indexOfRuneAfterPattern]
		if nextRune == '.' {
			// . is fine
			join := strings.Join(strings.Split(s.numbers, ",")[1:], ",")
			nextState := State{s.pattern[indexOfRuneAfterPattern:], join}
			return cacheIt(s, solve(nextState))
		}
		if nextRune == '?' {
			// ? is fine, but assume '.'
			join := strings.Join(strings.Split(s.numbers, ",")[1:], ",")
			dotState := State{"", join}
			dotState.pattern = s.pattern[indexOfRuneAfterPattern+1:]
			return cacheIt(s, solve(dotState))
		}
		if nextRune == '#' {
			// too long
			return cacheIt(s, 0)
		}
	}

	// if ? replace it with . and # and send both down
	if s.pattern[0] == '?' {
		dotState := State{s.pattern[1:], s.numbers}
		dotSum := solve(dotState)
		poundState := State{s.pattern[1:], s.numbers}
		poundState.pattern = "#" + poundState.pattern
		poundSum := solve(poundState)
		return dotSum + poundSum
	}

	panic("ho no")
}
