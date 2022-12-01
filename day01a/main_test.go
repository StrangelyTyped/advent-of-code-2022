package main

import (
	"reflect"
	"strings"
	"testing"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

const testInput string = `
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

var testParseResult = []int{
	6000,
	4000,
	11000,
	24000,
	10000,
}

const testResultPart1 = 24000
const testResultPart2 = 45000

func TestParsing(t *testing.T) {
	ch := parseElves(strings.NewReader(utils.CleanInput(testInput)))
	result := []int{}
	for x := range ch {
		result = append(result, x)
	}
	if !reflect.DeepEqual(result, testParseResult) {
		t.Errorf("expected %v, got %v", testParseResult, result)
	}
}

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
