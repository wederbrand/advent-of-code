package computer

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strconv"
)

type Computer struct {
	Name         string
	running      bool
	p            int // conveniently defaults to 0
	m            map[int]int
	relativeBase int
	in           chan int
	out          chan int
}

func NewComputer(input []string) Computer {
	instructions := make(map[int]int)
	for i, in := range util.MatchingNumbersAfterSplitOnAny(input[0], "", ",")[0] {
		instructions[i] = in
	}

	return Computer{
		p:            0,
		m:            instructions,
		relativeBase: 0,
		in:           make(chan int),
		out:          make(chan int),
		running:      true,
	}
}

func (c *Computer) IsRunning() bool {
	return c.running
}

func (c *Computer) SetMemory(addr int, val int) {
	c.m[addr] = val
}

func (c *Computer) GetMemory(addr int) int {
	return c.m[addr]
}

func (c *Computer) SetInput(newChan chan int) {
	c.in = newChan
}

func (c *Computer) GetInput() chan int {
	return c.in
}

func (c *Computer) GetOutput() chan int {
	return c.out
}

func (c *Computer) Run() {
	c.running = true
	for {
		if c.p >= len(c.m) {
			panic(c.Name + " out of bounds instruction " + strconv.Itoa(c.p) + " " + strconv.Itoa(len(c.m)))
		}

		op := c.m[c.p] % 100
		mode1 := (c.m[c.p] / 100) % 10
		mode2 := (c.m[c.p] / 1000) % 10
		mode3 := (c.m[c.p] / 10000) % 100000
		switch op {
		case 1: // addition
			c.m[c.addr(mode3, 3)] = c.m[c.addr(mode1, 1)] + c.m[c.addr(mode2, 2)]
			c.p += 4

		case 2: // multiply
			c.m[c.addr(mode3, 3)] = c.m[c.addr(mode1, 1)] * c.m[c.addr(mode2, 2)]
			c.p += 4

		case 3: // store input
			c.m[c.addr(mode1, 1)] = <-c.in
			c.p += 2

		case 4: // return output
			c.out <- c.m[c.addr(mode1, 1)]
			c.p += 2

		case 5: // jump-if-true
			value := c.m[c.addr(mode1, 1)]
			if value != 0 {
				c.p = c.m[c.addr(mode2, 2)]
			} else {
				c.p += 3
			}

		case 6: // jump-if-false
			value := c.m[c.addr(mode1, 1)]
			if value == 0 {
				c.p = c.m[c.addr(mode2, 2)]
			} else {
				c.p += 3
			}

		case 7: // less than
			value1 := c.m[c.addr(mode1, 1)]
			value2 := c.m[c.addr(mode2, 2)]
			if value1 < value2 {
				c.m[c.addr(mode3, 3)] = 1
			} else {
				c.m[c.addr(mode3, 3)] = 0
			}

			c.p += 4

		case 8: // equals
			value1 := c.m[c.addr(mode1, 1)]
			value2 := c.m[c.addr(mode2, 2)]
			if value1 == value2 {
				c.m[c.addr(mode3, 3)] = 1
			} else {
				c.m[c.addr(mode3, 3)] = 0
			}

			c.p += 4

		case 9: // adjust relative base
			c.relativeBase += c.m[c.addr(mode1, 1)]
			c.p += 2

		case 99:
			c.p += 1
			c.running = false
			close(c.out)
			return
		default:
			panic(c.Name + " invalid instruction " + strconv.Itoa(op))
		}
	}
}

func (c *Computer) Print() {
	for _, i := range c.m {
		fmt.Print(strconv.Itoa(i), " ")
	}
	fmt.Println()
}

func (c *Computer) addr(mode int, i int) int {
	if mode == 0 {
		// position mode
		return c.m[c.p+i]
	} else if mode == 1 {
		// immediate mode
		return c.p + i
	} else if mode == 2 {
		// relative mode
		return c.m[c.p+i] + c.relativeBase
	}

	panic("invalid mode")
}
