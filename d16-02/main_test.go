package main

import (
	"fmt"
	"testing"
)

func TestFFT(t *testing.T) {
	for _, tt := range []struct {
		input    []byte
		expected string
	}{
		{[]byte("80871224585914546619083218645595"), "24176176"},
		{[]byte("19617804207202209144916044189917"), "73745418"},
		{[]byte("69317163492948606335995924319873"), "52432133"},
	} {
		t.Run(string(tt.input), func(t *testing.T) {
			ints := bsToInts(tt.input)
			if got := TTF(ints); got != tt.expected {
				t.Errorf("TTF(%s) == %s, expected %s", tt.input, got, tt.expected)
			}
		})
	}
}

func TestPattern(t *testing.T) {
	for _, tt := range []struct {
		nr, i    int
		expected int
	}{
		{0, 0, 1},
		{0, 1, 0},
		{0, 6, -1},
		{0, 7, 0},
		{1, 0, 0},
		{1, 1, 1},
		{1, 6, -1},
		{1, 7, 0},
		{7, 0, 0},
		{7, 1, 0},
		{7, 6, 0},
		{7, 7, 1},
	} {
		t.Run(fmt.Sprintf("pattern(%d,%d)", tt.nr, tt.i), func(t *testing.T) {
			if got := pattern(tt.nr, tt.i); got != tt.expected {
				t.Errorf("pattern(%d, %d) == %d, expected %d", tt.nr, tt.i, got, tt.expected)
			}
		})
	}
}

func BenchmarkTTF(b *testing.B) {
	ints := bsToInts([]byte("80871224585914546619083218645595"))

	for n := 0; n < b.N; n++ {
		TTF(ints)
	}

}
