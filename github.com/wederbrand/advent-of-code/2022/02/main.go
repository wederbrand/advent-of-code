package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
)

func main() {
	inFile := util.GetFileContents("2022/02/input.txt", "\n")

	part1(inFile)
	part2(inFile)
}

func part2(inFile []string) {
	// A for Rock, B for Paper, and C for Scissors
	// X for Rock, Y for Paper, and Z for Scissors.
	// 1 for Rock, 2 for Paper, and 3 for Scissors
	// 0 if you lost, 3 if the round was a draw, and 6 if you won
	scores := map[uint8]int{
		'A': 1,
		'B': 2,
		'C': 3,
		'X': 1,
		'Y': 2,
		'Z': 3,
	}
	score := 0
	for _, s := range inFile {
		opponent := scores[s[0]]

		// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win.
		if s[2] == 'X' {
			// lose
			me := opponent - 1
			if me == 0 {
				me = 3
			}
			score += me
		}
		if s[2] == 'Y' {
			// draw
			score += 3
			score += opponent
		}
		if s[2] == 'Z' {
			// win
			score += 6
			me := opponent + 1
			if me == 4 {
				me = 1
			}
			score += me
		}
	}

	fmt.Println(score)
}

func part1(inFile []string) {
	// A for Rock, B for Paper, and C for Scissors
	// X for Rock, Y for Paper, and Z for Scissors.
	// 1 for Rock, 2 for Paper, and 3 for Scissors
	// 0 if you lost, 3 if the round was a draw, and 6 if you won
	scores := map[uint8]int{
		'A': 1,
		'B': 2,
		'C': 3,
		'X': 1,
		'Y': 2,
		'Z': 3,
	}
	score := 0
	for _, s := range inFile {
		opponent := scores[s[0]]
		me := scores[s[2]]
		score += me
		if opponent == me {
			// draw
			score += 3
		}
		if me-opponent == 1 || me-opponent == -2 {
			// win
			score += 6
		}
	}

	fmt.Println(score)
}
