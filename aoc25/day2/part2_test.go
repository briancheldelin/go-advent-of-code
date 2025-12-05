package day2

import (
	"fmt"
	"slices"
	"testing"
)

func Test_isRepeatRecursive(t *testing.T) {
	tests := []struct {
		number string
		want   bool
	}{
		{number: "11", want: true},
		{number: "22", want: true},
		{number: "99", want: true},
		{number: "111", want: true},
		{number: "999", want: true},
		{number: "1010", want: true},
		{number: "1188511885", want: true},
		{number: "222222", want: true},
		{number: "446446", want: true},
		{number: "38593859", want: true},
		{number: "565656", want: true},
		{number: "824824824", want: true},
		{number: "2121212121", want: true},
		{number: "1238", want: false},
		{number: "1698528", want: false},
		{number: "12", want: false},
		{number: "110", want: false},
		{number: "115", want: false},
		{number: "1011", want: false},
		{number: "446447", want: false},
		{number: "446448", want: false},
		{number: "446449", want: false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("number %s", tt.number), func(t *testing.T) {
			got := testNumeber(tt.number)
			if got != tt.want {
				t.Errorf("isRepeatRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_testRange(t *testing.T) {
	tests := []struct {
		start     int
		end       int
		wantSum   int
		wantCount int
		wantIDs   []int
	}{
		{start: 1, end: 21, wantSum: 11, wantCount: 1, wantIDs: []int{11}},
		{start: 11, end: 22, wantSum: 33, wantCount: 2, wantIDs: []int{11, 22}},
		{start: 95, end: 115, wantSum: 210, wantCount: 2, wantIDs: []int{99, 111}},
		{start: 998, end: 1012, wantSum: 2009, wantCount: 2, wantIDs: []int{999, 1010}},
		{start: 222220, end: 222224, wantSum: 222222, wantCount: 1, wantIDs: []int{222222}},
		{start: 446443, end: 446449, wantSum: 446446, wantCount: 1, wantIDs: []int{446446}},
		{start: 565653, end: 565659, wantSum: 565656, wantCount: 1, wantIDs: []int{565656}},
		{start: 1698522, end: 1698528, wantSum: 0, wantCount: 0, wantIDs: []int{}},
		{start: 38593856, end: 38593862, wantSum: 38593859, wantCount: 1, wantIDs: []int{38593859}},
		{start: 824824821, end: 824824827, wantSum: 824824824, wantCount: 1, wantIDs: []int{824824824}},
		{start: 1188511880, end: 1188511890, wantSum: 1188511885, wantCount: 1, wantIDs: []int{1188511885}},
		{start: 2121212118, end: 2121212124, wantSum: 2121212121, wantCount: 1, wantIDs: []int{2121212121}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Range %d_%d", tt.start, tt.end), func(t *testing.T) {
			gotSum, gotCount, gotIDs := testRange(tt.start, tt.end)
			if gotSum != tt.wantSum {
				t.Errorf("testRange() sum = %v, want %v", gotSum, tt.wantSum)
			}
			if gotCount != tt.wantCount {
				t.Errorf("testRange() count = %v, want %v", gotCount, tt.wantCount)
			}
			if !slices.Equal(gotIDs, tt.wantIDs) {
				t.Errorf("testRange() ids = %v, want %v", gotIDs, tt.wantIDs)
			}
		})
	}
}
