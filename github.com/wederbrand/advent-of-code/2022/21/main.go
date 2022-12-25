package main

import (
	"errors"
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strconv"
	"strings"
	"time"
)

type monkey struct {
	name       string
	value      int
	done       bool
	expression string
	one        string
	two        string
	you        bool
}

func newMonkey(input string) *monkey {
	m := monkey{}
	split := strings.Split(input, ":")
	m.name = split[0]
	atoi, err := strconv.Atoi(strings.TrimSpace(split[1]))
	if err == nil {
		m.value = atoi
		m.done = true
	} else {
		fmt.Sscanf(strings.TrimSpace(split[1]), "%s %s %s", &m.one, &m.expression, &m.two)
	}
	return &m
}

func (m *monkey) getValue(monkeys map[string]*monkey) (int, error) {
	if m.you {
		return 0, errors.New("it's me")
	}
	if !m.done {
		v1, err1 := monkeys[m.one].getValue(monkeys)
		if err1 != nil {
			return 0, err1
		}
		v2, err2 := monkeys[m.two].getValue(monkeys)
		if err2 != nil {
			return 0, err2
		}
		switch m.expression {
		case "+":
			m.value = v1 + v2
		case "-":
			m.value = v1 - v2
		case "*":
			m.value = v1 * v2
		case "/":
			m.value = v1 / v2
		default:
			panic("unknown expression")
		}
	}

	m.done = true
	return m.value, nil
}

func (m *monkey) setValue(v int, monkeys map[string]*monkey) {
	if m.you {
		m.value = v
		m.done = true
		return
	}

	m1 := monkeys[m.one]
	m2 := monkeys[m.two]
	v1, err1 := m1.getValue(monkeys)
	v2, err2 := m2.getValue(monkeys)

	switch m.expression {
	case "+":
		if err1 != nil {
			m1.setValue(v-v2, monkeys)
		}
		if err2 != nil {
			m2.setValue(v-v1, monkeys)
		}
	case "-":
		if err1 != nil {
			m1.setValue(v+v2, monkeys)
		}
		if err2 != nil {
			m2.setValue(v1-v, monkeys)
		}
	case "*":
		if err1 != nil {
			m1.setValue(v/v2, monkeys)
		}
		if err2 != nil {
			m2.setValue(v/v1, monkeys)
		}
	case "/":
		if err1 != nil {
			m1.setValue(v*v2, monkeys)
		}
		if err2 != nil {
			m2.setValue(v*v1, monkeys)
		}
	default:
		panic("unknown expression")

	}
}

func (m *monkey) setEquals(monkeys map[string]*monkey) {
	m1 := monkeys[m.one]
	m2 := monkeys[m.two]
	v1, err1 := m1.getValue(monkeys)
	v2, err2 := m2.getValue(monkeys)

	if err1 != nil {
		m1.setValue(v2, monkeys)
	}
	if err2 != nil {
		m2.setValue(v1, monkeys)
	}
}

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2022/21/input.txt", "\n")

	monkeys := make(map[string]*monkey)
	for _, s := range inFile {
		m := newMonkey(s)
		monkeys[m.name] = m
	}

	monkeys["humn"].you = true
	monkeys["root"].setEquals(monkeys)

	fmt.Println("part2:", monkeys["humn"].value, "in", time.Since(start))
}
