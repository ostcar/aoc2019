package main

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	for _, tt := range []struct {
		key   int
		valid bool
	}{
		{111111, false},
		{223450, false},
		{223150, false},
		{123789, false},
		{112233, true},
		{123444, false},
		{111122, true},
	} {
		t.Run(fmt.Sprint(tt.key), func(t *testing.T) {
			if got := IsValid(tt.key); got != tt.valid {
				t.Errorf("IsValid(%d) == %v, expected %v", tt.key, got, tt.valid)
			}
		})
	}
}

func TestSplitInt(t *testing.T) {
	for _, tt := range []struct {
		key    int
		expect []int
	}{
		{111111, ints(1, 1, 1, 1, 1, 1)},
		{223450, ints(2, 2, 3, 4, 5, 0)},
		{123789, ints(1, 2, 3, 7, 8, 9)},
	} {
		t.Run(fmt.Sprint(tt.key), func(t *testing.T) {
			if got := SplitInt(tt.key); !cmpInts(got, tt.expect) {
				t.Errorf("SplitInt(%d) == %v, expected %v", tt.key, got, tt.expect)
			}
		})
	}
}

func cmpInts(i1, i2 []int) bool {
	if len(i1) != len(i2) {
		return false
	}
	for idx := range i1 {
		if i1[idx] != i2[idx] {
			return false
		}
	}
	return true
}

func ints(i ...int) []int {
	return i
}
