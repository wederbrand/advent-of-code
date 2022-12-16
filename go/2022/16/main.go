package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

type vertex struct {
	target string
	price  int
}

type valve struct {
	index    int
	name     string
	flow     int
	tunnels  []string
	vertexes map[string]int // time per name
}

func (v *valve) setTunnels(tunnels string) {
	for _, s := range strings.Split(tunnels, ",") {
		name := strings.TrimSpace(s)
		v.tunnels = append(v.tunnels, name)
	}
}

func (v *valve) scanForVertexes(valves map[string]*valve) {
	q := make([]vertex, 0)
	// queue all tunnels
	for _, tunnel := range v.tunnels {
		q = append(q, vertex{
			target: tunnel,
			price:  2,
		})
	}

	for len(q) > 0 {
		v2 := q[0]
		q = q[1:]

		if v2.target == v.name {
			continue
		}

		i, found := v.vertexes[v2.target]
		if !found || v2.price < i {
			// replace it and queue all it's tunnels
			v.vertexes[v2.target] = v2.price
			for _, tunnel := range valves[v2.target].tunnels {
				q = append(q, vertex{
					target: tunnel,
					price:  v2.price + 1,
				})
			}
		}
	}

	// clear out zeros
	for s := range v.vertexes {
		v2 := valves[s]
		if v2.flow == 0 {
			delete(v.vertexes, s)
		}
	}
}

func main() {
	start := time.Now()
	readFile, err := os.ReadFile("16/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	valves := make(map[string]*valve)
	valveIndex := 1

	for _, s := range inFile {
		var name string
		var flow int
		fmt.Sscanf(s, "Valve %s has flow rate=%d; tunnels lead to valves", &name, &flow)

		_, after, found := strings.Cut(s, "to valves")
		if !found {
			_, after, _ = strings.Cut(s, "to valve")
		}

		v := &valve{
			index:    valveIndex,
			name:     name,
			flow:     flow,
			vertexes: make(map[string]int),
		}
		v.setTunnels(after)
		valves[name] = v
	}

	// create a graph with all nodes and "lines"
	for _, v := range valves {
		v.scanForVertexes(valves)
	}

	v := valves["AA"]
	notVisited := make([]string, 0)
	for s, v := range valves {
		if v.flow > 0 {
			notVisited = append(notVisited, s)
		}
	}
	max := getMax(v, valves, notVisited, 0, 0, 30)
	fmt.Println("part1:", max, "in", time.Since(start))

	all := allAlternatives(notVisited)

	max = math.MinInt
	for _, alternative := range all {
		meMax := getMax(v, valves, alternative[0], 0, 0, 26)
		elephantMax := getMax(v, valves, alternative[1], 0, 0, 26)
		if meMax+elephantMax > max {
			max = meMax + elephantMax
		}
	}

	fmt.Println("part2:", max, "in", time.Since(start))

	// split notVisited in all non overlapping permutations of all valves
	// test each
	// keep best
}

func allAlternatives(in []string) [][][]string {
	l := len(in)
	var result [][][]string
	for i := 0; i < (1 << uint(l)); i++ {
		var incl []string
		var excl []string
		for j := 0; j < l; j++ {
			if i&(1<<uint(j)) > 0 {
				incl = append(incl, in[j])
			} else {
				excl = append(excl, in[j])
			}
		}
		result = append(result, [][]string{incl, excl})
	}
	return result
}

func getMax(v *valve, valves map[string]*valve, notVisited []string, flow int, total int, timeLeft int) int {
	max := total + timeLeft*flow
	for i, tunnel := range notVisited {
		price := v.vertexes[tunnel]
		if price > timeLeft {
			continue
		}
		v2 := valves[tunnel]
		newNotVisited := make([]string, len(notVisited))
		copy(newNotVisited, notVisited)
		newNotVisited = append(newNotVisited[:i], newNotVisited[i+1:]...)
		value := getMax(v2, valves, newNotVisited, flow+v2.flow, total+flow*price, timeLeft-price)
		if value > max {
			max = value
		}
	}

	return max
}
