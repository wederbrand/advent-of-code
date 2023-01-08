package util

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func GetFileContents(fileName string, splitOn string) []string {
	readFile, err := os.ReadFile("github.com/wederbrand/advent-of-code/" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimRightFunc(string(readFile), unicode.IsSpace), splitOn)
	return inFile
}

// StringOrNumber returns a number and an empty string if possible, otherwise the string
func StringOrNumber(in string) (string, int) {
	atoi, err := strconv.Atoi(in)
	if err == nil {
		return "", atoi
	} else {
		return in, 0
	}
}

// Atoi returns the int from strconv.Atoi and ignores any errors
func Atoi(in string) int {
	i, _ := strconv.Atoi(in)
	return i
}

func MaxOf(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

var printed = false

func PrintOnce(a ...any) {
	if !printed {
		fmt.Println(a)
	}
	printed = true
}

func Keys(all map[string]bool) (keys []string) {
	for city := range all {
		keys = append(keys, city)
	}

	return
}

func AllBut(all []string, but string) (result []string) {
	for _, city := range all {
		if city != but {
			result = append(result, city)
		}
	}
	return result
}

// Permutations return all permutations of the input array
// written by ChatGPT
func Permutations(arr []string) [][]string {
	var helper func([]string, int)
	var res [][]string

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					arr[i], arr[n-1] = arr[n-1], arr[i]
				} else {
					arr[0], arr[n-1] = arr[n-1], arr[0]
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func Key(a string, b string) string {
	return fmt.Sprintf("%s|%s", a, b)
}
