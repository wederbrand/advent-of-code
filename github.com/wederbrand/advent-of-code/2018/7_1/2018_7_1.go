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
	file, err := os.Open("2018_7.input")
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

	workers := make([]int, 0)
	workers = append(workers, 0)
	workers = append(workers, 0)
	workers = append(workers, 0)
	workers = append(workers, 0)
	workers = append(workers, 0)

	timeFactor := 60
	time := 0
	events := make(map[int][]string)
	process := make(map[string]bool)

	for len(keys) > 0 {
		for _, key := range keys {
			if len(graph[key]) == 0 && !process[key] {
				// ready letter
				// take ready worker, if any
				for i, w := range workers {
					if w <= 0 {
						workers[i] = timeFactor + int(key[0]) - 'A' + 1
						events[time+workers[i]] = append(events[time+workers[i]], key)
						process[key] = true
						break
					}
				}

			}
		}

		// advance time
		time++
		for _, event := range events[time] {
			for _, key := range event {
				// clean it from the others
				for _, innerMap := range graph {
					delete(innerMap, string(key))
				}

				// delete it
				for i := range keys {
					if keys[i] == string(key) {
						keys = append(keys[:i], keys[i+1:]...)
						break
					}
				}
				process[string(key)] = false
			}
		}

		for i := range workers {
			workers[i]--
		}

	}

	fmt.Println(time)
}
