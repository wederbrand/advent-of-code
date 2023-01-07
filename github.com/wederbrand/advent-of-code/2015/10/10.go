package main

import (
	"fmt"
	"time"
)

type conway struct {
	cnt int
	val int
}

func main() {
	start := time.Now()

	results := make([]conway, 0)
	results = append(results, conway{1, 3})
	results = append(results, conway{2, 1})
	results = append(results, conway{2, 3})
	results = append(results, conway{2, 2})
	results = append(results, conway{2, 1})
	results = append(results, conway{1, 3})

	for i := 0; i < 40; i++ {
		results = lookAndSay(results)
	}
	fmt.Println("part1", count(results), "in", time.Since(start))

	for i := 0; i < 10; i++ {
		results = lookAndSay(results)
	}
	fmt.Println("part2", count(results), "in", time.Since(start))
}

func count(results []conway) int {
	cnt := 0
	for _, c := range results {
		cnt += c.cnt
	}

	return cnt
}

func lookAndSay(in []conway) (out []conway) {
	for _, c := range in {
		out = append(out, conway{1, c.cnt})
		out = append(out, conway{1, c.val})
	}

	out = merge(out)
	return
}

func merge(in []conway) (out []conway) {
	for i, c := range in {
		if i > 0 && out[len(out)-1].val == c.val {
			out[len(out)-1].cnt += c.cnt
		} else {
			out = append(out, c)
		}
	}

	return
}
