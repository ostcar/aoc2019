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

	var image [size]byte
	for i := range image {
		image[i] = '2'
	}

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

		merge(&image, buf)
	}
	draw(image)
}

func merge(image *[size]byte, new []byte) {
	for i := range image {
		if image[i] == '2' {
			image[i] = new[i]
		}
	}
}

func draw(image [size]byte) {
	for i := 0; i < 6; i++ {
		for j := 0; j < 25; j++ {
			switch image[25*i+j] {
			case '0':
				fmt.Print(" ")
			case '1':
				fmt.Print("▇")
			}
		}
		fmt.Println()
	}
}
