package main

import (
	"fmt"
	"testing"
)

func TestOrbitCount(t *testing.T) {
	for _, tt := range []struct {
		input    []string
		expected int
	}{
		{strs("COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"), 42},
	} {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			if got := OrbitCount(tt.input); got != tt.expected {
				t.Errorf("OrbitCount(%v) == %v, expected %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestStepCount(t *testing.T) {
	for _, tt := range []struct {
		input    []string
		expected int
	}{
		{strs("COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"), 4},
	} {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			if got := StepCount(tt.input); got != tt.expected {
				t.Errorf("StepCount(%v) == %v, expected %v", tt.input, got, tt.expected)
			}
		})
	}

}

func strs(s ...string) []string {
	return s
}
