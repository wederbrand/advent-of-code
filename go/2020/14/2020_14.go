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
			addr := submatch[1]
			value, _ := strconv.ParseInt(submatch[2], 10, 64)
			maskedAddresses := maskIt(addr, mask)
			for _, maskedAddress := range maskedAddresses {
				memory[maskedAddress] = value
			}
		}
	}

	var sum int64
	for _, i := range memory {
		sum += i
	}
	fmt.Println(sum)
}

func maskIt(addr string, mask string) []int64 {
	v, _ := strconv.ParseInt(addr, 10, 64)
	binaryStr := strconv.FormatInt(v, 2)
	binaryRunes := []rune(fmt.Sprintf("%036v", binaryStr))

	for i, r := range mask {
		if r == '0' {
			continue
		}
		binaryRunes[i] = r
	}

	return getAddresses(binaryRunes)
}

func getAddresses(runes []rune) []int64 {
	addresses := make([]int64, 0)

	found := false
	for i, r := range runes {
		if r == 'X' {
			found = true
			copy := append([]rune{}, runes...)
			copy[i] = '0'
			addresses = append(addresses, getAddresses(copy)...)

			copy = append([]rune{}, runes...)
			copy[i] = '1'
			addresses = append(addresses, getAddresses(copy)...)

			break // just look for one X, the rest are recursively fixed
		}
	}

	if !found {
		binaryStr := string(runes)
		parseInt, _ := strconv.ParseInt(binaryStr, 2, 64)
		addresses = append(addresses, parseInt)
	}

	return addresses
}
