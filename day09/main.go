package main

import (
	"bufio"
	"io"
	"strings"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

var dirMap = map[string][2]int{
	"U": {0, -1},
	"D": {0, 1},
	"L": {-1, 0},
	"R": {1, 0},
}

func solve(input io.Reader, tails int) int {

	scanner := bufio.NewScanner(input)

	hPos := [2]int{0, 0}
	tailSet := [][2]int{}
	for i := 0; i < tails; i++ {
		tailSet = append(tailSet, [2]int{0, 0})
	}

	tVisited := make(map[[2]int]bool)
	trackedTail := &tailSet[tails - 1]
	tVisited[*trackedTail] = true
	tVisitCount := 1

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		dir := tokens[0]
		distance := utils.AtoiOrPanic(tokens[1])

		dirVector := dirMap[dir]
		for i := 0; i < distance; i++ {
			hPos[0] += dirVector[0]
			hPos[1] += dirVector[1]

			pPos := &hPos

			for j := 0; j < tails; j++ {
				tPos := &tailSet[j]
				delta := [2]int{tPos[0] - pPos[0], tPos[1] - pPos[1]}

				if utils.Abs(delta[0]) > 1 || utils.Abs(delta[1]) > 1 {
					tPos[0] -= utils.Sign(delta[0])
					tPos[1] -= utils.Sign(delta[1])
				}
				pPos = tPos
			}

			_, visited := tVisited[*trackedTail]
			if !visited {
				tVisited[*trackedTail] = true
				tVisitCount ++
			}
		}
	}
	return tVisitCount
}

func Part1(input io.Reader) int {
	return solve(input, 1)
}

func Part2(input io.Reader) int {
	return solve(input, 9)
}

func main() {
	utils.Run("day09.txt", Part1, Part2)
}
