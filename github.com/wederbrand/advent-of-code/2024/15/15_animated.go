package main

import (
	"github.com/gdamore/tcell"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"log"
	"strings"
)

func main() {
	inFile := GetFileContents("2024/15/input.txt", "\n")
	chartData1 := make([]string, 0)
	chartData2 := make([]string, 0)

	spaceIndex := 0
	for i, s := range inFile {
		if s == "" {
			spaceIndex = i
			break
		}

		chartData1 = append(chartData1, s)
		// for part 2 duplicate all
		line := ""
		for _, r := range s {
			if r == '.' {
				line += ".."
			} else if r == '#' {
				line += "##"
			} else if r == 'O' {
				line += "[]"
			} else if r == '@' {
				line += "@."
			} else {
				panic("hoho")
			}
		}
		chartData2 = append(chartData2, line)
	}

	moves := strings.Join(inFile[spaceIndex+1:], "")

	c1 := MakeChart(chartData1, "")
	doItAnimated(c1, moves)

	c2 := MakeChart(chartData2, "")
	doItAnimated(c2, moves)
}

func doItAnimated(c Chart, moves string) {
	tCell, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	} else {
		if err := tCell.Init(); err != nil {
			log.Fatalf("%+v", err)
		}
		// Set default text style
		defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorMediumSpringGreen)
		tCell.SetStyle(defStyle)

		// Clear screen
		tCell.Clear()

		tCellMap(tCell, c)
	}

	robot := Coord{}

	for coord, s := range c {
		if s == "@" {
			robot = coord
			break
		}
	}

	for _, r := range moves {
		d := UP
		switch r {
		case '^':
			d = UP
		case 'v':
			d = DOWN
		case '<':
			d = LEFT
		case '>':
			d = RIGHT
		}

		if pushItAnimated(c, robot, d, false) {
			pushItAnimated(c, robot, d, true)
			robot = robot.Move(d)
			tCellMap(tCell, c)
		}
	}
}

func pushItAnimated(c Chart, object Coord, d Dir, move bool) bool {
	next := object.Move(d, 1)
	if c[next] == "#" {
		// wall, that's the end of it
		return false
	}

	if c[next] == "." {
		// empty space, no more pushing
		if move {
			c[object], c[next] = c[next], c[object]
		}
		return true
	}

	if c[next] == "O" || ((c[next] == "[" || c[next] == "]") && (d == LEFT || d == RIGHT)) {
		// push the food first
		if pushItAnimated(c, next, d, move) {
			if move {
				c[object], c[next] = c[next], c[object]
			}
			return true
		} else {
			return false
		}
	}

	if (c[next] == "[" || c[next] == "]") && (d == UP || d == DOWN) {
		// also push the othe half
		theOtherHalf := next.Move(RIGHT)
		if c[next] == "]" {
			theOtherHalf = next.Move(LEFT)
		}
		if pushItAnimated(c, next, d, move) && pushItAnimated(c, theOtherHalf, d, move) {
			if move {
				c[object], c[next] = c[next], c[object]
			}
			return true
		} else {
			return false
		}
	}

	panic("ho ho")
}

func tCellMap(cell tcell.Screen, m Chart) {
	robot := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorOrangeRed)
	food := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorSteelBlue)
	minC, maxC := GetChartMaxes(m)
	for y := minC.Y; y <= maxC.Y; y++ {
		for x := minC.X; x <= maxC.X; x++ {
			s, found := m[Coord{x, y}]
			if !found {
				s = "."
			}
			if s == "@" {
				cell.SetContent(x, y, rune(s[0]), nil, robot)
			} else if s == "O" || s == "[" || s == "]" {
				cell.SetContent(x, y, rune(s[0]), nil, food)
			} else {
				cell.SetContent(x, y, rune(s[0]), nil, tcell.StyleDefault)
			}
		}
	}

	cell.Show()
	// time.Sleep(1 * time.Millisecond)
}
