package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ostcar/aoc-2019/intcode"
)

const width = 58
const camLen = 58*49 + 1 //2843
const pause = 100 * time.Millisecond
const videoFeed = true

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Can not read input: %v", err)
	}

	input := make(chan int, 1)
	output := make(chan int, 1)

	c, err := intcode.NewFromReader(f, intcode.WithInputChan(input), intcode.WithOutputChan(output))
	if err != nil {
		log.Fatal(err)
	}
	c.Manipulate(0, 2)

	go c.Run()

	data := getCam(output)
	fmt.Println(getPath(data))

	done := make(chan struct{})
	screen := &screen{c: output, done: done}
	go screen.print()

	writeString(input, "A,B,A,C,A,A,C,B,C,B")
	writeString(input, "L,12,L,8,R,12")
	writeString(input, "L,10,L,8,L,12,R,12")
	writeString(input, "R,12,L,8,L,10")
	if videoFeed {
		writeString(input, "y")
	} else {
		writeString(input, "n")
	}

	<-screen.done
}

func getPath(data []byte) string {
	var pos int
	for i, v := range data {
		if v == '^' {
			pos = i
			break
		}
	}
	buf := new(bytes.Buffer)
	buf.WriteString("L")
	direction := 3
	var count int
	for {
		if next, ok := nextPos(pos, direction); ok && data[next] == '#' {
			count++
			pos = next
			continue
		}
		var b byte
		direction, b = nextDirection(data, pos, direction)
		if direction == -1 {
			fmt.Fprintf(buf, ",%d", count)
			break
		}
		fmt.Fprintf(buf, ",%d,%s", count, string(b))
		count = 0
	}
	return buf.String()
}

func inRange(pos int) bool {
	return pos >= 0 && pos <= camLen
}

func nextPos(pos int, direction int) (int, bool) {
	var next int
	switch direction {
	case 0:
		next = pos - width
	case 1:
		next = pos + 1
	case 2:
		next = pos + width
	case 3:
		next = pos - 1
	default:
		log.Fatalf("Invalid direction")
	}
	return next, inRange(next)
}

func nextDirection(data []byte, pos int, direction int) (int, byte) {
	r := (direction + 1) % 4
	if next, ok := nextPos(pos, r); ok && data[next] == '#' {
		return r, 'R'
	}

	l := (4 + direction - 1) % 4
	if next, ok := nextPos(pos, l); ok && data[next] == '#' {
		return l, 'L'
	}
	return -1, 0
}

func writeString(c chan<- int, s string) {
	for _, b := range s {
		c <- int(b)
	}
	c <- '\n'
}

func getCam(c <-chan int) []byte {
	var bs []byte
	var count int
	for {
		v := <-c
		count++
		if count == camLen {
			break
		}
		bs = append(bs, byte(v))
	}
	return bs
}

type screen struct {
	c    <-chan int
	done chan struct{}
}

func (s *screen) print() {
	var last int
	buf := new(bytes.Buffer)
	for v := range s.c {
		if v == '\n' && last == '\n' {
			fmt.Print(buf.String())
			buf.Reset()
			time.Sleep(pause)
		}
		last = v
		buf.WriteByte(byte(v))
	}
	fmt.Println(last)

	close(s.done)
}
