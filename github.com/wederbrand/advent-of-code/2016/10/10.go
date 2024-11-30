package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"sort"
	"strings"
	"time"
)

type Bot struct {
	lowOutput  string
	highOutput string
	lowTarget  int
	highTarget int
	values     []int
}

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2016/10/input.txt", "\n")

	part1, part2 := part1(inFile)
	fmt.Println("Part 1: ", part1, "in", time.Since(start))
	fmt.Println("Part 2: ", part2, "in", time.Since(start))
}

func part1(inFile []string) (int, int) {
	var part1 int

	bots := make(map[int]*Bot)
	for _, s := range inFile {
		if strings.HasPrefix(s, "bot") {
			var bot, low, high int
			var lowBotOrOutput, highBotOrOutput string
			fmt.Sscanf(s, "bot %d gives low to %s %d and high to %s %d", &bot, &lowBotOrOutput, &low, &highBotOrOutput, &high)
			bots[bot] = &Bot{lowBotOrOutput, highBotOrOutput, low, high, []int{}}
		}
	}

	for _, s := range inFile {
		if strings.HasPrefix(s, "value") {
			var value, bot int
			fmt.Sscanf(s, "value %d goes to bot %d", &value, &bot)
			b := bots[bot]
			b.values = append(b.values, value)
			sort.Ints(b.values)
		}
	}

	outputs := make(map[int][]int)
	found := true
	for found {
		// go forever, we'll break when we're done
		found = false
		for botName, bot := range bots {
			if len(bot.values) == 2 {
				found = true
				if bot.values[0] == 17 && bot.values[1] == 61 {
					part1 = botName
				}

				if bot.lowOutput == "output" {
					outputs[bot.lowTarget] = append(outputs[bot.lowTarget], bot.values[0])
				} else {
					bots[bot.lowTarget].values = append(bots[bot.lowTarget].values, bot.values[0])
					sort.Ints(bots[bot.lowTarget].values)
				}

				if bot.highOutput == "output" {
					outputs[bot.highTarget] = append(outputs[bot.highTarget], bot.values[1])
				} else {
					bots[bot.highTarget].values = append(bots[bot.highTarget].values, bot.values[1])
					sort.Ints(bots[bot.highTarget].values)
				}

				bot.values = []int{}
			}
		}

		if !found {
			break
		}
	}

	part2 := outputs[0][0] * outputs[1][0] * outputs[2][0]
	return part1, part2
}
