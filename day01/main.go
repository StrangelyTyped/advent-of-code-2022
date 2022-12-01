package main

import (
	"io"
	"sort"
	"strings"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

func parseElves(input string) []int {
	sum := 0
	result := []int{}

	for _, line := range strings.Split(input, "\n") {
		if len(strings.TrimSpace(line)) == 0 {
			result = append(result, sum)
			sum = 0
		} else {
			sum += utils.AtoiOrPanic(line)
		}
	}
	if sum != 0 {
		result = append(result, sum)
	}

	return result
}

func Part1(input io.Reader) int {
	maxElf := 0
	for _, elf := range parseElves(utils.ReadInput(input)) {
		if elf > maxElf {
			maxElf = elf
		}
	}
	return maxElf
}

func Part2(input io.Reader) int {
	elves := parseElves(utils.ReadInput(input))
	sort.Ints(elves)
	count := len(elves)
	return elves[count-1] + elves[count-2] + elves[count-3]
}

func main() {
	utils.Run("day01.txt", Part1, Part2)
}
