package main

import (
	"strings"
	"fmt"
	"os"
	"bufio"
)

type Box []rune

func (a Box) diffOne(b Box) bool {
	return len(a.common(b)) == len(a) - 1
}

func (a Box) common(b Box) string {
	var common []rune
	for i := 0; i<len(a); i++ {
		if a[i] == b[i] {
			common = append(common, a[i])
		}
	}

	return string(common)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	input := make([]Box, 0)
	for scanner.Scan() {
		sValue := scanner.Text()
		box := Box(strings.TrimSpace(sValue))
		input = append(input, box)
	}

	for indexA, boxA := range input {
		for indexB, boxB := range input[indexA:] {
			fmt.Println("testing", indexA, indexB)
			if boxA.diffOne(boxB) {
				fmt.Println(boxA.common(boxB))
				os.Exit(0)
			}
		} 
	}
}