package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func minus(a string, b string) string {
	for _, s := range b {
		a = strings.ReplaceAll(a, string(s), "")
	}

	return a
}

func sortString(value string) string {
	s := strings.Split(value, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	readFile, err := os.ReadFile("2021/08/2021_08.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	sum := 0
	for _, text := range inFile {
		split := strings.Split(text, " | ")
		inputValues := strings.Split(split[0], " ")
		outputValues := strings.Split(split[1], " ")

		for i, value := range inputValues {
			inputValues[i] = sortString(value)
		}

		for i, value := range outputValues {
			outputValues[i] = sortString(value)
		}

		var key = map[string]int{}
		var letters [10]string

		// first pick 1, 4, 7 and 8
		for _, value := range inputValues {
			if len(value) == 2 {
				key[value] = 1
				letters[1] = value
			}
			if len(value) == 3 {
				key[value] = 7
				letters[7] = value
			}
			if len(value) == 4 {
				key[value] = 4
				letters[4] = value
			}
			if len(value) == 7 {
				key[value] = 8
				letters[8] = value
			}
		}

		// then look for 3 and 6
		for _, value := range inputValues {
			_, found := key[value]
			if found {
				continue
			}
			if len(value) == 5 {
				if len(minus(value, letters[1])) == 3 {
					key[value] = 3
					letters[3] = value
					continue
				}
			}
			if len(value) == 6 {
				if len(minus(value, letters[1])) == 5 {
					key[value] = 6
					letters[6] = value
					continue
				}
			}
		}

		// finally, find 0, 2, 5 and 9
		for _, value := range inputValues {
			_, found := key[value]
			if found {
				continue
			}
			if len(value) == 5 {
				if len(minus(value, letters[4])) == 2 {
					key[value] = 5
					letters[5] = value
				}
				if len(minus(value, letters[4])) == 3 {
					key[value] = 2
					letters[2] = value
				}
			}
			if len(value) == 6 {
				if len(minus(value, letters[4])) == 2 {
					key[value] = 9
					letters[9] = value
				}
				if len(minus(value, letters[4])) == 3 {
					key[value] = 0
					letters[0] = value
				}
			}
		}

		result := ""
		for _, value := range outputValues {
			itoa := strconv.Itoa(key[value])
			result += itoa
		}

		atoi, _ := strconv.Atoi(result)
		sum += atoi
	}

	fmt.Println(sum)
	os.Exit(0)
}
