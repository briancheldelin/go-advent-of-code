package main

import (
	"fmt"
	"grid"
	"gui"
	"os"
	"utility"

	tea "github.com/charmbracelet/bubbletea"
)

type Position struct {
	facing direction
	X      int
	Y      int
}

type direction byte

const (
	up    = '^'
	right = '>'
	down  = 'v'
	left  = '<'
)

var trail map[byte]byte = map[byte]byte{
	'>': '-',
	'<': '-',
	'^': '|',
	'v': '|',
}

// A room is prepresented by a 2d grid. Feautres:
// * The 0, 0 position in the room is in the top left
// * Each point on the room grid is eaither a 1 or 0.
// * A `#` is an ocupied space,
// * A ` ` is an unocupied space that has not been searched by guard.
// * A `.` represenets and unocupied space that has been searched by guard.
type Challange struct {
	id               int
	grid             grid.Grid
	guard            Position
	searchedSpaces   int
	totalMoves       int
	done             bool
	part2Solution    bool
	movesSinceUnique int
}

func NewChallange(id int, inputGrid grid.Grid) *Challange {
	// input := utility.GetExampleInput()
	// input := utility.GetInput()

	// grid := grid.NewGrid(input)

	c := Challange{
		id:             id,
		grid:           inputGrid,
		searchedSpaces: 0,
		done:           false,
		part2Solution:  false,
	}

	c.findGuard()

	return &c
}

func (c *Challange) findGuard() {
	for y := 0; y < len(c.grid); y++ {
		for x := 0; x < len(c.grid[y]); x++ {
			space := c.grid[y][x].Space
			if space == byte(up) {
				c.guard = Position{facing: direction(space), X: x, Y: y}
				return
			}
		}
	}
}

// Satisfies Challange interface
func (c *Challange) DoWork() bool {
	// Check if we are done
	if c.done {
		return false
	}

	// Check if we are on a loop of a 1M unique Moves
	if c.movesSinceUnique > 200 {
		c.done = true
		c.part2Solution = true
		return false
	}

	if !c.checkPath() {
		if !c.done {
			c.TurnRight()
		} else {
			return false
		}
	} else {
		if !c.done {
			c.StepForword(trail[byte(c.guard.facing)])
			// c.StepForword(byte(c.guard.facing))
		} else {
			return false
		}
	}

	return true
}

// Satisfies Challange interface
func (c *Challange) Render() string {
	return c.grid.String()
}

// Satisfies Challange interface
func (c *Challange) Result() int {
	return c.searchedSpaces
}

//
// Position Functions
//

func (c *Challange) StepForword(leaveBehind byte) {
	// Leave behind a trail
	c.grid[c.guard.Y][c.guard.X].Space = leaveBehind
	c.grid[c.guard.Y][c.guard.X].Light = grid.Yellow

	switch c.guard.facing {
	case up:
		c.guard.Y--
	case right:
		c.guard.X++
	case down:
		c.guard.Y++
	case left:
		c.guard.X--
	}

	// Update Moves
	c.totalMoves++
	c.movesSinceUnique++

	// Check if we are on a new unexplorered space
	if c.grid[c.guard.Y][c.guard.X].Space == '.' {
		c.searchedSpaces++
		c.movesSinceUnique = 0 // Reset
	}

	// Place our Guard
	c.grid[c.guard.Y][c.guard.X].Space = byte(c.guard.facing)
	c.grid[c.guard.Y][c.guard.X].Light = grid.Red
}

func (c *Challange) TurnRight() {
	switch c.guard.facing {
	case up:
		c.guard.facing = right
	case right:
		c.guard.facing = down
	case down:
		c.guard.facing = left
	case left:
		c.guard.facing = up
	}
	c.grid[c.guard.Y][c.guard.X].Space = byte(c.guard.facing)
	// c.StepForword(byte(c.guard.facing))
	// c.StepForword('+')
}

// Check forword
func (c *Challange) checkPath() bool {
	lookX := c.guard.X
	lookY := c.guard.Y

	switch c.guard.facing {
	case up:
		lookY--
	case right:
		lookX++
	case down:
		lookY++
	case left:
		lookX--
	}
	// Check that we are withing the bounds
	if lookY < 0 || lookY > len(c.grid)-1 {
		c.done = true
		return false
	}
	if lookX < 0 || lookX > len(c.grid[lookY])-1 {
		c.done = true
		return false
	}

	// Check if we can move into the space
	s := c.grid[lookY][lookX].Space
	if s == '#' || s == 'O' {
		return false
	}

	// Check if we are about to walk into a path that will result in loop
	if s == byte(c.guard.facing) /*|| s == '|' || s == '-' /*||  s == '+'*/ {
		c.part2Solution = true
		c.done = true
		return false
	}

	return true
}

// Part 2 Functions
func getPossible(o *Challange) []*Challange {
	solutions := make([]*Challange, 0, o.searchedSpaces+1)

	id := 2

	for y := 0; y < len(o.grid); y++ {
		for x := 0; x < len(o.grid[y]); x++ {
			if isDirection(o.grid[y][x].Space) {

				n := NewChallange(id, grid.NewGrid(utility.InputString()))
				n.grid[y][x].Space = 'O'
				solutions = append(solutions, n)
				id++
			}
		}
	}

	return solutions
}

func isDirection(b byte) bool {
	switch b {
	case up, right, down, left:
		return true
	case '|', '-':
		return true
	}
	return false
}

func processChallange(c *Challange) bool {
	// fmt.Printf("Checking Solution #%d: ", c.id)
	for c.DoWork() {
		// Check if we are stuck in a loop
		// if c.totalMoves >= 10000 {
		// 	visulizeChallange(c)
		// }
		// Check if we are in a loop
		// if c.movesSinceUnique >= 1000 {
		// 	c.part2Solution = true
		// 	break
		// }

		continue
	}
	// fmt.Printf("moves=%d, unique=%d ", c.totalMoves, c.searchedSpaces)
	return true
}

func visulizeChallange(c *Challange) {
	p := tea.NewProgram(
		gui.NewGUI(fmt.Sprintf("Day 6 Solution %d", c.id), c, 5000),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func part1() *Challange {
	c := NewChallange(1, grid.NewGrid(utility.InputString()))
	// visulizeChallange(c)
	for c.DoWork() {
		continue
	}
	fmt.Printf("Part 1 Result is: %d\n", c.searchedSpaces+1)
	return c
}

func part2(c *Challange) {
	// Get a list of challanges that have an obstical placed in them.
	possibleSolutions := getPossible(c)

	fmt.Printf("Created %d solutions to test\n", len(possibleSolutions))

	var solutions int

	for _, p := range possibleSolutions {
		// visulizeChallange(p)
		processChallange(p)
		if p.part2Solution {
			// fmt.Printf("Good\n")
			solutions++
		} else {
			// fmt.Printf("Nope\n")
		}
	}

	fmt.Printf("Part 2 Solution is %d\n", solutions)
}

func main() {
	part2(part1())
}
