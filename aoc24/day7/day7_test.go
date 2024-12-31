package day7

import (
	"testing"
)

func TestNewCalibration(t *testing.T) {
	test := "190: 10 19"
	expect := calibration{
		190,
		[]int{10, 19},
	}

	result := NewCalibration(test)

	if result.sum != expect.sum {
		t.Errorf("calibration sum is unexpected, got %d, expected %d", result.sum, expect.sum)
	}

	if len(result.ops) != len(expect.ops) {
		t.Errorf("calibration ops length is unexpected, got %d, expected %d", len(result.ops), len(expect.ops))
	}
}

type SearchTest struct {
	input  string
	expect bool
}

func TestSearch(t *testing.T) {
	tests := []SearchTest{
		{"3: 1 1 1", true},
		{"190: 10 19", true},
		{"3267: 81 40 27", true},
		{"83: 17 5", false},
		{"156: 15 6", false},
		{"7290: 6 8 6 15", false},
		{"161011: 16 10 13", false},
		{"192: 17 8 14", false},
		{"21037: 9 7 18 13", false},
		{"292: 11 6 16 20", true},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			c := NewCalibration(test.input)
			addFirstSearch := c.Search('+', 1, c.ops[0])
			multiFirstSearch := c.Search('*', 1, c.ops[0])

			// If eaither search paths is true then don't fail the test
			if !(addFirstSearch == test.expect || multiFirstSearch == test.expect) {
				t.Error("We did not get the expected result")
			}
		})
	}
}

func TestSearchV2(t *testing.T) {
	tests := []SearchTest{
		{"3: 1 1 1", true},
		{"190: 10 19", true},
		{"3267: 81 40 27", true},
		{"83: 17 5", false},
		{"156: 15 6", true},
		{"7290: 6 8 6 15", true},
		{"161011: 16 10 13", false},
		{"192: 17 8 14", true},
		{"21037: 9 7 18 13", false},
		{"292: 11 6 16 20", true},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			c := NewCalibration(test.input)
			addFirstSearch := c.SearchV2('+', 1, c.ops[0])
			multiFirstSearch := c.SearchV2('*', 1, c.ops[0])
			concatFirstSearch := c.SearchV2('|', 1, c.ops[0])

			// If eaither search paths is true then don't fail the test
			if !(addFirstSearch == test.expect || multiFirstSearch == test.expect || concatFirstSearch == test.expect) {
				t.Error("We did not get the expected result")
			}
		})
	}
}

func TestPart1(t *testing.T) {
	input := "190: 10 19\n3267: 81 40 27\n83: 17 5\n156: 15 6\n7290: 6 8 6 15\n161011: 16 10 13\n192: 17 8 14\n21037: 9 7 18 13\n292: 11 6 16 20"
	expectedAnswer := 3749
	expectedCount := 3

	count, answer := part1(input)

	if count != expectedCount {
		t.Errorf("did not get expected count of correct calibrations: got=%d, expected=%d", count, expectedCount)
	}

	if answer != expectedAnswer {
		t.Errorf("did not get expected answer: got=%d, expected=%d", answer, expectedAnswer)
	}
}
