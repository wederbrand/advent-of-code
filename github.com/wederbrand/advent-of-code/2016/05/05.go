package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	part1, part2 := part1("cxdnnyjw")
	fmt.Println("Part 1: ", part1, "in", time.Since(start))
	fmt.Println("Part 2: ", part2, "in", time.Since(start))
}

func part1(in string) (string, string) {
	part1 := ""
	part2 := "________"
	index := 0
	for strings.Contains(part2, "_") {
		hash := md5.Sum([]byte(in + fmt.Sprintf("%d", index)))
		hashStr := hex.EncodeToString(hash[:])
		if hashStr[:5] == "00000" {
			posString := string(hashStr[5])
			pos := util.Atoi(posString)
			if len(part1) < 8 {
				part1 += posString
			}
			if posString >= "0" && posString <= "7" && part2[pos] == '_' {
				part2 = part2[:pos] + string(hashStr[6]) + part2[pos+1:]
			}
		}
		index++
	}

	return part1, part2
}
