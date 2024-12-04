package main

import (
	"bufio"
	"io"
	"log/slog"
	"regexp"
	"utility"
)

const FIND_REG_PART1 = `(mul\(\d+,\d+\))`
const FIND_REG_PART2 = `(mul\(\d+,\d+\))`
const PARSE_REG = `mul\((\d+),(\d+)\)`

func find(inputReader io.Reader, regexString string) []string {
	scanner := bufio.NewScanner(inputReader)
	re := regexp.MustCompile(regexString)
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

	for _, operation := range operations {
		matches := re.FindStringSubmatch(operation)
		matchesInt := utility.AtoiSlice(matches[1:])
		parsedOperations = append(parsedOperations, matchesInt)
	}

	return parsedOperations
}

func multiplySum(operations [][]int) (sum int) {
	for _, op := range operations {
		sum += multiply(op)
	}
	return
}

func multiply(op []int) (result int) {
	result = 1
	for _, op := range op {
		result *= op
	}
	return result
}

func partOne() {
	inputReader := utility.GetInputReader()
	foundOperations := find(inputReader, FIND_REG_PART1)
	multiplicationSets := parse(foundOperations)
	multplicationSum := multiplySum(multiplicationSets)

	slog.Info("calculated result for input", "sum", multplicationSum)
}

func partTwo() {
	inputReader := utility.GetInputReader()
	foundOperations := find(inputReader, FIND_REG_PART2)
	multiplicationSets := parse(foundOperations)
	multplicationSum := multiplySum(multiplicationSets)

	slog.Info("calculated result for input", "sum", multplicationSum)
}

func main() {
	partOne()
	partTwo()
}
