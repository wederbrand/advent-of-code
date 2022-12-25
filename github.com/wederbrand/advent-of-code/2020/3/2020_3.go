package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	forrest := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		forrest = append(forrest, scanner.Text())
	}

	// Start "off screen" to land on the right starting square
	hits := rideTheForrest(forrest, 1, 1)
	hits *= rideTheForrest(forrest, 3, 1)
	hits *= rideTheForrest(forrest, 5, 1)
	hits *= rideTheForrest(forrest, 7, 1)
	hits *= rideTheForrest(forrest, 1, 2)

	fmt.Println(hits)
}

func rideTheForrest(forrest []string, xOffset int, yOffset int) int {
	hits := 0
	x := 0
	y := 0
	for {
		x += xOffset
		y += yOffset

		if y >= len(forrest) {
			break
		}

		text := forrest[y]
		if text[x%len(text)] == '#' {
			hits++
		}
	}
	return hits
}
