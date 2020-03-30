package main

import (
	"fmt"
	"log"
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
	moons := real
	fmt.Println(findRepeat(moons))
}

func findRepeat(moons []*moon) uint64 {
	var found [3]uint64
	var success bool

	start := writeDataBackup(moons)
	for i := 0; i < 10000000; i++ {
		// calc
		next(moons)

		// compare
		if cmpRollup(moons, start, &found, i+1) {
			success = true
			break
		}
	}
	if !success {
		log.Fatalf("cound not find all :(")
	}

	return lcm(found[0], found[1], found[2])
}

func cmpRollup(moons []*moon, start [3][8]int, found *[3]uint64, step int) bool {
	values := writeDataBackup(moons)

	var foundCount int
	for i := 0; i < 3; i++ {
		if found[i] != 0 {
			foundCount++
			continue
		}
		if values[i] == start[i] {
			found[i] = uint64(step)
		}
	}
	return foundCount == 3
}

func writeDataBackup(moons []*moon) [3][8]int {
	var data [3][8]int
	for i := 0; i < 4; i++ {
		data[0][i] = moons[i].x
		data[0][i+4] = moons[i].vx
		data[1][i] = moons[i].y
		data[1][i+4] = moons[i].vy
		data[2][i] = moons[i].z
		data[2][i+4] = moons[i].vz
	}
	return data
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
