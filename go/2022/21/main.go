package main

import (
	"fmt"
	"log"
	"os"
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

func (m *monkey) getValue(monkeys map[string]*monkey) int {
	if !m.done {
		v1 := monkeys[m.one].getValue(monkeys)
		v2 := monkeys[m.two].getValue(monkeys)
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
	return m.value
}

func main() {
	start := time.Now()
	readFile, err := os.ReadFile("21/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	monkeys := make(map[string]*monkey)
	for _, s := range inFile {
		m := newMonkey(s)
		monkeys[m.name] = m
	}
	fmt.Println("part1:", monkeys["root"].getValue(monkeys), "in", time.Since(start))
}
