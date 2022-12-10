package main

import (
	"bufio"
	"fmt"
	"io"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

func cycle(cycleCount *int, x int) int {
	xCoord := *cycleCount % 40
	if xCoord == 0 {
		fmt.Print("\n")
	}
	if utils.Abs(x-xCoord) <= 1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}

	*cycleCount++

	if (*cycleCount-20)%40 == 0 {
		return *cycleCount * x
	}
	return 0
}

func Part1(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	x := 1
	cycleCount := 0
	strengthSum := 0

	for scanner.Scan() {
		line := scanner.Text()

		strengthSum += cycle(&cycleCount, x)
		if len(line) > 5 && line[0:5] == "addx " {
			strengthSum += cycle(&cycleCount, x)
			x += utils.AtoiOrPanic(line[5:])
		}
	}
	fmt.Print("\n")
	return strengthSum
}

func Part2(input io.Reader) int {

	return 0
}

func main() {
	utils.Run("day10.txt", Part1, Part2)
}
