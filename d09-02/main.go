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
		log.Fatalf("Can not open input: %v", err)
	}
	c := intcode.New(string(code), intcode.WithInput(2))
	fmt.Println(c.Run())
}
