package intcode

import (
	"fmt"
	"strconv"
	"testing"
)

func TestComputer(t *testing.T) {
	for _, tt := range []struct {
		code     string
		input    int
		expected []int
	}{
		{"1,0,0,0,4,0,99", 0, ints(2)},
		{"01,0,0,0,4,0,99", 0, ints(2)},
		{"101,0,7,0,4,0,99,1", 0, ints(1)},
		{"2,3,0,0,4,0,99", 0, ints(0)},
		{"2,6,6,0,4,0,99,0", 0, ints(9801)},
		{"1,1,1,4,99,5,6,0,4,0,99", 0, ints(30)},
		{"3,9,8,9,10,9,4,9,99,-1,8", 8, ints(1)},
		{"3,9,8,9,10,9,4,9,99,-1,8", 7, ints(0)},
		{"3,9,7,9,10,9,4,9,99,-1,8", 7, ints(1)},
		{"3,9,7,9,10,9,4,9,99,-1,8", 8, ints(0)},
		{"3,3,1108,-1,8,3,4,3,99", 8, ints(1)},
		{"3,3,1108,-1,8,3,4,3,99", 7, ints(0)},
		{"3,3,1107,-1,8,3,4,3,99", 7, ints(1)},
		{"3,3,1107,-1,8,3,4,3,99", 8, ints(0)},
		{"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", 0, ints(0)},
		{"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", 1, ints(1)},
		{"3,3,1105,-1,9,1101,0,0,12,4,12,99,1", 0, ints(0)},
		{"3,3,1105,-1,9,1101,0,0,12,4,12,99,1", 1, ints(1)},
		{`3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,
		1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,
		999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99`, 7, ints(999)},
		{`3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,
		1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,
		999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99`, 8, ints(1000)},
		{`3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,
		1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,
		999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99`, 9, ints(1001)},
		{"109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99", 0, ints(109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99)},
		{"1102,34915192,34915192,7,4,7,99,0", 0, ints(1219070632396864)},
		{"104,1125899906842624,99", 0, ints(1125899906842624)},
	} {
		t.Run(fmt.Sprintf("%s(%d)", tt.code, tt.input), func(t *testing.T) {
			c := New(tt.code, WithInput(tt.input))
			if got := c.Run(); !cmpInts(got, tt.expected) {
				t.Errorf("Program %s returned %d, expected %d", tt.code, got, tt.expected)
			}
		})
	}
}

func TestPow10(t *testing.T) {
	for _, tt := range []struct {
		nr       int
		expected int
	}{
		{0, 1},
		{1, 10},
		{2, 100},
	} {
		t.Run(strconv.Itoa(tt.nr), func(t *testing.T) {
			if got := pow10(tt.nr); got != tt.expected {
				t.Errorf("pow10(%d)==%d, expected %d", tt.nr, got, tt.expected)
			}
		})
	}
}

func TestJoinInt(t *testing.T) {
	for _, tt := range []struct {
		ints     []int
		expected int
	}{
		{ints(1, 0), 10},
		{ints(1), 1},
		{ints(1, 2, 3), 123},
	} {
		t.Run(fmt.Sprint(tt.ints), func(t *testing.T) {
			if got := joinInt(tt.ints); got != tt.expected {
				t.Errorf("joinInt(%v)==%d, expected %d", tt.ints, got, tt.expected)
			}
		})
	}
}

func TestSplitInt(t *testing.T) {
	for _, tt := range []struct {
		nr       int
		expected []int
	}{
		{10, ints(1, 0)},
		{1, ints(1)},
		{123, ints(1, 2, 3)},
	} {
		t.Run(strconv.Itoa(tt.nr), func(t *testing.T) {
			if got := splitInt(tt.nr); !cmpInts(got, tt.expected) {
				t.Errorf("splitInt(%d)==%v, expected %v", tt.nr, got, tt.expected)
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
