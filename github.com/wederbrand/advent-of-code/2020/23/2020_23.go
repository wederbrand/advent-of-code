package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

type cup struct {
	value int
	next  *cup
}

func main() {
	start := time.Now()
	readFile, err := ioutil.ReadFile("23/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.TrimSpace(string(readFile)), "")
	prev := new(cup)
	current := prev
	one := prev
	max := 0
	all := make(map[int]*cup)
	for _, s := range input {
		atoi, _ := strconv.Atoi(s)
		if atoi > max {
			max = atoi
		}
		nextCup := new(cup)
		nextCup.value = atoi
		all[nextCup.value] = nextCup

		prev.next = nextCup
		prev = nextCup

		if atoi == 1 {
			one = nextCup
		}
	}

	for i := 10; i <= 1_000_000; i++ {
		if i > max {
			max = i
		}

		nextCup := new(cup)
		nextCup.value = i
		all[nextCup.value] = nextCup
		prev.next = nextCup
		prev = nextCup
	}

	prev.next = current.next
	current = current.next

	for i := 0; i < 10_000_005; i++ {
		current = rotate(current, max, all)
	}

	fmt.Println(one.next.value * one.next.next.value)
	fmt.Println(time.Since(start))
}

func rotate(current *cup, max int, all map[int]*cup) *cup {
	lift := current.next                       // don't care for how many it is, it's all
	current.next = current.next.next.next.next // shortcut 3 of them

	insertTarget := current.value - 1

	for {
		for insertTarget == lift.value || insertTarget == lift.next.value || insertTarget == lift.next.next.value {
			insertTarget--
			if insertTarget == 0 {
				insertTarget = max
			}
		}

		if insertTarget == 0 {
			insertTarget = max
			continue
		}

		break
	}

	// insert the lifted ones
	insertCup := all[insertTarget]
	insertCup.next, lift.next.next.next = lift, insertCup.next

	return current.next
}
