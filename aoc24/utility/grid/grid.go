package grid

import (
	"bytes"
	"strings"
)

// Cell colors
const Reset = "\033[0m"

type Color string

const (
	Red     Color = "\033[31m"
	Green   Color = "\033[32m"
	Yellow  Color = "\033[33m"
	Blue    Color = "\033[34m"
	Magenta Color = "\033[35m"
	Cyan    Color = "\033[36m"
	Gray    Color = "\033[37m"
	White   Color = "\033[97m"
)

// A tile is a point in a room
type Cell struct {
	Space byte
	Light Color
}

func (c Color) String() string {
	return string(c)
}

type Grid [][]Cell

// Creates a new Grid from []byte with eual length lines split by newlines
// Default color of each cell is Gray
func NewGrid(input []byte) Grid {
	inputCells := bytes.Split(input, []byte("\n"))

	newGrid := make(Grid, len(inputCells))

	for y := range inputCells {
		newGrid[y] = make([]Cell, len(inputCells[y]))
		for x := range inputCells[y] {
			newGrid[y][x] = Cell{inputCells[y][x], Gray}
		}
	}

	return newGrid
}

// Check if two grids are equal
func (g Grid) equal(c Grid) bool {
	if g == nil || c == nil {
		return false // Grids are not initialized
	}

	if len(g) != len(c) {
		return false // Grids are not equal in lines
	}

	for i := range g {
		if len(g[i]) != len(c[i]) {
			return false // Line length is not equal
		}
		for j := range g[i] {
			if g[i][j].Space != c[i][j].Space {
				return false // Cells don't match
			}
			if g[i][j].Light != c[i][j].Light {
				return false // Lighting of cell dons't match
			}
		}
	}

	return true // Everything is equal
}

// String() Converts a grid into a String version of the grid with colors.
func (g Grid) String() string {
	var builder strings.Builder
	for y := range g {
		for x := range g[y] {
			builder.WriteString(g[y][x].Light.String()) // Add Color
			builder.WriteByte(g[y][x].Space)            // Add Space Contents
		}
		builder.WriteRune('\n') // Add Newline at end of line
	}
	return builder.String()
}
