package main

import (
	"strings"
	"testing"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

const testInput string = `
R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`

const testResultPart1 = 13
const testResultPart2 = 1

const testInput2 string = `
R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`
const testResult2Part2 = 36

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

func TestPart2_2(t *testing.T) {
	result := Part2(strings.NewReader(utils.CleanInput(testInput2)))
	if result != testResult2Part2 {
		t.Errorf("expected %v, got %v", testResult2Part2, result)
	}
}
