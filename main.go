package main

import (
	"fmt"
	"io"
	"os"

	"github.com/strangelytyped/advent-of-code-2022/day01"
	"github.com/strangelytyped/advent-of-code-2022/utils"
)

const inputFolder = "inputs"

func openOrPanic(filename string) io.Reader {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}

func main() {
	fmt.Println("Day 1")
	utils.PrintOutput(day01.Main(openOrPanic(inputFolder + "/day01.txt")))
}
