package main

import (
	"strings"
	"testing"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

const testInput string = `
....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..
`

const testInput2 string = `
.....
..##.
..#..
.....
..##.
.....
`

const testResultPart1 = 110
const testResultPart2 = 20

func TestPart1(t *testing.T) {
	result := Part1(strings.NewReader(utils.CleanInput(testInput)))
	if result != testResultPart1 {
		t.Errorf("expected %v, got %v", testResultPart1, result)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(strings.NewReader(utils.CleanInput(testInput)))
	if result != testResultPart2 {
		t.Errorf("expected %v, got %v", testResultPart2, result)
	}
}
