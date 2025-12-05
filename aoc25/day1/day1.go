package day1

import (
	"log/slog"
	"math"
	"strconv"
	"strings"
)

func RunChallange(input string, part int) {
	slog.Info("Hello from day1")

	switch part {
	case 1:
		Part1(input)
	case 2:
		Part2(input)
	}

}

func Part1(input string) {
	var position float64 = 50
	var count int

	for line := range strings.Lines(input) {
		direction := line[0]
		line = strings.Trim(line, "RL\n")
		amount, _ := strconv.ParseFloat(line, 64)
		slog.Debug("Got", "direction", string(direction), "amount", line)

		switch direction {
		case 'L':
			slog.Debug("Going Left", "position", position, "amount", amount)
			position -= amount
		case 'R':
			slog.Debug("Going Right", "position", position, "amount", amount)
			position += amount
		}

		remainder := math.Mod(math.Abs(position), 100)
		slog.Info("Mod result", "remainder", remainder, "position", position)
		if remainder == 0 {
			count++
		}
	}
	slog.Info("Challange Solved", "count", count)
}

func Part2(input string) {
	var position float64 = 50
	var count int
	var found int

	for line := range strings.Lines(input) {
		direction := line[0]
		line = strings.Trim(line, "RL\n")
		amount, _ := strconv.ParseFloat(line, 64)
		slog.Debug("Got", "direction", string(direction), "amount", line)

		switch direction {
		case 'L':
			slog.Debug("Going Left", "position", position, "amount", amount)
			found, position = rotate(position, amount, left)
		case 'R':
			slog.Debug("Going Right", "position", position, "amount", amount)
			found, position = rotate(position, amount, right)
		}

		remainder := math.Mod(math.Abs(position), 100)
		slog.Info("Mod result", "remainder", remainder, "position", position)
		// if remainder == 0 {
		// 	count++
		// }
		count += found
	}
	slog.Info("Challange Solved", "count", count)
}

func rotate(startPosition float64, amount float64, d func(i float64) float64) (count int, endPosition float64) {
	position := startPosition
	for range int(amount) {
		position = d(position)
		remainder := math.Mod(math.Abs(position), 100)
		if remainder == 0 {
			count++
		}
	}
	endPosition = position

	return
}

func left(i float64) float64 {
	return i - 1
}

func right(i float64) float64 {
	return i + 1
}
