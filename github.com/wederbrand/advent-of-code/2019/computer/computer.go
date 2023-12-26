package computer

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strconv"
)

type Computer struct {
	p int // conveniently defaults to 0
	m []int
}

func NewComputer(input []string) Computer {
	instructions := make([]int, 0)
	for _, i := range input {
		result := util.MatchingNumbersAfterSplitOnAny(i, "", ",")
		instructions = append(instructions, result[0]...)
	}
	return Computer{m: instructions}
}

func (c *Computer) SetMemory(addr int, val int) {
	c.m[addr] = val
}

func (c *Computer) GetMemory(addr int) int {
	return c.m[addr]
}

func (c *Computer) Run() {
	for {
		switch c.m[c.p] {
		case 1:
			c.m[c.m[c.p+3]] = c.m[c.m[c.p+1]] + c.m[c.m[c.p+2]]
			c.p += 4
		case 2:
			c.m[c.m[c.p+3]] = c.m[c.m[c.p+1]] * c.m[c.m[c.p+2]]
			c.p += 4
		case 99:
			c.p += 1
			return
		}
	}
}

func (c *Computer) Print() {
	for _, i := range c.m {
		fmt.Print(strconv.Itoa(i), " ")
	}
	fmt.Println()
}
