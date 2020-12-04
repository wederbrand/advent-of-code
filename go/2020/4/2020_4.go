package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p passport) valid() bool {
	if p.byr == "" {
		return false
	}
	atoi, err := strconv.Atoi(p.byr)
	if err != nil {
		return false
	}
	if atoi < 1920 || atoi > 2002 {
		return false
	}

	if p.iyr == "" {
		return false
	}
	atoi, err = strconv.Atoi(p.iyr)
	if err != nil {
		return false
	}
	if atoi < 2010 || atoi > 2020 {
		return false
	}

	if p.eyr == "" {
		return false
	}
	atoi, err = strconv.Atoi(p.eyr)
	if err != nil {
		return false
	}
	if atoi < 2020 || atoi > 2030 {
		return false
	}

	if p.hgt == "" {
		return false
	}
	compile := regexp.MustCompile("(\\d+)(.*)")
	if !compile.MatchString(p.hgt) {
		return false
	}
	submatch := compile.FindStringSubmatch(p.hgt)
	atoi, err = strconv.Atoi(submatch[1])
	if err != nil {
		return false
	}
	switch submatch[2] {
	case "cm":
		{
			if atoi < 150 || atoi > 193 {
				return false
			}
		}
	case "in":
		{
			if atoi < 59 || atoi > 76 {
				return false
			}
		}
	default:
		return false
	}

	if p.hcl == "" {
		return false
	}
	matchString, _ := regexp.MatchString("#[0-9a-f]{6}", p.hcl)
	if !matchString {
		return false
	}

	if p.ecl == "" {
		return false
	}

	if p.ecl != "amb" && p.ecl != "blu" && p.ecl != "brn" && p.ecl != "gry" && p.ecl != "grn" && p.ecl != "hzl" && p.ecl != "oth" {
		return false
	}

	if p.pid == "" {
		return false
	}

	matched, _ := regexp.MatchString("^\\d{9}$", p.pid)
	if !matched {
		return false
	}

	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	passports := make([]*passport, 0)
	scanner := bufio.NewScanner(file)

	current := new(passport)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			passports = append(passports, current)
			current = new(passport)
			continue
		}
		for _, s := range strings.Split(text, " ") {
			pair := strings.Split(s, ":")
			switch pair[0] {
			case "byr":
				current.byr = pair[1]
			case "iyr":
				current.iyr = pair[1]
			case "eyr":
				current.eyr = pair[1]
			case "hgt":
				current.hgt = pair[1]
			case "hcl":
				current.hcl = pair[1]
			case "ecl":
				current.ecl = pair[1]
			case "pid":
				current.pid = pair[1]
			case "cid":
				current.cid = pair[1]
			}
		}
	}
	passports = append(passports, current)

	count := 0
	for _, p := range passports {
		if p.valid() {
			count++
			fmt.Println(p.pid)
		}
	}

	fmt.Println(count)

	// 225 too high

}
