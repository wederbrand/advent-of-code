package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"slices"
	"strings"
	"time"
)

type Node struct {
	name        string
	connections []*Node
}

func getOrCreate(s string, m map[string]*Node) *Node {
	node, found := m[s]
	if !found {
		node = &Node{s, make([]*Node, 0)}
		m[s] = node
	}

	return node
}

func main() {
	startTimer := time.Now()
	inFile := util.GetFileContents("2023/25/input.txt", "\n")

	allNodes := make(map[string]*Node)
	for _, s := range inFile {
		split := strings.Split(s, ":")
		src := getOrCreate(split[0], allNodes)
		for _, s2 := range strings.Split(strings.TrimSpace(split[1]), " ") {
			dst := getOrCreate(s2, allNodes)
			src.connections = append(src.connections, dst)
			dst.connections = append(dst.connections, src)
		}
	}

	// convert input to dot file
	// graph xmax {
	// a: b c d becomes a -- {b c d}
	// }

	// graphically see that I need to cut
	// neato github.com/wederbrand/advent-of-code/2023/25/input.txt -Tjpg > output.jpg
	// sqh - jbz
	// nvg - vfj
	// fvh - fch
	// remove these
	cut(allNodes, "sqh", "jbz")
	cut(allNodes, "jbz", "sqh")
	cut(allNodes, "nvg", "vfj")
	cut(allNodes, "vfj", "nvg")
	cut(allNodes, "fvh", "fch")
	cut(allNodes, "fch", "fvh")

	one := sizeOfCluster(allNodes["sqh"])
	two := sizeOfCluster(allNodes["jbz"])
	part1 := one * two
	fmt.Println("part1: ", part1, "in", time.Since(startTimer))
}

func sizeOfCluster(node *Node) int {
	seen := make(map[string]bool)
	q := make([]*Node, 0)
	q = append(q, node)
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		_, found := seen[n.name]
		if found {
			continue
		}
		seen[n.name] = true
		for _, connection := range n.connections {
			q = append(q, connection)
		}
	}

	return len(seen)
}

func cut(allNodes map[string]*Node, a string, b string) {
	allNodes[a].connections = slices.DeleteFunc(allNodes[a].connections, func(node *Node) bool {
		return node.name == b
	})
}
