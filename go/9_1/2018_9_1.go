package main

import (
	"fmt"
	"github.com/tebeka/deque"
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

	circle := deque.New()
	circle.Append(0)
	for i := 1; i <= max; i++ {
		if i % 23 == 0 {
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
