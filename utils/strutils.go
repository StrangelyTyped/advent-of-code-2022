package utils

import (
	"fmt"
	"io"
	"strings"
)

func CleanInput(input string) string {
	return strings.TrimSpace(input)
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

func MapToInt(in []string) []int {
	out := []int{}
	for _, x := range in {
		out = append(out, AtoiOrPanic(x))
	}
	return out
}