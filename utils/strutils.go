package utils

import (
	"fmt"
	"io"
	"strings"
)

func CleanInput(input string) string {
	return strings.TrimRight(strings.TrimLeft(input, "\n"), "\n")
}



func ReadInput(reader io.Reader) string {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	return CleanInput(string(bytes))
}

func PrintOutput(result1, result2 int) {
	fmt.Printf("Part 1: %d\nPart 2: %d\n\n", result1, result2)
}

func PrintOutputStr(result1, result2 string) {
	fmt.Printf("Part 1: %s\nPart 2: %s\n\n", result1, result2)
}

func MapToInt(in []string) []int {
	out := []int{}
	for _, x := range in {
		out = append(out, AtoiOrPanic(x))
	}
	return out
}