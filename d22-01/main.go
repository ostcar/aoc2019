package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open input: %v", err)
	}
	defer f.Close()

	deck := initDeck(10007)

	if err := applyShuffle(deck, f); err != nil {
		log.Fatalf("Can not apply shuffle: %v", err)
	}

	for i := 0; i < len(deck); i++ {
		if deck[i] == 2019 {
			fmt.Println(i)
			return
		}
	}
}

func applyShuffle(deck []int, r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var i int

		_, err := fmt.Sscanf(line, "cut %d", &i)
		if err == nil {
			cut(deck, i)
			continue
		}

		_, err = fmt.Sscanf(line, "deal with increment %d", &i)
		if err == nil {
			increment(deck, i)
			continue
		}

		if line == "deal into new stack" {
			reverse(deck)
			continue
		}

		log.Fatalf("Invalid line: %s", line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func initDeck(l int) []int {
	deck := make([]int, l)
	for i := 0; i < l; i++ {
		deck[i] = i
	}
	return deck
}

func reverse(deck []int) {
	l := len(deck)
	for i := 0; i < l/2; i++ {
		deck[i], deck[l-1-i] = deck[l-1-i], deck[i]
	}
}

func cut(deck []int, c int) {
	l := len(deck)
	if c < 0 {
		c = (c + l) % l
	}

	t := make([]int, c)
	copy(t, deck)
	for i := 0; i < len(deck)-c; i++ {
		deck[i] = deck[i+c]
	}
	for i := 0; i < c; i++ {
		deck[i+l-c] = t[i]
	}

}

func increment(deck []int, n int) {
	nDeck := make([]int, len(deck))

	for i := 0; i < len(deck); i++ {
		ni := (i * n) % len(deck)
		nDeck[ni] = deck[i]
	}
	copy(deck, nDeck)
}
