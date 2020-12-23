package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	readFile, err := ioutil.ReadFile("23/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.TrimSpace(string(readFile)), "")
	var ring [9]int
	for i, s := range input {
		atoi, _ := strconv.Atoi(s)
		ring[i] = atoi
	}

	for i := 0; i < 100; i++ {
		ring = rotate(ring)
		fmt.Println(i+1, ring)
	}

	// collect cups after 1
	result := ""
	oneIndex := 0
	for i, i2 := range ring {
		if i2 == 1 {
			oneIndex = i
			break
		}
	}
	for i := oneIndex + 1; i < len(ring); i++ {
		result += strconv.Itoa(ring[i])
	}

	for i := 0; i < oneIndex; i++ {
		result += strconv.Itoa(ring[i])
	}

	fmt.Println(result)

}

func rotate(in [9]int) [9]int {
	out := make([]int, 0)
	lift := in[1:4]
	rest := append(in[4:], in[0])
	insertTarget := in[0] - 1

	for insertTarget == lift[0] || insertTarget == lift[1] || insertTarget == lift[2] {
		insertTarget--
		if insertTarget == 0 {
			insertTarget = 9
		}
	}

	if insertTarget == 0 {
		insertTarget = 9
	}

	for insertTarget == lift[0] || insertTarget == lift[1] || insertTarget == lift[2] {
		insertTarget--
		if insertTarget == 0 {
			insertTarget = 9
		}
	}

	for _, cup := range rest {
		out = append(out, cup)
		if cup == insertTarget {
			out = append(out, lift...)
		}
	}

	var output [9]int
	copy(output[:], out)
	return output
}
