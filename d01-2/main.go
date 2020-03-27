package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Can not read input: %v", err)
	}

	fmt.Println(Calc(f))
}

// Calc does the requested calculation
func Calc(r io.Reader) string {
	s := bufio.NewScanner(r)
	var sum int
	for s.Scan() {
		mass, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatalf("Can not convert %s to int: %v", s.Text(), err)
		}
		sum += FuelCalc(mass)
	}
	if err := s.Err(); err != nil {
		log.Fatalf("Can not scan text: %v", err)
	}
	return strconv.Itoa(sum)
}

// FuelCalc gets the required fuel for a mass
func FuelCalc(mass int) int {
	step := mass
	var sum int
	for {
		step = step/3 - 2
		if step <= 0 {
			break
		}
		sum += step
	}
	return sum
}
