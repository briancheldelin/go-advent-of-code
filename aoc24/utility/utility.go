package utility

import (
	"io"
	"log/slog"
	"os"
	"strconv"
)

func NewInputReader(filename string) *os.File {
	inputReader, err := os.Open(filename)

	if err != nil {
		slog.Error("got an error while trying to open input file", "Error", err)
	}

	return inputReader
}

func InputString() []byte {
	inputFile := NewInputReader("input.txt")

	if b, err := io.ReadAll(inputFile); err == nil {
		return b
	}

	return nil
}

// Gets the problem input from the same directory as package
func GetInput() []byte {
	inputFile := NewInputReader("input.txt")

	if b, err := io.ReadAll(inputFile); err == nil {
		return b
	}

	return nil
}

// Gets the example input from the same directory as package.
func GetExampleInput() []byte {
	inputFile := NewInputReader("input-example.txt")

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
