package main

import "testing"

func TestSteps(t *testing.T) {
	moons := []*moon{
		&moon{x: -1, y: 0, z: 2},
		&moon{x: 2, y: -10, z: -7},
		&moon{x: 4, y: -8, z: 8},
		&moon{x: 3, y: 5, z: -1},
	}
	steps(moons, 100)
	if got := energy(moons); got != 179 {
		t.Errorf("Got energy %d, expected 179", got)
	}
	t.Error()
}

func BenchmarkNext(b *testing.B) {
	moons := []*moon{
		&moon{x: -1, y: 0, z: 2},
		&moon{x: 2, y: -10, z: -7},
		&moon{x: 4, y: -8, z: 8},
		&moon{x: 3, y: 5, z: -1},
	}

	for i := 0; i < b.N; i++ {
		next(moons)
	}
}
