package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Can not open input: %v", err)
	}
	stars := newStarList(f)
	fmt.Println(shootNr(stars, 200))
}

func shootNr(stars []*star, nr int) int {
	idx, best := findBest(stars)

	// Remove best from stars
	stars[idx] = stars[len(stars)-1]
	stars = stars[:len(stars)-1]

	// sort stars by angle
	sort.Slice(stars, func(i, j int) bool {
		return best.angle(stars[i]) < best.angle(stars[j])
	})

	var counter int
	for _, star := range stars {
		if star.hidden[[2]int{best.x, best.y}] {
			continue
		}
		counter++
		if counter == nr {
			return star.x*100 + star.y
		}
	}
	return -1
}

func findBest(stars []*star) (int, *star) {
	min := stars[0]
	var minI int
	for i, star := range stars[1:] {
		if star.hiddenCount() < min.hiddenCount() {
			min = star
			minI = i
		}
	}
	return minI, min
}

type star struct {
	x, y   int
	hidden map[[2]int]bool
}

func (s *star) String() string {
	return fmt.Sprintf("star{x: %d y: %d hidden: %d)", s.x, s.y, s.hiddenCount())
}

func (s *star) angle(other *star) float64 {
	x := float64(other.x - s.x)
	y := float64(other.y - s.y)
	c := math.Sqrt(x*x + y*y)
	asin := math.Asin(x / c)
	if y > 0 {
		asin = math.Pi - asin
	}
	asin = math.Mod(asin+math.Pi*2, math.Pi*2)
	return asin
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
