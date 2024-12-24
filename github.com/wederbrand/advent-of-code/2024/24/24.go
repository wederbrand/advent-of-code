package main

import (
	"fmt"
	"github.com/awalterschulze/gographviz"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2024/24/input.txt", "\n")

	wires := make(map[string]bool)
	gatesStrings := make([]string, 0)

	for _, s := range inFile {
		if strings.Contains(s, ":") {
			// wire
			name := strings.Split(s, ": ")[0]
			value := Atoi(strings.Split(s, ": ")[1])

			wires[name] = value == 1
		}

		if strings.Contains(s, "->") {
			gatesStrings = append(gatesStrings, s)
		}
	}

	p1 := part1(CloneMap(wires), CloneSlice(gatesStrings))
	fmt.Println("Part 1:", p1, "in", time.Since(start))

	p2 := part2(CloneMap(wires), CloneSlice(gatesStrings))
	fmt.Println("Part 2:", p2, "in", time.Since(start))
}

func part1(wires map[string]bool, gatesStrings []string) int64 {
	matchAll(gatesStrings, wires)
	zString := getBinaryString(wires, "z")
	decimalValue, _ := strconv.ParseInt(zString, 2, 64)
	return decimalValue
}

func part2(wires map[string]bool, gatesStrings []string) string {
	swaps := []string{}
	for i := 0; i < len(gatesStrings); i++ {
		// swaps guessed from graphviz
		// z19 was the first incorrect one
		if gatesStrings[i] == "csn AND nmn -> z19" {
			swaps = append(swaps, "z19")
			gatesStrings[i] = "csn AND nmn -> vwp"
		}
		if gatesStrings[i] == "csn XOR nmn -> vwp" {
			swaps = append(swaps, "vwp")
			gatesStrings[i] = "csn XOR nmn -> z19"
		}

		// z25 was the second incorrect one
		if gatesStrings[i] == "mqj XOR pqn -> mps" {
			swaps = append(swaps, "mps")
			gatesStrings[i] = "mqj XOR pqn -> z25"
		}
		if gatesStrings[i] == "vbw OR qkk -> z25" {
			swaps = append(swaps, "z25")
			gatesStrings[i] = "vbw OR qkk -> mps"
		}

		// z33 was the third incorrect one
		if gatesStrings[i] == "x33 AND y33 -> cqm" {
			swaps = append(swaps, "cqm")
			gatesStrings[i] = "x33 AND y33 -> vjv"
		}
		if gatesStrings[i] == "x33 XOR y33 -> vjv" {
			swaps = append(swaps, "vjv")
			gatesStrings[i] = "x33 XOR y33 -> cqm"
		}

		// this one is weird, I get the right answer with and without it
		if gatesStrings[i] == "mks XOR bhr -> vcv" {
			swaps = append(swaps, "vcv")
			gatesStrings[i] = "mks XOR bhr -> z13"
		}
		if gatesStrings[i] == "x13 AND y13 -> z13" {
			swaps = append(swaps, "z13")
			gatesStrings[i] = "x13 AND y13 -> vcv"
		}
	}

	// Create a new graph
	graph := gographviz.NewGraph()
	graph.SetName("LogicalGates")
	graph.SetDir(true) // Directed graph

	// Add nodes
	seen := make(map[string]bool)

	// Add edges with logical gates as labels
	for _, s := range gatesStrings {
		var a, b, op, out string
		fmt.Sscanf(s, "%s %s %s -> %s", &a, &op, &b, &out)

		if !seen[a] {
			graph.AddNode("LogicalGates", a, nil)
			seen[a] = true
		}
		if !seen[b] {
			graph.AddNode("LogicalGates", b, nil)
			seen[b] = true
		}
		if !seen[out] {
			graph.AddNode("LogicalGates", out, nil)
			seen[out] = true
		}

		graph.AddEdge(a, out, true, map[string]string{"label": op})
		graph.AddEdge(b, out, true, map[string]string{"label": op})
	}

	// Output the graph in DOT format
	err := os.WriteFile("github.com/wederbrand/advent-of-code/2024/24/output.dot", []byte(graph.String()), 0644)
	if err != nil {
		panic(err)
	}

	matchAll(gatesStrings, wires)

	xString := getBinaryString(wires, "x")
	yString := getBinaryString(wires, "y")
	zString := getBinaryString(wires, "z")

	xDecimalValue, _ := strconv.ParseInt(xString, 2, 64)
	yDecimalValue, _ := strconv.ParseInt(yString, 2, 64)

	correct := xDecimalValue + yDecimalValue
	correctString := strconv.FormatInt(correct, 2)

	fmt.Println("Correct   :", correctString)
	fmt.Println("Calculated:", zString)
	for i := 0; i < len(correctString); i++ {
		if correctString[len(correctString)-1-i] != zString[len(correctString)-1-i] {
			fmt.Println("Mismatch at", i)
		}
	}

	slices.Sort(swaps)
	return strings.Join(swaps, ",")
}

func matchAll(gatesStrings []string, wires map[string]bool) {
	for len(gatesStrings) > 0 {
		// pop gates from gatesStrings that can be processed
		for i, s := range gatesStrings {
			var a, b, op, out string
			fmt.Sscanf(s, "%s %s %s -> %s", &a, &op, &b, &out)
			w1, ok1 := wires[a]
			w2, ok2 := wires[b]
			if ok1 && ok2 {
				// both inputs are known
				if op == "AND" {
					wires[out] = w1 && w2
				} else if op == "OR" {
					wires[out] = w1 || w2
				} else if op == "XOR" {
					wires[out] = w1 != w2
				} else {
					panic("Unknown operation")
				}
				// remove gate from gatesStrings
				gatesStrings = append(gatesStrings[:i], gatesStrings[i+1:]...)
				break
			}
		}
	}
}

func getBinaryString(wires map[string]bool, s string) string {
	binaryString := ""
	for i := 0; ; i++ {
		wireName := fmt.Sprintf("%s%02d", s, i)
		wire, exists := wires[wireName]
		if !exists {
			break
		}
		if wire {
			binaryString = "1" + binaryString
		} else {
			binaryString = "0" + binaryString
		}
	}
	return binaryString
}
