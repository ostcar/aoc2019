package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"testing"
)

func TestReverse(t *testing.T) {
	for _, tt := range []struct {
		len    int64
		value  int64
		expect int64
	}{
		{10, 3, 6},
	} {
		t.Run(fmt.Sprintf("%d %d", tt.len, tt.value), func(t *testing.T) {
			len := big.NewInt(tt.len)
			value := big.NewInt(tt.value)
			expect := big.NewInt(tt.expect)

			reverse(value, len)
			if value.Cmp(expect) != 0 {
				t.Errorf("reverse() returnd %d, expected %d", value, expect)
			}
		})
	}
}

func TestCut(t *testing.T) {
	for _, tt := range []struct {
		len    int64
		value  int64
		cut    int64
		expect int64
	}{
		{10, 3, 3, 0},
		{10, 7, 3, 4},
		{10, 1, 3, 8},
		{10, 3, -4, 7},
		{10, 7, -4, 1},
		{10, 1, -4, 5},
	} {
		t.Run(fmt.Sprintf("%d %d", tt.len, tt.value), func(t *testing.T) {
			len := big.NewInt(tt.len)
			value := big.NewInt(tt.value)
			cutV := big.NewInt(tt.cut)
			expect := big.NewInt(tt.expect)

			cut(value, len, cutV)
			if value.Cmp(expect) != 0 {
				t.Errorf("cut() returnd %d, expected %d", value, expect)
			}
		})
	}
}

func TestIncrement(t *testing.T) {
	for _, tt := range []struct {
		len    int64
		value  int64
		inc    int64
		expect int64
	}{
		{10, 3, 3, 9},
		{10, 7, 3, 1},
		{10, 1, 3, 3},
	} {
		t.Run(fmt.Sprintf("%d %d", tt.len, tt.value), func(t *testing.T) {
			len := big.NewInt(tt.len)
			value := big.NewInt(tt.value)
			inc := big.NewInt(tt.inc)
			expect := big.NewInt(tt.expect)

			increment(value, len, inc)
			if value.Cmp(expect) != 0 {
				t.Errorf("increment() returnd %d, expected %d", value, expect)
			}
		})
	}
}

func TestMulti(t *testing.T) {
	deckLen := big.NewInt(10)
	rawInst := `
	deal into new stack
	cut -2
	deal with increment 7
	cut 8
	cut -4
	deal with increment 7
	cut 3
	deal with increment 9
	deal with increment 3
	cut -1`
	startValue := int64(3)

	instructions, err := readInstructions(strings.NewReader(rawInst), deckLen)
	if err != nil {
		t.Errorf("Can not read instructions: %v", err)
	}

	v1 := big.NewInt(startValue)
	for i := 0; i < 101; i++ {
		applyShuffle(v1, instructions, deckLen)
	}

	v2 := big.NewInt(startValue)
	instructions = multi(instructions, 101, deckLen)
	applyShuffle(v2, instructions, deckLen)

	if v1.Cmp(v2) != 0 {
		t.Errorf("multi retuend different result :(")
	}

}

func TestApplyShuffle(t *testing.T) {
	for _, tt := range []struct {
		len          int64
		value        int64
		instructions string
		expect       int64
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
			len := big.NewInt(tt.len)
			value := big.NewInt(tt.value)
			expect := big.NewInt(tt.expect)

			instructions, err := readInstructions(strings.NewReader(tt.instructions), len)
			if err != nil {
				t.Errorf("Can not read instructions: %v", err)
			}

			applyShuffle(value, instructions, len)

			if value.Cmp(expect) != 0 {
				t.Errorf("applyShulle returned %d, expected %d", value, expect)
			}
		})
	}
}

func BenchmarkApplyShuffle(b *testing.B) {
	f, err := os.Open("input.txt")
	if err != nil {
		b.Fatalf("can not open input: %v", err)
	}
	defer f.Close()

	deckLen := big.NewInt(deckLenV)
	value := big.NewInt(2020)

	instructions, err := readInstructions(f, deckLen)
	if err != nil {
		log.Fatalf("Can not read instructions: %v", err)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		applyShuffle(value, instructions, deckLen)
	}
}
