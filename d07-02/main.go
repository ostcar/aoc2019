package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"

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
	return Perm([]int{5, 6, 7, 8, 9}, func(setting []int) int {
		var ch [5]chan int
		for i := 0; i < 5; i++ {
			ch[i] = make(chan int, 1)
			ch[i] <- setting[i]
		}

		var wg sync.WaitGroup
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func(i int) {
				oChan := (i + 1) % 5
				c := intcode.New(
					code,
					intcode.WithInputChan(ch[i]),
					intcode.WithOutputChan(ch[oChan]),
					intcode.WithName(fmt.Sprintf("C%d", i)),
				)
				c.Run()
				wg.Done()
			}(i)
		}
		ch[0] <- 0

		wg.Wait()
		return <-ch[0]
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
