package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/ostcar/aoc-2019/intcode"
)

func main() {
	code, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Can not read input: %v", err)
	}
	input := make(chan int)
	output := make(chan int)
	c := intcode.New(string(code), intcode.WithInputChan(input), intcode.WithOutputChan(output))

	go c.Run()

	pos := position{}
	space := make(map[position]int)

	steps := lookAround(space, pos, input, output)
	paint(space, pos)
	fmt.Println(steps)
}

func step(space map[position]int, pos position, input chan<- int, output <-chan int, direction int) (position, bool) {
	oldpos := pos
	switch direction {
	case dUp:
		pos.y--
	case dRight:
		pos.x++
	case dDown:
		pos.y++
	case dLeft:
		pos.x--
	default:
		log.Fatalf("Unkown direction %d", direction)
	}
	input <- direction

	var found bool
	switch <-output {
	case 0:
		space[pos] = tWall
		pos = oldpos
	case 1:
		space[pos] = tFree
	case 2:
		space[pos] = tOxygen
		found = true
	default:
		log.Fatalf("Unknown status")
	}
	return pos, found
}

func lookDirection(space map[position]int, pos position, input chan<- int, output <-chan int, lookPos position, direction int, vsDirection int) int {
	origPos := pos
	if space[lookPos] == tUnknown {
		var found bool
		pos, found = step(space, pos, input, output, direction)

		if pos != origPos {
			var c int
			if !found {
				c = lookAround(space, pos, input, output)
				if c == 0 {
					c = -1
				}
			}
			step(space, pos, input, output, vsDirection)
			return c + 1
		}
	}
	return 0
}

func lookAround(space map[position]int, pos position, input chan<- int, output <-chan int) int {
	steps := lookDirection(space, pos, input, output, position{pos.x, pos.y - 1}, dUp, dDown)
	if steps > 0 {
		return steps
	}
	steps = lookDirection(space, pos, input, output, position{pos.x, pos.y + 1}, dDown, dUp)
	if steps > 0 {
		return steps
	}
	steps = lookDirection(space, pos, input, output, position{pos.x - 1, pos.y}, dLeft, dRight)
	if steps > 0 {
		return steps
	}
	return lookDirection(space, pos, input, output, position{pos.x + 1, pos.y}, dRight, dLeft)
}

func paint(space map[position]int, my position) {
	minPos, maxPos := getMinMax(space)
	width := maxPos.x - minPos.x
	heigh := maxPos.y - minPos.y

	var buf strings.Builder
	for i := 0; i < heigh; i++ {
		for j := 0; j < width; j++ {
			pos := position{x: minPos.x + j, y: minPos.y + i}

			if pos == my {
				buf.WriteString("x")
				continue
			}

			v := space[pos]
			switch v {
			case tUnknown:
				buf.WriteString("▒")
			case tWall:
				buf.WriteString("█")
			case tFree:
				buf.WriteString(" ")
			case tOxygen:
				buf.WriteString("O")
			default:
				log.Fatalf("Invalid value %d", v)
			}
		}
		buf.WriteByte('\n')
	}
	fmt.Println(buf.String())
}

func getMinMax(screen map[position]int) (position, position) {
	minPos := position{x: 100, y: 100}
	maxPos := position{x: -100, y: -100}
	for pos := range screen {
		if pos.x < minPos.x {
			minPos.x = pos.x
		}
		if pos.x > maxPos.x {
			maxPos.x = pos.x
		}
		if pos.y < minPos.y {
			minPos.y = pos.y
		}
		if pos.y > maxPos.y {
			maxPos.y = pos.y
		}
	}
	minPos.x -= 4
	minPos.y -= 4
	maxPos.x += 4
	maxPos.y += 4
	return minPos, maxPos
}

const (
	dUp = iota + 1
	dDown
	dLeft
	dRight
)

const (
	tUnknown = iota
	tFree
	tWall
	tOxygen
)

type position struct {
	x, y int
}
