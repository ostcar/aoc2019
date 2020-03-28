package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var infos []string
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Can not open input file: %v", err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		infos = append(infos, strings.TrimSpace(scanner.Text()))
	}
	if scanner.Err() != nil {
		log.Fatalf("Can not scan input: %v", scanner.Err())
	}

	count := OrbitCount(infos)
	fmt.Println(count)
}

// OrbitCount returns the sum of all direct and indirect orbits.
func OrbitCount(input []string) int {
	objects := make(map[string]*object)
	for _, info := range input {
		o1, o2 := splitObject(info)

		// create objects, if they not exist
		if _, ok := objects[o1]; !ok {
			objects[o1] = &object{}
		}
		if _, ok := objects[o2]; !ok {
			objects[o2] = &object{}
		}

		objects[o2].inOrbit = objects[o1]
	}

	var count int
	for _, object := range objects {
		count += object.count()
	}
	return count
}

type object struct {
	inOrbit *object
}

func (o *object) count() int {
	if o.inOrbit == nil {
		return 0
	}
	return 1 + o.inOrbit.count()
}

func splitObject(s string) (string, string) {
	parts := strings.Split(s, ")")
	if len(parts) != 2 {
		log.Fatalf("Invalid object info %s", s)
	}
	return parts[0], parts[1]
}
