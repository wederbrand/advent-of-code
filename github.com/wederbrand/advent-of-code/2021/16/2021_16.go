package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type packet struct {
	input   string
	sub     []*packet
	version int
	typeId  int
	value   int
}

func (p *packet) getTotalVersion() int {
	result := p.version

	for _, p2 := range p.sub {
		result += p2.getTotalVersion()
	}

	return result
}

func (p *packet) getResultingValue() int {
	switch p.typeId {

	//Packets with type ID 0 are sum packets - their value is the sum of the values of their sub-packets. If they only have a single sub-packet, their value is the value of the sub-packet.
	case 0:
		sum := 0
		for _, p2 := range p.sub {
			sum += p2.getResultingValue()
		}
		return sum
	//Packets with type ID 1 are product packets - their value is the result of multiplying together the values of their sub-packets. If they only have a single sub-packet, their value is the value of the sub-packet.
	case 1:
		prod := 1
		for _, p2 := range p.sub {
			prod *= p2.getResultingValue()
		}
		return prod
	//Packets with type ID 2 are minimum packets - their value is the minimum of the values of their sub-packets.
	case 2:
		min := math.MaxInt
		for _, p2 := range p.sub {
			value := p2.getResultingValue()
			if value < min {
				min = value
			}
		}
		return min
	//Packets with type ID 3 are maximum packets - their value is the maximum of the values of their sub-packets.
	case 3:
		max := math.MinInt
		for _, p2 := range p.sub {
			value := p2.getResultingValue()
			if value > max {
				max = value
			}
		}
		return max
	case 4:
		return p.value
	//Packets with type ID 5 are greater than packets - their value is 1 if the value of the first sub-packet is greater than the value of the second sub-packet; otherwise, their value is 0. These packets always have exactly two sub-packets.
	case 5:
		if p.sub[0].getResultingValue() > p.sub[1].getResultingValue() {
			return 1
		}
		return 0
	//Packets with type ID 6 are less than packets - their value is 1 if the value of the first sub-packet is less than the value of the second sub-packet; otherwise, their value is 0. These packets always have exactly two sub-packets.
	case 6:
		if p.sub[0].getResultingValue() < p.sub[1].getResultingValue() {
			return 1
		}
		return 0
	//Packets with type ID 7 are equal to packets - their value is 1 if the value of the first sub-packet is equal to the value of the second sub-packet; otherwise, their value is 0. These packets always have exactly two sub-packets.
	case 7:
		if p.sub[0].getResultingValue() == p.sub[1].getResultingValue() {
			return 1
		}
		return 0
	}
	return -1
}

func parsePacket(in string) (*packet, string) {
	p := new(packet)
	p.version = str2bin(in[0:3])
	p.typeId = str2bin(in[3:6])
	rest := in[6:]
	if p.typeId == 4 {
		p.value, rest = decodeLiteral(rest)
	} else {
		p.sub, rest = decodeOperation(rest)
	}

	return p, rest
}

func decodeOperation(s string) ([]*packet, string) {
	lengthTypeId := str2bin(s[0:1])
	result := make([]*packet, 0)
	rest := ""
	if lengthTypeId == 0 {
		totalLength := str2bin(s[1:16])
		rest = s[16 : 16+totalLength]
		for {
			var p *packet
			p, rest = parsePacket(rest)
			result = append(result, p)
			if len(rest) == 0 {
				break
			}
		}
		rest = s[16+totalLength:]
	} else {
		nbrOfPackets := str2bin(s[1:12])
		rest = s[12:]
		for i := 0; i < nbrOfPackets; i++ {
			var p *packet
			p, rest = parsePacket(rest)
			result = append(result, p)
		}
	}
	return result, rest
}

func decodeLiteral(s string) (int, string) {
	binString := ""
	i := 0
	for {
		binString += s[i+1 : i+5]

		if s[i] == '0' {
			break
		}
		i += 5
	}
	return str2bin(binString), s[i+5:]
}

func str2bin(in string) int {
	parseInt, _ := strconv.ParseInt(in, 2, 64)
	return int(parseInt)
}

func main() {
	readFile, err := os.ReadFile("2021/16/2021_16.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), "")
	all := ""
	for _, str := range inFile {
		parseInt, _ := strconv.ParseInt(str, 16, 64)
		binStr := strconv.FormatInt(parseInt, 2)
		padded := fmt.Sprintf("%04s", binStr)
		all += padded
	}
	p, _ := parsePacket(all)

	// part 1
	part1 := p.getTotalVersion()
	fmt.Println(part1)
	part2 := p.getResultingValue()
	fmt.Println(part2)
}
