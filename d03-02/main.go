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
	var steps [2]map[[2]int]int

	wires := parseInput(input)
	for wireIdx, wire := range wires {
		steps[wireIdx] = make(map[[2]int]int)
		pos := [2]int{0, 0}
		var count int
		for _, move := range wire {
			for i := range move {
				for move[i] != 0 {
					// s = 1 or -1
					s := move[i] / abs(move[i])
					pos[i] = pos[i] + s
					move[i] = move[i] - s
					count++
					if _, ok := steps[wireIdx][pos]; !ok {
						steps[wireIdx][pos] = count
					}
				}
			}
		}
	}
	var intersections []int
	for pos, count0 := range steps[0] {
		if count1, ok := steps[1][pos]; ok {
			intersections = append(intersections, count0+count1)
		}
	}
	return closest(intersections)
}

func closest(steps []int) int {
	min := int(^uint(0) >> 1)
	for _, count := range steps {
		if count < min {
			min = count
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
