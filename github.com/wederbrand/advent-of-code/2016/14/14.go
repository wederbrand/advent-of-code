package main

import (
	"crypto/md5"
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	salt := "zpqevtbw"

	// 25_000 comes from trail and error.
	hashes := make([]string, 25_000)
	for i := range hashes {
		go func(i int) {
			hashes[i] = generate(salt, i, 0)
		}(i)
	}

	part1 := doIt(hashes)
	fmt.Println("Part 1: ", part1, "in", time.Since(start))

	for i := range hashes {
		go func(i int) {
			hashes[i] = generate(salt, i, 2016)
		}(i)
	}

	part2 := doIt(hashes)
	fmt.Println("Part 2: ", part2, "in", time.Since(start))
}

func doIt(hashes []string) int {
	var keys []string
	index := 0
	for len(keys) < 64 {
		hash := hashes[index]
		found := hasTriplet(hash)
		if found != "" {
			// triplet found
			for i := 1; i <= 1000; i++ {
				totalIndex := index + i
				testHash := hashes[totalIndex]
				if strings.Contains(testHash, strings.Repeat(found, 5)) {
					keys = append(keys, hash)
				}
			}
		}
		index++
	}
	return index - 1
}

func generate(salt string, index int, stretches int) string {
	hash := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", salt, index))))
	for i := 0; i < stretches; i++ {
		hash = fmt.Sprintf("%x", md5.Sum([]byte(hash)))
	}
	return hash
}

func hasTriplet(s string) string {
	for i := 0; i < len(s)-2; i++ {
		// Compare current character with next two
		if s[i] == s[i+1] && s[i] == s[i+2] {
			return string(s[i])
		}
	}
	return ""
}
