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

type AoCFuncStr func(io.Reader) string

func RunStr(inputFile string, part1 AoCFuncStr, part2 AoCFuncStr) {
	inputDir := flag.String("input-dir", "../inputs", "Directory containing input files")
	flag.Parse()
	inFile := path.Join(*inputDir, inputFile)
	PrintOutputStr(part1(OpenOrPanic(inFile)), part2(OpenOrPanic(inFile)))
}
