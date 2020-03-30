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

	output := make(chan int, 1)
	computer := intcode.New(string(code), intcode.WithOutputChan(output))

	go computer.Run()

	screen := make(map[position]int)
	for {
		x, ok := <-output
		if !ok {
			break
		}
		y := <-output
		v := <-output
		screen[position{x, y}] = v
	}

	var counter int
	for _, v := range screen {
		if v == 2 {
			counter++
		}
	}
	fmt.Println(counter)
}

type position struct {
	x, y int
}
