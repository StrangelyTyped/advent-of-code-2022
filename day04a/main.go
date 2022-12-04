package main

import (
	"bufio"
	"io"
	"regexp"
	"math/big"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

var lineRe = regexp.MustCompile(`([\d]+)-([\d]+),([\d]+)-([\d]+)`)

func toBitmap(r1, r2 int) big.Int {
	var out big.Int
	for i := r1; i <= r2; i++ {
		out.SetBit(&out, i, 1)
	}
	return out
}

func solve(input io.Reader, impl func(a, b big.Int) int) int {
	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		matches := utils.MapToInt(lineRe.FindStringSubmatch(line)[1:])
		sum += impl(toBitmap(matches[0],matches[1]), toBitmap(matches[2],matches[3]))
	}
	return sum
}

func Part1(input io.Reader) int {
	return solve(input, func(a, b big.Int) int {
		var x big.Int
		x.And(&a, &b)
		if (x.Cmp(&a) == 0 || x.Cmp(&b) == 0) {
			return 1
		}
		return 0
	})
}

func Part2(input io.Reader) int {
	return solve(input, func(a, b big.Int) int {
		var x, y big.Int
		x.And(&a, &b)
		if x.Cmp(&y) == 0 {
			return 0
		}
		return 1
	})
}

func main() {
	utils.Run("day04.txt", Part1, Part2)
}
