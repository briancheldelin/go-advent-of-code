package main

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

func startSearch(x, y int, matrix *matrix, letter int) int {

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

func searchDirection(x int, y int, d direction, matrix *matrix, letter int) bool {
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

//
// Lipgloss
//

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

//
// Buble Tea fun
///

type matrix [][]byte

func initModel() model {

	inputMatrix := matrix(bytes.Split([]byte(INPUT_EXAMPLE), []byte("\n")))

	rows := []table.Row{}
	for _, row := range inputMatrix {
		rows = append(rows, table.Row{string(row)})
	}

	t := table.New(
		table.WithRows(rows),
	)

	s := table.DefaultStyles()

	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	t.SetStyles(s)

	return model{
		grid:   matrix(bytes.Split([]byte(INPUT_EXAMPLE), []byte("\n"))),
		xFocus: 0,
		yFocus: 0,
		done:   false,
		view: searchView{
			header: "Lets save XMAS?\n\n",
			table:  t,
			footer: "\nPress q to quit.\n",
		},
	}
}

type searchView struct {
	header string
	table  table.Model
	footer string
}

type model struct {
	grid   matrix
	xFocus int
	yFocus int
	done   bool
	view   searchView
}

func (m model) Init() tea.Cmd {
	return searchTick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case frameMsg:
		if m.done {
			return m, searchTick()
		}

		found := startSearch(m.xFocus, m.yFocus, &m.grid, 0)

		if found > 0 {
			m.grid[m.yFocus][m.xFocus] = '*'
		}

		if m.xFocus == len(m.grid[m.xFocus])-1 && m.yFocus == len(m.grid)-1 {
			// We are at the end of the grid
			m.done = true
		} else if m.xFocus < len(m.grid[m.xFocus])-1 {
			// Stay on same line
			m.xFocus++
		} else {
			// Move to next line
			m.xFocus = 0
			m.yFocus++
		}

		return m, searchTick()
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := "Lets save XMAS?\n\n"

	// Iterate over our choices
	s += baseStyle.Render(m.view.table.View()) + "\n"

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

const FPS = 5

type frameMsg struct{}

func searchTick() tea.Cmd {
	return tea.Tick(time.Second/FPS, func(_ time.Time) tea.Msg {
		return frameMsg{}
	})
}

func main() {
	// slog.SetLogLoggerLevel(slog.LevelDebug)
	// input := []byte(INPUT_EXAMPLE)
	// matrix := bytes.Split(input, []byte("\n"))
	// hight := len(matrix)
	// width := len(matrix[0])
	// var found = 0
	// for y := 0; y < hight; y++ {
	// 	for x := 0; x < width; x++ {
	// 		// Start the search and go negative so we search from start
	// 		found += startSearch(x, y, &matrix, 0)
	// 	}
	// }
	// slog.Info("done searching for XMAS", "count", found)

	p := tea.NewProgram(initModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}
