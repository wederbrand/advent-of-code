package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"math"
	"strings"
	"time"
)

type Ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func main() {
	start := time.Now()
	inFile := GetFileContents("2015/15/input.txt", "\n")

	ingredients := make([]Ingredient, 0)
	for _, s := range inFile {
		split := strings.Split(s, ":")
		i := Ingredient{name: split[0]}
		fmt.Sscanf(split[1], " capacity %d, durability %d, flavor %d, texture %d, calories %d", &i.capacity, &i.durability, &i.flavor, &i.texture, &i.calories)
		ingredients = append(ingredients, i)
	}

	part1 := findMax(ingredients, 100, Ingredient{}, false)
	fmt.Println("Part 1: ", part1, "in", time.Since(start))

	part2 := findMax(ingredients, 100, Ingredient{}, true)
	fmt.Println("Part 2: ", part2, "in", time.Since(start))
}

func findMax(ingredients []Ingredient, rem int, total Ingredient, calorieLimit bool) int {
	if len(ingredients) == 1 {
		total = Ingredient{
			capacity:   total.capacity + ingredients[0].capacity*rem,
			durability: total.durability + ingredients[0].durability*rem,
			flavor:     total.flavor + ingredients[0].flavor*rem,
			texture:    total.texture + ingredients[0].texture*rem,
			calories:   total.calories + ingredients[0].calories*rem,
		}
		if total.capacity < 0 || total.durability < 0 || total.flavor < 0 || total.texture < 0 {
			return 0
		}
		if calorieLimit && total.calories != 500 {
			return 0
		}
		return total.capacity * total.durability * total.flavor * total.texture
	}

	best := math.MinInt
	for i := 0; i <= rem; i++ {
		newTotal := Ingredient{
			capacity:   total.capacity + ingredients[0].capacity*i,
			durability: total.durability + ingredients[0].durability*i,
			flavor:     total.flavor + ingredients[0].flavor*i,
			texture:    total.texture + ingredients[0].texture*i,
			calories:   total.calories + ingredients[0].calories*i,
		}

		val := findMax(ingredients[1:], rem-i, newTotal, calorieLimit)
		if val > best {
			best = val
		}
	}

	return best
}
