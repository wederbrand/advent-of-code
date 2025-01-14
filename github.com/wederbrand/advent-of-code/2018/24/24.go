package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"golang.org/x/exp/maps"
	"slices"
	"strings"
	"time"
)

type Group struct {
	groupType  string
	name       string
	units      int
	hp         int
	attack     int
	attackType string
	initiative int
	weaknesses []string
	immunities []string
}

func main() {
	start := time.Now()
	inFile := GetFileContents("2018/24/input.txt", "\n")

	armies := getArmies(inFile, 0)
	doIt(armies)

	part1 := 0
	for _, group := range armies {
		part1 += group.units
	}
	fmt.Println("Part 1:", part1, "in", time.Since(start))

	boostMin := 1
	boostMax := 1
	for {
		armies = getArmies(inFile, boostMax)
		winner := doIt(armies)
		if winner == "Immune System" {
			break
		} else {
			boostMin = boostMax
			boostMax *= 2
		}
	}

	for {
		boostMid := (boostMin + boostMax) / 2
		armies = getArmies(inFile, boostMid)
		winner := doIt(armies)
		if winner == "Immune System" {
			boostMax = boostMid
		} else {
			boostMin = boostMid
		}
		if boostMax == boostMin+1 {
			break
		}
	}

	part2 := 0
	for _, group := range armies {
		part2 += group.units
	}
	fmt.Println("Part 2:", part2, "in", time.Since(start))
}

func getArmies(inFile []string, immuneBoost int) map[string]*Group {
	armies := map[string]*Group{}
	currentArmy := ""

	groupId := 0
	for _, line := range inFile {
		groupId++
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "Infection") {
			currentArmy = "Infection"
			groupId = 0
			continue
		}
		if strings.HasPrefix(line, "Immune") {
			currentArmy = "Immune System"
			groupId = 0
			continue
		}
		group := Group{groupType: currentArmy, name: currentArmy + " " + fmt.Sprint(groupId)}
		immuneOrWeak := ""
		if strings.Contains(line, "(") {
			immuneOrWeak = line[strings.Index(line, "(") : strings.Index(line, ")")+1]
			line = strings.Replace(line, immuneOrWeak+" ", "", -1)
		}
		fmt.Sscanf(line, "%d units each with %d hit points with an attack that does %d %s damage at initiative %d", &group.units, &group.hp, &group.attack, &group.attackType, &group.initiative)
		if group.groupType == "Immune System" {
			group.attack += immuneBoost
		}
		if immuneOrWeak != "" {
			imOrWeak := strings.Split(immuneOrWeak[1:len(immuneOrWeak)-1], "; ")
			for _, iw := range imOrWeak {
				if strings.HasPrefix(iw, "weak") {
					group.weaknesses = strings.Split(iw[8:], ", ")
				} else {
					group.immunities = strings.Split(iw[10:], ", ")
				}
			}
		}
		armies[group.name] = &group
	}
	return armies
}

func doIt(armies map[string]*Group) string {
	seenHP := map[int]bool{}
	for {
		// sort all groups by effective power, then by initiative
		all := maps.Keys(armies)
		slices.SortFunc(all, func(a, b string) int {
			groupA := armies[a]
			groupB := armies[b]
			if groupA.units*groupA.attack == groupB.units*groupB.attack {
				return groupB.initiative - groupA.initiative
			} else {
				return groupB.units*groupB.attack - groupA.units*groupA.attack
			}
		})

		totalHP := 0
		for _, group := range armies {
			totalHP += group.units
		}

		if seenHP[totalHP] {
			return ""
		}
		seenHP[totalHP] = true

		// Target selection
		targets := make(map[string]*Group)
		picked := make(map[string]*Group)

		for _, attackerName := range all {
			attacker := *armies[attackerName]
			if attacker.units <= 0 {
				continue
			}
			var target *Group = nil
			for _, enemy := range armies {
				if enemy.groupType == attacker.groupType {
					continue
				}
				if picked[enemy.name] != nil {
					continue
				}
				if target == nil {
					target = enemy
				} else if damage(attacker, enemy) == damage(attacker, target) {
					if enemy.units*enemy.attack == target.units*target.attack {
						if enemy.initiative > target.initiative {
							target = enemy
						}
					} else if enemy.units*enemy.attack > target.units*target.attack {
						target = enemy
					}
				} else if damage(attacker, enemy) > damage(attacker, target) {
					target = enemy
				}
			}
			if target != nil && damage(attacker, target) > 0 {
				targets[attacker.name] = target
				picked[target.name] = target
			}
		}

		if len(picked) == 0 {
			break
		}

		slices.SortFunc(all, func(a, b string) int {
			return armies[b].initiative - armies[a].initiative
		})

		for _, attackerName := range all {
			attacker := armies[attackerName]
			if attacker.units <= 0 {
				continue
			}

			if targets[attacker.name] != nil {
				enemy := targets[attacker.name]
				d := damage(*attacker, enemy)
				deaths := d / enemy.hp
				enemy.units -= deaths
			}
		}

		for _, group := range armies {
			if group.units <= 0 {
				delete(armies, group.name)
			}
		}
	}
	for _, group := range armies {
		return group.groupType
	}

	return ""
}

func damage(attacker Group, enemy *Group) int {
	if slices.Contains(enemy.immunities, attacker.attackType) {
		return 0
	} else if slices.Contains(enemy.weaknesses, attacker.attackType) {
		return attacker.units * attacker.attack * 2
	} else {
		return attacker.units * attacker.attack
	}
}
