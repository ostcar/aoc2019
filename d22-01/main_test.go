package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestCut(t *testing.T) {
	for _, tt := range []struct {
		len    int
		cut    int
		expect []int
	}{
		{10, 3, ints(3, 4, 5, 6, 7, 8, 9, 0, 1, 2)},
		{10, -4, ints(6, 7, 8, 9, 0, 1, 2, 3, 4, 5)},
	} {
		t.Run(fmt.Sprintf("%d %d", tt.len, tt.cut), func(t *testing.T) {
			deck := initDeck(tt.len)
			if cut(deck, tt.cut); !cmpInts(deck, tt.expect) {
				t.Errorf("cut returnd %v, expected %v", deck, tt.expect)
			}
		})
	}
}

func TestIncrement(t *testing.T) {
	for _, tt := range []struct {
		len    int
		n      int
		expect []int
	}{
		{10, 3, ints(0, 7, 4, 1, 8, 5, 2, 9, 6, 3)},
	} {
		t.Run(fmt.Sprintf("%d %d", tt.len, tt.n), func(t *testing.T) {
			deck := initDeck(tt.len)
			if increment(deck, tt.n); !cmpInts(deck, tt.expect) {
				t.Errorf("increment returnd %v, expected %v", deck, tt.expect)
			}
		})
	}
}

func TestApplyShuffle(t *testing.T) {
	for _, tt := range []struct {
		len          int
		instructions string
		expect       []int
	}{
		{
			10,
			`
			deal with increment 7
			deal into new stack
			deal into new stack`,
			ints(0, 3, 6, 9, 2, 5, 8, 1, 4, 7),
		},
		{
			10,
			`
			cut 6
			deal with increment 7
			deal into new stack`,
			ints(3, 0, 7, 4, 1, 8, 5, 2, 9, 6),
		},
		{
			10,
			`
			deal with increment 7
			deal with increment 9
			cut -2`,
			ints(6, 3, 0, 7, 4, 1, 8, 5, 2, 9),
		},
		{
			10,
			`
			deal into new stack
			cut -2
			deal with increment 7
			cut 8
			cut -4
			deal with increment 7
			cut 3
			deal with increment 9
			deal with increment 3
			cut -1`,
			ints(9, 2, 5, 8, 1, 4, 7, 0, 3, 6),
		},
	} {
		t.Run(fmt.Sprintf("%d", tt.len), func(t *testing.T) {
			deck := initDeck(tt.len)
			if applyShuffle(deck, strings.NewReader(tt.instructions)); !cmpInts(deck, tt.expect) {
				t.Errorf("applyShullfe returnd %v, expected %v", deck, tt.expect)
			}
		})
	}
}

func cmpInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func ints(i ...int) []int {
	return i
}
