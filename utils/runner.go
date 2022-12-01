package utils

import (
	"flag"
	"io"
	"path"
)

type AoCFunc func(io.Reader) int

func Run(inputFile string, part1 AoCFunc, part2 AoCFunc) {
	inputDir := flag.String("input-dir", "../inputs", "Directory containing input files")
	flag.Parse()
	inFile := path.Join(*inputDir, inputFile)
	PrintOutput(part1(OpenOrPanic(inFile)), part2(OpenOrPanic(inFile)))
}
