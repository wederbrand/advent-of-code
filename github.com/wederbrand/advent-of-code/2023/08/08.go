package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strings"
	"time"
)

type Node struct {
	name     string
	leftRaw  string
	rightRaw string
	left     *Node
	right    *Node
}

func main() {
	startTimer := time.Now()
	inFile := util.GetFileContents("2023/08/input.txt", "\n")

	inst := inFile[0]

	nodes := make(map[string]*Node)
	start := make([]*Node, 0)
	for i, s := range inFile {
		if i < 2 {
			continue
		}

		node := new(Node)

		fmt.Sscanf(s, "%s = (%3s, %3s)", &node.name, &node.leftRaw, &node.rightRaw)
		nodes[node.name] = node
		if strings.HasSuffix(node.name, "A") {
			start = append(start, node)
		}
	}

	for _, node := range nodes {
		node.left = nodes[node.leftRaw]
		node.right = nodes[node.rightRaw]
	}

	fmt.Println("parsing:", time.Since(startTimer))
	startTimer = time.Now()

	part1 := solveOne(nodes["AAA"], inst)
	fmt.Println("part1: ", part1, "in", time.Since(startTimer))

	startTimer = time.Now()
	parts := make([]int, 0)
	for i := range start {
		one := solveOne(start[i], inst)
		if one%len(inst) != 0 {
			panic("ho ho, it would have been good if this was a full multiple of the instructions")
		}
		parts = append(parts, one)
	}
	part2 := util.Lcd(parts)
	fmt.Println("part2: ", part2, "in", time.Since(startTimer))
}

func solveOne(current *Node, inst string) int {
	one := 0
	for i := 0; !strings.HasSuffix(current.name, "Z"); i++ {
		one++
		if inst[i%len(inst)] == 'L' {
			current = current.left
		} else {
			current = current.right
		}
	}
	return one
}
