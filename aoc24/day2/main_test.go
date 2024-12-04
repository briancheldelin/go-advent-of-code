package main

import (
	"testing"
)

type testLevels struct {
	name     string
	level    []int
	expected bool
}

// 7 6 4 2 1
// 1 2 7 8 9
// 9 7 6 2 1
// 1 3 2 4 5
// 8 6 4 4 1
// 1 3 6 7 9

func TestCheckLevel(t *testing.T) {
	testCases := []testLevels{
		{name: "safe lvl 1", level: []int{7, 6, 4, 2, 1}, expected: true},
		{name: "unsafe lvl 2", level: []int{1, 2, 7, 8, 9}, expected: false},
		{name: "unsafe lvl 3", level: []int{9, 7, 6, 2, 1}, expected: false},
		{name: "unsafe lvl 4", level: []int{1, 3, 2, 4, 5}, expected: false},
		{name: "unsafe lvl 5", level: []int{8, 6, 4, 4, 1}, expected: false},
		{name: "safe lvl 6", level: []int{1, 3, 6, 7, 9}, expected: true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := checkLevel(tc.level)
			if result != tc.expected {
				t.Errorf("level safty not expected: name=%s, level=%v, result=%t, expected=%t", tc.name, tc.level, result, tc.expected)
			}
		})
	}
}
