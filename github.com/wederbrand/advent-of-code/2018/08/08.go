package main

import (
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strings"
	"time"

	"fmt"
)

type Node struct {
	children []Node
	metadata []int
}

func parseInput(index int) (Node, int) {
	numberOfChildren := input[index]
	numberOfMetadata := input[index+1]
	result := Node{}
	result.children = make([]Node, numberOfChildren)
	nextIndex := index + 2
	var childNode = Node{}
	for i := 0; i < numberOfChildren; i++ {
		childNode, nextIndex = parseInput(nextIndex)
		result.children[i] = childNode
	}

	result.metadata = input[nextIndex : nextIndex+numberOfMetadata]

	return result, nextIndex + numberOfMetadata
}

func sum(n Node) int {
	total := 0
	for _, c := range n.children {
		total += sum(c)
	}

	for i := 0; i < len(n.metadata); i++ {
		total += n.metadata[i]
	}

	return total
}

func sum2(n Node) int {
	total := 0

	if len(n.children) == 0 {
		for i := 0; i < len(n.metadata); i++ {
			total += n.metadata[i]
		}
	}

	for i := 0; i < len(n.metadata); i++ {
		childIndex := n.metadata[i] - 1
		if childIndex < len(n.children) {
			total += sum2(n.children[childIndex])
		}
	}

	return total
}

var input = make([]int, 0)

func main() {
	start := time.Now()
	inFile := GetFileContents("2018/08/input.txt", "\n")

	split := strings.Split(inFile[0], " ")

	for _, s := range split {
		input = append(input, Atoi(s))
	}

	rootNode, _ := parseInput(0)

	fmt.Println("Part 1:", sum(rootNode), "in", time.Since(start))
	fmt.Println("Part 2:", sum2(rootNode), "in", time.Since(start))
}
