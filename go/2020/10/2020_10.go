package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	readFile, err := ioutil.ReadFile("10/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.TrimSpace(string(readFile)), "\n")
	jolts := make([]int, 0)
	jolts = append(jolts, 0)
	for _, s := range input {
		atoi, _ := strconv.Atoi(s)
		jolts = append(jolts, atoi)
	}

	sort.Ints(jolts)
	jolts = append(jolts, jolts[len(jolts)-1]+3)

	i := 0
	combinations := 1
	for i < len(jolts) {
		// look for the next 3-step
		j := i + 1
		for j < len(jolts) && jolts[j]-jolts[j-1] != 3 {
			j++
		}
		combinations *= smallChecker(i, j-1, jolts)
		i = j
	}

	fmt.Println(combinations)
}

func smallChecker(i int, j int, jolts []int) int {
	if j - i == 0 {
		return 1
	}

	combinations := 0
	for jump := 1; jump <= 3 && i+jump <= j; jump++ {
		combinations += smallChecker(i+jump, j, jolts)
	}
	return combinations
}
