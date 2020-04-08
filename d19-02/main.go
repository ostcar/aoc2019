package main

import (
	"fmt"
)

// Smalest x on y:       100 =       114
//                     1_000 =     1_137
//                    10_000 =    11_365
//                   100_000 =   113_650
//                 1_000_000 = 1_136_498
//
// Highest x on y        100 =       144
//                     1_000 =     1_442
//                    10_000 =    14_420
//                   100_000 =   144_205
//                 1_000_000 = 1_442_052

func main() {

	x := 0
	y := 99
	for {
		if !inBeam(x, y) {
			x++
			continue
		}

		if !inBeam(x+99, y-99) {
			y++
			continue
		}
		break
	}
	fmt.Println(x*10_000 + y - 99)
}

func f1(x int) int {
	return int(float64(x) * 1_136_497.5 / 1_000_000)
}

func f2(x int) int {
	return int(float64(x) * 1_442_052 / 1_000_000)
}

func inBeam(x, y int) bool {
	v1 := y > f1(x)
	v2 := y <= f2(x)
	return v1 && v2
}
