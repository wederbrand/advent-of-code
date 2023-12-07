package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"slices"
	"sort"
	"strings"
	"time"
	"unicode"
)

type Hand struct {
	cards [5]int
	bid   int
	score int
}

func newHand(s string) *Hand {
	h := new(Hand)

	split := strings.Split(s, " ")
	seen := make(map[int]int)
	jokers := 0
	for i, r := range split[0] {
		if unicode.IsNumber(r) {
			h.cards[i] = int(r - '0')
		} else if r == 'A' {
			h.cards[i] = 14
		} else if r == 'K' {
			h.cards[i] = 13
		} else if r == 'Q' {
			h.cards[i] = 12
		} else if r == 'J' {
			h.cards[i] = 1
			jokers++
		} else if r == 'T' {
			h.cards[i] = 10
		}

		seen[h.cards[i]]++
	}

	// pretend we didn't see the jokers
	h.bid = util.Atoi(split[1])

	count := make([]int, 0)
	for cardValue, cardCount := range seen {
		if cardValue != 1 /* ignore jokers */ {
			count = append(count, cardCount)
		}
	}

	sort.Ints(count)
	slices.Reverse(count)

	if jokers == 5 || count[0]+jokers == 5 {
		h.score = 7
	} else if count[0]+jokers == 4 {
		h.score = 6
	} else if count[0] == 3 && count[1] == 2 {
		h.score = 5
	} else if jokers == 1 && count[0] == 2 && count[1] == 2 {
		h.score = 5
	} else if count[0]+jokers == 3 {
		h.score = 4
	} else if count[0] == 2 && count[1] == 2 {
		h.score = 3
	} else if count[0]+jokers == 2 {
		h.score = 2
	} else if len(count) == 5 {
		h.score = 1
	} else {
		panic("oh no")
	}

	return h
}

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2023/07/input.txt", "\n")

	hands := make([]*Hand, 0)
	for _, s := range inFile {
		hands = append(hands, newHand(s))
	}

	sort.Slice(hands, func(a, b int) bool {
		if hands[a].score == hands[b].score {
			for i := range hands[a].cards {
				aCard := hands[a].cards[i]
				bCard := hands[b].cards[i]
				if aCard == bCard {
					continue
				} else {
					return aCard > bCard
				}
			}
		} else {
			return hands[a].score > hands[b].score
		}

		panic("ho no")
	})

	part2 := 0
	for i, hand := range hands {
		part2 += hand.bid * (len(hands) - i)
	}

	fmt.Println("part2: ", part2, "in", time.Since(start))
}
