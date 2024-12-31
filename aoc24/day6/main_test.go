package main

import (
	"testing"

	"github.com/briancheldelin/go-advent-of-code/aoc24/utility/grid"
)

type ChallangeTestCase struct {
	testGrid      []byte
	expected      bool
	part2Solution bool
}

var exampleGridPart1 = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestProcessChallange(t *testing.T) {
	testCases := map[string]ChallangeTestCase{
		"Example Grid": {[]byte(exampleGridPart1), true, false},
	}

	for i, test := range testCases {
		t.Run(i, func(t *testing.T) {
			c := NewChallange(1, grid.NewGrid(test.testGrid))
			processChallange(c)
			if c.done != true {
				t.Errorf("challange done state is unexpected")
			}
			if c.part2Solution != test.part2Solution {
				t.Errorf("challanage part2 solution is unexpected: got %t, wanted %t", c.part2Solution, test.part2Solution)
			}
		})
	}
}
