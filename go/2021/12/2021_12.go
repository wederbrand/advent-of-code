package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type cave struct {
	name     string
	small    bool
	smallish bool
	pathsTo  map[*cave]bool
}

func newCave(name string) *cave {
	c := new(cave)
	c.name = name
	if strings.ToLower(name) == name {
		c.small = true
	}
	c.pathsTo = make(map[*cave]bool)
	return c
}

type visited []string

func (a visited) contains(b string) int {
	i := 0
	for _, s := range a {
		if s == b {
			i++
		}
	}
	return i
}

func main() {
	readFile, err := os.ReadFile("2021/12/2021_12.txt")
	if err != nil {
		log.Fatal(err)
	}

	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")
	caves := make(map[string]*cave)
	for _, s := range inFile {
		split := strings.Split(s, "-")
		c1, found := caves[split[0]]
		if !found {
			c1 = newCave(split[0])
			caves[c1.name] = c1
		}
		c2, found := caves[split[1]]
		if !found {
			c2 = newCave(split[1])
			caves[c2.name] = c2
		}
		c1.pathsTo[c2] = true
		c2.pathsTo[c1] = true
	}

	result := make(map[string]bool)
	explore(caves["start"], visited{}, result)
	fmt.Println("part 1", len(result))

	result = make(map[string]bool)
	for _, c := range caves {
		if c.small && c.name != "start" && c.name != "end" {
			c.small = false
			c.smallish = true
			explore(caves["start"], visited{}, result)
			c.small = true
			c.smallish = false
		}
	}
	fmt.Println("part 2", len(result))

	os.Exit(0)
}

func explore(c *cave, v visited, result map[string]bool) {
	localVisited := make(visited, len(v))
	copy(localVisited, v)
	localVisited = append(localVisited, c.name)
	if c.name == "end" {
		onePath := ""
		for i, s := range localVisited {
			if i != 0 {
				onePath += ","
			}
			onePath += s
		}
		result[onePath] = true
		return
	}

	for pathTo := range c.pathsTo {
		if pathTo.small && localVisited.contains(pathTo.name) == 1 {
			continue
		}
		if pathTo.smallish && localVisited.contains(pathTo.name) == 2 {
			continue
		}
		explore(pathTo, localVisited, result)
	}
}
