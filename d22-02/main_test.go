package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func TestReverse(t *testing.T) {
	for _, tt := range []struct {
		len    int
		value  int
		expect int
	}{
		{10, 3, 6},
	} {
		t.Run(fmt.Sprintf("%d %d", tt.len, tt.value), func(t *testing.T) {

			if got := reverse(tt.value, tt.len); got != tt.expect {
				t.Errorf("reverse() returnd %d, expected %d", got, tt.expect)
			}
		})
	}
}

func TestCut(t *testing.T) {
	for _, tt := range []struct {
		len    int
		index  int
		cut    int
		expect int
	}{
		{10, 3, 3, 0},
		{10, 7, 3, 4},
		{10, 1, 3, 8},
		{10, 3, -4, 7},
		{10, 7, -4, 1},
		{10, 1, -4, 5},
	} {
		t.Run(fmt.Sprintf("%d %d", tt.len, tt.index), func(t *testing.T) {
			if got := cut(tt.index, tt.len, tt.cut); got != tt.expect {
				t.Errorf("cut() returnd %d, expected %d", got, tt.expect)
			}
		})
	}
}

func TestIncrement(t *testing.T) {
	for _, tt := range []struct {
		len    int
		index  int
		inc    int
		expect int
	}{
		{10, 3, 3, 9},
		{10, 7, 3, 1},
		{10, 1, 3, 3},
	} {
		t.Run(fmt.Sprintf("%d %d", tt.len, tt.index), func(t *testing.T) {
			if got := increment(tt.index, tt.len, tt.inc); got != tt.expect {
				t.Errorf("increment() returnd %d, expected %d", got, tt.expect)
			}
		})
	}
}

func TestApplyShuffle(t *testing.T) {
	for _, tt := range []struct {
		len          int
		index        int
		instructions string
		expect       int
	}{
		{
			10,
			3,
			`
			deal with increment 7
			deal into new stack
			deal into new stack`,
			1,
		},
		{
			10,
			3,
			`
			cut 6
			deal with increment 7
			deal into new stack`,
			0,
		},
		{
			10,
			3,
			`
			deal with increment 7
			deal with increment 9
			cut -2`,
			1,
		},
		{
			10,
			3,
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
			8,
		},
	} {
		t.Run(fmt.Sprintf("%d", tt.len), func(t *testing.T) {
			instructions, err := readInstructions(strings.NewReader(tt.instructions), tt.len)
			if err != nil {
				t.Errorf("Can not read instructions: %v", err)
			}

			value := applyShuffle(tt.index, instructions)

			if value != tt.expect {
				t.Errorf("applyShulle returned %d, expected %d", value, tt.expect)
			}
		})
	}
}

func BenchmarkApplyShuffle(b *testing.B) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open input: %v", err)
	}
	defer f.Close()

	instructions, err := readInstructions(f, deckLen)
	if err != nil {
		log.Fatalf("Can not read instructions: %v", err)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		applyShuffle(2020, instructions)
	}
}
