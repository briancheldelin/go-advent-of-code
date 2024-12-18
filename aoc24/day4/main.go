package main

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"utility"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
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

// var Reset = "\033[0m"
// var Red = "\033[31m"
// var Green = "\033[32m"
// var Yellow = "\033[33m"
// var Blue = "\033[34m"
// var Magenta = "\033[35m"
// var Cyan = "\033[36m"
// var Gray = "\033[37m"
// var White = "\033[97m"

type direction func(int, int) (int, int)

func startSearch(x, y int, matrix *matrixV2, letter int) int {

	// Check if we are an the begenning
	testChar := XMAS[letter]
	if (*matrix)[y][x].character != testChar {
		return 0
	}
	slog.Debug("We found an X")

	letter++ // Now we searh for the next letter

	right := func(x, y int) (int, int) { slog.Debug("searching right"); x++; return x, y }
	downRight := func(x, y int) (int, int) { slog.Debug("searching down right"); x++; y++; return x, y }
	down := func(x, y int) (int, int) { slog.Debug("searching down"); y++; return x, y }
	downLeft := func(x, y int) (int, int) { slog.Debug("searching down left"); x--; y++; return x, y }
	left := func(x, y int) (int, int) { slog.Debug("searching left"); x--; return x, y }
	upLeft := func(x, y int) (int, int) { slog.Debug("searching up left"); x--; y--; return x, y }
	up := func(x, y int) (int, int) { slog.Debug("searching up"); y--; return x, y }
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

func searchDirection(x int, y int, d direction, matrix *matrixV2, letter int) bool {

	// Mutate the x, y for the direction we are searching.
	x, y = d(x, y)

	// Check that we are in bounds of where we are looking
	hight := len(*matrix)      // Check how many rows we have for hight
	width := len((*matrix)[0]) // Check the first line of matrix for width
	if outsideBoundry(x, y, hight, width) {
		return false
	}

	if (*matrix)[y][x].character == XMAS[letter] {
		// check if we are at the end
		if letter == len(XMAS)-1 {
			(*matrix)[y][x].color = Green
			return true
		}
		// Search for the next letter in the direction
		if searchDirection(x, y, d, matrix, letter+1) {
			(*matrix)[y][x].color = Green
			return true
		} else {
			return false
		}
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

type matrixV2 [][]cell

type cell struct {
	character byte
	color     string
}

func initModel() *model {
	// input := bytes.Split([]byte(INPUT_EXAMPLE), []byte("\n"))
	input := bytes.Split(utility.InputString(), []byte("\n"))

	g := make(matrixV2, len(input))

	for y := range input {
		g[y] = make([]cell, len(input[y]))
		for x := range input[y] {
			g[y][x] = cell{input[y][x], ""}
		}
	}

	return &model{
		grid:   &g,
		xFocus: 0,
		yFocus: 0,
		total:  0,
		done:   false,
	}
}

type model struct {
	grid     *matrixV2
	xFocus   int
	yFocus   int
	total    int
	done     bool
	viewport viewport.Model
	ready    bool
}

func visWork() {
	p := tea.NewProgram(
		initModel(),
		tea.WithAltScreen(),       // use the full size of the terminal in its "alternate screen buffer"
		tea.WithMouseCellMotion(), // turn on mouse support so we can track the mouse wheel
	)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func main() {
	visWork()
}
