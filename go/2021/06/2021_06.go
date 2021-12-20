package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.ReadFile("2021/06/2021_06.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), ",")

	var fish [9]int
	for _, text := range inFile {
		atoi, _ := strconv.Atoi(text)
		fish[atoi]++
	}

	for i := 0; i < 256; i++ {
		var newFish [9]int
		for age, num := range fish {
			if age == 0 {
				newFish[8] += num
				newFish[6] += num
			} else {
				newFish[age-1] += num
			}
		}
		fish = newFish

		if i == 79 || i == 255 {
			total := 0
			for _, i := range fish {
				total += i
			}

			fmt.Println("after", i+1, "days", total)
		}
	}

	os.Exit(0)
}
