package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestFuelCalc(t *testing.T) {
	for _, tt := range []struct {
		mass int
		fuel int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	} {
		t.Run(strconv.Itoa(tt.mass), func(t *testing.T) {
			if got := FuelCalc(tt.mass); got != tt.fuel {
				t.Errorf("FuelCalc(%d)==%d, expected %d", tt.mass, got, tt.fuel)
			}
		})
	}
}

func TestCalc(t *testing.T) {
	input := strings.NewReader(`14
1969
100756`)
	expect := strconv.Itoa(2 + 966 + 50346)
	if got := Calc(input); got != expect {
		t.Errorf("Calc()==%s, expected %s", got, expect)
	}
}
