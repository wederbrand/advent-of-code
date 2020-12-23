package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type cup struct {
	value int
	next  *cup
}

func main() {
	readFile, err := ioutil.ReadFile("23/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.TrimSpace(string(readFile)), "")
	prev := new(cup)
	current := prev
	one := prev
	max := 0
	for _, s := range input {
		atoi, _ := strconv.Atoi(s)
		if atoi > max {
			max = atoi
		}
		nextCup := new(cup)
		nextCup.value = atoi
		prev.next = nextCup
		prev = nextCup

		if atoi == 1 {
			one = nextCup
		}
	}

	prev.next = current.next
	current = current.next

	for i := 0; i < 100; i++ {
		current = rotate(current, max)
	}

	result := ""
	for i := 0; i < 8; i++ {
		one = one.next
		result += strconv.Itoa(one.value)
	}

	fmt.Println(result)
}

func rotate(current *cup, max int) *cup {
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
	insertCup := current
	for {
		insertCup = insertCup.next
		if insertCup.value == insertTarget {
			break
		}
	}

	insertCup.next, lift.next.next.next = lift, insertCup.next

	return current.next
}
