package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"strings"
)

const (
	deckLen      = 119_315_717_514_047
	shuffleCount = 101_741_582_076_661
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open input: %v", err)
	}
	defer f.Close()

	instructions, err := readInstructions(f, deckLen)
	if err != nil {
		log.Fatalf("Can not read instructions: %v", err)
	}

	value := 2020

	// for i := 0; i < shuffleCount; i++ {
	// 	if i%1_000_000 == 0 {
	// 		fmt.Println(i)
	// 	}
	// 	value = applyShuffle(value, instructions)
	// }
	value = applyShuffle(value, instructions)
	fmt.Println(value)
	value = applyShuffle(value, instructions)
	fmt.Println(value)
}

func removeReverse(instr []instruction) []instruction {
	var nInstr []instruction
	var inReverse bool
	for _, inst := range instr {
		if inst.iType == iReverse {
			inReverse = !inReverse
			continue
		}

		if !inReverse {
			nInstr = append(nInstr, inst)
			continue
		}

		if inst.iType == iCut {
			inst.value *= -1
			nInstr = append(nInstr, inst)
			continue
		}

		if inst.iType == iIncrement {
			nInstr = append(nInstr, inst)
			nInstr = append(nInstr, instruction{iType: iCut, value: -(inst.value - 1)})
			continue
		}

		panic("You should not be here")
	}
	if inReverse {
		nInstr = append(nInstr, instruction{iType: iReverse})
	}
	return nInstr
}

func moveIncrement(instr []instruction, deckLen int) []instruction {
	//var nInstr []instruction
	for i := 0; i < len(instr); i++ {
		//Switch cut with increment
		if i != len(instr)-1 && instr[i].iType == iCut && instr[i+1].iType == iIncrement {
			instr[i].value = (instr[i].value * instr[i+1].value) % deckLen
			instr[i], instr[i+1] = instr[i+1], instr[i]
			i = -1
			continue
		}

		// Merge double cut
		if i != len(instr)-1 && instr[i].iType == iCut && instr[i+1].iType == iCut {
			value := (deckLen + instr[i].value + instr[i+1].value) % deckLen
			instr[i] = instruction{iType: iCut, value: value}
			instr = append(instr[:i+1], instr[i+2:]...)
			i--
			continue
		}

		//Merge double increment
		if i != len(instr)-1 && instr[i].iType == iIncrement && instr[i+1].iType == iIncrement {
			value := (instr[i].value * instr[i+1].value) % deckLen
			instr[i] = instruction{iType: iIncrement, value: value}
			instr = append(instr[:i+1], instr[i+2:]...)
			i--
			continue
		}
	}
	return instr
}

func readInstructions(r io.Reader, deckLen int) ([]instructionF, error) {
	var instructions []instruction
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		var i int

		_, err := fmt.Sscanf(line, "cut %d", &i)
		if err == nil {
			instructions = append(instructions, instruction{iType: iCut, value: i})
			continue
		}

		_, err = fmt.Sscanf(line, "deal with increment %d", &i)
		if err == nil {
			instructions = append(instructions, instruction{iType: iIncrement, value: i})
			continue
		}

		if line == "deal into new stack" {
			instructions = append(instructions, instruction{iType: iReverse})
			continue
		}

		return nil, fmt.Errorf("invalid line: %s", line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	instructions = removeReverse(instructions)
	instructions = moveIncrement(instructions, deckLen)
	printInstructions(instructions)
	return instFuncs(instructions, deckLen), nil
}

func applyShuffle(value int, instr []instructionF) int {
	for _, instruction := range instr {
		value = instruction(value)
	}
	return value
}

func reverse(value, len int) int {
	return len - 1 - value
}

func cut(value, len, c int) int {
	return (len + value - c) % len
}

func increment(value, len, n int) int {
	v1 := big.NewInt(int64(value))
	v2 := big.NewInt(int64(n))
	v1 = v1.Mul(v1, v2)
	v1.Mod(v1, big.NewInt(int64(len)))
	return int(v1.Int64())
}

type instructionF func(value int) int

const (
	iReverse = iota
	iCut
	iIncrement
	iIncrementCut
)

type instruction struct {
	iType int
	value int
}

func instFuncs(instr []instruction, l int) []instructionF {
	instF := make([]instructionF, len(instr))
	for i, v := range instr {
		switch v.iType {
		case iReverse:
			instF[i] = func(value int) int {
				return reverse(value, l)
			}
		case iCut:
			vi := v.value
			instF[i] = func(value int) int {
				return cut(value, l, vi)
			}
		case iIncrement:
			vi := v.value
			instF[i] = func(value int) int {
				return increment(value, l, vi)
			}
		default:
			panic("You should not be here :(")
		}
	}
	return instF
}

func printInstructions(instr []instruction) {
	for _, v := range instr {
		switch v.iType {
		case iReverse:
			fmt.Println("deal into new stack")
		case iCut:
			fmt.Printf("cut %d\n", v.value)
		case iIncrement:
			fmt.Printf("deal with increment %d\n", v.value)
		default:
			panic("You should not be here")
		}
	}
}
