package main

import (
	"bufio"
	"log/slog"
	"strings"

	"github.com/bcheldelin/go-advent-of-code/aoc24/utility"
)

func checkLevel(level []int) bool {
	var direction string
	if level[0] > level[1] {
		direction = "up"
	} else {
		direction = "down"
	}
	for i := 1; i < len(level); i++ {
		diff := level[i-1] - level[i]
		// Check that the distance between floors is not more than 3
		if diff > 3 || diff < -3 || diff == 0 {
			return false
		}
		// Check if we started going up that we still are
		if direction == "up" && level[i-1] < level[i] {
			return false
		}
		// Check if we started going down that we still are
		if direction == "down" && level[i-1] > level[i] {
			return false
		}
	}

	return true
}

func saftyCheck(report []int) bool {
	// Check to see if the level is safe with simple rules
	if checkLevel(report) {
		return true
	}

	// Try removing each level util one passes
	for i := range report {
		newReport := make([]int, 0, len(report)-1)
		newReport = append(newReport, report[:i]...)
		newReport = append(newReport, report[i+1:]...)
		if checkLevel(newReport) {
			return true // The new report passed
		}
	}

	// Removing a level did not result in a safty pass
	return false
}

func main() {
	inputReader := utility.NewInputReader("input.txt")
	defer inputReader.Close()

	safeLevels := 0
	safeLevelsDamped := 0
	slog.Info("got here")

	scanner := bufio.NewScanner(inputReader)
	for scanner.Scan() {
		line := scanner.Text()
		slog.Info("read line", "line", line)
		stringSlice := strings.Split(line, " ")
		slog.Info("read line", "slice", stringSlice)
		report := utility.AtoiSlice(stringSlice)
		if checkLevel(report) {
			safeLevels++
		}
		if saftyCheck(report) {
			safeLevelsDamped++
		}
	}
	slog.Info("safe reports calculated", "count", safeLevels)
	slog.Info("safe reports calculated with dampening", "count", safeLevelsDamped)
}
