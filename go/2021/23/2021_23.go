package main

import (
	"fmt"
	"log"
	"os"
)

const A = 2
const B = 4
const C = 6
const D = 8

func getHome(in string) int {
	switch in {
	case "A":
		return A
	case "B":
		return B
	case "C":
		return C
	case "D":
		return D
	}
	log.Fatal("no such home")
	return 0
}

func getPrice(in string) int {
	switch in {
	case "A":
		return 1
	case "B":
		return 10
	case "C":
		return 100
	case "D":
		return 1000
	}
	log.Fatal("no such price")
	return 0
}

type state struct {
	pos   [11][5]string
	price int
}

func (s state) move(fromX int, fromY int, toX int, toY int) state {
	s2 := state{
		pos:   s.pos,
		price: s.price,
	}

	s2.pos[toX][toY] = s2.pos[fromX][fromY]
	s2.pos[fromX][fromY] = "."

	return s2
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
	// add state to queue and sort it by total price
	if len(q.states) == 0 {
		q.states = append(q.states, s)
	} else {
		for i, p2 := range q.states {
			if s.price < p2.price {
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

func (q *queue) empty() bool {
	return len(q.states) == 0
}

func (q *queue) dequeue() *state {
	p := q.states[0]
	q.states = q.states[1:len(q.states)]
	return p
}

func main() {
	s := state{
		price: 0,
	}

	s.pos[0][0] = "."
	s.pos[1][0] = "."
	s.pos[2][0] = "."
	s.pos[3][0] = "."
	s.pos[4][0] = "."
	s.pos[5][0] = "."
	s.pos[6][0] = "."
	s.pos[7][0] = "."
	s.pos[8][0] = "."
	s.pos[9][0] = "."
	s.pos[10][0] = "."

	s.pos[A][1] = "B"
	s.pos[B][1] = "C"
	s.pos[C][1] = "A"
	s.pos[D][1] = "D"

	s.pos[A][2] = "D"
	s.pos[B][2] = "C"
	s.pos[C][2] = "B"
	s.pos[D][2] = "A"

	s.pos[A][3] = "D"
	s.pos[B][3] = "B"
	s.pos[C][3] = "A"
	s.pos[D][3] = "C"

	s.pos[A][4] = "B"
	s.pos[B][4] = "C"
	s.pos[C][4] = "D"
	s.pos[D][4] = "A"

	q := newQueue()
	q.add(&s)

	lowestProcessedState := make(map[[11][5]string]int)
	for !q.empty() {
		theState := q.dequeue()
		lowest, found := lowestProcessedState[theState.pos]
		if found && theState.price >= lowest {
			// we've seen this already, and cheaper
			continue
		} else {
			lowestProcessedState[theState.pos] = theState.price
		}
		if theState.done() {
			fmt.Println("part 2", theState.price)
			os.Exit(0)
		}

		// find all possible movements
		findMoves(theState, q)
	}

}

func findMoves(s *state, q *queue) {
	// check top row to see if any can get home
	// if so, do
	for x := 0; x < 11; x++ {
		animal := s.pos[x][0]
		if animal != "." {
			// see if it can get home from x to
			home := getHome(animal)
			canGetHome, dist, depth := s.canGoHome(animal, x, home)
			if canGetHome {
				newState := s.move(x, 0, home, depth)
				newState.price += dist * getPrice(animal)
				//printIt(&newState)
				q.add(&newState)
			}
		}
	}

	// move all top ones from homes to top row, all permutations
	for _, r := range [4]string{"A", "B", "C", "D"} {
		x := getHome(r)
		if !s.doneHome(x, r) {
			for y := 1; y <= 4; y++ {
				if s.pos[x][y] != "." {
					// add all possible permutations
					for _, i := range []int{0, 1, 3, 5, 7, 9, 10} {
						if s.freeFromTo(x, i) {
							newState := s.move(x, y, i, 0)
							newState.price += (y + abs(x-i)) * getPrice(s.pos[x][y])
							//printIt(&newState)
							q.add(&newState)
						}
					}
					break // one found
				}
			}
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (s *state) freeFromTo(from int, to int) bool {
	for dx := from + 1; dx <= to; dx++ {
		if s.pos[dx][0] != "." {
			return false
		}
	}
	for dx := from - 1; dx >= to; dx-- {
		if s.pos[dx][0] != "." {
			return false
		}
	}

	return true
}

func (s *state) canGoHome(animal string, x int, home int) (bool, int, int) {
	// check free path
	if !s.freeFromTo(x, home) {
		return false, -1, -1
	}

	dist := 0
	if home > x {
		dist += home - x
	} else {
		dist += x - home
	}

	depth := 0
	for d := 1; d <= 4; d++ {
		occupant := s.pos[home][d]
		if occupant != animal && occupant != "." {
			return false, -1, -1
		}
		if occupant == "." {
			dist++
			depth++
		}
	}

	return true, dist, depth
}

func (s state) done() bool {
	return s.doneHome(A, "A") && s.doneHome(B, "B") && s.doneHome(C, "C") && s.doneHome(D, "D")
}

func (s state) doneHome(x int, home string) bool {
	return s.pos[x][1] == home && s.pos[x][2] == home && s.pos[x][3] == home && s.pos[x][4] == home
}
