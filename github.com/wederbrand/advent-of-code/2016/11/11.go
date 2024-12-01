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

type Element struct {
	name           string
	generatorIndex int
	microchipIndex int
}

type Floor struct {
	generators *[]string
	microchips *[]string
}

var elements = make(map[string]Element)

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
	fmt.Println("Part 1: ", part1, "in", time.Since(start))

	line := inFile[0] + " an elerium generator, an elerium-compatible microchip, a dilithium generator, and a dilithium-compatible microchip."
	generators := generatorRE.FindAllString(line, -1)
	microchips := microChipRE.FindAllString(line, -1)
	floors[0] = Floor{&generators, &microchips}

	part2 := doIt(floors)
	fmt.Println("Part 2: ", part2, "in", time.Since(start))
}

type FloorState struct {
	elevator int
	floors   []Floor
}

func (fs FloorState) key() string {
	key := strconv.Itoa(fs.elevator) + " "
	for _, floor := range fs.floors {
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

func (fs FloorState) isValid() bool {
	for _, floor := range fs.floors {
		if floor.generators == nil || len(*floor.generators) == 0 {
			// all floors without generators are valid
			continue
		}

		// find any unprotected microchips
		for _, microchip := range *floor.microchips {
			protected := false
			generatorName := strings.Replace(microchip, "-compatible", " generator", -1)
			for _, generator := range *floor.generators {
				if generator == generatorName {
					protected = true
				}
			}

			if !protected {
				return false
			}
		}
	}

	return true
}

func doIt(floors []Floor) int {
	q := priorityqueue.NewQueue()
	initial := priorityqueue.State{
		Data:     FloorState{0, floors},
		Priority: 0,
	}
	q.Add(&initial)

	seen := make(map[string]int)
	seen[initial.Data.(FloorState).key()] = 0
	for q.HasNext() {
		qs := q.Next()

		// pick all possible moves
		fs := qs.Data.(FloorState)

		if strings.HasPrefix(fs.key(), "3 |||") {
			return qs.Priority
		}

		moves := possibleMoves(fs)
		for _, nextS := range moves {
			// if not seen or cheaper, add to queue
			oldVal, found := seen[nextS.key()]
			nextNumberOfSteps := qs.Priority + 1
			if found && oldVal <= nextNumberOfSteps {
				continue
			}
			seen[nextS.key()] = nextNumberOfSteps

			// add to queue
			q.Add(&priorityqueue.State{Data: nextS, Priority: nextNumberOfSteps})
		}
	}
	panic("ho ho")
}

func possibleMoves(s FloorState) []FloorState {
	result := make([]FloorState, 0)

	var permutations [][]string
	// The permutations are either only generators, only microchips or a pair.
	// Never move a microchip with a non-matching generator
	// no need to ever move another pair
	floor := s.floors[s.elevator]
	generators := *floor.generators
	microchips := *floor.microchips

	for i := 0; i < len(generators); i++ {
		generator := generators[i]
		permutations = append(permutations, []string{generator})
		for j := i + 1; j < len(generators); j++ {
			generator2 := generators[j]
			permutations = append(permutations, []string{generator, generator2})
		}
	}

	for i := 0; i < len(microchips); i++ {
		microchip := microchips[i]
		permutations = append(permutations, []string{microchip})
		for j := i + 1; j < len(microchips); j++ {
			microchip2 := microchips[j]
			permutations = append(permutations, []string{microchip, microchip2})
		}
	}

	for _, generator := range generators {
		microchipName := strings.Replace(generator, " generator", "-compatible", -1)
		if slices.Contains(*floor.microchips, microchipName) {
			permutations = append(permutations, []string{generator, microchipName})
			break
		}
	}

	for _, permutation := range permutations {
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

func floorsBelowAreEmpty(s FloorState) bool {
	for i := 0; i < s.elevator; i++ {
		if len(*s.floors[i].generators) > 0 || len(*s.floors[i].microchips) > 0 {
			return false
		}
	}
	return true
}

func moveIfValid(fs FloorState, permutation []string, result []FloorState, up bool) []FloorState {
	newState := deepCopyState(fs)
	if up {
		newState.elevator++
	} else {
		newState.elevator--
	}
	moveElevator(fs, permutation, newState)

	// validate
	if newState.isValid() {
		result = append(result, newState)
	}
	return result
}

func deepCopyState(s FloorState) FloorState {
	newState := FloorState{
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

func moveElevator(s FloorState, permutation []string, newState FloorState) {
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
