package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type board struct {
	values [5][5]int
	marks  [5][5]bool
	done   bool
}

func (b *board) mark(v int) {
	for j := 0; j < 5; j++ {
		for k := 0; k < 5; k++ {
			if b.values[j][k] == v {
				b.marks[j][k] = true
			}
		}
	}
}

func (b *board) set(j int, k int, v int) {
	b.values[j][k] = v
}

func (b *board) win() bool {
	for j := 0; j < 5; j++ {
		if b.marks[j][0] && b.marks[j][1] && b.marks[j][2] && b.marks[j][3] && b.marks[j][4] || b.marks[0][j] && b.marks[1][j] && b.marks[2][j] && b.marks[3][j] && b.marks[4][j] {
			return true
		}
	}

	return false
}

func (b *board) score(v int) int {
	unmarked := 0
	for j := 0; j < 5; j++ {
		for k := 0; k < 5; k++ {
			if !b.marks[j][k] {
				unmarked += b.values[j][k]
			}
		}
	}

	return unmarked * v
}

func main() {
	readFile, err := os.ReadFile("2021/2021_4.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	re := regexp.MustCompile("\\d+")

	numbers := inFile[0]
	boards := make([]*board, 0)

	for i := 2; i < len(inFile); i += 6 {
		b := new(board)
		for j := 0; j < 5; j++ {
			split := re.FindAllString(inFile[i+j], -1)
			for k := 0; k < 5; k++ {
				v, _ := strconv.Atoi(split[k])
				b.set(j, k, v)
			}
		}

		boards = append(boards, b)
	}

	for _, num := range strings.Split(numbers, ",") {
		atoi, _ := strconv.Atoi(num)
		for _, b := range boards {
			b.mark(atoi)
			if !b.done && b.win() {
				b.done = true
				fmt.Println(b.score(atoi))
			}
		}
	}

	os.Exit(1)
}
