package main

import (
	"bufio"
	"fmt"
	"os"
)

type Cart struct {
	direction rune
	turns     int
}

func (cart *Cart) turn() Cart {
	cart.turns += 1
	switch cart.turns % 3 {
	case 1:
		switch cart.direction {
		case '<':
			cart.direction = 'v'
		case '>':
			cart.direction = '^'
		case '^':
			cart.direction = '<'
		case 'v':
			cart.direction = '>'
		}
	case 0: // 3
		switch cart.direction {
		case '<':
			cart.direction = '^'
		case '>':
			cart.direction = 'v'
		case '^':
			cart.direction = '>'
		case 'v':
			cart.direction = '<'
		}
	}
	return *cart
}

var tracks [155][155]rune
var carts [155][155]Cart
var carts2 [155][155]Cart

func main() {
	file, err := os.Open("13_1/2018_13.input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for y, row := range tracks {
		for x := range row {
			tracks[y][x] = ' '
			carts[y][x] = Cart{}
			carts2[y][x] = Cart{}
		}
	}

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, r := range line {
			switch r {
			case '-':
				tracks[y][x] = '-'
			case '|':
				tracks[y][x] = '|'
			case '/':
				tracks[y][x] = '/'
			case '\\':
				tracks[y][x] = '\\'
			case '+':
				tracks[y][x] = '+'
			case 'v':
				tracks[y][x] = '|'
				carts[y][x] = Cart{direction: 'v'}
			case '^':
				tracks[y][x] = '|'
				carts[y][x] = Cart{direction: '^'}
			case '<':
				tracks[y][x] = '-'
				carts[y][x] = Cart{direction: '<'}
			case '>':
				tracks[y][x] = '-'
				carts[y][x] = Cart{direction: '>'}
			}
		}
		y++
	}

	// printIt()
	for {
		x, y, collision := moveOne()
		// printIt()
		if collision {
			fmt.Println("done", x, y)
			break
		}
	}
}

func moveOne() (x int, y int, collision bool) {
	for y, row := range tracks {
		for x := range row {
			carts2[y][x] = Cart{}
		}
	}

	for y = range carts {
		for x = range carts[y] {
			collision = false
			cart := carts[y][x]
			switch cart.direction {
			case 0:
				continue
			case 'v':
				carts[y][x] = Cart{}
				carts2[y+1][x], collision = getDirection(x, y+1, cart)
				if collision {
					carts2[y+1][x] = Cart{}
				}
			case '^':
				carts[y][x] = Cart{}
				carts2[y-1][x], collision = getDirection(x, y-1, cart)
				if collision {
					carts2[y-1][x] = Cart{}
				}
			case '<':
				carts[y][x] = Cart{}
				carts2[y][x-1], collision = getDirection(x-1, y, cart)
				if collision {
					carts2[y][x-1] = Cart{}
				}
			case '>':
				carts[y][x] = Cart{}
				carts2[y][x+1], collision = getDirection(x+1, y, cart)
				if collision {
					carts2[y][x+1] = Cart{}
				}
			}
		}
	}
	carts = carts2

	// count them
	count := 0
	finalX, finalY := 0, 0
	for y := range tracks {
		for x := range carts[y] {
			if carts[y][x].direction != 0 {
				count++
				finalX = x
				finalY = y
			}

		}
	}

	if count == 1 {
		return finalX, finalY, true
	}

	return 0, 0, false
}

func getDirection(x int, y int, cart Cart) (Cart, bool) {
	track := tracks[y][x]

	if carts[y][x].direction != 0 || carts2[y][x].direction != 0 {
		return Cart{direction: 'X'}, true
	}
	if track == '|' || track == '-' {
		return cart, false
	}
	if track == '+' {
		return cart.turn(), false
	}

	if track == '/' && cart.direction == 'v' {
		return Cart{direction: '<', turns: cart.turns}, false
	}
	if track == '/' && cart.direction == '>' {
		return Cart{direction: '^', turns: cart.turns}, false
	}
	if track == '/' && cart.direction == '^' {
		return Cart{direction: '>', turns: cart.turns}, false
	}
	if track == '/' && cart.direction == '<' {
		return Cart{direction: 'v', turns: cart.turns}, false
	}

	if track == '\\' && cart.direction == 'v' {
		return Cart{direction: '>', turns: cart.turns}, false
	}
	if track == '\\' && cart.direction == '>' {
		return Cart{direction: 'v', turns: cart.turns}, false
	}
	if track == '\\' && cart.direction == '^' {
		return Cart{direction: '<', turns: cart.turns}, false
	}
	if track == '\\' && cart.direction == '<' {
		return Cart{direction: '^', turns: cart.turns}, false
	}

	panic("should never happen")
}

func printIt() {
	for y := range tracks {
		for x := range carts[y] {
			if carts[y][x].direction != 0 {
				fmt.Print(string(carts[y][x].direction))
			} else if carts2[y][x].direction != 0 {
				fmt.Print(string(carts2[y][x].direction))
			} else {
				fmt.Print(string(tracks[y][x]))
			}
		}
		fmt.Println()
	}
}
