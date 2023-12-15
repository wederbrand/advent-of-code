package util

import (
	"fmt"
	"log"
	"math"
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

type Dir [2]int

var N = Dir{0, -1}
var UP = Dir{0, -1}
var S = Dir{0, +1}
var DOWN = Dir{0, +1}
var E = Dir{+1, 0}
var RIGHT = Dir{+1, 0}
var W = Dir{-1, 0}
var LEFT = Dir{-1, 0}

type Coord struct {
	X int
	Y int
}

type Map map[Coord]string

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

// Contains returns true if b contains a
func Contains(a int, b []int) bool {
	for _, i := range b {
		if a == i {
			return true
		}
	}

	return false
}

// Permutations return all permutations of the input array
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

func DeKey(key string) (a int, b int) {
	split := strings.Split(key, "|")
	return Atoi(split[0]), Atoi(split[1])
}

func Manhattan(ax int, ay int, bx int, by int) int {
	return int(math.Abs(float64(ax-bx)) + math.Abs(float64(ay-by)))
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

// Gcd returns the greatest common divisor
func Gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// Lcd returns the Least Common Denominator
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

func BinarySearch(start int, end int, fn func(int) bool) (last int, first int) {
	if fn(start) {
		panic("the start value must produce false")
	}
	if !fn(end) {
		panic("the end value must produce false")
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

func MakeMap(in []string) Map {
	m := make(map[Coord]string)

	for y, s := range in {
		for x, r := range s {
			if r == '.' {
				continue
			}
			m[Coord{x, y}] = string(r)
		}
	}

	return m
}

func RotateClockWise(in Map) Map {
	// For my reversed Y clockwise is the same a normal counterclockwise
	// 90° counterclockwise rotation: (𝑥,𝑦) becomes (−𝑦,𝑥)

	out := make(Map)

	for key, value := range in {
		out[Coord{-key.Y, key.X}] = value
	}

	return out
}

func RotateCounterClockWise(in Map) Map {
	// For my reversed Y counterclockwise is the same a normal clockwise
	// 90° clockwise rotation: (𝑥,𝑦) becomes (𝑦,-𝑥)

	out := make(Map)

	for key, value := range in {
		out[Coord{key.Y, -key.X}] = value
	}

	return out
}

func GetMapMaxes(m Map) (minC Coord, maxC Coord) {
	minX := math.MaxInt
	minY := math.MaxInt
	maxX := math.MinInt
	maxY := math.MinInt
	for k := range m {
		minX = min(minX, k.X)
		minY = min(minY, k.Y)
		maxX = max(maxX, k.X)
		maxY = max(maxY, k.Y)
	}
	return Coord{minX, minY}, Coord{maxX, maxY}
}

func PrintMap(m Map) {
	minC, maxC := GetMapMaxes(m)
	for y := minC.Y; y <= maxC.Y; y++ {
		for x := minC.X; x <= maxC.X; x++ {
			s, found := m[Coord{x, y}]
			if !found {
				fmt.Print(".")
			} else {
				fmt.Print(s)
			}
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println()
}

func MapAsString(m Map) string {
	out := ""
	minC, maxC := GetMapMaxes(m)
	for y := minC.Y; y <= maxC.Y; y++ {
		for x := minC.X; x <= maxC.X; x++ {
			s, found := m[Coord{x, y}]
			if !found {
				out += "."
			} else {
				out += s
			}
		}
		out += "|"
	}

	return out
}
