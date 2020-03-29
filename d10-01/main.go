package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Can not open input: %v", err)
	}
	stars := newStarList(f)
	best := findBest(stars)
	fmt.Println(len(stars) - best.hiddenCount() - 1)
}

func findBest(stars []*star) *star {
	min := stars[0]
	for _, star := range stars[1:] {
		if star.hiddenCount() < min.hiddenCount() {
			min = star
		}
	}
	return min
}

type star struct {
	x, y   int
	hidden map[[2]int]bool
}

func (s *star) String() string {
	return fmt.Sprintf("star{x: %d y: %d hidden: %d)", s.x, s.y, s.hiddenCount())
}

func (s *star) hiddenCount() int {
	if s == nil {
		return -1
	}
	return len(s.hidden)
}

func (s *star) setHidden(other *star) {
	if s.hidden == nil {
		s.hidden = make(map[[2]int]bool)
	}
	s.hidden[[2]int{other.x, other.y}] = true
}

func newStarList(r io.Reader) []*star {
	var stars []*star
	scanner := bufio.NewScanner(r)
	for lineNr := 0; scanner.Scan(); lineNr++ {
		line := bytes.TrimSpace(scanner.Bytes())
		for ColumnNr, v := range line {
			if v == '#' {
				stars = append(stars, &star{x: ColumnNr, y: lineNr})
			}
		}
	}
	setHiddenList(stars)
	return stars
}

func setHiddenList(stars []*star) {
	for i, s1 := range stars[:len(stars)-2] {
		for j, s2 := range stars[i+1 : len(stars)-1] {
			for _, s3 := range stars[i+j+2 : len(stars)] {
				if isHidden(s1, s2, s3) {
					s1.setHidden(s3)
					s3.setHidden(s1)
				}
			}
		}
	}
}

func isHidden(s1, s2, s3 *star) bool {
	vector1_2 := float32(s2.x-s1.x) / float32(s2.y-s1.y)
	vector1_3 := float32(s3.x-s1.x) / float32(s3.y-s1.y)
	return floatEqual(vector1_2, vector1_3)
}

func floatEqual(f1, f2 float32) bool {
	return f1 == f2
}
