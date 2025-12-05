package day2

import (
	"log/slog"
	"math"
	"slices"
	"strconv"
	"strings"
)

func part2(input string) {
	var invalidIDSumm int
	for line := range strings.SplitSeq(input, ",") {
		start, end := parseRange(line)
		sum, _, _ := testRange(start, end)
		invalidIDSumm += sum
	}
	slog.Info("Sum of invalid IDs", "sum", invalidIDSumm)
}

func testRange(start int, end int) (sum int, count int, invalidIDs []int) {
	for i := start; i <= end; i++ {
		if testNumeber(strconv.Itoa(i)) {
			sum += i
			count++
			invalidIDs = append(invalidIDs, i)
		}
	}
	return
}

func testNumeber(number string) bool {
	if len(number) == 1 {
		return false
	}

	if strings.Count(number, string(number[0])) == len(number) {
		return true // Found simple repeating number
	}

	primes := []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	if slices.Contains(primes, len(number)) {
		return false
	}

	if isRepeatRecursive(number, 2) {
		return true
	}

	return false
}

func isRepeatRecursive(number string, segments int) bool {
	if (len(number)) < segments {
		return false // End condition
	}

	if math.Mod(float64(len(number)), float64(segments)) != 0 {
		return isRepeatRecursive(number, segments+1)
	}

	segLength := len(number) / segments
	firstSegment := number[0:segLength]
	count := strings.Count(number, firstSegment)
	if count == segments {
		return true
	}

	return isRepeatRecursive(number, segments+1)
}
