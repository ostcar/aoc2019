package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ostcar/aoc-2019/intcode"
)

func main() {
	code, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Can not read input: %v", err)
	}
	trust := MaxSetting(string(code))
	fmt.Println(trust)
}

// MaxSetting finds the best setting and returns the output.
func MaxSetting(code string) int {
	return Perm([]int{0, 1, 2, 3, 4}, func(setting []int) int {
		var input int
		for i := 0; i < 5; i++ {
			c := intcode.New(code, intcode.WithInput(setting[i], input))
			input = c.Run()[0]

		}
		return input
	})
}

// Perm calls f with all permutations of a
func Perm(a []int, f func([]int) int) int {
	return perm(a, f, 0)
}

func perm(a []int, f func([]int) int, i int) int {
	if i > len(a) {
		return f(a)
	}
	max := perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		v := perm(a, f, i+1)
		if v > max {
			max = v
		}
		a[i], a[j] = a[j], a[i]
	}
	return max
}
