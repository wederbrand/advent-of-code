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

func MatchingNumbersAfterSplitOnAny(in string, splitOn string, separator string) (result [][]int) {
	fieldsFunc := strings.FieldsFunc(in, func(r rune) bool {
		return strings.ContainsRune(splitOn, r)
	})

	for i, s := range fieldsFunc {
		result = append(result, []int{})
		raw := strings.Split(strings.TrimSpace(s), separator)
		for _, number := range raw {
			if number != "" {
				atoi, err := strconv.Atoi(number)
				if err == nil {
					result[i] = append(result[i], atoi)
				}
			}
		}
	}

	return
}

// Permutations return all permutations of the input array
func Permutations[T any](arr []T) [][]T {
	var helper func([]T, int)
	var res [][]T

	helper = func(arr []T, n int) {
		if n == 1 {
			tmp := make([]T, len(arr))
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

func DeKey(key string) (a int, b int) {
	split := strings.Split(key, "|")
	return Atoi(split[0]), Atoi(split[1])
}

func IntKey(x int, y int) string {
	return fmt.Sprintf("%d|%d", x, y)
}

func FirstRune(s string) rune {
	for _, r := range s {
		return r
	}
	return 0
}

// IntAbs returns the absolut value
func IntAbs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

// Gcd returns the greatest common divisor
func Gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// Lcd returns the Least Common Denominator
func Lcd(numbers []int) int {
	// initialize least common denominator to the first number in the slice
	lcd := numbers[0]

	// find the least common denominator by taking the gcd of the lcd and each subsequent number in the slice
	for _, number := range numbers[1:] {
		lcd = (lcd * number) / Gcd(lcd, number)
	}

	return lcd
}

// BinarySearch takes a start and end value to run a binary search on
// the returned values is the last value that still produces false and the first value that produces true
func BinarySearch(start int, end int, fn func(int) bool) (last int, first int) {
	if fn(start) {
		panic("the start value must produce false")
	}
	if !fn(end) {
		panic("the end value must produce true")
	}

	last = start // last that still produces false
	first = end  // first tha produces true

	for first-last != 1 {
		candidate := (first + last) / 2
		if fn(candidate) {
			first = candidate
		} else {
			last = candidate
		}
	}

	return last, first
}

func CloneMap[T any](original map[string]T) map[string]T {
	out := make(map[string]T)
	for key, value := range original {
		out[key] = value
	}
	return out
}

func CloneSlice[T any](in []T) (out []T) {
	out = make([]T, len(in))
	copy(out, in)
	return
}

func CloneSliceKeep[T any](in []T, keep func(int, T) bool) (out []T) {
	out = make([]T, 0)
	for i, item := range in {
		if keep(i, item) {
			out = append(out, item)
		}
	}
	return
}

func CloneSliceDelete[T any](in []T, delete func(int, T) bool) (out []T) {
	out = make([]T, 0)
	for i, item := range in {
		if !delete(i, item) {
			out = append(out, item)
		}
	}
	return
}

// Equation represents a linear equation in the form of a*x + b*y = c
type Equation struct {
	A int
	B int
	C int
}

// WikiCramer solves a system of linear equations using Cramer's rule
// the equations are in the form of:
// a1*x + b1*y = c1 (as the type Equation)
// a2*x + b2*y = c2 (as the type Equation)
func WikiCramer(one Equation, two Equation) (x int, y int, err error) {
	a1, b1, c1 := one.A, one.B, one.C
	a2, b2, c2 := two.A, two.B, two.C

	determinant := a1*b2 - a2*b1
	if determinant == 0 {
		return 0, 0, fmt.Errorf("cramer hates zeroes")
	}

	isXInteger := (c1*b2-b1*c2)%determinant == 0
	isYInteger := (a1*c2-c1*a2)%determinant == 0
	if !isXInteger || !isYInteger {
		return 0, 0, fmt.Errorf("the formula assumes floats, forcing integer division loses some precision, detected here")
	}

	x = (c1*b2 - b1*c2) / determinant
	y = (a1*c2 - c1*a2) / determinant

	return x, y, nil
}
