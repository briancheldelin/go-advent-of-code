package day3

import (
	"strconv"
	"testing"
)

func Test_joltage(t *testing.T) {
	tests := []struct {
		name      string
		bank      string
		wantJolts int
	}{
		{name: "example basic", bank: "98", wantJolts: 98},
		{name: "example line 1", bank: "987654321111111", wantJolts: 98},
		{name: "example line 2", bank: "811111111111119", wantJolts: 89},
		{name: "example line 3", bank: "234234234234278", wantJolts: 78},
		{name: "example line 4", bank: "818181911112111", wantJolts: 92},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := joltageTwoDigit(tt.bank)
			if got != tt.wantJolts {
				t.Errorf("joltage() bank = %v, want %v", got, tt.wantJolts)
			}
		})
	}
}

func Test_joltageTwelveDigit(t *testing.T) {
	tests := []struct {
		name      string
		bank      string
		wantJolts int
	}{
		{name: "example basic", bank: "123456789123", wantJolts: 123456789123},
		{name: "example line 1", bank: "987654321111111", wantJolts: 987654321111},
		{name: "example line 2", bank: "811111111111119", wantJolts: 811111111119},
		{name: "example line 3", bank: "234234234234278", wantJolts: 434234234278},
		{name: "example line 4", bank: "818181911112111", wantJolts: 888911112111},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bankInt, _ := strconv.Atoi(tt.bank)
			got := joltageTwelveDigit(bankInt, len(tt.bank))
			if got != tt.wantJolts {
				t.Errorf("joltageTwelveDigit() = %v, want %v", got, tt.wantJolts)
			}
		})
	}
}

func Test_digit(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		bank   int
		index  int
		length int
		want   int
	}{
		// {name: "Example 12", bank: 812345612345, index: 0, length: 12, want: 800000000000},
		// {name: "Example 11", bank: 812345612345, index: 1, length: 12, want: 10000000000},
		// {name: "Example 10", bank: 812345612345, index: 2, length: 12, want: 2000000000},
		// {name: "Example 9", bank: 810345612345, index: 2, length: 12, want: 0},
		{name: "Example longer", bank: 81374561234544, index: 2, length: 14, want: 300000000000},
		// {name: "Example 8", bank: 812345612345, index: 8, length: 12, want: 800000000000},
		// {name: "Example 7", bank: 812345612345, index: 7, length: 12, want: 800000000000},
		// {name: "Example 6", bank: 812345612345, index: 6, length: 12, want: 800000000000},
		// {name: "Example 5", bank: 812345612345, index: 5, length: 12, want: 800000000000},
		// {name: "Example 4", bank: 812345612345, index: 4, length: 12, want: 800000000000},
		// {name: "Example 3", bank: 812345612345, index: 3, length: 12, want: 800000000000},
		// {name: "Example 2", bank: 812345612345, index: 2, length: 12, want: 800000000000},
		// {name: "Example 1", bank: 812345612345, index: 1, length: 12, want: 800000000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := digit(tt.bank, tt.index, tt.length, 12)
			if got != tt.want {
				t.Errorf("digit() = %v, want %v", got, tt.want)
			}
		})
	}
}
