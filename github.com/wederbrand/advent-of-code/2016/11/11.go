package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/priorityqueue"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Floor struct {
	generators *[]string
	microchips *[]string
}

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2016/11/input.txt", "\n")

	// parsing
	generatorRE := regexp.MustCompile(`(\w+?) generator`)
	microChipRE := regexp.MustCompile(`(\w+?)-compatible`)

	floors := make([]Floor, 4)
	for i, line := range inFile {
		generators := generatorRE.FindAllString(line, -1)
		microchips := microChipRE.FindAllString(line, -1)
		floors[i] = Floor{&generators, &microchips}
	}

	part1 := doIt(floors)

	var part2 = 0
	fmt.Println("Part 1: ", part1, "in", time.Since(start))
	fmt.Println("Part 2: ", part2, "in", time.Since(start))
}

type State struct {
	elevator int
	floors   []Floor
}

func (s State) key() string {
	key := strconv.Itoa(s.elevator) + " "
	for _, floor := range s.floors {
		slices.Sort(*floor.generators)
		for _, generator := range *floor.generators {
			key += generator + " "
		}

		slices.Sort(*floor.microchips)
		for _, microchip := range *floor.microchips {
			key += microchip + " "
		}

		key += "|"
	}

	return key
}

func (s State) isValid() bool {
	for _, floor := range s.floors {
		if floor.generators == nil || len(*floor.generators) == 0 {
			continue
		}

		// find any unprotected microchips
	outer:
		for _, microchip := range *floor.microchips {
			generatorName := strings.Replace(microchip, "-compatible", " generator", -1)
			for _, generator := range *floor.generators {
				if generator == generatorName {
					// unprotected microchip and we know there is at least one generator. It's fried.
					continue outer
				}
			}
			return false
		}
	}

	return true
}

func doIt(floors []Floor) int {
	q := priorityqueue.NewQueue()
	initial := priorityqueue.State{
		Data:     State{0, floors},
		Priority: 0,
	}
	q.Add(&initial)

	seen := make(map[string]int)
	seen[initial.Data.(State).key()] = 0
	for q.HasNext() {
		qs := q.Next()

		// pick all possible moves
		s := qs.Data.(State)

		if strings.HasPrefix(s.key(), "3 |||") {
			return qs.Priority
		}

		for _, nextS := range possibleMoves(s) {
			// if not seen or cheaper, add to queue
			oldVal, found := seen[nextS.key()]
			nextNumberOfSteps := qs.Priority + 1
			if found && oldVal < nextNumberOfSteps {
				continue
			}

			// add to queue
			q.Add(&priorityqueue.State{Data: nextS, Priority: nextNumberOfSteps})
			//fmt.Println("added", nextS.key())
			seen[nextS.key()] = nextNumberOfSteps
		}
	}
	panic("ho ho")
}

func possibleMoves(s State) []State {
	result := make([]State, 0)
	// pick any combinations of 1 or 2 items to move
	// things doesn't have to be IN the elevator, just move to another floor
	thingsOnFloor := make([]string, 0)
	const EMPTY = ""
	thingsOnFloor = append(thingsOnFloor, EMPTY) // we need the empty string meaning move nothing with one thing
	thingsOnFloor = append(thingsOnFloor, *s.floors[s.elevator].generators...)
	thingsOnFloor = append(thingsOnFloor, *s.floors[s.elevator].microchips...)
	// these permutations are way too naive in that they give back things ordered.
	// I hope that doesn't matter much as it will have been seen on the outside
	permutations := util.Permutations(thingsOnFloor)

	m := make(map[string][]string)
	for _, permutation := range permutations {
		// pick first 2 (there will always be at least one + EMPTY)
		permutation = permutation[:2]
		slices.Sort(permutation)
		m[strings.Join(permutation, "|")] = permutation
	}

	for _, permutation := range m {
		permutation = slices.DeleteFunc(permutation, func(s string) bool { return s == EMPTY })

		// move up
		if s.elevator < 3 {
			result = moveIfValid(s, permutation, result, true)
		}

		// move down
		if s.elevator > 0 && !floorsBelowAreEmpty(s) {
			result = moveIfValid(s, permutation, result, false)
		}
	}

	return result
}

func floorsBelowAreEmpty(s State) bool {
	for i := 0; i < s.elevator; i++ {
		if len(*s.floors[i].generators) > 0 || len(*s.floors[i].microchips) > 0 {
			return false
		}
	}
	return true
}

func moveIfValid(s State, permutation []string, result []State, up bool) []State {
	newState := deepCopyState(s)
	if up {
		newState.elevator++
	} else {
		newState.elevator--
	}
	moveElevator(s, permutation, newState)

	//fmt.Println("old", s.key())

	// validate
	if newState.isValid() {
		//fmt.Println("valid", newState.key())
		result = append(result, newState)
	}
	return result
}

func deepCopyState(s State) State {
	newState := State{
		elevator: s.elevator,
		floors:   make([]Floor, 4),
	}

	for i, floor := range s.floors {
		newState.floors[i] = Floor{
			generators: &[]string{},
			microchips: &[]string{},
		}
		*newState.floors[i].generators = append(*newState.floors[i].generators, *floor.generators...)
		*newState.floors[i].microchips = append(*newState.floors[i].microchips, *floor.microchips...)
	}
	return newState
}

func moveElevator(s State, permutation []string, newState State) {
	for _, thing := range permutation {
		// move thing
		nextFloor := newState.floors[newState.elevator]
		lastFloor := newState.floors[s.elevator]
		del := func(s string) bool {
			return s == thing
		}
		if strings.Contains(thing, "compatible") {
			// microchip
			*nextFloor.microchips = append(*nextFloor.microchips, thing)
			*lastFloor.microchips = slices.DeleteFunc(*lastFloor.microchips, del)
		} else {
			// generator
			*nextFloor.generators = append(*nextFloor.generators, thing)
			*lastFloor.generators = slices.DeleteFunc(*lastFloor.generators, del)
		}
	}
}
