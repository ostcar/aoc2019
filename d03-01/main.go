package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	raw, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Can not read input: %v", err)
	}
	value := Intersect(string(raw))
	fmt.Println(value)
}

// Intersect finds the intersection of the two input lines
func Intersect(input string) int {
	var steps [2]map[[2]int]bool

	wires := parseInput(input)
	for wireIdx, wire := range wires {
		steps[wireIdx] = make(map[[2]int]bool)
		pos := [2]int{0, 0}
		for _, move := range wire {
			for i := range move {
				for move[i] != 0 {
					// s = 1 or -1
					s := move[i] / abs(move[i])
					pos[i] = pos[i] + s
					move[i] = move[i] - s
					steps[wireIdx][pos] = true
				}
			}
		}
	}
	var intersections [][2]int
	for pos := range steps[0] {
		if steps[1][pos] {
			intersections = append(intersections, pos)
		}
	}
	return closest(intersections)
}

func closest(position [][2]int) int {
	min := int(^uint(0) >> 1)
	for _, pos := range position {
		d := abs(pos[0]) + abs(pos[1])
		if d < min {
			min = d
		}
	}
	return min
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func parseInput(input string) [2][][2]int {
	lines := strings.Split(input, "\n")
	var out [2][][2]int
	for idx, line := range lines[:2] {
		out[idx] = make([][2]int, 0)
		commands := strings.Split(line, ",")
		for _, raw := range commands {
			raw = strings.TrimSpace(raw)
			var move [2]int
			switch raw[0] {
			case 'U':
				move[1] = mustInt(raw[1:])
			case 'D':
				move[1] = -mustInt(raw[1:])
			case 'R':
				move[0] = mustInt(raw[1:])
			case 'L':
				move[0] = -mustInt(raw[1:])
			}
			out[idx] = append(out[idx], move)
		}
	}
	return out
}

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
