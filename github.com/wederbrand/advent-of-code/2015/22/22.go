package main

import (
	"fmt"
	"math"
	"time"
)

type fighter struct {
	hp, damage, mana int
}

type spell struct {
	cost, damage, heal, turns, shield, poison, recharge int
}

var spells = []spell{
	{cost: 53, damage: 4},                // magic missile
	{cost: 73, damage: 2, heal: 2},       // drain
	{cost: 113, turns: 6, shield: 7},     // shield
	{cost: 173, turns: 6, poison: 3},     // poison
	{cost: 229, turns: 5, recharge: 101}, // recharge
}

func main() {
	start := time.Now()

	player := fighter{hp: 50, mana: 500}
	boss := fighter{hp: 55, damage: 8}

	_, part1 := doIt(true, player, boss, 0, 0, 0, false)
	fmt.Println("Part 1: ", part1, "in", time.Since(start))

	_, part2 := doIt(true, player, boss, 0, 0, 0, true)
	fmt.Println("Part 2: ", part2, "in", time.Since(start))
}

type cacheKey struct {
	playerTurn    bool
	player        fighter
	boss          fighter
	shieldTurns   int
	poisonTurns   int
	rechargeTurns int
	hardMode      bool
}

var cache = make(map[cacheKey]int)

func doIt(playerTurn bool, player fighter, boss fighter, shieldTurns int, poisonTurns int, rechargeTurns int, hardMode bool) (bool, int) {
	key := cacheKey{playerTurn: playerTurn, player: player, boss: boss, shieldTurns: shieldTurns, poisonTurns: poisonTurns, rechargeTurns: rechargeTurns, hardMode: hardMode}
	if val, ok := cache[key]; ok {
		return val >= 0, val
	}

	if hardMode && playerTurn {
		player.hp--
		if player.hp <= 0 {
			cache[key] = -1
			return false, -1
		}
	}

	if rechargeTurns > 0 {
		player.mana += 101
		rechargeTurns--
	}

	if poisonTurns > 0 {
		boss.hp -= 3
		poisonTurns--
		if boss.hp <= 0 {
			cache[key] = 0
			return true, 0
		}
	}

	shieldActive := false
	if shieldTurns > 0 {
		shieldTurns--
		if shieldTurns == 0 {
			shieldActive = false
		} else {
			shieldActive = true
		}
	}

	cheapest := math.MaxInt
	if playerTurn {
		// player
		if player.mana < 53 {
			cache[key] = -1
			return false, -1
		}

		for _, s := range spells {
			if s.cost > player.mana {
				continue
			}

			if s.shield > 0 && shieldTurns > 0 {
				continue
			}

			if s.poison > 0 && poisonTurns > 0 {
				continue
			}

			if s.recharge > 0 && rechargeTurns > 0 {
				continue
			}

			newPlayer := player
			newBoss := boss

			newPlayer.mana -= s.cost
			newPlayer.hp += s.heal
			newBoss.hp -= s.damage
			if newBoss.hp <= 0 {
				cache[key] = s.cost
				return true, s.cost
			}

			var win bool
			var price int
			if s.recharge > 0 {
				win, price = doIt(false, newPlayer, newBoss, shieldTurns, poisonTurns, s.turns, hardMode)
			} else if s.poison > 0 {
				win, price = doIt(false, newPlayer, newBoss, shieldTurns, s.turns, rechargeTurns, hardMode)
			} else if s.shield > 0 {
				win, price = doIt(false, newPlayer, newBoss, s.turns, poisonTurns, rechargeTurns, hardMode)
			} else {
				win, price = doIt(false, newPlayer, newBoss, shieldTurns, poisonTurns, rechargeTurns, hardMode)
			}

			if win && price+s.cost < cheapest {
				cheapest = price + s.cost
			}
		}
	} else {
		// boss
		newPlayer := player

		damage := boss.damage
		if shieldActive {
			damage -= 7
		}
		if damage < 1 {
			damage = 1
		}
		newPlayer.hp -= damage
		if newPlayer.hp <= 0 {
			cache[key] = -1
			return false, -1
		}

		win, price := doIt(true, newPlayer, boss, shieldTurns, poisonTurns, rechargeTurns, hardMode)

		if win && price < cheapest {
			cheapest = price
		}
	}

	if cheapest == math.MaxInt {
		cache[key] = -1
		return false, 1
	} else {
		cache[key] = cheapest
		return true, cheapest
	}
}
