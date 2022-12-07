package main

import (
	"bufio"
	"io"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

var primes = []int{
	2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
	73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173,
}

func chrIndex(b byte) int {
	return primes[b - 'a']
}

func solve(input io.Reader, chunkSize int) int {
	scanner := bufio.NewScanner(input)
	
	scanner.Scan()
	line := scanner.Text()
	
	windowedProduct := 1
	windowStart, windowEnd := 0, 0
	for (windowEnd - windowStart) < chunkSize && windowEnd < len(line) {
		x := chrIndex(line[windowEnd])
		for windowedProduct % x == 0 {
			windowedProduct /= chrIndex(line[windowStart])
			windowStart++
		}
		windowedProduct *= x
		windowEnd++
	}
	return windowEnd
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
