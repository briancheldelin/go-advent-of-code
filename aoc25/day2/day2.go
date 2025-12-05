package day2

import (
	"log/slog"
	"strconv"
	"strings"
)

func RunChallange(input string, part int) {
	slog.Info("Hello from day2")

	switch part {
	case 1:
		part1(input)
	case 2:
		part2(input)
	}

}

func parseRange(line string) (start int, end int) {
	split := strings.Split(line, "-")
	start, _ = strconv.Atoi(split[0])
	end, _ = strconv.Atoi(split[1])
	return
}
