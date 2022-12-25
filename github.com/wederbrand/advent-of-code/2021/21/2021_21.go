package main

import (
	"fmt"
)

type player struct {
	pos   int
	score int
}

type state struct {
	players       [2]player
	currentPlayer int
}

type wins [2]int64

var cachedResults map[state]wins

func (s *state) nextPlayer() {
	s.currentPlayer = (s.currentPlayer + 1) % 2
}

func main() {
	start := [2]int{4, 9}
	fmt.Println("part 1", part1(start))

	s := state{
		players: [2]player{
			{pos: start[0], score: 0},
			{pos: start[1], score: 0}},
		currentPlayer: 0,
	}

	cachedResults = make(map[state]wins)
	w := part2(s)

	if w[0] > w[1] {
		fmt.Println("part 2", w[0])
	} else {
		fmt.Println("part 2", w[1])
	}
}

func part2(s state) wins {
	cachedWins, found := cachedResults[s]
	if found {
		return cachedWins
	}
	w := wins{0, 0}

	// do one player turn, then send it on recursively
	// for each turn spawn 27 universe (3^3 dice results)
	for d1 := 1; d1 <= 3; d1++ {
		for d2 := 1; d2 <= 3; d2++ {
			for d3 := 1; d3 <= 3; d3++ {
				newState := s
				totalDistance := d1 + d2 + d3
				newState.players[newState.currentPlayer].pos += totalDistance
				if newState.players[newState.currentPlayer].pos > 10 {
					newState.players[newState.currentPlayer].pos %= 10
				}
				newState.players[newState.currentPlayer].score += newState.players[newState.currentPlayer].pos
				if newState.players[newState.currentPlayer].score >= 21 {
					w[newState.currentPlayer]++
				} else {
					newState.nextPlayer()
					newWins := part2(newState)
					w[0] += newWins[0]
					w[1] += newWins[1]
				}
			}
		}
	}

	cachedResults[s] = w
	return w
}

func part1(start [2]int) int {
	players := start
	scores := [2]int{0, 0}

	dice := 100 // next roll is 1
	totalRolls := 0
	currentPlayer := 0
	for {
		throw := throw3(&dice)
		totalRolls += 3
		players[currentPlayer] += throw
		if players[currentPlayer] > 10 {
			players[currentPlayer] %= 10
			if players[currentPlayer] == 0 {
				players[currentPlayer] = 10
			}
		}

		scores[currentPlayer] += players[currentPlayer]
		if scores[currentPlayer] >= 1000 {
			// winner
			theOtherPlayer := (currentPlayer + 1) % 2
			return totalRolls * scores[theOtherPlayer]
		}

		currentPlayer = (currentPlayer + 1) % 2
	}
}

func throw3(dice *int) int {
	throw := 0
	for i := 0; i < 3; i++ {
		*dice++
		if *dice > 100 {
			*dice = 1
		}
		throw += *dice
	}

	return throw
}
