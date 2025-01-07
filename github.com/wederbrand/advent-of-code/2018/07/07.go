package main

import (
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"golang.org/x/exp/maps"
	"slices"
	"time"

	"fmt"
	"regexp"
)

type Letters map[string]bool

func main() {
	start := time.Now()
	inFile := GetFileContents("2018/07/input.txt", "\n")

	lineMatcher := regexp.MustCompile(`^Step (.) must be finished before step (.) can begin.$`)

	letters := make(map[string]bool)
	graph1 := make(map[string]map[string]bool)
	graph2 := make(map[string]map[string]bool)

	for _, s := range inFile {
		lineMatch := lineMatcher.FindStringSubmatch(s)
		a := lineMatch[1]
		b := lineMatch[2]
		letters[a] = true
		letters[b] = true
		if graph1[b] == nil {
			graph1[b] = make(Letters)
			graph2[b] = make(Letters)
		}
		graph1[b][a] = true
		graph2[b][a] = true
	}

	// create a slice if keys
	keys := maps.Keys(letters)
	slices.Sort(keys)

	workers := make([]int, 0)
	workers = append(workers, 0)
	workers = append(workers, 0)
	workers = append(workers, 0)
	workers = append(workers, 0)
	workers = append(workers, 0)

	timeFactor := 60
	clock := 0
	events := make(map[int][]string)
	process := make(map[string]bool)

	part1 := ""
	for len(keys) > 0 {
		for _, key := range keys {
			if len(graph1[key]) == 0 {
				part1 += key
				// clean it from the others
				for _, innerMap := range graph1 {
					delete(innerMap, key)
				}

				// delete it
				keys = slices.DeleteFunc(keys, func(s string) bool {
					return s == key
				})
				break
			}
		}
	}

	keys = maps.Keys(letters)
	slices.Sort(keys)
	for len(keys) > 0 {
		for _, key := range keys {
			if len(graph2[key]) == 0 && !process[key] {
				// ready letter
				// take ready worker, if any
				for i, w := range workers {
					if w <= 0 {
						workers[i] = timeFactor + int(key[0]) - 'A' + 1
						events[clock+workers[i]] = append(events[clock+workers[i]], key)
						process[key] = true
						break
					}
				}
			}
		}

		// advance clock
		clock++
		for _, event := range events[clock] {
			for _, key := range event {
				// clean it from the others
				for _, innerMap := range graph2 {
					delete(innerMap, string(key))
				}

				// delete it
				keys = slices.DeleteFunc(keys, func(s string) bool {
					return s == string(key)
				})
				process[string(key)] = false
			}
		}

		for i := range workers {
			workers[i]--
		}
	}

	fmt.Println("Part 1:", part1, "in", time.Since(start))
	fmt.Println("Part 2:", clock, "in", time.Since(start))
}
