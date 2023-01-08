package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"math"
	"time"
)

type deer struct {
	name              string
	speed             int
	flyTime           int
	restTime          int
	distance          int
	flying            bool
	remainingFlyTime  int
	remainingRestTime int
	score             int
}

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2015/14/input.txt", "\n")

	deers := make([]*deer, 0)
	for _, s := range inFile {
		d := new(deer)
		fmt.Sscanf(s, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &d.name, &d.speed, &d.flyTime, &d.restTime)
		d.flying = true
		d.remainingFlyTime = d.flyTime
		deers = append(deers, d)
	}

	for i := 0; i < 2503; i++ {
		scores := make(map[int][]*deer)
		max := math.MinInt
		for _, d := range deers {
			if d.flying {
				d.distance += d.speed
				d.remainingFlyTime--
				if d.remainingFlyTime == 0 {
					d.flying = false
					d.remainingRestTime = d.restTime
				}
			} else {
				d.remainingRestTime--
				if d.remainingRestTime == 0 {
					d.flying = true
					d.remainingFlyTime = d.flyTime
				}
			}
			scores[d.distance] = append(scores[d.distance], d)
			max = util.MaxOf(max, d.distance)
		}
		for _, d := range scores[max] {
			d.score++
		}
	}

	part1 := math.MinInt
	for _, d := range deers {
		part1 = util.MaxOf(part1, d.distance)
	}

	fmt.Println("part1", part1, "in", time.Since(start))

	part2 := math.MinInt
	for _, d := range deers {
		part2 = util.MaxOf(part2, d.score)
	}

	fmt.Println("part2", part2, "in", time.Since(start))
}
