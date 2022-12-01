package day01

import (
	"reflect"
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

var testOutputPart1 = []int{
	6000,
	4000,
	11000,
	24000,
	10000,
}



func TestPart1(t *testing.T) {
	result := Part1(utils.CleanInput(testInput))
	if !reflect.DeepEqual(result, testOutputPart1) {
		t.Errorf("Part 1, expected %v, got %v", testOutputPart1, result)
	}
}