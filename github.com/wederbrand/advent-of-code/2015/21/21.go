package main

import (
	"fmt"
	"math"
	"time"
)

type fighter struct {
	hp, damage, armor int
}

type gear struct {
	cost, damage, armor int
}

func main() {
	start := time.Now()

	weapons := []gear{
		{8, 4, 0},
		{10, 5, 0},
		{25, 6, 0},
		{40, 7, 0},
		{74, 8, 0},
	}

	armours := []gear{
		{0, 0, 0}, // one faked free armor
		{13, 0, 1},
		{31, 0, 2},
		{53, 0, 3},
		{75, 0, 4},
		{102, 0, 5},
	}

	rings := []gear{
		{0, 0, 0}, // one faked free ring
		{0, 0, 0}, // another faked free ring
		{25, 1, 0},
		{25, 1, 0},
		{50, 2, 0},
		{100, 3, 0},
		{20, 0, 1},
		{40, 0, 2},
		{80, 0, 3},
	}

	cheapest := math.MaxInt
	expenisvest := math.MinInt
	for _, weapon := range weapons {
		for _, armour := range armours {
			for _, ring1 := range rings {
				for _, ring2 := range rings {
					if ring1 == ring2 {
						continue
					}
					totalCost := weapon.cost + armour.cost + ring1.cost + ring2.cost
					player := fighter{100, weapon.damage + armour.damage + ring1.damage + ring2.damage, armour.armor + ring1.armor + ring2.armor}
					boss := fighter{104, 8, 1}

					if doIt(player, boss) {
						if totalCost < cheapest {
							cheapest = totalCost
						}
					} else {
						if totalCost > expenisvest {
							expenisvest = totalCost
						}
					}
				}
			}
		}
	}

	fmt.Println("Part 1: ", cheapest, "in", time.Since(start))
	fmt.Println("Part 2: ", expenisvest, "in", time.Since(start))
}

func doIt(player fighter, boss fighter) bool {
	for {
		damage := player.damage - boss.armor
		if damage < 1 {
			damage = 1
		}
		boss.hp -= damage

		if boss.hp <= 0 {
			return true
		}

		damage = boss.damage - player.armor
		if damage < 1 {
			damage = 1
		}
		player.hp -= damage
		if player.hp <= 0 {
			return false
		}
	}
}
