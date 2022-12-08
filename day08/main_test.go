package main

import (
	"strings"
	"testing"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

const testInput string = `
30373
25512
65332
33549
35390
`

const testResultPart1 = 21
const testResultPart2 = 8

func TestPart1(t *testing.T) {
	result := Part1(strings.NewReader(utils.CleanInput(testInput)))
	if result != testResultPart1 {
		t.Errorf("expected %v, got %v", testResultPart1, result)
	}
}

func TestPart2(t *testing.T) {
	//result := Part2(strings.NewReader(utils.CleanInput(testInput)))
	result := Part2(utils.OpenOrPanic("/home/alexis/advent-of-code-2022/inputs/day08.txt"))
	if result != testResultPart2 {
		t.Errorf("expected %v, got %v", testResultPart2, result)
	}
}
