package main

import (
	"bufio"
	"io"
	"fmt"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)


func Part1(input io.Reader) string {
	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += snafuToInt(line)
	}
	fmt.Printf("Base10 result: %d\n", sum)
	return intToSnafu(sum)
}

var snafuMap = map[string]int{
	"2": 2,
	"1": 1,
	"0": 0,
	"-": -1,
	"=": -2,
}
var snafuRMap = map[int]string{
	2: "2",
	1: "1",
	0: "0",
	-1: "-",
	-2: "=",
}

func snafuToInt(snafu string) int {
	sum := 0
	x := 1
	for i := len(snafu)-1; i >= 0; i-- {
		sum += snafuMap[snafu[i:i+1]] * x
		x*=5
	}
	return sum
}

func intToSnafu(i int) string {
	x := 1
	for (x * 3) < i {
		x *= 5
	}
	outDigits := []int{}
	for ; x >= 1; x /= 5 {
		delta := i / x
		outDigits = append(outDigits, delta)
		
		i -= delta * x
	}

	for i := len(outDigits) - 1; i >= 0; i-- {
		if outDigits[i] > 2 {
			outDigits[i] -= 5
			if i == 0 {
				i = 1
				outDigits = append([]int{0}, outDigits...)
			}
			outDigits[i - 1] += 1
		}
	}
	outStr := ""
	for _, digit := range outDigits {
		outStr += snafuRMap[digit]
	}
	return outStr
}

func Part2(input io.Reader) string {
	return ""
}

func main() {
	utils.RunStr("day25.txt", Part1, Part2)
}
