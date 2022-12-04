package main

import (
	"bufio"
	"io"
	"math"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

func itemPriority(chr rune) int {
	if chr < 'a' {
		return int((chr - 'A') + 27)
	}
	return int((chr - 'a') + 1)
}

func lineToBitmap(line string) uint64 {
	out := uint64(0)
	for _, chr := range line {
		out |= 1 << itemPriority(chr)
	}
	return out
}

func Part1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		packLen := len(line)/2

		sack1 := lineToBitmap(line[0:packLen])
		sack2 := lineToBitmap(line[packLen:])

		sum += int(math.Log2(float64(sack1 & sack2)))

	}
	return sum
}

func Part2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		sack1 := lineToBitmap(scanner.Text())
		scanner.Scan()
		sack2 := lineToBitmap(scanner.Text())
		scanner.Scan()
		sack3 := lineToBitmap(scanner.Text())

		sum += int(math.Log2(float64(sack1 & sack2 & sack3)))
	}
	return sum
}

func main() {
	utils.Run("day03.txt", Part1, Part2)
}
