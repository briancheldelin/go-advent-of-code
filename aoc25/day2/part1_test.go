package day2

import (
	"testing"
)

func Test_isRepeat(t *testing.T) {
	tests := []struct {
		number string
		want   bool
	}{
		{number: "1010", want: true},
		{number: "22", want: true},
		{number: "11", want: true},
		{number: "1188511885", want: true},
		{number: "38593859", want: true},
		{number: "446446", want: true},
		{number: "222222", want: true},
		{number: "1238", want: false},
		{number: "1698528", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.number, func(t *testing.T) {
			got := isRepeat(tt.number)
			if got != tt.want {
				t.Errorf("isMirror() = %v, want %v", got, tt.want)
			}
		})
	}
}
