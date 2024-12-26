package grid

import (
	"testing"
)

func TestString(t *testing.T) {
	testGrid := []byte("...\n...\n...")
	expectedGrid := string("\x1b[37m.\x1b[37m.\x1b[37m.\n\x1b[37m.\x1b[37m.\x1b[37m.\n\x1b[37m.\x1b[37m.\x1b[37m.\n")

	grid := NewGrid(testGrid)
	result := grid.String()

	if result != expectedGrid {
		t.Error("Did not get expected grid")
	}

}

func TestNewGrid(t *testing.T) {
	testGrid := []byte("...\n...\n...")

	expecedGrid := Grid{
		{
			{Space: '.', Light: Gray},
			{Space: '.', Light: Gray},
			{Space: '.', Light: Gray},
		}, {
			{Space: '.', Light: Gray},
			{Space: '.', Light: Gray},
			{Space: '.', Light: Gray},
		}, {
			{Space: '.', Light: Gray},
			{Space: '.', Light: Gray},
			{Space: '.', Light: Gray},
		},
	}

	grid := NewGrid(testGrid)
	// result := grid.String()

	if !grid.equal(expecedGrid) {
		t.Error("Did not get expected grid")
	}

}
