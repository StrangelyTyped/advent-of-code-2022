package utils

import "strconv"

func AtoiOrPanic(input string) int {
	entry, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return entry
}