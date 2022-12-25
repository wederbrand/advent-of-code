package main

import (
	"fmt"
	"io/ioutil"
	"log"
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
	return p.byrValid() && p.iyrValid() && p.eyrValid() && p.hgtValid() && p.hclValid() && p.eclValid() && p.pidValid()
}

func (p passport) byrValid() bool {
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

	return true
}
func (p passport) iyrValid() bool {
	if p.iyr == "" {
		return false
	}
	atoi, err := strconv.Atoi(p.iyr)
	if err != nil {
		return false
	}
	if atoi < 2010 || atoi > 2020 {
		return false
	}

	return true
}

func (p passport) eyrValid() bool {
	if p.eyr == "" {
		return false
	}
	atoi, err := strconv.Atoi(p.eyr)
	if err != nil {
		return false
	}
	if atoi < 2020 || atoi > 2030 {
		return false
	}

	return true
}

func (p passport) hgtValid() bool {
	if p.hgt == "" {
		return false
	}
	compile := regexp.MustCompile("(\\d+)(.*)")
	if !compile.MatchString(p.hgt) {
		return false
	}
	submatch := compile.FindStringSubmatch(p.hgt)
	atoi, err := strconv.Atoi(submatch[1])
	if err != nil {
		return false
	}
	switch submatch[2] {
	case "cm":
		if atoi < 150 || atoi > 193 {
			return false
		}
	case "in":
		if atoi < 59 || atoi > 76 {
			return false
		}
	default:
		return false
	}

	return true
}

func (p passport) hclValid() bool {
	if p.hcl == "" {
		return false
	}
	matchString, _ := regexp.MatchString("#[0-9a-f]{6}", p.hcl)
	if !matchString {
		return false
	}

	return true
}

func (p passport) eclValid() bool {
	if p.ecl == "" {
		return false
	}

	if p.ecl != "amb" && p.ecl != "blu" && p.ecl != "brn" && p.ecl != "gry" && p.ecl != "grn" && p.ecl != "hzl" && p.ecl != "oth" {
		return false
	}

	return true
}

func (p passport) pidValid() bool {
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
	readFile, err := ioutil.ReadFile("4/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	passports := make([]*passport, 0)

	allPassports := strings.Split(strings.TrimSpace(string(readFile)), "\n\n")
	for _, onePassport := range allPassports {
		replace := strings.Replace(onePassport, "\n", " ", -1)
		theFields := strings.Split(replace, " ")
		passport := new(passport)
		for _, oneField := range theFields {
			pair := strings.Split(oneField, ":")
			switch pair[0] {
			case "byr":
				passport.byr = pair[1]
			case "iyr":
				passport.iyr = pair[1]
			case "eyr":
				passport.eyr = pair[1]
			case "hgt":
				passport.hgt = pair[1]
			case "hcl":
				passport.hcl = pair[1]
			case "ecl":
				passport.ecl = pair[1]
			case "pid":
				passport.pid = pair[1]
			case "cid":
				passport.cid = pair[1]
			}
		}
		passports = append(passports, passport)
	}

	count := 0
	for _, p := range passports {
		if p.valid() {
			count++
		}
	}

	fmt.Println(count)
}
