package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ostcar/aoc-2019/intcode"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Can not read input: %v", err)
	}

	output := make(chan int)

	c, err := intcode.NewFromReader(f, intcode.WithOutputChan(output))
	if err != nil {
		log.Fatal(err)
	}

	go c.Run()

	var ints []int
	var width int
	for v := range output {
		ints = append(ints, v)
		if v == '\n' && width == 0 {
			width = len(ints)
		}
	}

	var count int
	for i := width; i < len(ints)-width-1; i++ {
		if i%width == 0 || i%width == width-2 || i%width == width-1 {
			continue
		}
		if ints[i] == '#' && ints[i+1] == '#' && ints[i-1] == '#' && ints[i-width] == '#' && ints[i+width] == '#' {
			ints[i] = 'O'
			count += (i / width) * (i % width)
		}
	}

	for _, v := range ints {
		fmt.Print(string(v))
	}
	fmt.Println(count)
}
