package main

import (
	"fmt"
	"testing"
)

func TestIntersect(t *testing.T) {
	for idx, tt := range []struct {
		input    string
		distance int
	}{
		{`R75,D30,R83,U83,L12,D49,R71,U7,L72
		U62,R66,U55,R34,D71,R55,D58,R83`, 159},
		{`R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
		U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`, 135},
	} {
		t.Run(fmt.Sprint(idx), func(t *testing.T) {
			if got := Intersect(tt.input); got != tt.distance {
				t.Errorf("%d. Intersect() got %d, expected %d", idx, got, tt.distance)
			}
		})
	}
}
