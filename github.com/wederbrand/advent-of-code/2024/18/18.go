package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/priorityqueue"
	"log"
	"slices"
	"sync"
	"time"
)

var animation = true

func main() {
	start := time.Now()
	inFile := GetFileContents("2024/18/input.txt", "\n")

	m := Chart{}

	c := Coord{0, 0}
	exit := Coord{70, 70}
	minC := Coord{0, 0}
	maxC := Coord{70, 70}

	tCell, err := tcell.NewScreen()
	if err != nil {
		animation = false
	} else {
		if err := tCell.Init(); err != nil {
			log.Fatalf("%+v", err)
		}
		// Set default text style
		defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorMediumSpringGreen)
		tCell.SetStyle(defStyle)

		// Clear screen
		tCell.Clear()

		tCellMap(tCell, m, minC, maxC)
	}

	var currentChart Chart
	for i, s := range inFile {
		if i == 1024 {
			path := walkIt(c, exit, minC, maxC, m)

			if !animation {
				fmt.Println("Part 1:", len(path.steps), "in", time.Since(start))
			}

			currentChart = Chart{}
			for _, step := range path.steps {
				currentChart[step] = "O"
				if animation {
					tCellPath(tCell, step, false)
				}
			}
			if animation {
				tCell.Show()
			}
		}
		var x, y int
		fmt.Sscanf(s, "%d,%d", &x, &y)
		nextByte := Coord{x, y}
		m[nextByte] = "#"

		if animation {
			tCellByte(tCell, nextByte, false)
		}

		if i > 1024 {
			if currentChart[nextByte] == "O" {
				// This is where we used to walk, take a new route
				path := walkIt(c, exit, minC, maxC, m)

				if len(path.steps) == 0 {
					if !animation {
						fmt.Println("Part 2:", s, "in", time.Since(start))
					} else {
						tCellByte(tCell, nextByte, true)
						wg := sync.WaitGroup{}
						wg.Add(1)
						go func() {
							time.Sleep(5 * time.Second)
							wg.Done()
						}()
						wg.Wait()
					}

					break
				}

				if animation {
					for coord := range currentChart {
						tCellPath(tCell, coord, true)
					}
				}
				currentChart = Chart{}
				for _, step := range path.steps {
					currentChart[step] = "O"
				}
				if animation {
					for coord := range currentChart {
						tCellPath(tCell, coord, false)
					}
					tCell.Show()
				}

			}
		}
	}
}

type Path struct {
	steps []Coord
}

func walkIt(start Coord, exit Coord, minC Coord, maxC Coord, m Chart) Path {
	q := priorityqueue.NewQueue()
	q.Add(&priorityqueue.State{Data: Path{[]Coord{start}}, Priority: 0})

	seen := make(map[Coord]bool)
	for q.HasNext() {
		s := q.Next()
		p := s.Data.(Path)
		c := p.steps[len(p.steps)-1]

		if c == exit {
			return p
		}

		if seen[c] {
			continue
		}
		seen[c] = true

		for _, dir := range ALL {
			next := c.Move(dir)
			if m[next] == "#" || next.X < minC.X || next.Y < minC.Y || next.X > maxC.X || next.Y > maxC.Y {
				continue
			}
			newPath := Path{slices.Clone(p.steps)}
			newPath.steps = append(newPath.steps, next)
			q.Add(&priorityqueue.State{Data: newPath, Priority: s.Priority + 1})
		}
	}

	return Path{[]Coord{}}
}

func tCellMap(cell tcell.Screen, m Chart, minC Coord, maxC Coord) {
	for y := minC.Y; y <= maxC.Y; y++ {
		for x := minC.X; x <= maxC.X; x++ {
			s, found := m[Coord{x, y}]
			if !found {
				s = "."
			}
			cell.SetContent(x, y, rune(s[0]), nil, tcell.StyleDefault)
		}
	}

	cell.Show()
	//time.Sleep(1 * time.Millisecond)
}

func tCellByte(cell tcell.Screen, byte Coord, final bool) {
	byteStyle := tcell.StyleDefault
	if final {
		byteStyle = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorDarkGoldenrod).Blink(true)
	}
	cell.SetContent(byte.X, byte.Y, '#', nil, byteStyle)

	cell.Show()
	time.Sleep(3 * time.Millisecond)
}

func tCellPath(cell tcell.Screen, path Coord, old bool) {
	pathStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorOrangeRed)
	if old {
		pathStyle = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorPink)
	}
	cell.SetContent(path.X, path.Y, '0', nil, pathStyle)
}
