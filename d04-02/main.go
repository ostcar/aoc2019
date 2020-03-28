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
	count := make(map[int]int)

	for i := 0; i < 6; i++ {
		count[digets[i]]++
		if i < 5 && digets[i] > digets[i+1] {
			return false
		}
	}
	for _, c := range count {
		if c == 2 {
			return true
		}
	}
	return false
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
