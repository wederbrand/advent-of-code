package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2016/18/input.txt", "\n")

	maxX := len(inFile[0])
	lines := make([]string, 0)
	lines = append(lines, inFile[0])
	safe := strings.Count(lines[0], ".")
	for {
		lastLine := "." + lines[len(lines)-1] + "."
		newLine := ""

		for i := 0; i < maxX; i++ {
			testString := lastLine[i : i+3]
			if testString == "^^." || testString == ".^^" || testString == "^.." || testString == "..^" {
				newLine += "^"
			} else {
				newLine += "."
			}
		}

		lines = append(lines, newLine)
		safe += strings.Count(newLine, ".")
		if len(lines) == 40 {
			fmt.Println("Part 1: ", safe, "in", time.Since(start))
		} else if len(lines) == 400000 {
			fmt.Println("Part 2: ", safe, "in", time.Since(start))
			break
		}
	}
}
