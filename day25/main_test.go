package main

import (
	"strings"
	"testing"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

const testInput string = `
1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122
`

const testResultPart1 = "2=-1=0"

func TestPart1(t *testing.T) {
	result := Part1(strings.NewReader(utils.CleanInput(testInput)))
	if result != testResultPart1 {
		t.Errorf("expected %v, got %v", testResultPart1, result)
	}
}

var snafuToIntInputs = []string{
	"1=-0-2",
	"12111",
	"2=0=",
	"21",
	"2=01",
	"111",
	"20012",
	"112",
	"1=-1=",
	"1-12",
	"12",
	"1=",
	"122",
}

var snafuToIntOutputs  = []int{
	1747,
	906,
	198,
	11,
	201,
	31,
	1257,
	32,
	353,
	107,
	7,
	3,
	37,
}

func TestSnafuToInt(t *testing.T) {
	for idx, snafu := range snafuToIntInputs {
		result := snafuToInt(snafu)
		if result != snafuToIntOutputs[idx] {
			t.Errorf("Input %v, expected %v, got %v", snafu, snafuToIntOutputs[idx], result)
		}
	}
}

func TestIntToSnafu(t *testing.T) {
	for idx, snafu := range snafuToIntOutputs {
		result := intToSnafu(snafu)
		if result != snafuToIntInputs[idx] {
			t.Errorf("Input %v, expected %v, got %v", snafu, snafuToIntInputs[idx], result)
		}
	}
}

