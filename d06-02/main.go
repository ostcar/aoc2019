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

	count := StepCount(infos)
	fmt.Println(count)
}

func buildObjects(input []string) map[string]*object {
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
	return objects
}

// StepCount counts the steps between you and san
func StepCount(input []string) int {
	objects := buildObjects(input)
	objects["YOU"].inOrbit.mark()
	objects["SAN"].inOrbit.mark()

	var count int
	for _, object := range objects {
		if object.tagCount == 1 {
			count++
		}
	}
	return count
}

// OrbitCount returns the sum of all direct and indirect orbits.
func OrbitCount(input []string) int {
	objects := buildObjects(input)

	var count int
	for _, object := range objects {
		count += object.count()
	}
	return count
}

type object struct {
	inOrbit  *object
	tagCount int
}

func (o *object) count() int {
	if o.inOrbit == nil {
		return 0
	}
	return 1 + o.inOrbit.count()
}

func (o *object) mark() {
	if o == nil {
		return
	}
	o.tagCount++
	o.inOrbit.mark()
}

func splitObject(s string) (string, string) {
	parts := strings.Split(s, ")")
	if len(parts) != 2 {
		log.Fatalf("Invalid object info %s", s)
	}
	return parts[0], parts[1]
}
