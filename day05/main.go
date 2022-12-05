package main

import (
	"bufio"
	"io"
	"math"
	"strings"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func addState(state *[][]string, line string) *[][]string {
	for len(*state) < ((len(line) + 1) / 4) {
		newState := append(*state, []string{})
		state = &newState
	}
	col := 0
	for i := 0; i < (len(line) + 1); i += 4 {
		segment := line[i:i+3]
		idx := strings.Index(segment, "[")
		if idx != -1 {
			(*state)[col] = append((*state)[col], segment[idx + 1:idx + 2])
		}
		col++
	}
	return state
}

func solve(input io.Reader, maxStack int) string {
	scanner := bufio.NewScanner(input)
	stacks := [][]string{}
	readingState := true
	for scanner.Scan() {
		line := scanner.Text()
		if readingState {
			if len(strings.TrimSpace(line)) == 0 {
				readingState = false
				continue
			} else {
				stacks = *addState(&stacks, line)
			}
		} else {
			tokens := strings.Split(line, " ")
			count := utils.AtoiOrPanic(tokens[1])
			source := utils.AtoiOrPanic(tokens[3]) - 1
			dest := utils.AtoiOrPanic(tokens[5]) - 1
			for i := 0; i < count; i+= maxStack {
				chunk := Min(maxStack, count)
				move := stacks[source][0:chunk]
				stacks[source] = stacks[source][chunk:]
				stacks[dest] = append(append([]string{}, move...), stacks[dest]...)
			}
		}
	}
	str := ""
	for i := 0; i < len(stacks); i++ {
		str = str + stacks[i][0]
	}
	return str
}

func Part1(input io.Reader) string {
	return solve(input, 1)
}

func Part2(input io.Reader) string {
	return solve(input, math.MaxInt32)
}

func main() {
	utils.RunStr("day05.txt", Part1, Part2)
}
