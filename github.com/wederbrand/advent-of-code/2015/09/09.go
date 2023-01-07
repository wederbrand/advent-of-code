package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/priorityqueue"
	"math"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2015/09/input.txt", "\n")

	q := priorityqueue.NewQueue()

	routes := make(map[string]int)
	all := make(map[string]bool)
	for _, in := range inFile {
		var src string
		var dst string
		var distance int
		fmt.Sscanf(in, "%s to %s = %d", &src, &dst, &distance)
		routes[key(src, dst)] = distance
		routes[key(dst, src)] = distance
		all[src] = true
		all[dst] = true
	}

	// add all starting points to the queue
	for city := range all {
		s := priorityqueue.State{
			Data:     append([]string{city}, util.AllBut(util.Keys(all), city)...),
			Priority: 0,
		}
		q.Add(&s)
	}

	max := math.MinInt
	for q.HasNext() {
		s := q.Next()
		// first element is current city, all others are remaining, visit all

		cities := s.Data.([]string)
		if len(cities) == 1 {
			// done
			max = util.MaxOf(max, s.Priority)
			util.PrintOnce("part1", s.Priority, "in", time.Since(start))
		}
		src := cities[0]
		cities = util.AllBut(cities, src)
		for _, dst := range cities {
			s2 := priorityqueue.State{
				Data:     append([]string{dst}, util.AllBut(cities, dst)...),
				Priority: s.Priority + routes[key(src, dst)],
			}
			q.Add(&s2)
		}
	}

	fmt.Println("part2", max, "in", time.Since(start))
}

func key(src string, dst string) string {
	return fmt.Sprintf("%s|%s", src, dst)
}
