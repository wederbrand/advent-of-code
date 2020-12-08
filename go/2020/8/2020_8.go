package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	readFile, err := ioutil.ReadFile("8/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	computer := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	for i, _ := range computer {
		split := strings.Split(computer[i], " ")
		if split[0] == "nop" {
			computer[i] = "jmp " + split[1]
		} else if split[0] == "jmp" {
			computer[i] = "nop " + split[1]
		} else {
			continue
		}

		acc, err := doIt(computer)
		if err == nil {
			fmt.Println(acc)
			break
		}

		// restore the computer
		computer[i] = split[0] + " " + split[1]
	}
}

func doIt(computer []string) (int, error) {
	acc := 0
	pointer := 0
	pointerHistory := make(map[int]bool)

	for {
		if pointer >= len(computer) {
			return acc, nil
		}
		if pointerHistory[pointer] {
			break
		}
		pointerHistory[pointer] = true
		split := strings.Split(computer[pointer], " ")
		atoi, _ := strconv.Atoi(split[1])
		switch split[0] {
		case "acc":
			acc += atoi
			pointer++
		case "jmp":
			pointer += atoi
		case "nop":
			pointer++
		}
	}
	return acc, errors.New("loop")
}
