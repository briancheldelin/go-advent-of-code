package day7

import (
	"log/slog"
	"strconv"
	"strings"
	"time"
	"utility"
)

type calibration struct {
	sum int
	ops []int
}

func NewCalibration(line string) calibration {
	lineSplit := strings.Split(line, ":")
	sum, err := strconv.Atoi(lineSplit[0])
	if err != nil {
		slog.Error("unable to parse sum to int")
	}
	ops := utility.AtoiSlice(strings.Split(lineSplit[1], " ")[1:])

	return calibration{
		sum,
		ops,
	}
}

// This is a search tree, we search each potential path of operator combinations.
// As soon as the stepSum is ovrer the calibration desired sum then we fail the seach path.
// Other paths will contine to search until the first path that return true.
func (c *calibration) Search(o byte, i int, stepSum int) bool {
	// Check if we are over and fail the search path
	if stepSum > c.sum {
		return false
	}

	// Check that we are finished and we have a good path
	if len(c.ops) < i+1 {
		if stepSum == c.sum {
			return true
		} else {
			return false
		}
	}

	// For addition add to our sum then search for + and *
	if o == '+' {
		stepSum += c.ops[i]
		i++
		if c.Search('+', i, stepSum) {
			return true
		}
		if c.Search('*', i, stepSum) {
			return true
		}
		return false
	}

	// For multiplication multiply to our sum then search for + and *
	if o == '*' {
		stepSum *= c.ops[i]
		i++
		if c.Search('+', i, stepSum) {
			return true
		}
		if c.Search('*', i, stepSum) {
			return true
		}
		return false
	}
	slog.Error("We should have never gotten here")
	return false
}

func part1(input string) (count int, answer int) {
	var calibrations []calibration
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		calibrations = append(calibrations, NewCalibration(line))
	}

	for _, cal := range calibrations {

		if cal.Search('+', 1, cal.ops[0]) || cal.Search('*', 1, cal.ops[0]) {
			count++
			answer += cal.sum
		}
	}
	return
}

func Challange(input string, part int) {
	slog.Info("Running Challange", "day", 7, "part", part)
	start := time.Now()
	part1Count, part1Answer := part1(input)
	elapsed := time.Since(start)
	slog.Info("finished challange part1", "count", part1Count, "time", elapsed, "answer", part1Answer)
}