package main

import (
	"bufio"
	"io"
	"regexp"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

var lineRe = regexp.MustCompile(`([\d]+)-([\d]+),([\d]+)-([\d]+)`)

func solve(input io.Reader, impl func(line []int) int) int {
	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		matches := utils.MapToInt(lineRe.FindStringSubmatch(line)[1:])
		sum += impl(matches)
	}
	return sum
}

func Part1(input io.Reader) int {
	return solve(input, func(line []int) int {
		if (line[0] <= line[2] && line[1] >= line[3]) ||
			(line[2] <= line[0] && line[3] >= line[1]) {
			return 1
		}
		return 0
	})
}

func Part2(input io.Reader) int {
	return solve(input, func(line []int) int {
		if (line[0] >= line[2] && line[0] <= line[3]) ||
			(line[1] >= line[2] && line[1] <= line[3]) ||
			(line[2] >= line[0] && line[2] <= line[1]) ||
			(line[3] >= line[0] && line[3] <= line[1]) {
			return 1
		}
		return 0
	})
}

func main() {
	utils.Run("day04.txt", Part1, Part2)
}
