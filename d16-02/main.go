package main

import (
	"bytes"
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
	ints := bsToInts(bytes.TrimSpace(data))
	fmt.Println(TTF(repeated(ints)))
}

// TTF implements the algo
func TTF(ints []int) string {
	newInts := make([]int, len(ints))
	for phaseNr := 0; phaseNr < 100; phaseNr++ {
		for i := 0; i < len(ints); i++ {
			var sum int
			for j := 0; j < len(ints); j++ {
				idx := ((j + 1) / (i + 1)) % 4
				if idx == 0 || idx == 2 {
					continue
				}
				if idx == 1 {
					sum += ints[j]
				} else {
					sum -= ints[j]
				}
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

const repeatedCount = 10000

func repeated(ints []int) []int {
	r := make([]int, len(ints)*repeatedCount)
	for i := 0; i < len(ints); i++ {
		v := ints[i]
		for j := 0; j < repeatedCount; j++ {
			idx := j + i*repeatedCount
			r[idx] = v
		}
	}
	return r
}
