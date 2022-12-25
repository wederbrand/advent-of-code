package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func atoi(in string) int {
	slask, _ := strconv.Atoi(in)
	return slask
}

func getRotations() [24]rot {
	return [24]rot{
		{'x', 'y', 'z'},
		{'Y', 'x', 'z'},
		{'X', 'Y', 'z'},
		{'y', 'X', 'z'},
		{'X', 'y', 'Z'},
		{'y', 'x', 'Z'},
		{'x', 'Y', 'Z'},
		{'Y', 'X', 'Z'},
		{'Z', 'y', 'x'},
		{'Z', 'x', 'Y'},
		{'Z', 'Y', 'X'},
		{'Z', 'X', 'y'},
		{'z', 'y', 'X'},
		{'z', 'x', 'y'},
		{'z', 'Y', 'x'},
		{'z', 'X', 'Y'},
		{'x', 'Z', 'y'},
		{'Y', 'Z', 'x'},
		{'X', 'Z', 'Y'},
		{'y', 'Z', 'X'},
		{'x', 'z', 'Y'},
		{'Y', 'z', 'X'},
		{'X', 'z', 'y'},
		{'y', 'z', 'x'},
	}
}

type xyz [3]int

type rot [3]rune

type scannerOfScanners []*scanner

type scanner struct {
	name    string
	pos     xyz
	beacons []*beacon
}

type beacon struct {
	pos     xyz
	vectors []*xyz
}

func (s scannerOfScanners) match(target *scanner) *scanner {
	for _, candidate := range s {
		// for each candidate in scanners of scanner
		sameAs := candidate.sameAs(target)
		if sameAs {
			return candidate
		}
	}
	return nil
}

func (s *scanner) sameAs(target *scanner) bool {
	for _, candidateBeacon := range s.beacons {
		for _, targetBeacon := range target.beacons {
			vectorCount := 0
			for _, candidateVector := range candidateBeacon.vectors {
				for _, targetVector := range targetBeacon.vectors {
					if *candidateVector == *targetVector {
						vectorCount++
						if vectorCount >= 11 {
							// align the scanners with the target
							s.pos[0] = target.pos[0] + targetBeacon.pos[0] - candidateBeacon.pos[0]
							s.pos[1] = target.pos[1] + targetBeacon.pos[1] - candidateBeacon.pos[1]
							s.pos[2] = target.pos[2] + targetBeacon.pos[2] - candidateBeacon.pos[2]

							return true
						}
					}
				}
			}
		}
	}

	return false
}

func (s scanner) calculateAllRotations() []*scanner {
	out := make([]*scanner, 0)

	// for each of the 24 rotations
	for _, r := range getRotations() {
		rotatedScanner := new(scanner)
		rotatedScanner.name = s.name
		out = append(out, rotatedScanner)

		// for each beacon
		for _, b := range s.beacons {
			// rotate it
			b2 := rotateBeacon(b, r)
			rotatedScanner.beacons = append(rotatedScanner.beacons, b2)
		}

		for _, b1 := range rotatedScanner.beacons {
			// then calculate all the vectors for all beacons
			for _, b2 := range rotatedScanner.beacons {
				if b1 == b2 {
					continue
				}

				vector := new(xyz)
				(*vector)[0] = b1.pos[0] - b2.pos[0]
				(*vector)[1] = b1.pos[1] - b2.pos[1]
				(*vector)[2] = b1.pos[2] - b2.pos[2]
				b1.vectors = append(b1.vectors, vector)
			}
		}
	}
	return out
}

func rotateBeacon(b *beacon, r rot) *beacon {
	b2 := new(beacon)
	b2.pos[0] = rotateOne(b, r, 0)
	b2.pos[1] = rotateOne(b, r, 1)
	b2.pos[2] = rotateOne(b, r, 2)

	return b2
}

func rotateOne(b *beacon, r rot, position int) int {
	switch r[position] {
	case 'x':
		return b.pos[0]
	case 'X':
		return -b.pos[0]
	case 'y':
		return b.pos[1]
	case 'Y':
		return -b.pos[1]
	case 'z':
		return b.pos[2]
	case 'Z':
		return -b.pos[2]
	}
	log.Fatal("oh no")
	return 0
}

func main() {
	readFile, err := os.ReadFile("2021/19/2021_19.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")
	re := regexp.MustCompile("--- scanner (\\d+) ---")
	scanners := make([]*scanner, 0)
	var currentScanner *scanner
	for _, s := range inFile {
		if len(s) == 0 {
			continue
		}
		if re.MatchString(s) {
			// new scanner
			submatch := re.FindStringSubmatch(s)
			currentScanner = new(scanner)
			currentScanner.name = submatch[1]
			scanners = append(scanners, currentScanner)
		} else {
			// new beacon
			split := strings.Split(s, ",")
			b := &beacon{
				pos: xyz{
					atoi(split[0]),
					atoi(split[1]),
					atoi(split[2]),
				},
				vectors: make([]*xyz, 0),
			}
			currentScanner.beacons = append(currentScanner.beacons, b)
		}
	}

	allScannersOfScanners := make([]scannerOfScanners, 0)
	for _, s := range scanners {
		allScannersOfScanners = append(allScannersOfScanners, s.calculateAllRotations())
	}

	// by now we should have as many scanner of scanners as we had scanners, just that they are all rotated
	// pick 0 and put in all
	// check all remaining to see if they match
	// if the do, remove from allScannersOfScanners and put in all with the picked alignment

	found := make([]*scanner, 0)
	found = append(found, allScannersOfScanners[0][0])
	allScannersOfScanners = allScannersOfScanners[1:]

outer:
	for len(allScannersOfScanners) > 0 {
		for i, sos := range allScannersOfScanners {
			for _, f := range found {
				match := sos.match(f)
				if match != nil {
					found = append(found, match)
					allScannersOfScanners = append(allScannersOfScanners[:i], allScannersOfScanners[i+1:]...)
					continue outer
				}
			}
		}
	}

	// all aligned and positions

	allBeacons := make(map[string]bool)
	for _, s := range found {
		for _, b := range s.beacons {
			key := ""
			key += strconv.Itoa(s.pos[0] + b.pos[0])
			key += ","
			key += strconv.Itoa(s.pos[1] + b.pos[1])
			key += ","
			key += strconv.Itoa(s.pos[2] + b.pos[2])

			allBeacons[key] = true
		}
	}
	fmt.Println("part 1", len(allBeacons))

	maxManhattan := math.MinInt
	for _, s1 := range found {
		for _, s2 := range found {
			if s1 == s2 {
				continue
			}

			dx := s1.pos[0] - s2.pos[0]
			dy := s1.pos[1] - s2.pos[1]
			dz := s1.pos[2] - s2.pos[2]
			manhattan := abs(dx) + abs(dy) + abs(dz)

			if manhattan > maxManhattan {
				maxManhattan = manhattan
			}
		}
	}
	fmt.Println("part 2", maxManhattan)

}

func abs(in int) int {
	if in < 0 {
		return -in
	}
	return in
}
