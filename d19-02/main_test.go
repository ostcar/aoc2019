package main

import (
	"fmt"
	"testing"
)

func TestInBeam(t *testing.T) {
	for _, tt := range []struct {
		x, y   int
		expect bool
	}{
		{10, 12, true},
		{1_000_000, 1_136_498, true},
		{1_000_000, 1_136_497, false},
	} {
		t.Run(fmt.Sprintf("%d, %d", tt.x, tt.y), func(t *testing.T) {
			if got := inBeam(tt.x, tt.y); got != tt.expect {
				t.Errorf("inBeam(%d, %d) == %t, expected %t", tt.x, tt.y, got, tt.expect)
			}
		})
	}
}
