package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"sort"
	"strings"
	"time"
	"unicode"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2023/01/input.txt", "\n")

	sum := 0
	for _, s := range inFile {
		numbers := make(map[int]int)

		for i, r := range s {
			if unicode.IsNumber(r) {
				numbers[i] = int(r - '0')
			} else {
				for j := i; j <= len(s); j++ {
					if strings.HasPrefix(s[i:j], "one") {
						numbers[i] = 1
						break
					} else if strings.HasPrefix(s[i:j], "two") {
						numbers[i] = 2
						break
					} else if strings.HasPrefix(s[i:j], "three") {
						numbers[i] = 3
						break
					} else if strings.HasPrefix(s[i:j], "four") {
						numbers[i] = 4
						break
					} else if strings.HasPrefix(s[i:j], "five") {
						numbers[i] = 5
						break
					} else if strings.HasPrefix(s[i:j], "six") {
						numbers[i] = 6
						break
					} else if strings.HasPrefix(s[i:j], "seven") {
						numbers[i] = 7
						break
					} else if strings.HasPrefix(s[i:j], "eight") {
						numbers[i] = 8
						break
					} else if strings.HasPrefix(s[i:j], "nine") {
						numbers[i] = 9
						break
					}
				}
			}
		}

		keys := make([]int, 0)
		for k := range numbers {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		value := numbers[keys[0]]*10 + numbers[keys[len(keys)-1]]
		sum += value
	}

	fmt.Println("part2: ", sum, "in", time.Since(start))
}
