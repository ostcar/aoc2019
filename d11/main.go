package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/ostcar/aoc-2019/intcode"
)

const firstColor = colorWhite

func main() {
	code, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Can not read input file: %v", err)
	}
	cam := make(chan int, 1)
	command := make(chan int, 1)
	rs := &robotSpace{
		cam:     cam,
		command: command,
		space:   map[position]int{position{0, 0}: firstColor},
	}
	p := intcode.New(string(code), intcode.WithInputChan(cam), intcode.WithOutputChan(command))

	go p.Run()

	for rs.next() {

	}

	var count int
	for range rs.space {
		count++
	}
	fmt.Println(rs)
	fmt.Println(count)
}

const (
	dUp = iota
	dLeft
	dDown
	dRight
)

const (
	colorBlack = iota
	colorWhite
)

const colorDefaul = colorBlack

type position struct {
	x, y int
}

type robotSpace struct {
	direction int
	pos       position
	space     map[position]int
	cam       chan<- int
	command   <-chan int
}

func (rs *robotSpace) next() bool {
	rs.cam <- rs.read()

	color, ok := <-rs.command
	if !ok {
		return false
	}
	rs.paint(color)
	command, ok := <-rs.command
	if !ok {
		log.Fatalf("Only one command :(")
	}
	rs.turn(command)
	return true
}

func (rs *robotSpace) turn(direction int) {
	directions := map[int][2]int{
		dUp:    [2]int{0, -1},
		dDown:  [2]int{0, 1},
		dRight: [2]int{-1, 0},
		dLeft:  [2]int{1, 0},
	}
	switch direction {
	case 0:
		rs.direction = (4 + rs.direction - 1) % 4
	case 1:
		rs.direction = (rs.direction + 1) % 4
	default:
		log.Fatalf("Unknown direction %d", direction)
	}
	move := directions[rs.direction]
	rs.pos.x += move[0]
	rs.pos.y += move[1]
}

func (rs *robotSpace) read() int {
	return rs.readAt(rs.pos)
}

func (rs *robotSpace) readAt(pos position) int {
	if rs.space == nil {
		rs.space = make(map[position]int)
		return colorDefaul
	}

	color, ok := rs.space[pos]
	if !ok {
		return colorDefaul
	}
	return color
}

func (rs *robotSpace) paint(color int) {
	if rs.space == nil {
		rs.space = make(map[position]int)
	}

	rs.space[rs.pos] = color
}

func (rs *robotSpace) String() string {
	minPos := position{x: 100, y: 100}
	maxPos := position{x: -100, y: -100}
	for pos := range rs.space {
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

	width := maxPos.x - minPos.x
	heigh := maxPos.y - minPos.y
	space := make([][]int, heigh)

	var buf strings.Builder
	for i := 0; i < heigh; i++ {
		space[i] = make([]int, width)
		for j := 0; j < width; j++ {
			pos := position{x: minPos.x + j, y: minPos.y + i}
			color := rs.readAt(pos)
			switch color {
			case colorBlack:
				buf.WriteString("░")
			case colorWhite:
				buf.WriteString("█")
			default:
				log.Fatalf("123 Invalid color %d", color)
			}
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}
