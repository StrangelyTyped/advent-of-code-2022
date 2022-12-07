package utils

import "strconv"

func AtoiOrPanic(input string) int {
	entry, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return entry
}

func Atoi64OrPanic(input string) uint64 {
	entry, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		panic(err)
	}
	return entry
}