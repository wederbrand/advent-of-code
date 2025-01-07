package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"sort"
	"time"
)

type Cart struct {
	c        Coord
	d        Dir
	nextTurn int // 0 is left, 1 is straight, 2 is right
	deleted  bool
}

func main() {
	start := time.Now()
	inFile := GetFileContents("2018/13/input.txt", "\n")

	m := MakeChart(inFile, " ")

	carts := make([]*Cart, 0)
	for c, s := range m {
		if s == ">" {
			carts = append(carts, &Cart{c: c, d: RIGHT})
			m[c] = "-"
		} else if s == "<" {
			carts = append(carts, &Cart{c: c, d: LEFT})
			m[c] = "-"
		} else if s == "^" {
			carts = append(carts, &Cart{c: c, d: UP})
			m[c] = "|"
		} else if s == "v" {
			carts = append(carts, &Cart{c: c, d: DOWN})
			m[c] = "|"
		}
	}

	for {
		// Remove deleted carts
		newCarts := make([]*Cart, 0)
		for _, cart := range carts {
			if !cart.deleted {
				newCarts = append(newCarts, cart)
			}
		}
		carts = newCarts

		if len(carts) == 1 {
			fmt.Println("Part 2:", carts[0].c, "in", time.Since(start))
			break
		}

		sort.Slice(carts, func(i, j int) bool {
			if carts[i].c.Y != carts[j].c.Y {
				return carts[i].c.Y < carts[j].c.Y
			}
			return carts[i].c.X < carts[j].c.X
		})

		for _, cart := range carts {
			next := cart.c.Move(cart.d)
			nextTrack := m[next]
			if nextTrack == "+" {
				if cart.nextTurn == 0 {
					cart.d = cart.d.Left()
				} else if cart.nextTurn == 2 {
					cart.d = cart.d.Right()
				}
				cart.nextTurn = (cart.nextTurn + 1) % 3
			} else if nextTrack == "\\" {
				if cart.d == UP {
					cart.d = LEFT
				} else if cart.d == LEFT {
					cart.d = UP
				} else if cart.d == DOWN {
					cart.d = RIGHT
				} else if cart.d == RIGHT {
					cart.d = DOWN
				}
			} else if nextTrack == "/" {
				if cart.d == UP {
					cart.d = RIGHT
				} else if cart.d == RIGHT {
					cart.d = UP
				} else if cart.d == DOWN {
					cart.d = LEFT
				} else if cart.d == LEFT {
					cart.d = DOWN
				}
			}

			cart.c = next

			for _, otherCart := range carts {
				if cart == otherCart {
					continue
				}
				if otherCart.c == cart.c && !otherCart.deleted && !cart.deleted {
					PrintOnce("Part 1:", cart.c, "in", time.Since(start))
					otherCart.deleted = true
					cart.deleted = true
				}
			}
		}
	}
}
