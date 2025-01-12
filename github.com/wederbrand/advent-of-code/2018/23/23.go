package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"math"
	"time"
)

type Bot struct {
	x, y, z, r int
}

func main() {
	start := time.Now()
	inFile := GetFileContents("2018/23/input.txt", "\n")

	bots := make([]Bot, 0)
	strongest := Bot{}
	minX, minY, minZ := math.MaxInt, math.MaxInt, math.MaxInt
	maxX, maxY, maxZ := math.MinInt, math.MinInt, math.MinInt
	for _, line := range inFile {
		var bot Bot
		fmt.Sscanf(line, "pos=<%d,%d,%d>, r=%d", &bot.x, &bot.y, &bot.z, &bot.r)
		bots = append(bots, bot)
		if bot.r > strongest.r {
			strongest = bot
		}
		if bot.x < minX {
			minX = bot.x
		}
		if bot.y < minY {
			minY = bot.y
		}
		if bot.z < minZ {
			minZ = bot.z
		}
		if bot.x > maxX {
			maxX = bot.x
		}
		if bot.y > maxY {
			maxY = bot.y
		}
		if bot.z > maxZ {
			maxZ = bot.z
		}
	}

	part1 := inRange(bots, strongest, false)

	fmt.Println("Part 1:", part1, "in", time.Since(start))

	position := checkAll(bots, minX, minY, minZ, maxX, maxY, maxZ)

	fmt.Println("Part 2:", manhattan3d(position, Bot{0, 0, 0, 0}), "in", time.Since(start))
}

func checkAll(bots []Bot, minX int, minY int, minZ int, maxX int, maxY int, maxZ int) Bot {
	xInc := 1
	if maxX-minX > 10 {
		xInc = (maxX - minX) / 10
	}
	yInc := 1
	if maxY-minY > 10 {
		yInc = (maxY - minY) / 10
	}
	zInc := 1
	if maxZ-minZ > 10 {
		zInc = (maxZ - minZ) / 10
	}

	maxCount := 0
	bestBot := Bot{}
	for x := minX; x < maxX; x += xInc {
		for y := minY; y < maxY; y += yInc {
			for z := minZ; z < maxZ; z += zInc {
				bot := Bot{x, y, z, 0}
				count := inRange(bots, bot, true)
				if count > maxCount {
					maxCount = count
					bestBot = bot
				} else if count == maxCount && manhattan3d(bot, Bot{0, 0, 0, 0}) < manhattan3d(bestBot, Bot{0, 0, 0, 0}) {
					bestBot = bot
				}
			}
		}
	}

	if xInc == 1 && yInc == 1 && zInc == 1 {
		return bestBot
	}

	return checkAll(bots, bestBot.x-xInc, bestBot.y-yInc, bestBot.z-zInc, bestBot.x+xInc, bestBot.y+yInc, bestBot.z+zInc)
}

func inRange(bots []Bot, target Bot, reverse bool) int {
	result := 0
	for _, bot := range bots {
		if !reverse {
			if manhattan3d(bot, target) <= target.r {
				result++
			}
		} else {
			if manhattan3d(bot, target) <= bot.r {
				result++
			}
		}
	}
	return result
}

func manhattan3d(a Bot, b Bot) int {
	return IntAbs(a.x-b.x) + IntAbs(a.y-b.y) + IntAbs(a.z-b.z)
}
