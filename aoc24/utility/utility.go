package utility

import (
	"io"
	"log/slog"
	"os"
	"strconv"
)

func NewInputReader() *os.File {
	inputReader, err := os.Open("input.txt")

	if err != nil {
		slog.Error("got an error while trying to open input file", "Error", err)
	}

	return inputReader
}

func InputString() []byte {
	inputFile := NewInputReader()

	if b, err := io.ReadAll(inputFile); err == nil {
		return b
	}

	return nil
}

func AtoiSlice(d []string) []int {
	intSlice := make([]int, len(d))
	for i, e := range d {
		convertedElement, err := strconv.Atoi(e)
		if err != nil {
			slog.Error("error converting string to int", "error", err)
		}
		intSlice[i] = convertedElement
	}
	return intSlice
}
