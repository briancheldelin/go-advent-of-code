package main

import (
	"fmt"
	"grid"
	"gui"
	"os"
	"utility"

	tea "github.com/charmbracelet/bubbletea"
)

type position struct {
	x int
	y int
}

type direction int

const (
	up direction = iota
	right
	down
	left
)

// A room is prepresented by a 2d grid. Feautres:
// * The 0, 0 position in the room is in the top left
// * Each point on the room grid is eaither a 1 or 0.
// * A `#` is an ocupied space,
// * A ` ` is an unocupied space that has not been searched by guard.
// * A `.` represenets and unocupied space that has been searched by guard.
type Challange struct {
	grid   grid.Grid
	guard  position
	facing direction
	moves  int
}

func NewChallange() *Challange {
	input := utility.GetExampleInput()
	return &Challange{
		grid.NewGrid(input),
		position{0, 0},
		up,
		0,
	}
}

func (c *Challange) DoWork() bool {
	c.moves++
	return true
}

func (c *Challange) Render() string {
	return c.grid.String()
}

func (c *Challange) Result() int {
	return c.moves
}

func main() {
	c := NewChallange()

	p := tea.NewProgram(
		gui.NewGUI("Day 6", c, 1),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}
