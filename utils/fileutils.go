package utils

import (
	"io"
	"os"
)

func OpenOrPanic(filename string) io.Reader {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}