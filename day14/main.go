package main

import (
	"bufio"
	"io"
	"strings"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

type Position [2]int

func readInput(input io.Reader) (map[Position]bool, int) {
	cave := map[Position]bool{}
	bottomOfCave := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, " -> ")
		start := parsePosition(coords[0])
		bottomOfCave = utils.Max(bottomOfCave, start[1])
		for _, pos := range(coords[1:]) {
			end := parsePosition(pos)

			for y := utils.Min(start[1],end[1]); y <= utils.Max(start[1],end[1]); y ++ {
				for x := utils.Min(start[0],end[0]); x <= utils.Max(start[0],end[0]); x ++ {
					cave[Position{x, y}] = true
				}
			}
			bottomOfCave = utils.Max(bottomOfCave, end[1])
			start = end
		}
	}
	return cave, bottomOfCave
}

func parsePosition(coord string) Position {
	comma := strings.Index(coord, ",")
	return Position{utils.AtoiOrPanic(coord[0:comma]), utils.AtoiOrPanic(coord[comma+1:])}
}

func isOccupied(cave map[Position]bool, floor int, hardFloor bool, pos Position) bool {
	if cave[pos] {
		return true
	}
	if hardFloor && pos[1] == floor {
		return true
	}
	return false
}

// return value is whether this unit settled or not
func simulate(cave map[Position]bool, floor int, hardFloor bool) bool {
	sandOrigin := Position{500, 0}

	if isOccupied(cave, floor, hardFloor, sandOrigin) {
		return false
	}

	unitLoc := sandOrigin

	for {
		newUnitLoc := Position{unitLoc[0],unitLoc[1] + 1}		

		if !hardFloor && newUnitLoc[1] > floor {
			return false
		}

		if isOccupied(cave, floor, hardFloor, newUnitLoc) {
			newUnitLoc[0] -= 1
			if isOccupied(cave, floor, hardFloor, newUnitLoc) {
				newUnitLoc[0] += 2
			}
		}

		if isOccupied(cave, floor, hardFloor, newUnitLoc) {
			cave[unitLoc] = true
			return true
		}
		unitLoc = newUnitLoc
	}
}

func Part1(input io.Reader) int {

	cave, bottomOfCave := readInput(input)

	unitsAtRest := 0
	for simulate(cave, bottomOfCave, false) {
		unitsAtRest++
	}
	return unitsAtRest

}

func Part2(input io.Reader) int {
	cave, bottomOfCave := readInput(input)

	unitsAtRest := 0
	for simulate(cave, bottomOfCave+2, true) {
		unitsAtRest++
	}
	return unitsAtRest}



func main() {
	utils.Run("day14.txt", Part1, Part2)
}
