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
	deckLenV     = 119_315_717_514_047
	shuffleCount = 101_741_582_076_661
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open input: %v", err)
	}
	defer f.Close()

	deckLen := big.NewInt(deckLenV)

	instructions, err := readInstructions(f, deckLen)
	if err != nil {
		log.Fatalf("Can not read instructions: %v", err)
	}
	printInstructions(instructions)

	// var reminder []instruction
	// for i := 101_741_582_076_661; i > 1; {
	// 	if i%2 == 1 {
	// 		reminder = add(instructions, reminder, deckLen)
	// 	}
	// 	i /= 2
	// 	instructions = double(instructions, deckLen)
	// }

	value := big.NewInt(2020)
	// value = applyShuffle(value, instructions, deckLen)
	// value = applyShuffle(value, reminder, deckLen)
	// fmt.Println(value)

	// for i := 0; i < shuffleCount; i++ {
	// 	if i%1_000_000 == 0 {
	// 		fmt.Println(i)
	// 	}
	// 	value = applyShuffle(value, instructions)
	// }

	applyShuffle(value, instructions, deckLen)
	fmt.Println(value)
	applyShuffle(value, instructions, deckLen)
	fmt.Println(value)
}

func double(inst []instruction, deckLen *big.Int) []instruction {
	return add(inst, inst, deckLen)
}

func add(inst1, inst2 []instruction, deckLen *big.Int) []instruction {
	t := append([]instruction(nil), inst1...)
	t = append(t, inst2...)
	t = normalize(t, deckLen)
	return t
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
			inst.value.Mul(inst.value, big.NewInt(-1))
			nInstr = append(nInstr, inst)
			continue
		}

		if inst.iType == iIncrement {
			nInstr = append(nInstr, inst)
			t := new(big.Int).Set(inst.value)
			t.Sub(t, big.NewInt(1))
			t.Mul(t, big.NewInt(-1))
			nInstr = append(nInstr, instruction{iType: iCut, value: t})
			continue
		}

		panic("You should not be here")
	}
	if inReverse {
		nInstr = append(nInstr, instruction{iType: iReverse})
	}
	return nInstr
}

func moveIncrement(instr []instruction, deckLen *big.Int) []instruction {
	for i := 0; i < len(instr); i++ {
		//Switch cut with increment
		if i != len(instr)-1 && instr[i].iType == iCut && instr[i+1].iType == iIncrement {
			instr[i].value.Mul(instr[i].value, instr[i+1].value)
			instr[i].value.Mod(instr[i].value, deckLen)

			//instr[i].value = (instr[i].value * instr[i+1].value) % deckLen
			instr[i], instr[i+1] = instr[i+1], instr[i]
			i = -1
			continue
		}

		// Merge double cut
		if i != len(instr)-1 && instr[i].iType == iCut && instr[i+1].iType == iCut {
			//value := (deckLen + instr[i].value + instr[i+1].value) % deckLen
			//instr[i] = instruction{iType: iCut, value: value}
			instr[i].value.Add(instr[i].value, instr[i+1].value)
			instr[i].value.Add(instr[i].value, deckLen)
			instr[i].value.Mod(instr[i].value, deckLen)

			instr = append(instr[:i+1], instr[i+2:]...)
			i--
			continue
		}

		//Merge double increment
		if i != len(instr)-1 && instr[i].iType == iIncrement && instr[i+1].iType == iIncrement {
			//value := (instr[i].value * instr[i+1].value) % deckLen
			//instr[i] = instruction{iType: iIncrement, value: value}
			instr[i].value.Mul(instr[i].value, instr[i+1].value)
			instr[i].value.Mod(instr[i].value, deckLen)

			instr = append(instr[:i+1], instr[i+2:]...)
			i--
			continue
		}
	}
	return instr
}

func readInstructions(r io.Reader, deckLen *big.Int) ([]instruction, error) {
	var instructions []instruction
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		var i int64

		_, err := fmt.Sscanf(line, "cut %d", &i)
		if err == nil {
			instructions = append(instructions, instruction{iType: iCut, value: big.NewInt(i)})
			continue
		}

		_, err = fmt.Sscanf(line, "deal with increment %d", &i)
		if err == nil {
			instructions = append(instructions, instruction{iType: iIncrement, value: big.NewInt(i)})
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

	instructions = normalize(instructions, deckLen)
	return instructions, nil
}

func normalize(instructions []instruction, deckLen *big.Int) []instruction {
	instructions = removeReverse(instructions)
	return moveIncrement(instructions, deckLen)
}

func reverse(value, len *big.Int) {
	value.Sub(len, value)
	value.Sub(value, big.NewInt(1))
	//return len - 1 - value
}

func cut(value, len, c *big.Int) {
	value.Sub(value, c)
	value.Add(value, len)
	value.Mod(value, len)
}

func increment(value, len, n *big.Int) {
	value.Mul(value, n)
	value.Mod(value, len)
}

const (
	iReverse = iota
	iCut
	iIncrement
)

type instruction struct {
	iType int
	value *big.Int
}

type instructionF func(value *big.Int)

func instFuncs(instr []instruction, l *big.Int) []instructionF {
	instF := make([]instructionF, len(instr))
	for i, v := range instr {
		switch v.iType {
		case iReverse:
			instF[i] = func(value *big.Int) {
				reverse(value, l)
			}
		case iCut:
			vi := v.value
			instF[i] = func(value *big.Int) {
				cut(value, l, vi)
			}
		case iIncrement:
			vi := v.value
			instF[i] = func(value *big.Int) {
				increment(value, l, vi)
			}
		default:
			panic("You should not be here :(")
		}
	}
	return instF
}

func applyShuffle(value *big.Int, instr []instruction, deckLen *big.Int) {
	for _, instruction := range instFuncs(instr, deckLen) {
		instruction(value)
	}
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
