package main

import (
	"fmt"
	"log"
	"sort"
)

var example1 = []*moon{
	&moon{x: -1, y: 0, z: 2},
	&moon{x: 2, y: -10, z: -7},
	&moon{x: 4, y: -8, z: 8},
	&moon{x: 3, y: 5, z: -1},
}

var example2 = []*moon{
	&moon{x: -8, y: -10, z: 0},
	&moon{x: 5, y: 5, z: 10},
	&moon{x: 2, y: -7, z: 3},
	&moon{x: 9, y: -8, z: -3},
}

var real = []*moon{
	&moon{x: 1, y: 2, z: -9},
	&moon{x: -1, y: -9, z: -4},
	&moon{x: 17, y: 6, z: 8},
	&moon{x: 12, y: 4, z: 2},
}

func main() {
	moons := example1
	fmt.Println(findRepeat(moons))
}

func findRepeat(moons []*moon) uint64 {
	var rollup [2][4][3]int
	var start [2][4][3]int
	var found [4][3]int
	var success bool
	for i := 0; i < 10000000; i++ {
		if i == 0 {
			writeRollup(moons, &start)
		}
		// calc
		next(moons)

		// set rollup
		writeRollup(moons, &rollup)

		// set second value
		if i == 0 {
			writeRollup(moons, &start)
			continue
		}

		// compare
		if cmpRollup(&start, &rollup, &found, i) {
			success = true
			break
		}
	}
	if !success {
		log.Fatalf("cound not find all :(")
	}

	fmt.Println(found)

	seen := make(map[uint64]bool)
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			seen[uint64(found[i][j])] = true
		}
	}

	var ints []uint64
	for i := range seen {
		ints = append(ints, i)
	}

	sort.Slice(ints, func(i, j int) bool {
		return ints[i] < ints[j]
	})

	fmt.Println(ints)
	return lcm(ints...)
}

func cmpRollup(r1, r2 *[2][4][3]int, found *[4][3]int, step int) bool {
	var foundCount int
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if found[i][j] != 0 {
				foundCount++
				continue
			}

			if r1[0][i][j] == r2[0][i][j] && r1[1][i][j] == r2[1][i][j] {
				found[i][j] = step
			}
		}
	}
	return foundCount == 4*3
}

func writeRollup(moons []*moon, data *[2][4][3]int) {
	data[0] = data[1]
	for i := 0; i < 4; i++ {
		data[1][i][0] = moons[i].vx
		data[1][i][1] = moons[i].vy
		data[1][i][2] = moons[i].vz
	}
}

func prepare(moons []*moon) {
	for i, m := range moons {
		m.id = i
	}
}

func steps(moons []*moon, nr int) {
	for i := 0; i < nr; i++ {
		next(moons)
	}
}

func next(moons []*moon) {
	pairs(moons, func(m1 *moon, m2 *moon) {
		m1.gravity(m2)
	})
	for _, m := range moons {
		m.move()
	}
}

func energy(moons []*moon) int {
	var count int
	for _, moon := range moons {
		count += moon.energy()
	}
	return count
}

type moon struct {
	id         int
	x, y, z    int
	vx, vy, vz int
}

func (m *moon) String() string {
	return fmt.Sprintf("pos=<x=%3d, y=%3d, z=%3d>, vel=<x=%3d, y=%3d, z=%3d>", m.x, m.y, m.z, m.vx, m.vy, m.vz)
}

func (m *moon) energy() int {
	potential := abs(m.x) + abs(m.y) + abs(m.z)
	kinetic := abs(m.vx) + abs(m.vy) + abs(m.vz)
	return potential * kinetic
}

func (m *moon) move() {
	m.x += m.vx
	m.y += m.vy
	m.z += m.vz
}

func (m *moon) gravity(o *moon) {
	moons, different := sortMoons(mx, m, o)
	if different {
		moons[0].vx++
		moons[1].vx--
	}

	moons, different = sortMoons(my, m, o)
	if different {
		moons[0].vy++
		moons[1].vy--
	}

	moons, different = sortMoons(mz, m, o)
	if different {
		moons[0].vz++
		moons[1].vz--
	}
}

func sortMoons(f func(*moon) int, m ...*moon) ([]*moon, bool) {
	a := f(m[0])
	b := f(m[1])
	if a == b {
		return nil, false
	}
	if a > b {
		m[1], m[0] = m[0], m[1]
	}

	return m, true
}

func mx(m *moon) int { return m.x }
func my(m *moon) int { return m.y }
func mz(m *moon) int { return m.z }

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -1 * i
}

func pairs(moons []*moon, f func(*moon, *moon)) {
	if len(moons) < 2 {
		return
	}

	for _, m := range moons[1:] {
		f(moons[0], m)
	}
	pairs(moons[1:], f)
}

// greatest common divisor (gcd) via Euclidean algorithm
func gcd(a, b uint64) uint64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (lcm) via GCD
func lcm(ints ...uint64) uint64 {
	result := ints[0] * ints[1] / gcd(ints[0], ints[1])

	for i := 2; i < len(ints); i++ {
		result = lcm(result, ints[i])
	}

	return result
}
