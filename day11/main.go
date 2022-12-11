package main

import (
	"bufio"
	"io"
	"sort"
	"strings"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

type Monkey struct {
	Items []int
	Modulus int
	ZeroModuloMonkey int
	NonzeroModuloMonkey int
	Modifier func(int) int
}

func grabTrailingInt(line string) int {
	return utils.AtoiOrPanic(line[strings.LastIndex(line, " ") + 1:])
}

func solve(input io.Reader, rounds, divisor int) int {
	scanner := bufio.NewScanner(input)

	monkeys := []*Monkey{}

	currentMonkey := &Monkey{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			monkeys = append(monkeys, currentMonkey)
			currentMonkey = &Monkey{}
		} else {
			if strings.HasPrefix(line, "Starting") {
				items := strings.Split(line[strings.Index(line, ":") + 2:], ", ")
				for _, x := range items {
					currentMonkey.Items = append(currentMonkey.Items, utils.AtoiOrPanic(x))
				}
			} else if strings.HasPrefix(line, "Test") {
				currentMonkey.Modulus = grabTrailingInt(line)
			} else if strings.HasPrefix(line, "If true") {
				currentMonkey.ZeroModuloMonkey = grabTrailingInt(line)
			} else if strings.HasPrefix(line, "If false") {
				currentMonkey.NonzeroModuloMonkey = grabTrailingInt(line)
			} else if strings.HasPrefix(line, "Operation") {
				tokens := strings.Split(line, " ")
				tokens = tokens[len(tokens) - 3:]
				arg2 := 0
				if tokens[2] != "old" {
					arg2 = utils.AtoiOrPanic(tokens[2])
				}
				currentMonkey.Modifier = func(i int) int {
					rhs := arg2
					if tokens[2] == "old" {
						rhs = i
					}
					if tokens[1] == "+" {
						return i + rhs
					}
					if tokens[1] == "*" {
						return i * rhs
					}
					panic("Shouldn't get here")
				}
			}
		}
	}
	monkeys = append(monkeys, currentMonkey)

	inspectCounts := make([]int, len(monkeys))

	globalModulus := 1
	for _, m := range monkeys {
		globalModulus *= m.Modulus
	}

	for i := 0; i < rounds; i++ {
		for mIdx, m := range monkeys {
			// Making the somewhat executive assumption a monkey won't throw to itself
			inspectCounts[mIdx] += len(m.Items)
			for len(m.Items) > 0 {
				item := (m.Modifier(m.Items[0]) / divisor) % globalModulus
				m.Items = m.Items[1:]
				monkey2 := monkeys[m.ZeroModuloMonkey]
				if item % m.Modulus != 0 {
					monkey2 = monkeys[m.NonzeroModuloMonkey]
				}
				monkey2.Items = append(monkey2.Items, item)
			}
		}
	}
	sort.Slice(inspectCounts, func(i, j int) bool { return inspectCounts[i] > inspectCounts[j]})
	return inspectCounts[0] * inspectCounts[1]
}

func Part1(input io.Reader) int {
	return solve(input, 20, 3)
}

func Part2(input io.Reader) int {
	return solve(input, 10000, 1)
}

func main() {
	utils.Run("day11.txt", Part1, Part2)
}
