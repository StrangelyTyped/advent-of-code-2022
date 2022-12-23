package main

import (
	"bufio"
	"fmt"
	"io"
	"math"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

type Position [2]int

func readInput(input io.Reader) map[Position]int {
	scanner := bufio.NewScanner(input)

	elfMap := map[Position]int{}
	y := 0
	elf := 0
	for scanner.Scan() {
		line := scanner.Text()

		for x, v := range line {
			if v == '#' {
				elfMap[Position{x, y}] = elf
				elf++
			}
		}
		y++
	}
	return elfMap
}

var moveDirections = []Position{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

var checkOffsets = [][]Position {
	{{0, -1}, {-1, -1}, {1, -1}},
	{{0, 1}, {-1, 1}, {1, 1}},
	{{-1, 0}, {-1, -1}, {-1, 1}},
	{{1, 0}, {1, -1}, {1, 1}},
}

func Part1(input io.Reader) int {
	elfMap := readInput(input)
	moveDirection := 0
	for i := 0; i < 10; i++ {
		elfMap, moveDirection, _ = simulate(elfMap, moveDirection)
	}

	maxX, minX, maxY, minY := 0, math.MaxInt, 0, math.MaxInt
	for coord := range elfMap {
		maxX = utils.Max(maxX, coord[0])
		minX = utils.Min(minX, coord[0])
		maxY = utils.Max(maxY, coord[1])
		minY = utils.Min(minY, coord[1])
	}

	return (((1+maxX) - minX) * ((1+maxY) - minY)) - len(elfMap)
}

func Part2(input io.Reader) int {
	elfMap := readInput(input)
	moveDirection := 0
	for i := 0; ; i++ {
		var moveCount int
		elfMap, moveDirection, moveCount = simulate(elfMap, moveDirection)

		if moveCount == 0 {
			fmt.Printf("No movement detected, steady state after %d iterations\n", i)
			return i + 1
		}
	}

}

func simulate(elfMap map[Position]int, moveDirection int) (map[Position]int, int, int) {
	movementProposals := make([]int, len(elfMap))
	for idx := range movementProposals {
		movementProposals[idx] = -1
	}

	for coord, elf := range elfMap {
		willMove := false
		for _, x := range checkOffsets {
			for _, y := range x {
				dest := Position{coord[0] + y[0], coord[1] + y[1]}
				_, has := elfMap[dest]
				if has {
					willMove = true;
					break
				}
			}
			if willMove {
				break
			}
		}
		if willMove {
			for j := 0; j < len(moveDirections); j++ {
				checkDirection := (j + moveDirection) % len(moveDirections)
				canMove := true
				for _, k := range checkOffsets[checkDirection] {
					_, has := elfMap[Position{coord[0] + k[0], coord[1] + k[1]}]
					if has {
						canMove = false
						break
					}
				}
				if canMove {
					movementProposals[elf] = checkDirection
					break	
				}
			}
		}
	}

	destClearing := map[Position]int{}
	for coord, elf := range elfMap {
		dir := movementProposals[elf]
		if dir != -1 {
			dest := Position{coord[0] + moveDirections[dir][0], coord[1] + moveDirections[dir][1]}
			otherElf, occupied := destClearing[dest]
			if occupied {
				movementProposals[otherElf] = -1
				movementProposals[elf] = -1
			} else {
				destClearing[dest] = elf
			}
		}
	}

	newElfMap := map[Position]int{}
	movedElves := 0
	for coord, elf := range elfMap {
		dir := movementProposals[elf]
		dest := Position{coord[0], coord[1]}
		if dir != -1 {
			dest[0] += moveDirections[dir][0]
			dest[1] += moveDirections[dir][1]
			movedElves++
		}
		newElfMap[dest] = elf
	}

	moveDirection = (moveDirection + 1) % len(moveDirections)
	return newElfMap, moveDirection, movedElves
}


func main() {
	utils.Run("day23.txt", Part1, Part2)
}
