package main

import (
	"bufio"
	"io"
	"strings"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

func solve(input io.Reader, chunkSize int) int {
	scanner := bufio.NewScanner(input)
	
	scanner.Scan()
	line := scanner.Text()
	
	for i := chunkSize; i < len(line); i++ {
		slice := line[i-chunkSize:i]
		good := true
		for j := 1; j < len(slice); j++ {
			if strings.IndexByte(slice, slice[j]) != j {
				good = false
				break
			}
		}
		if good {
			return i
		}
	}

	return -1
}

func Part1(input io.Reader) int {
	return solve(input, 4)
}

func Part2(input io.Reader) int {
	return solve(input, 14)
}

func main() {
	utils.Run("day06.txt", Part1, Part2)
}
