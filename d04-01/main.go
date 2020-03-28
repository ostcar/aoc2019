package main

import "fmt"

func main() {
	var count int
	for i := 240920; i < 789857; i++ {
		count += toInt(IsValid(i))
	}
	fmt.Println(count)
}

// IsValid returns, if the given number is a possible key
func IsValid(i int) bool {
	digets := SplitInt(i)

	// two are the same
	same := false
	for i := 0; i < 5; i++ {
		if digets[i] == digets[i+1] {
			same = true
		}
		if digets[i] > digets[i+1] {
			return false
		}
	}
	if !same {
		return false
	}
	return digets[4] <= digets[5]
}

// SplitInt returns a list of digets for an int
func SplitInt(number int) []int {
	digets := make([]int, 6)

	for i := 0; i < 6; i++ {
		digets[5-i] = number % 10
		number /= 10
	}
	return digets
}

func toInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
