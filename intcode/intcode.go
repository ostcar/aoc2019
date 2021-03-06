package intcode

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

const (
	pModePosition = iota
	pModeImmediate
	pModeRelative
)

// Computer runs a intcode program.
type Computer struct {
	name         string
	memory       []int
	done         bool
	operations   map[int]operation
	mode         int
	input        <-chan int
	output       chan int
	returnOutput bool
	pos          int
	relativeBase int
	inputFunc    func() int
	debug        bool
	ioDebug      bool
}

// NewFromReader returns a computer from an io.Reader
func NewFromReader(r io.Reader, ops ...Option) (*Computer, error) {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("can not read code: %v", err)
	}
	return New(string(bs), ops...), nil
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
	if c.output == nil {
		c.output = make(chan int, 99)
		c.returnOutput = true
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
	close(c.output)

	if !c.returnOutput {
		return nil
	}

	var out []int
	for v := range c.output {
		out = append(out, v)
	}
	return out
}

// Manipulate sets a value in memory.
func (c *Computer) Manipulate(index int, value int) {
	c.set(index, value)
}

func (c *Computer) getOpArgs(nr int) (operation, []int) {
	opArgs := splitInt(c.get(nr))
	opCode, argsMode := splitOpArgs(opArgs)

	op, ok := c.operations[opCode]
	if !ok {
		panic(fmt.Sprintf("Computer %s: Unknown operation %d on position %d", c.name, c.get(nr), nr))
	}

	var args []int
	for i := 0; i < op.argCount; i++ {
		switch get(argsMode, i) {
		case pModePosition:
			args = append(args, c.get(nr+i+1))
		case pModeImmediate:
			args = append(args, nr+i+1)
		case pModeRelative:
			args = append(args, c.get(nr+i+1)+c.relativeBase)
		default:
			panic(fmt.Sprintf("Unknown parameter mode %d on position %d", get(argsMode, i), nr))
		}
	}

	return op, args
}

func (c *Computer) get(nr int) int {
	if len(c.memory) <= nr {
		return 0
	}
	return c.memory[nr]
}

func (c *Computer) set(nr, value int) {
	if len(c.memory) <= nr {
		buf := make([]int, nr+1)
		copy(buf, c.memory)
		c.memory = buf
	}
	c.memory[nr] = value
}

func (c *Computer) log(s string, format ...interface{}) {
	fmt.Printf("Computer %s: %s\n", c.name, fmt.Sprintf(s, format...))
}
