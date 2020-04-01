package main

import (
	"testing"
)

func TestFFT(t *testing.T) {
	for _, tt := range []struct {
		input    []byte
		expected string
	}{
		{[]byte("12345678"), "23845678"},
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

func BenchmarkTTF(b *testing.B) {
	ints := bsToInts([]byte("1234567890"))

	for n := 0; n < b.N; n++ {
		TTF(ints)
	}

}
