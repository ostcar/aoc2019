package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const size = 25 * 6

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Can not read input file: %v", err)
	}

	var min [3]int
	min[0] = int(^uint(0) >> 1)

	buf := make([]byte, size)
	for {
		n, err := f.Read(buf)
		if n != size {
			break
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Can not read image: %v", err)
		}
		numbers := count(buf)
		if numbers[0] < min[0] {
			min = numbers
		}
	}

	value := min[1] * min[2]
	fmt.Println(value)
}

func count(s []byte) [3]int {
	var numbers [3]int
	for _, v := range s {
		switch v {
		case '0':
			numbers[0]++
		case '1':
			numbers[1]++
		case '2':
			numbers[2]++
		default:
			log.Println(v)
		}
	}
	return numbers
}
