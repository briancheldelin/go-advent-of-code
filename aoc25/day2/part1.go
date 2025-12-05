package day2

import (
	"log/slog"
	"strconv"
	"strings"
)

func part1(input string) {
	var invalidIDSumm int

	for _, line := range strings.Split(input, ",") {
		start, end := parseRange(line)
		for i := start; i <= end; i++ {
			if isRepeat(strconv.Itoa(i)) {
				invalidIDSumm += i
			}
		}
	}
	slog.Info("Sum of invalid IDs", "sum", invalidIDSumm)
}

func isRepeat(number string) bool {
	mid := len(number) / 2
	firstHalf := number[0:mid]
	secondtHalf := number[mid:]
	return firstHalf == secondtHalf
}
