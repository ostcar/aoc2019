package main

import (
	"fmt"
	"testing"
)

func TestIntCode(t *testing.T) {
	for idx, tt := range []struct {
		code   []int
		result []int
	}{
		{ints(1, 0, 0, 0, 99), ints(2, 0, 0, 0, 99)},
		{ints(2, 3, 0, 3, 99), ints(2, 3, 0, 6, 99)},
		{ints(2, 4, 4, 5, 99, 0), ints(2, 4, 4, 5, 99, 9801)},
		{ints(1, 1, 1, 4, 99, 5, 6, 0, 99), ints(30, 1, 1, 4, 2, 5, 6, 0, 99)},
	} {
		t.Run(fmt.Sprint(idx), func(t *testing.T) {
			if IntCode(tt.code); !cmpInts(tt.code, tt.result) {
				t.Errorf("%d. IntCode() got %v, expected %v", idx, tt.code, tt.result)
			}
		})
	}
}

func TestReadInts(t *testing.T) {
	for _, tt := range []struct {
		input  string
		result []int
	}{
		{"1,2,3", ints(1, 2, 3)},
		{"1,1,1", ints(1, 1, 1)},
	} {
		t.Run(tt.input, func(t *testing.T) {
			got, err := readInts(tt.input)
			if err != nil {
				t.Errorf("readInts(%s) returned the unexpected error %v", tt.input, err)
			}
			if !cmpInts(got, tt.result) {
				t.Errorf("readInts(%s)==%v, expected %v", tt.input, got, tt.result)
			}
		})
	}
}

func ints(i ...int) []int {
	return i
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
