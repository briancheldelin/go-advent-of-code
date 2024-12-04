package main

import (
	"bufio"
	"io"
	"regexp"
	"utility"
)

const FIND_REG = `(mul\(\d+,\d+\))`
const PARSE_REG = `mul\((\d+),(\d+)\)`

func find(inputReader io.Reader) []string {

	scanner := bufio.NewScanner(inputReader)

	re := regexp.MustCompile(FIND_REG)

	hits := make([]string, 0, 10)

	for scanner.Scan() {
		line := scanner.Text()
		hits = append(hits, re.FindAllString(line, -1)...)
	}

	return hits
}

func parse(operations []string) [][]int {
	parsedOperations := make([][]int, 0, len(operations))
	re := regexp.MustCompile(PARSE_REG)

	for i, operation := range operations {
		matches := re.FindAllString(operation, -1)
		matchesInt := utility.AtoiSlice(matches)
		parsedOperations[i] = append(parsedOperations[i], matchesInt...)

	}

}

func main() {
	// inputReader := utility.GetInputReader()
}
