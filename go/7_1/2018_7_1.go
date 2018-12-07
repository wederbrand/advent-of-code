package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
)

type Letters map[string]struct{}

func main() {
	file, err := os.Open("2018_7.input");
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	lineMatcher := regexp.MustCompile(`^Step (.) must be finished before step (.) can begin.$`)

	letters := make(map[string]struct{}, 0)
	graph := make(map[string]map[string]struct{})

	for scanner.Scan() {
		lineMatch := lineMatcher.FindStringSubmatch(scanner.Text())
		a := lineMatch[1]
		b := lineMatch[2]
		letters[a] = struct{}{}
		letters[b] = struct{}{}
		if graph[b] == nil {
			graph[b] = make(Letters, 0)
		}
		graph[b][a] = struct{}{}
	}
	
	// create a slice if keys
	keys := make([]string, 0, len(letters))
	for k := range letters {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	result := make([]string, 0, len(letters))

	outer:
	for len(keys) > 0 {
		for i, key := range keys {
			if len(graph[key]) == 0 {
				// store it
				result = append(result, key)
				// clean it from the others
				for _, innerMap := range graph {
					delete (innerMap, key)
				}
				// delete it
				keys = append(keys[:i], keys[i+1:]...)
				// start from the beginning
				continue outer
			}
		}
	}

	for _, r := range result {
		fmt.Print(r)
	}
	fmt.Println()
}
