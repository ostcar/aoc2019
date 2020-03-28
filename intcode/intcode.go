package intcode

import (
	"fmt"
	"strings"
)

const (
	pModePosition = iota
	pModeImmediate
)

// Computer runs a intcode program.
type Computer struct {
	memory     []int
	done       bool
	operations map[int]operation
	mode       int
	input      int
	output     []int
	pos        int
}

// New creates a new computer from an input code.
func New(input string, ops ...Option) *Computer {
	vals := strings.Split(input, ",")
	ints := make([]int, len(vals))
	for i := range vals {
		ints[i] = mustInt(vals[i])
	}

	c := &Computer{memory: ints}

	for _, op := range ops {
		op(c)
	}
	c.setOperations()
	return c
}

// Run calls the computer.
func (c *Computer) Run() []int {
	for !c.done {
		op, args := c.getOpArgs(c.pos)
		op.run(c, args)
	}
	return c.output
}

func (c *Computer) getOpArgs(nr int) (operation, []int) {
	opArgs := splitInt(c.memory[nr])
	opCode, argsMode := splitOpArgs(opArgs)

	op, ok := c.operations[opCode]
	if !ok {
		panic(fmt.Sprintf("Unknown operation %d on position %d", c.memory[nr], nr))
	}

	var args []int
	for i := 0; i < op.argCount; i++ {
		switch get(argsMode, i) {
		case pModePosition:
			args = append(args, c.memory[nr+i+1])
		case pModeImmediate:
			args = append(args, nr+i+1)
		default:
			panic(fmt.Sprintf("Unknown parameter mode %d on position %d", get(argsMode, i), nr))
		}
	}

	return op, args
}