package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	readFile, err := ioutil.ReadFile("13/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.TrimSpace(string(readFile)), "\n")
	bussesStr := strings.Split(input[1], ",")
	busses := make(map[int]int, 0)
	for i, s := range bussesStr {
		atoi, err := strconv.Atoi(s)
		if err == nil {
			busses[i] = atoi
		}
	}

	// pick the first two and reduce into one
	currentTime := 0
	currentMultiple := 0
	for offset, busId := range busses {
		if currentTime == 0 {
			currentTime = busId - offset
			currentMultiple = busId
		} else {
			currentTime, currentMultiple = reduce(currentTime, currentMultiple, offset, busId)
		}
	}

	fmt.Println(currentTime)
}

func reduce(time int, multiple int, offset int, busId int) (int, int) {
	for {
		time += multiple
		if (time+offset)%busId == 0 {
			return time, multiple * busId
		}
	}
}
