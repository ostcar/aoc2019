package main

import "fmt"

func main() {
	moons := []*moon{
		&moon{x: -1, y: 0, z: 2},
		&moon{x: 2, y: -10, z: -7},
		&moon{x: 4, y: -8, z: 8},
		&moon{x: 3, y: 5, z: -1},
	}
	// moons := []*moon{
	// 	&moon{x: 1, y: 2, z: -9},
	// 	&moon{x: -1, y: -9, z: -4},
	// 	&moon{x: 17, y: 6, z: 8},
	// 	&moon{x: 12, y: 4, z: 2},
	// }
	start := make([]moon, len(moons))
	for i := range moons {
		start[i] = *moons[i]
	}

	var count int
	for count < 10000000 {
		next(moons)
		count++
		if equal(start, moons) {
			fmt.Println("Fount!!")
			break
		}
	}
	fmt.Println(count)
}

func equal(moons1 []moon, moons2 []*moon) bool {
	for i := range moons1 {
		if moons1[i] != *moons2[i] {
			return false
		}
	}
	return true
}

func steps(moons []*moon, nr int) {
	for i := 0; i < nr; i++ {
		//fmt.Printf("After %d steps:\n", i)
		next(moons)
	}
}

func next(moons []*moon) {
	pairs(moons, func(m1 *moon, m2 *moon) {
		m1.gravity(m2)
	})
	for _, m := range moons {
		m.move()
		//fmt.Println(m)
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
	return fmt.Sprintf("pos=<x=%2d, y=%2d, z=%2d>, vel=<x=%2d, y=%2d, z=%2d>", m.x, m.y, m.z, m.vx, m.vy, m.vz)
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
	moons, different := sort(mx, m, o)
	if different {
		moons[0].vx++
		moons[1].vx--
	}

	moons, different = sort(my, m, o)
	if different {
		moons[0].vy++
		moons[1].vy--
	}

	moons, different = sort(mz, m, o)
	if different {
		moons[0].vz++
		moons[1].vz--
	}
}

func sort(f func(*moon) int, m ...*moon) ([]*moon, bool) {
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
