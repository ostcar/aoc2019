package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ostcar/aoc-2019/intcode"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Can not open input: %v", err)
	}
	defer f.Close()

	code, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("Can not read code: %v", err)
	}

	var count int
	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {
			output := make(chan int)
			c := intcode.New(string(code), intcode.WithInput(i, j), intcode.WithOutputChan(output))
			go c.Run()
			v := <-output
			count += v
		}
	}
	fmt.Println(count)
}
