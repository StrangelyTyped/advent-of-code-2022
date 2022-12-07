package main

import (
	"strings"
	"testing"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

type TestCase struct {
	Input string
	Part1 int
	Part2 int
}


var tests = []TestCase {
	{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7, 19},
	{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5, 23},
	{"nppdvjthqldpwncqszvftbrmjlhg", 6, 23},
	{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10, 29},
	{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11, 26},
}


func TestPart1(t *testing.T) {
	for idx, test := range tests {
		result := Part1(strings.NewReader(utils.CleanInput(test.Input)))
		if result != test.Part1 {
			t.Errorf("expected %v, got %v for test case %d", test.Part1, result, idx)
		}
	}
}

func TestPart2(t *testing.T) {
	for idx, test := range tests {
		result := Part2(strings.NewReader(utils.CleanInput(test.Input)))
		if result != test.Part2 {
			t.Errorf("expected %v, got %v for test case %d", test.Part2, result, idx)
		}
	}
}
