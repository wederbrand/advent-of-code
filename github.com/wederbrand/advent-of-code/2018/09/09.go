package main

import (
	"fmt"
	"github.com/tebeka/deque"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("Part 1:", doIt(428, 72061), "in", time.Since(start))
	fmt.Println("Part 2:", doIt(428, 7206100), "in", time.Since(start))
}

func doIt(numberOfPlayers int, max int) int {
	players := make(map[int]int)

	circle := deque.New()
	circle.Append(0)
	for i := 1; i <= max; i++ {
		if i%23 == 0 {
			// collect score and pop that ball
			circle.Rotate(7)
			score, _ := circle.Pop()
			circle.Rotate(-1)
			playerIndex := i % numberOfPlayers
			players[playerIndex] += i + score.(int)
		} else {
			circle.Rotate(-1)
			circle.Append(i)
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
