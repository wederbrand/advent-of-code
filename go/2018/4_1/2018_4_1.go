package main

import (
	"strconv"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"time"
)

type input struct {
	logTime time.Time
	action  string
}

type By func(p1, p2 *input) bool

func (by By) Sort(inputs []input) {
	ps := &inputSorter{
		inputs: inputs,
		by:     by,
	}
	sort.Sort(ps)
}

type inputSorter struct {
	inputs []input
	by     func(p1, p2 *input) bool
}

func (s *inputSorter) Len() int {
	return len(s.inputs)
}

func (s *inputSorter) Swap(i, j int) {
	s.inputs[i], s.inputs[j] = s.inputs[j], s.inputs[i]
}

func (s *inputSorter) Less(i, j int) bool {
	return s.by(&s.inputs[i], &s.inputs[j])
}

type guard struct {
	id int
	minutes map[int]int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lineMatcher := regexp.MustCompile(`^\[(.+)\] (.*)$`)

	lines := make([]input, 0)
	for scanner.Scan() {
		sValue := scanner.Text()
		lineMatch := lineMatcher.FindStringSubmatch(sValue)
		logTime, _ := time.Parse("2006-01-02 15:04", lineMatch[1])
		action := lineMatch[2]
		i := input{logTime, action}
		lines = append(lines, i)
	}

	timeValue := func(p1, p2 *input) bool {
		return p1.logTime.Before(p2.logTime)
	}

	By(timeValue).Sort(lines)

	guards := make(map[int]guard, 0)
	var currentGuard guard
	
	guardMatcher := regexp.MustCompile(`^Guard #(\d+) begins shift$`)
	sleepSince := 0
	for _, line := range lines {
		guardMatch := guardMatcher.FindStringSubmatch(line.action)
		if guardMatch != nil {
			// new guard
			guardIndex, _ := strconv.Atoi(guardMatch[1])
			currentGuard = guards[guardIndex]
			if (currentGuard.minutes == nil) {
				currentGuard.id = guardIndex
				currentGuard.minutes = make(map[int]int)
			}			
		} else {
			// no new guard
			if line.action[0] == 'f' {
				// falls asleep
				sleepSince = line.logTime.Minute()
			} else {
				// wakes up
				for i := sleepSince ; i < line.logTime.Minute() ; i++ {
					currentGuard.minutes[i]++
				}				
				// this is good time to store the guard back
				guards[currentGuard.id] = currentGuard
			}
		}
	}

	sleepyGuard := currentGuard
	maxSleep := 0
	for _, currentGuard = range guards {
		totalSleep := 0
		for _, minutes := range currentGuard.minutes {
			totalSleep += minutes
		}
		if (totalSleep > maxSleep) {
			sleepyGuard = currentGuard
			maxSleep = totalSleep
		}
	}

	sleepyTime := 0
	maxMins := 0
	for i, m := range sleepyGuard.minutes {
		if (m > maxMins) {
			sleepyTime = i
			maxMins = m
		}
	}

	fmt.Println("part 1:", sleepyGuard.id, sleepyTime, sleepyGuard.id*sleepyTime)

	maxFreq := 0
	maxMinute := 0
	for _, currentGuard := range guards {
		for minute, freq := range currentGuard.minutes {
			if (freq > maxFreq) {
				maxMinute = minute;
				maxFreq = freq
				sleepyGuard = currentGuard
			}
		}
	}

	fmt.Println("part 2:", sleepyGuard.id, maxMinute, sleepyGuard.id*maxMinute)

}
