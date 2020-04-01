package main

import (
	"testing"
)

func TestRepeated(t *testing.T) {
	for _, tt := range []struct {
		input    []int
		count    int
		expected string
	}{
		{[]int{1, 2, 3}, 3, "123123123"},
	} {
		t.Run(tt.expected, func(t *testing.T) {
			if got := string(intsToBs(repeated(tt.input, tt.count))); got != tt.expected {
				t.Errorf("repated(%v, %d) == %s, expected %s", tt.input, tt.count, got, tt.expected)
			}
		})
	}
}
