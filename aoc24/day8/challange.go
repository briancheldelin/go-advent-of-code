package day8

import "github.com/briancheldelin/go-advent-of-code/aoc24/utility/grid"

type Challange struct {
	// Satisifies Interface
	id   int
	grid grid.Grid
	done bool

	// day8 attributes
	frequencies map[byte][]location
	antinodes   int
}

func NewChallange(input string) Challange {
	return Challange{}
}

// Satisfies Challange interface
func (c *Challange) DoWork() bool {
	// Todo
	return true
}

// Satisfies Challange interface
func (c *Challange) Render() string {
	return c.grid.String()
}

// Satisfies Challange interface
func (c *Challange) Result() int {
	return c.antinodes
}
