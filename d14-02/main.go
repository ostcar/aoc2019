package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const firstSteps = 1000
const oreCount = 1000000000000
const fastSteps = 1850

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Can not open input: %v", err)
	}
	defer f.Close()

	fmt.Println(calcMaxFuel(f))
}

func calcMaxFuel(r io.Reader) int {
	reaktions := readReactions(r)
	buffer := make(map[string]int)
	buffer["ORE"] = oreCount

	for i := 0; i < firstSteps; i++ {
		calcOres(reaktions["FUEL"], reaktions, buffer)
	}

	calcOresFast(buffer, fastSteps)

	count := firstSteps + firstSteps*fastSteps
	for {
		if !calcOres(reaktions["FUEL"], reaktions, buffer) {
			break
		}
		count++
		if count%10000 == 0 {
			fmt.Println(buffer["ORE"])
			fmt.Println(buffer)
		}
	}
	fmt.Println()
	fmt.Println(buffer)
	return count
}

func calcOresFast(buffer map[string]int, amount int) {
	for key := range buffer {
		if key == "ORE" {
			used := oreCount - buffer[key]
			buffer[key] -= used * amount
			continue
		}
		buffer[key] += buffer[key] * amount
	}
}

// calcOres returns the required ores for an reaction
func calcOres(r *reaktion, reaktions map[string]*reaktion, buffer map[string]int) bool {
	for _, input := range r.input {
		if buffer[input.name] >= input.amount {
			buffer[input.name] -= input.amount
			continue
		}

		if input.name == "ORE" {
			return false
		}

		needed := input.amount - buffer[input.name]
		for needed > 0 {
			if !calcOres(reaktions[input.name], reaktions, buffer) {
				return false
			}
			needed -= reaktions[input.name].output.amount
		}
		buffer[input.name] = -needed
	}
	return true
}

func calcFuel(r io.Reader) (int, map[string]int) {
	reaktions := readReactions(r)
	buffer := make(map[string]int)
	v := calcOresOld(reaktions["FUEL"], reaktions, buffer)
	return v, buffer
}

// calcOres returns the required ores for an reaction
func calcOresOld(r *reaktion, reaktions map[string]*reaktion, buffer map[string]int) int {
	if len(r.input) == 1 && r.input[0].name == "ORE" {
		return r.input[0].amount
	}

	var count int
	for _, input := range r.input {
		if buffer[input.name] >= input.amount {
			buffer[input.name] -= input.amount
			continue
		}

		needed := input.amount - buffer[input.name]
		for needed > 0 {
			count += calcOresOld(reaktions[input.name], reaktions, buffer)
			needed -= reaktions[input.name].output.amount
		}
		buffer[input.name] = -needed
	}
	return count
}

func readReactions(r io.Reader) map[string]*reaktion {
	scanner := bufio.NewScanner(r)
	reaktions := make(map[string]*reaktion)
	for scanner.Scan() {
		r := newReaktion(scanner.Text())
		reaktions[r.output.name] = r

	}
	if scanner.Err() != nil {
		log.Fatalf("Can not read from input: %v", scanner.Err())
	}
	return reaktions
}

type typeAmount struct {
	name   string
	amount int
}

func newTypeAmount(s string) typeAmount {
	var ta typeAmount
	v := strings.Split(strings.TrimSpace(s), " ")
	a, err := strconv.Atoi(strings.TrimSpace(v[0]))
	if err != nil {
		log.Fatalf("can not convert typeAmout value `%s`", s)
	}
	ta.amount = a
	ta.name = strings.TrimSpace(v[1])
	return ta
}

func (s typeAmount) String() string {
	return fmt.Sprintf("%d %s", s.amount, s.name)
}

type reaktion struct {
	output typeAmount
	input  []typeAmount
}

func newReaktion(s string) *reaktion {
	r := reaktion{}
	inputOutout := strings.Split(s, "=>")

	inputElements := strings.Split(inputOutout[0], ",")
	for _, e := range inputElements {
		r.input = append(r.input, newTypeAmount(e))
	}
	r.output = newTypeAmount(inputOutout[1])
	return &r
}

func (r *reaktion) String() string {
	buf := strings.Builder{}

	var inputs []string
	for _, input := range r.input {
		inputs = append(inputs, input.String())
	}
	buf.WriteString(strings.Join(inputs, ", "))
	buf.WriteString(" => ")
	buf.WriteString(r.output.String())
	buf.WriteByte('\n')
	return buf.String()
}
