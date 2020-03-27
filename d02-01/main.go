package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	exit  = 99
	add   = 1
	multi = 2
)

func main() {
	raw, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Can not read input: %v", err)
	}
	content := strings.TrimSpace(string(raw))
	ints, err := readInts(content)
	if err != nil {
		log.Fatalf("Can not read ints: %v", err)
	}
	ints[1] = 12
	ints[2] = 2
	IntCode(ints)
	fmt.Println(ints[0])
}

func readInts(s string) ([]int, error) {
	vals := strings.Split(s, ",")
	ints := make([]int, len(vals))
	var err error
	for i, v := range vals {
		ints[i], err = strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("can not convert int from %s: %w", v, err)
		}
	}
	return ints, nil
}

// IntCode runs the code, manipulatig it.
func IntCode(code []int) {
	var pos int
	for code[pos] != exit {
		p1 := code[pos+1]
		p2 := code[pos+2]
		p3 := code[pos+3]
		switch code[pos] {
		case add:
			code[p3] = code[p1] + code[p2]
		case multi:
			code[p3] = code[p1] * code[p2]
		}
		pos += 4
	}
}
