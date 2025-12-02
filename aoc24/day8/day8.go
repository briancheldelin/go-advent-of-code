package day8

import (
	"log/slog"
	"time"
)

func RunChallange(input string, part int) {
	slog.Info("Running Challange", "day", 7, "part", part)
	start := time.Now()

	var answer int

	switch part {
	case 1:
		answer = part1(input)
	case 2:
		answer = part2(input)
	}

	elapsed := time.Since(start)
	slog.Info("finished challange", "part", part, "time", elapsed, "answer", answer)
}

func part1(input string) int {
	// Parse input

	//

	return 0
}

func part2(input string) int { return 0 }

func parseInput(intput string) {

}

//
// Day8 Challange Locations
//

type location struct {
	x         int
	y         int
	antenna   bool
	freqency  byte
	antinodes int
}

func NewLocation() location {
	return location{}
}

// For a list of antenas create all combinations of paries
func findParies(a []location) [][]location {
	return nil
}

// For a pare of attenas find the antinodes location
func findAntinodes(p [][]location) {}
