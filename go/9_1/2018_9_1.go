package main

import (
	"fmt"
)

func main() {
	fmt.Println(doIt(9, 25))
	fmt.Println(doIt(10, 1618))
	fmt.Println(doIt(13, 7999))
	fmt.Println(doIt(17, 1104))
	fmt.Println(doIt(21, 6111))
	fmt.Println(doIt(30, 5807))
	fmt.Println(doIt(428, 72061))
	fmt.Println(doIt(428, 7206100))
}

func doIt(numberOfPlayers int, max int) int {
	players := make(map[int]int)

	balls := make([]int, 0)
	balls = append(balls, 0)
	lastUsedIndex := 0
	for i := 1; i <= max; i++ {
		if i % 100000 == 0 {
			fmt.Println(i)
		}
		if i % 23 == 0 {
			// collect score and pop that ball
			score := i
			sevenIndex := lastUsedIndex - 7
			sevenIndex = (sevenIndex + len(balls)) % len(balls)
			score += balls[sevenIndex]
			balls = append(balls[:sevenIndex], balls[sevenIndex+1:]...)
			lastUsedIndex = sevenIndex
			playerIndex := i % numberOfPlayers
			players[playerIndex] += score
		} else {
			nextIndex := lastUsedIndex + 2
			if lastUsedIndex == len(balls) - 1 {
				// last position used
				nextIndex = 1
			}
			balls = append(balls, 0)
			copy(balls[nextIndex+1:], balls[nextIndex:])
			balls[nextIndex] = i
			lastUsedIndex = nextIndex
		}
	}

	result := 0
	for _, value := range players {
		if value > result {
			result = value
		}
	}

	return result
}
