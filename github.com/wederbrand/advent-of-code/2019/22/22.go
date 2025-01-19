package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"math/big"
	"slices"
	"strings"
	"time"
)

const DEAL = "deal"
const CUT = "cut"
const INC = "inc"

type Shuffle struct {
	operation string
	value     int64
}

func main() {
	start := time.Now()

	inFile := GetFileContents("2019/22/input.txt", "\n")

	deckSize := int64(10007)
	shuffles := getShuffles(inFile, deckSize)
	deck := make([]int, deckSize)
	for i := 0; int64(i) < deckSize; i++ {
		deck[i] = i
	}

	for _, shuffle := range shuffles {
		switch shuffle.operation {
		case DEAL:
			slices.Reverse(deck)
		case CUT:
			deck = append(deck[shuffle.value:], deck[:shuffle.value]...)
		case INC:
			newDeck := make([]int, deckSize)
			for i, card := range deck {
				newDeck[(int64(i)*shuffle.value)%deckSize] = card
			}
			deck = newDeck
		}
	}

	part1 := 0
	for i, card := range deck {
		if card == 2019 {
			part1 = i
			break
		}
	}
	fmt.Println("part1: ", part1, "in", time.Since(start))

	deckSize = int64(119315717514047)
	shuffles = make([]Shuffle, 0)

	iterations := int64(101741582076661)
	factor := getShuffles(inFile, deckSize)

	for iterationsLeft := deckSize - iterations - 1; iterationsLeft != 0; iterationsLeft /= 2 {
		if iterationsLeft%2 == 1 {
			shuffles = append(shuffles, factor...)
			shuffles = compact(shuffles, deckSize)
		}
		factor = append(factor, factor...)
		factor = compact(factor, deckSize)
	}

	part2 := big.NewInt(2020)
	for _, shuffle := range shuffles {
		if shuffle.operation == INC {
			increment := shuffle.value
			part2.Mul(part2, big.NewInt(increment))
			part2.Mod(part2, big.NewInt(deckSize))

		} else if shuffle.operation == DEAL {
			part2.Sub(big.NewInt(deckSize-1), part2)

		} else if shuffle.operation == CUT {
			cut := shuffle.value

			if part2.Int64() < cut {
				part2.Add(part2, big.NewInt(deckSize-cut))
			} else {
				part2.Sub(part2, big.NewInt(cut))
			}
		}
	}

	fmt.Println("part2: ", part2.Int64(), "in", time.Since(start))
}

func getShuffles(inFile []string, deckSize int64) []Shuffle {
	shuffles := make([]Shuffle, 0)
	for _, line := range inFile {
		if strings.HasPrefix(line, "deal into new stack") {
			shuffles = append(shuffles, Shuffle{DEAL, 0})
		}
		if strings.HasPrefix(line, "cut") {
			cutIndex := int64(0)
			fmt.Sscanf(line, "cut %d cards", &cutIndex)
			if cutIndex < 0 {
				cutIndex = deckSize + cutIndex
			}
			shuffles = append(shuffles, Shuffle{CUT, cutIndex})
		}
		if strings.HasPrefix(line, "deal with increment") {
			increment := int64(0)
			fmt.Sscanf(line, "deal with increment %d", &increment)
			shuffles = append(shuffles, Shuffle{INC, increment})
		}
	}

	shuffles = compact(shuffles, deckSize)
	return shuffles
}

func compact(input []Shuffle, count int64) []Shuffle {
	{
		compacted := make([]Shuffle, 0, len(input))
		reverse := false
		for _, shuffle := range input {
			if shuffle.operation == DEAL {
				reverse = !reverse
				continue
			}
			if !reverse {
				compacted = append(compacted, shuffle)
				continue
			}
			switch shuffle.operation {
			case INC:
				compacted = append(compacted, shuffle)
				compacted = append(compacted, Shuffle{CUT, count + 1 - shuffle.value})

			case CUT:
				cut := (shuffle.value + count) % count // normalize negative values
				cut = count - cut                      // reverse cut
				compacted = append(compacted, Shuffle{CUT, cut})
			}
		}
		if reverse {
			compacted = append(compacted, Shuffle{DEAL, 0})
		}
		input = compacted
	}

	{
		compacted := make([]Shuffle, 0, len(input))
		cut := big.NewInt(0)
		for _, shuffle := range input {
			switch shuffle.operation {
			case DEAL:
				if value := cut.Int64(); value != 0 {
					compacted = append(compacted, Shuffle{CUT, value})
					cut.SetInt64(0)
				}
				compacted = append(compacted, shuffle)

			case INC:
				compacted = append(compacted, shuffle)
				cut.Mul(cut, big.NewInt(shuffle.value))
				cut.Mod(cut, big.NewInt(count))

			case CUT:
				cut.Add(cut, big.NewInt(shuffle.value))
				cut.Mod(cut, big.NewInt(count))
			}
		}
		if value := cut.Int64(); value != 0 {
			compacted = append(compacted, Shuffle{CUT, value})
			cut.SetInt64(0)
		}
		input = compacted
	}

	{
		compacted := make([]Shuffle, 0, len(input))
		increment := big.NewInt(1)
		for _, shuffle := range input {
			switch shuffle.operation {
			case INC:
				increment.Mul(increment, big.NewInt(shuffle.value))
				increment.Mod(increment, big.NewInt(count))

			default:
				if value := increment.Int64(); value != 1 {
					compacted = append(compacted, Shuffle{INC, value})
					increment.SetInt64(1)
				}
				compacted = append(compacted, shuffle)
			}
		}
		if value := increment.Int64(); value != 1 {
			compacted = append(compacted, Shuffle{INC, value})
			increment.SetInt64(1)
		}
		input = compacted
	}

	return input
}
