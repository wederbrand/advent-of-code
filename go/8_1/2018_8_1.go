package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	file, err := os.Open("2018_8.input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		input = append(input, i)
	}

	rootNode, _ := parseInput(0)

	fmt.Println("part1", sum(rootNode))

	fmt.Println("part2", sum2(rootNode))
}
