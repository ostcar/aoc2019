package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/ostcar/aoc-2019/intcode"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Can not open input: %v", err)
	}
	defer f.Close()

	input := make(chan int, 1)
	output := make(chan int, 1)

	c, err := intcode.NewFromReader(f, intcode.WithInputChan(input), intcode.WithOutputChan(output))
	if err != nil {
		log.Fatalf("Can not start computer: %v", err)
	}

	go c.Run()
	done := make(chan struct{})
	screen := &screen{c: output, done: done}
	go screen.print()

	// 1. is emtpy
	writeString(input, "NOT A J")

	// 3. is emtpy
	writeString(input, "NOT C T")

	// 4. is full
	writeString(input, "AND D T")

	// Jump if (1. is empty) OR (3. is empty AND 4. is full)
	writeString(input, "OR T J")
	writeString(input, "WALK")

	<-done

}

func writeString(c chan<- int, s string) {
	for _, b := range s {
		c <- int(b)
	}
	c <- '\n'
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
			//time.Sleep(pause)
		}
		last = v
		buf.WriteByte(byte(v))
	}
	fmt.Println(last)

	close(s.done)
}
