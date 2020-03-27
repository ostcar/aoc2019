package main

import (
	"errors"
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

const expect = 19690720

func main() {
	raw, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Can not read input: %v", err)
	}
	content := strings.TrimSpace(string(raw))
	noun, verb, err := findParameters(content)
	if err != nil {
		log.Fatalf("Can not find parameters: %v", err)
	}
	fmt.Printf("noun: %d, verb: %d = %d\n", noun, verb, 100*noun+verb)
}

func findParameters(rawCode string) (int, int, error) {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			ints, err := readInts(rawCode)
			if err != nil {
				return 0, 0, err
			}

			ints[1] = noun
			ints[2] = verb
			IntCode(ints)
			if ints[0] == expect {
				return noun, verb, nil
			}
		}
	}
	return 0, 0, errors.New("no valid inputs")
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
