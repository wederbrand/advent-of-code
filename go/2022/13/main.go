package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type node struct {
	i     int
	left  []node
	right []node
}

func main() {
	readFile, err := os.ReadFile("13/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n\n")

	sum := 0
	for i, pair := range inFile {
		split := strings.Split(strings.TrimSpace(pair), "\n")
		var result1 []interface{}
		json.Unmarshal([]byte(split[0]), &result1)
		var result2 []interface{}
		json.Unmarshal([]byte(split[1]), &result2)

		if correct(result1, result2) == -1 {
			sum += i + 1
		}
	}

	fmt.Println("part 1:", sum)

	all := make([][]interface{}, 0)
	for _, pair := range inFile {
		split := strings.Split(strings.TrimSpace(pair), "\n")
		var result1 []interface{}
		json.Unmarshal([]byte(split[0]), &result1)
		var result2 []interface{}
		json.Unmarshal([]byte(split[1]), &result2)
		all = append(all, result1)
		all = append(all, result2)
	}

	// extras
	var result1 []interface{}
	s1 := "[[2]]"
	json.Unmarshal([]byte(s1), &result1)
	all = append(all, result1)

	var result2 []interface{}
	s2 := "[[6]]"
	json.Unmarshal([]byte(s2), &result2)
	all = append(all, result2)

	sort.Slice(all, func(l, r int) bool {
		return correct(all[l], all[r]) == -1
	})

	key := 1
	for i, i2 := range all {
		marshal, _ := json.Marshal(i2)
		if string(marshal) == s1 || string(marshal) == s2 {
			key *= i + 1
		}
	}

	fmt.Println("part 2:", key)
}

// return -1 if left is smaller, 1 if right is smaller and 0 if they are equal
func correct(left []interface{}, right []interface{}) int {

	for i, l := range left {
		if i > len(right)-1 {
			// more on the left, order is incorrect
			return 1
		}
		r := right[i]

		switch lv := l.(type) {
		case float64:
			switch rv := r.(type) {
			case float64:
				if lv < rv {
					return -1
				} else if lv > rv {
					return 1
				}
			case []interface{}:
				slask := make([]interface{}, 1)
				slask[0] = lv
				result := correct(slask, rv)
				if result != 0 {
					return result
				}
			}
		case []interface{}:
			switch rv := r.(type) {
			case float64:
				slask := make([]interface{}, 1)
				slask[0] = rv
				result := correct(lv, slask)
				if result != 0 {
					return result
				}

			case []interface{}:
				result := correct(lv, rv)
				if result != 0 {
					return result
				}
			}
		}
	}

	if len(right) > len(left) {
		// more on the right, return -1
		return -1
	}
	return 0
}
