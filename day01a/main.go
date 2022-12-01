package main

import (
	"bufio"
	"io"
	"sort"
	"strings"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

type NMaximums struct {
	Values []int
	N      int
}

func (m *NMaximums) Push(val int) {
	if len(m.Values) < m.N {
		m.Values = append(m.Values, val)
		sort.Ints(m.Values)
	} else if val > m.Values[0] {
		m.Values = append(m.Values[1:m.N], val)
		sort.Ints(m.Values)
	}
}

func parseElves(input io.Reader) <-chan int {
	sum := 0

	result := make(chan int)

	scanner := bufio.NewScanner(input)
	go (func(){
		defer close(result)
		for scanner.Scan() {
			line := scanner.Text()
			if len(strings.TrimSpace(line)) == 0 {
				result <- sum
				sum = 0
			} else {
				sum += utils.AtoiOrPanic(line)
			}
		}
		if sum != 0 {
			result <- sum
		}
	})()

	return result
}

func solveForN(input io.Reader, n int) int {
	maximiser := NMaximums{N: n}
	for elf := range parseElves(input) {
		maximiser.Push(elf)
	}
	sum := 0
	for _, i := range maximiser.Values {
		sum += i
	}
	return sum
}

func Part1(input io.Reader) int {
	return solveForN(input, 1)
}

func Part2(input io.Reader) int {
	return solveForN(input, 3)
}

func main() {
	utils.Run("day01.txt", Part1, Part2)
}
