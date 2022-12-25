package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items       []int
	operation   func(int) int
	div         int
	ifTrue      int
	ifFalse     int
	inspections int
}

func (m *monkey) setOperation(operation string) {
	var sign string
	var value string
	fmt.Sscanf(operation, "new = old %s %s", &sign, &value)
	atoi, err := strconv.Atoi(value)
	if err != nil {
		switch sign {
		case "*":
			m.operation = func(item int) int { return item * item }
		case "+":
			m.operation = func(item int) int { return item + item }
		}
	} else {
		switch sign {
		case "*":
			m.operation = func(item int) int { return item * atoi }
		case "+":
			m.operation = func(item int) int { return item + atoi }
		}
	}
}

func main() {
	inFile := util.GetFileContents("2022/11/input.txt", "\n")

	monkeys := make(map[int]*monkey, 0)
	globalDiv := 1
	for i := 0; i < len(inFile); i++ {
		m := monkey{
			items: make([]int, 0),
		}

		var monkeyNbr int
		fmt.Sscanf(inFile[i], "Monkey %d:", &monkeyNbr)
		i++

		startingItems := strings.Split(strings.TrimSpace(strings.Split(inFile[i], ":")[1]), ",")
		i++
		for _, startingItem := range startingItems {
			atoi, _ := strconv.Atoi(strings.TrimSpace(startingItem))
			m.items = append(m.items, atoi)
		}

		m.setOperation(strings.TrimSpace(strings.Split(inFile[i], ":")[1]))
		i++

		fmt.Sscanf(strings.TrimSpace(inFile[i]), "Test: divisible by %d", &m.div)
		i++

		globalDiv *= m.div

		fmt.Sscanf(strings.TrimSpace(inFile[i]), "If true: throw to monkey  %d", &m.ifTrue)
		i++
		fmt.Sscanf(strings.TrimSpace(inFile[i]), "If false: throw to monkey  %d", &m.ifFalse)
		i++

		monkeys[monkeyNbr] = &m
	}

	for turn := 0; turn < 10000; turn++ {
		// monkeys throw
		for i := 0; i < len(monkeys); i++ {
			// monkey i throws
			m := monkeys[i]
			for _, item := range m.items {
				m.inspections++
				newValue := m.operation(item) % globalDiv
				target := monkeys[m.ifTrue]
				if newValue%m.div != 0 {
					target = monkeys[m.ifFalse]
				}
				target.items = append(target.items, newValue)
			}
			// clear out the array, all has been thrown
			m.items = m.items[:0]
		}
	}

	inspections := make([]int, len(monkeys))
	for i, m := range monkeys {
		inspections[i] = m.inspections
	}

	sort.Ints(inspections)

	fmt.Println(inspections[len(inspections)-2] * inspections[len(inspections)-1])
}
