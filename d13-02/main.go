package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/ostcar/aoc-2019/intcode"
)

func main() {
	code, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Can not read input: %v", err)
	}

	output := make(chan int, 1)
	joystick := make(chan int)
	computer := intcode.New(string(code), intcode.WithInputChan(joystick), intcode.WithOutputChan(output))

	s := new(screen)

	computer.Manipulate(0, 2)
	go computer.Run()

	go func() {
		for {
			ballPaddleDiff := s.ballPaddleDiff()

			var direction int
			switch {
			case ballPaddleDiff < 0:
				direction = -1
			case ballPaddleDiff > 0:
				direction = 1
			default:
				direction = 0
			}

			joystick <- direction
			time.Sleep(time.Millisecond)
			fmt.Println(s)
		}
	}()

	var score int
	for {
		x, ok := <-output
		if !ok {
			break
		}

		y := <-output
		v := <-output
		if x == -1 {
			score = v
			continue
		}
		s.set(x, y, v)
	}
	fmt.Println(score)
}

type screen struct {
	mu   sync.Mutex
	data map[position]int
}

func (s *screen) set(x, y, v int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.data == nil {
		s.data = make(map[position]int)
	}
	s.data[position{x, y}] = v
}

func (s *screen) ballPaddleDiff() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	var ballX int
	var paddleX int

	for pos, v := range s.data {
		if v == vBall {
			ballX = pos.x
		}
		if v == vPaddle {
			paddleX = pos.x
		}
	}
	return ballX - paddleX
}

func (s *screen) String() string {
	s.mu.Lock()
	defer s.mu.Unlock()

	minPos, maxPos := getMinMax(s.data)
	width := maxPos.x - minPos.x
	heigh := maxPos.y - minPos.y

	var buf strings.Builder
	for i := 0; i < heigh; i++ {
		for j := 0; j < width; j++ {
			pos := position{x: minPos.x + j, y: minPos.y + i}

			v := s.data[pos]
			switch v {
			case vEmpty:
				buf.WriteString(" ")
			case vWall:
				buf.WriteString("█")
			case vBlock:
				buf.WriteString("▬")
			case vPaddle:
				buf.WriteString("_")
			case vBall:
				buf.WriteString("○")
			default:
				log.Fatalf("Invalid value %d", v)
			}
		}
		buf.WriteByte('\n')
	}
	return buf.String()
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

type position struct {
	x, y int
}

const (
	vEmpty = iota
	vWall
	vBlock
	vPaddle
	vBall
)
