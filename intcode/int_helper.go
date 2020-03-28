package intcode

import (
	"strconv"
	"strings"
)

// getOp returns the last to elements of the input slice as op and the arguments left to right.
func splitOpArgs(ints []int) (int, []int) {
	if len(ints) == 0 {
		return 0, nil
	}
	if len(ints) == 1 {
		return ints[0], nil
	}
	args := ints[:len(ints)-2]
	op := joinInt(ints[len(ints)-2:])

	// Reverse
	for i := len(args)/2 - 1; i >= 0; i-- {
		opp := len(args) - 1 - i
		args[i], args[opp] = args[opp], args[i]
	}
	return op, args
}

// splitInt returns a list of digets for an int.
func splitInt(number int) []int {
	var digets []int
	for number > 0 {
		digets = append(digets, number%10)
		number /= 10
	}

	// Reverse
	for i := len(digets)/2 - 1; i >= 0; i-- {
		opp := len(digets) - 1 - i
		digets[i], digets[opp] = digets[opp], digets[i]
	}
	return digets
}

func joinInt(ints []int) int {
	var number int
	for i, diget := range ints {
		number += pow10(len(ints)-i-1) * diget
	}
	return number
}

func pow10(nr int) int {
	out := 1
	for i := 0; i < nr; i++ {
		out *= 10
	}
	return out
}

func mustInt(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return i
}

// get returns the value of the index or 0.
func get(ints []int, idx int) int {
	if len(ints) <= idx {
		return 0
	}
	return ints[idx]
}
