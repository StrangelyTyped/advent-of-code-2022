package day01

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

func Part1(input string) int {
	maxElf := 0
	for _, elf := range parseElves(input) {
		if elf > maxElf {
			maxElf = elf
		}
	}
	return maxElf
}

func Part2(input string) int {
	elves := parseElves(input)
	sort.Ints(elves)
	count := len(elves)
	return elves[count-1] + elves[count-2] + elves[count-3]
}

func Main(instream io.Reader) (int, int) {
	input := utils.ReadInput(instream)
	return Part1(input), Part2(input)
}
