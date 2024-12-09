package main

import (
	"bytes"
	"log/slog"
)

const INPUT_EXAMPLE = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

const X, M, A, S byte = 'X', 'M', 'A', 'S'

const XMAS = `XMAS`

type direction func(int, int) (int, int)

func startSearch(x, y int, matrix *[][]byte, letter int) int {

	// Check if we are an the begenning
	testChar := XMAS[letter]
	if (*matrix)[y][x] != testChar {
		return 0
	}
	slog.Debug("We found an X")

	letter++ // Now we searh for the next letter

	right := func(x, y int) (int, int) { slog.Debug("searching right"); x++; return x, y }
	downRight := func(x, y int) (int, int) { slog.Debug("searching down right"); x++; y++; return x, y }
	down := func(x, y int) (int, int) { slog.Debug("searching down"); y++; return x, y }
	downLeft := func(x, y int) (int, int) { slog.Debug("searching down left"); x--; y++; return x, y }
	left := func(x, y int) (int, int) { slog.Debug("searching left"); x--; return x, y }
	upLeft := func(x, y int) (int, int) { slog.Debug("searching up left"); x--; y++; return x, y }
	up := func(x, y int) (int, int) { slog.Debug("searching up"); y++; return x, y }
	upRight := func(x, y int) (int, int) { slog.Debug("searching up right"); x++; y--; return x, y }

	found := 0

	if searchDirection(x, y, right, matrix, letter) {
		found++
	}
	if searchDirection(x, y, downRight, matrix, letter) {
		found++
	}
	if searchDirection(x, y, down, matrix, letter) {
		found++
	}
	if searchDirection(x, y, downLeft, matrix, letter) {
		found++
	}
	if searchDirection(x, y, left, matrix, letter) {
		found++
	}
	if searchDirection(x, y, upLeft, matrix, letter) {
		found++
	}
	if searchDirection(x, y, up, matrix, letter) {
		found++
	}
	if searchDirection(x, y, upRight, matrix, letter) {
		found++
	}

	return found
}

func searchDirection(x int, y int, d direction, matrix *[][]byte, letter int) bool {
	// Check if we are finished
	if letter == len(XMAS)-1 {
		slog.Debug("we found the entire XMAS")
		return true // We found all of XMAS!
	}

	// Mutate the x, y for the direction we are searching.
	x, y = d(x, y)

	// Check that we are in bounds of where we are looking
	hight := len(*matrix)      // Check how many rows we have for hight
	width := len((*matrix)[0]) // Check the first line of matrix for width
	if outsideBoundry(x, y, hight, width) {
		return false
	}

	if (*matrix)[y][x] == XMAS[letter] {
		// Search for the next letter in the direction
		return searchDirection(x, y, d, matrix, letter+1)
	}
	return false
}

func outsideBoundry(x, y, h, w int) bool {
	// Check that we are positive
	if x < 0 || y < 0 {
		return true
	}

	if x >= w || y >= h {
		return true
	}

	return false // We are not outside the boundry
}

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	input := []byte(INPUT_EXAMPLE)

	matrix := bytes.Split(input, []byte("\n"))

	hight := len(matrix)
	width := len(matrix[0])

	var found = 0

	for y := 0; y < hight; y++ {
		for x := 0; x < width; x++ {
			// Start the search and go negative so we search from start
			found += startSearch(x, y, &matrix, 0)
		}
	}

	slog.Info("done searching for XMAS", "count", found)

}
