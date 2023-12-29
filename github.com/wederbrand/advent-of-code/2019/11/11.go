package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/2019/computer"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"time"
)

type Robot struct {
	p chart.Coord
	d chart.Dir
}

const WHITE = "#"
const BLACK = "."

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2019/11/input.txt", "\n")

	computer := NewComputer(inFile)
	robot := Robot{chart.Coord{0, 0}, chart.UP}
	m := chart.Chart{}
	runComputerUntilExit(computer, m, robot)

	part1 := len(m)
	fmt.Println("part1: ", part1, "in", time.Since(start))

	computer = NewComputer(inFile)
	robot = Robot{chart.Coord{0, 0}, chart.UP}
	m = chart.Chart{}
	m[robot.p] = WHITE
	runComputerUntilExit(computer, m, robot)

	chart.PrintChart(m)
}

func runComputerUntilExit(computer Computer, m chart.Chart, robot Robot) {
	go computer.Run()

	for computer.IsRunning() {
		s, found := m[robot.p]
		// I assume a race condition here, where the computer has been stopped after we checked, but before we came to sending input
		if found && s == WHITE {
			computer.GetInput() <- 1
		} else {
			computer.GetInput() <- 0
		}
		newColor, colorOK := <-computer.GetOutput()
		if !colorOK {
			return
		}
		newDir, dirOK := <-computer.GetOutput()
		if !dirOK {
			return
		}
		if newColor == 0 {
			m[robot.p] = BLACK
		} else {
			m[robot.p] = WHITE
		}
		if newDir == 0 {
			robot.d = robot.d.Left()
		} else {
			robot.d = robot.d.Right()
		}
		robot.p = robot.p.Move(robot.d, 1)
		// This is sad but the computer doesn't get to a none running state before I check it again
		time.Sleep(time.Microsecond)
	}
}
