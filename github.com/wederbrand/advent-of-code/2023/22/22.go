package main

import (
	"cmp"
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"slices"
	"strings"
	"time"
)

type Pos3D struct {
	x int
	y int
	z int
}

type Block struct {
	name  int
	cubes []Pos3D

	minX int
	maxX int
	minY int
	maxY int
	minZ int
	maxZ int
}

func (in Block) possiblyFall(m []Block) (int, Block) {
	var hitting Block
	for _, collisionCandidate := range m {
		if collisionCandidate.name == in.name {
			// this is me
			continue
		}

		if collisionCandidate.maxZ > in.minZ {
			// candidate is higher and can never collide
			continue
		}

		// can it collide?
		if in.overlapsXY(collisionCandidate) {
			// can collide, find point before collision
			if hitting.maxZ == 0 || hitting.maxZ < collisionCandidate.maxZ {
				// the new candidate is higher up and must be hit first
				hitting = collisionCandidate
			}
		}
	}

	// hitting.maxZ is one z below our new minZ
	steps := in.minZ - hitting.maxZ - 1
	if steps == 0 {
		return 0, in
	}
	nextBlock := in.copyWithLowerZ(steps)

	return steps, nextBlock
}

func (in Block) overlapsXY(candidate Block) bool {
	switch {
	case in.maxX < candidate.minX:
		return false
	case in.minX > candidate.maxX:
		return false
	case in.maxY < candidate.minY:
		return false
	case in.minY > candidate.maxY:
		return false
	}

	return true
}

func (in Block) copyWithLowerZ(steps int) Block {
	out := Block{
		name:  in.name,
		cubes: make([]Pos3D, 0),
		minX:  in.minX,
		maxX:  in.maxX,
		minY:  in.minY,
		maxY:  in.maxY,
		minZ:  in.minZ - steps,
		maxZ:  in.maxZ - steps,
	}

	for _, cube := range in.cubes {
		out.cubes = append(out.cubes, Pos3D{cube.x, cube.y, cube.z - steps})
	}
	return out
}

func newBlock(index int, s string) Block {
	split := strings.Split(s, "~")
	start := strings.Split(split[0], ",")
	end := strings.Split(split[1], ",")

	startX := util.Atoi(start[0])
	endX := util.Atoi(end[0])
	startY := util.Atoi(start[1])
	endY := util.Atoi(end[1])
	startZ := util.Atoi(start[2])
	endZ := util.Atoi(end[2])

	b := Block{index, make([]Pos3D, 0), startX, endX, startY, endY, startZ, endZ}
	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			for z := startZ; z <= endZ; z++ {
				b.cubes = append(b.cubes, Pos3D{x, y, z})
			}
		}
	}
	return b
}

func main() {
	startTimer := time.Now()
	inFile := util.GetFileContents("2023/22/input.txt", "\n")

	blocks := make([]Block, 0)
	for i, s := range inFile {
		blocks = append(blocks, newBlock(i, s))
	}
	fmt.Println("parsing in", time.Since(startTimer))
	startTimer = time.Now()

	slices.SortFunc(blocks, func(a, b Block) int {
		return cmp.Compare(a.minZ, b.minZ)
	})

	// let all fall until stable
	settle(blocks)

	part1 := 0
	part2 := 0
	for i := range blocks {
		checker := slices.Clone(blocks)
		checker = slices.Delete(checker, i, i+1)

		settle(checker)
		// compare blocks with checker
		// part1 is the number of times they are the same
		// part2 is the number of blocks that moved otherwise
		moved := 0
		for i2, a := range blocks {
			if i == i2 {
				// this is the one we removed; it doesn't count
				continue
			}
			for _, b := range checker {
				if a.name == b.name {
					if a.minZ != b.minZ {
						moved++
					}
				}
			}
		}
		if moved == 0 {
			part1++
		} else {
			part2 += moved
		}
	}

	fmt.Println("part1: ", part1, "in", time.Since(startTimer))
	fmt.Println("part2: ", part2, "in", time.Since(startTimer))
}

func settle(blocks []Block) {
	for {
		totalSteps := 0
		for i, block := range blocks {
			steps, nextBlock := block.possiblyFall(blocks)
			blocks[i] = nextBlock
			totalSteps += steps
		}
		if totalSteps == 0 {
			break
		}
	}
}
