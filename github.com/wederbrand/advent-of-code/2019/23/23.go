package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/2019/computer"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"sync"
	"time"
)

type IOComputer struct {
	name   int
	lock   sync.Mutex
	comp   Computer
	input  chan int
	output chan int
}

func main() {
	start := time.Now()
	inFile := GetFileContents("2019/23/input.txt", "\n")

	in, out := make([]chan int, 50), make([]chan int, 50)
	computers := make([]*IOComputer, 50)

	for i := 0; i < 50; i++ {
		in[i], out[i] = make(chan int), make(chan int)

		computers[i] = makeComputer(i, inFile, in[i], out[i])
		go computers[i].comp.Run()
		in[i] <- i
		in[i] <- -1

	}

	idle := 0

	lastNatY := 0
	natX, natY := 0, 0

	for i := 0; ; i = (i + 1) % 50 {
		select {
		case addr := <-out[i]:
			if addr == 255 {
				nextNatX := <-out[i]
				nextNatY := <-out[i]
				if natY == 0 {
					fmt.Println("part1:", nextNatY, "in", time.Since(start))
				}
				natX, natY = nextNatX, nextNatY
			} else {
				in[addr] <- <-out[i]
				in[addr] <- <-out[i]
			}
			idle = 0
		case in[i] <- -1:
			idle++
		}

		if idle >= 50 {
			if natY == lastNatY {
				fmt.Println("part2:", natY, "in", time.Since(start))
				return
			}
			in[0] <- natX
			in[0] <- natY
			lastNatY = natY
			idle = 0
		}
	}
}

func makeComputer(name int, inFile []string, input chan int, output chan int) *IOComputer {
	ioc := IOComputer{name: name, input: input, output: output}

	in := func() int {
		return <-ioc.input
	}

	out := func(i int) {
		ioc.output <- i
	}

	ioc.comp = NewComputer(inFile, in, out)
	return &ioc
}
