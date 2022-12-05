package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

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

func Part1(input io.Reader) string {
	scanner := bufio.NewScanner(input)
	stacks := [][]string{}
	readingState := true
	for scanner.Scan() {
		line := scanner.Text()
		if readingState {
			if len(strings.TrimSpace(line)) == 0 {
				fmt.Printf("%v\n", stacks)
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
			for i := 0; i < count; i++ {
				token := stacks[source][0]
				stacks[source] = stacks[source][1:]
				stacks[dest] = append([]string{token}, stacks[dest]...)
			}
		}
	}
	str := ""
	for i := 0; i < len(stacks); i++ {
		str = str + stacks[i][0]
	}
	return str
}

func Part2(input io.Reader) string {
	scanner := bufio.NewScanner(input)
	stacks := [][]string{}
	readingState := true
	for scanner.Scan() {
		line := scanner.Text()
		if readingState {
			if len(strings.TrimSpace(line)) == 0 {
				fmt.Printf("%v\n", stacks)
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

			move := stacks[source][0:count]
			stacks[source] = stacks[source][count:]
			stacks[dest] = append(append([]string{}, move...), stacks[dest]...)

		}
	}
	str := ""
	for i := 0; i < len(stacks); i++ {
		str = str + stacks[i][0]
	}
	return str
}

func main() {
	utils.RunStr("day05.txt", Part1, Part2)
}
