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
	maxBus := 0
	maxBusOffset := 0
	for i, s := range bussesStr {
		atoi, err := strconv.Atoi(s)
		if err == nil {
			busses[i] = atoi
			if atoi > maxBus {
				maxBus = atoi
				maxBusOffset = i
			}
		}
	}

	i := 100000000000000
	for (i+maxBusOffset)%maxBus != 0 {
		i++
	}
	for {
		i += maxBus
		if i%100000 == 0 {
			fmt.Println(i)
		}
		found := 0
		for offset, bus := range busses {
			if (i+offset)%bus != 0 {
				continue
			}
			found++
		}
		if found == len(busses) {
			break
		}
	}

	fmt.Println(i)
}
