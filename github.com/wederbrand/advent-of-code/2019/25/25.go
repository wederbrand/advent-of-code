package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/2019/computer"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2019/25/input.txt", "\n")

	input := []string{
		"east", // sick bay
		"take whirled peas",
		"east", // gift wrapping center
		"east", // observatory
		// "take escape pod",  // don't take yet
		"east",             // stables
		"east",             // kitchen
		"take dark matter", // kitchen
		"west",             // stables,
		"west",             // observatory,
		// "north",            // Corridor, dead end
		// "take infinite loop", // don't take
		"west",  // gift wrapping center
		"north", // Hallway, dead end
		"take prime number",
		"south", // gift wrapping center
		"west",  // sick bay
		"north", // passage
		"take coin",
		"west", // science lab
		// "take molten lava", // don't take
		"north", // crew quarters
		// "north", // warp drive maintenance, dead end
		// "take photons", // don't take
		"west", // arcade
		"take astrolabe",
		"east",  // crew quarters
		"south", // science lab
		"south", // storage
		"take antenna",
		"north", // science lab
		"east",  // passage
		"south", // sick bay
		"west",  // hull breach
		"north", // holodeck
		"take fixed point",
		"north", // navigation
		"take weather machine",
		"east", // security checkpoint
		"inv",
		"drop dark matter",
		"drop coin",
		"drop whirled peas",
		"drop fixed point",
		"drop astrolabe",
		"drop prime number",
		"drop antenna",
		"drop weather machine",
	}

	all := []string{"dark matter", "coin", "whirled peas", "fixed point", "prime number", "weather machine", "antenna", "astrolabe"}
	combinations := Combinations(all)
	for _, comb := range combinations {
		for _, item := range comb {
			input = append(input, "take "+item)
		}
		input = append(input, "inv")
		input = append(input, "south")
		for _, item := range comb {
			input = append(input, "drop "+item)
		}
	}

	in := func() int {
		if len(input) == 0 {
			panic("no input")
		} else {
			row := input[0]
			if len(row) == 0 {
				input = input[1:]
				return 10
			}
			char := row[0]
			input[0] = row[1:]
			return int(char)
		}

		panic("no input")
	}

	out := func(i int) {
		fmt.Print(string(i))
	}

	computer := NewComputer(inFile, in, out)
	computer.Run()

	fmt.Println("part1: ", "in", time.Since(start))
}
