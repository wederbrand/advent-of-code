package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	readFile, err := ioutil.ReadFile("14/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.TrimSpace(string(readFile)), "\n")
	memory := make(map[int64]int64)
	mask := ""

	memPattern := regexp.MustCompile("mem\\[(\\d+)\\] = (.+)")
	for _, s := range input {
		if strings.HasPrefix(s, "mask") {
			mask = strings.Split(s, " = ")[1]
		} else {
			submatch := memPattern.FindStringSubmatch(s)
			addr, _ := strconv.ParseInt(submatch[1], 10, 64)
			value := submatch[2]
			maskedValue := maskIt(value, mask)
			memory[addr] = maskedValue
		}
	}

	var sum int64
	for _, i := range memory {
		sum += i
	}
	fmt.Println(sum)
}

func maskIt(value string, mask string) int64 {
	v, _ := strconv.ParseInt(value, 10, 64)
	binaryStr := strconv.FormatInt(v, 2)
	binaryRunes := []rune(fmt.Sprintf("%036v", binaryStr))

	for i, r := range mask {
		if r == 'X' {
			continue
		}
		binaryRunes[i] = r
	}

	binaryStr = string(binaryRunes)
	parseInt, _ := strconv.ParseInt(binaryStr, 2, 64)
	return parseInt
}
