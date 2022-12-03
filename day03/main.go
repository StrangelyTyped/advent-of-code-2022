package main

import (
	"bufio"
	"io"
	"strings"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

func itemPriority(chr byte) int {
	if chr < 'a' {
		return int((chr - 'A') + 27)
	}
	return int((chr - 'a') + 1)
}

func Part1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		packLen := len(line)/2

		sack1 := line[0:packLen]
		sack2 := line[packLen:]
		for i := 0; i < packLen; i++  {
			// screw unicode (not really, love you usually)
			chr := sack1[i]
			if strings.IndexByte(sack2, chr) != -1 {
				sum += itemPriority(chr)
				break
			}
		}
	}
	return sum
}

func Part2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()

		for i := 0; i < len(line1); i++  {
			// screw unicode (not really, love you usually)
			chr := line1[i]
			if strings.IndexByte(line2, chr) != -1 && strings.IndexByte(line3, chr) != -1 {
				sum += itemPriority(chr)
				break
			}
		}
	}
	return sum
}

func main() {
	utils.Run("day03.txt", Part1, Part2)
}
