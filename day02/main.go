package main

import (
	"bufio"
	"io"
	"strings"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

var moveMap = map[string]int {
	"A": 1,
	"X": 1,
	"B": 2,
	"Y": 2,
	"C": 3,
	"Z": 3,
}

func solve(input io.Reader, scorer func(int, int)int) int {
	scanner := bufio.NewScanner(input)
	score := 0

	for scanner.Scan() {
		line := scanner.Text()

		moves := strings.Split(line, " ")
		score += scorer(moveMap[moves[0]], moveMap[moves[1]])
	}
	return score
}

func Part1(input io.Reader) int {
	return solve(input, func(theirMove, myMove int) int {
		score := 0
		if theirMove == myMove {
			score = 3
		} else if (myMove - theirMove == 1 || myMove - theirMove == -2) {
			score = 6
		}
		return myMove + score
	})
}

func Part2(input io.Reader) int {
	return solve(input, func(theirMove, result int) int {
		myMove := 0
		result = (result - 1) * 3

		if result == 0 {
			myMove = ((theirMove + 1) % 3) + 1
		}else if result == 3{
			myMove = theirMove
		} else {
			myMove = (theirMove % 3) + 1
		}
		return myMove + result
	})
}

func main() {
	utils.Run("day02.txt", Part1, Part2)
}
