package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestShootNr(t *testing.T) {
	starmap := `.#..##.###...#######
				##.############..##.
				.#.######.########.#
				.###.#######.####.#.
				#####.##.#.##.###.##
				..#####..#.#########
				####################
				#.####....###.#.#.##
				##.#################
				#####.##.###..####..
				..######..##.#######
				####.##.####...##..#
				.#####..#.######.###
				##...#.##########...
				#.##########.#######
				.####.#.###.###.#.##
				....##.##.###..#####
				.#.#.###########.###
				#.#.#.#####.####.###
				###.##.####.##.#..##`

	// starmap = ` .#..#
	// 			.....
	// 			#####
	// 			....#
	// 			...##`

	stars := newStarList(strings.NewReader(starmap))
	if got := shootNr(stars, 1); got != 1112 {
		t.Errorf("shootNr(%d) = %d, expected %d", 1, got, 1112)
	}
}

func TestNewStarList(t *testing.T) {
	for idx, tt := range []struct {
		starMap string
		expect  []*star
	}{
		{`......#.#.
		#..#.#....`,
			bStars(6, 0, 8, 0, 0, 1, 3, 1, 5, 1),
		},
	} {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			if got := newStarList(strings.NewReader(tt.starMap)); !cmpStars(got, tt.expect) {
				t.Errorf("Got %d stars, expected %d", len(got), len(tt.expect))
			}
		})
	}
}

func TestFindBest(t *testing.T) {
	for idx, tt := range []struct {
		starMap string
		expect  [3]int
	}{
		{
			`.#..#
			 .....
			 #####
			 ....#
			 ...##`,
			[3]int{3, 4, 8},
		},
		{
			`......#.#.
			#..#.#....
			..#######.
			.#.#.###..
			.#..#.....
			..#....#.#
			#..#....#.
			.##.#..###
			##...#..#.
			.#....####`,
			[3]int{5, 8, 33},
		},
		{
			`#.#...#.#.
			.###....#.
			.#....#...
			##.#.#.#.#
			....#.#.#.
			.##..###.#
			..#...##..
			..##....##
			......#...
			.####.###.`,
			[3]int{1, 2, 35},
		},
		{
			`.#..#..###
			####.###.#
			....###.#.
			..###.##.#
			##.##.#.#.
			....###..#
			..#.#..#.#
			#..#.#.###
			.##...##.#
			.....#.#..`,
			[3]int{6, 3, 41},
		},
		{
			`.#..##.###...#######
			##.############..##.
			.#.######.########.#
			.###.#######.####.#.
			#####.##.#.##.###.##
			..#####..#.#########
			####################
			#.####....###.#.#.##
			##.#################
			#####.##.###..####..
			..######..##.#######
			####.##.####...##..#
			.#####..#.######.###
			##...#.##########...
			#.##########.#######
			.####.#.###.###.#.##
			....##.##.###..#####
			.#.#.###########.###
			#.#.#.#####.####.###
			###.##.####.##.#..##`,
			[3]int{11, 13, 210},
		},
	} {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			stars := newStarList(strings.NewReader(tt.starMap))
			if _, got := findBest(stars); got == nil || got.x != tt.expect[0] || got.y != tt.expect[1] {
				t.Errorf("Got %v, expected star(x: %d, y: %d, hidden: %d)", got, tt.expect[0], tt.expect[1], len(stars)-tt.expect[2])
			}
		})
	}
}

func TestIsHidden(t *testing.T) {
	for idx, tt := range []struct {
		s1, s2, s3 *star
		expect     bool
	}{
		{
			&star{x: 0, y: 0},
			&star{x: 1, y: 1},
			&star{x: 2, y: 2},
			true,
		},
		{
			&star{x: 0, y: 0},
			&star{x: 1, y: 1},
			&star{x: 2, y: 1},
			false,
		},
		{
			&star{x: 1, y: 0},
			&star{x: 3, y: 2},
			&star{x: 4, y: 3},
			true,
		},
	} {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			if got := isHidden(tt.s1, tt.s2, tt.s3); got != tt.expect {
				t.Errorf("isHidden(%v, %v, %v) == %v, expected %v", tt.s1, tt.s2, tt.s3, got, tt.expect)
			}
		})
	}
}

func bStars(s ...int) []*star {
	var stars []*star
	for i := 0; i < len(s); i += 2 {
		stars = append(stars, &star{x: s[i], y: s[i+1]})
	}
	return stars
}

func cmpStars(s1, s2 []*star) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i].x != s2[i].x {
			return false
		}
		if s1[i].y != s2[i].y {
			return false
		}
	}
	return true
}
