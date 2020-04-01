package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

/*
I could not solve it and used this solution:
https://github.com/Gravitar64/Advent-of-Code-2019/blob/master/AoC_Tag%2016a.py
*/
func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Can not read input: %v", err)
	}

	offset, err := strconv.Atoi(string(data[:7]))
	if err != nil {
		log.Fatal(err)
	}

	// Create ints
	ints := bsToInts(bytes.TrimSpace(data))
	ints = repeated(ints, 10_000)
	ints = ints[offset:]

	for i := 0; i < 100; i++ {
		ps := sum(ints)
		for j := 0; j < len(ints); j++ {
			t := ps
			ps -= ints[j]
			ints[j] = t % 10
		}
	}
	fmt.Println(string(intsToBs(ints[:8])))

}

func sum(ints []int) int {
	var count int
	for _, v := range ints {
		count += v
	}
	return count
}

func bsToInts(data []byte) []int {
	ints := make([]int, len(data))
	for idx := range data {
		i, err := strconv.Atoi(string(data[idx]))
		if err != nil {
			log.Fatalf("Invalid char")
		}
		ints[idx] = i
	}
	return ints
}

func intsToBs(ints []int) []byte {
	bs := make([]byte, len(ints))
	for idx := range ints {
		bs[idx] = strconv.Itoa(ints[idx])[0]
	}
	return bs
}

func repeated(ints []int, count int) []int {
	r := make([]int, 0, len(ints)*count)
	for i := 0; i < count; i++ {
		r = append(r, ints...)

	}
	return r
}
