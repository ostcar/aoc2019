package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Can not read input: %v", err)
	}
	fmt.Println(TTF(data))
}

// TTF implements the algo
func TTF(input []byte) string {
	ints := bsToInts(input)

	for phaseNr := 0; phaseNr < 100; phaseNr++ {
		newInts := make([]int, len(ints))
		for i := 0; i < len(ints); i++ {
			var sum int
			for j := 0; j < len(ints); j++ {
				sum += ints[j] * pattern(i, j)
			}
			newInts[i] = simplify(sum)
		}
		ints = newInts
	}
	return string(intsToBs(ints)[:8])
}

func simplify(nr int) int {
	if nr < 0 {
		nr *= -1
	}
	return nr % 10
}

func pattern(phaseNr, i int) int {
	pattern := []int{0, 1, 0, -1}
	idx := ((i + 1) / (phaseNr + 1)) % len(pattern)
	return pattern[idx]
}

func bsToInts(data []byte) []int {
	ints := make([]int, len(data))
	for idx := range data {
		i, err := strconv.Atoi(string(data[idx]))
		if err != nil {
			break
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
