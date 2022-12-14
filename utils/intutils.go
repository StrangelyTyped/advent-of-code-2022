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

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(in int) int {
	if in < 0 {
		return -in
	}
	return in
}

func Sign(in int) int {
	if in < 0 {
		return -1
	} else if in > 0 {
		return 1
	}
	return 0
}