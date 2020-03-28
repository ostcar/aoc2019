package intcode

func (c *Computer) setOperations() {
	c.operations = map[int]operation{
		99: operation{"exit", opDone, 0},
		1:  operation{"add", opAdd, 3},
		2:  operation{"multi", opMul, 3},
		3:  operation{"input", opInput, 1},
		4:  operation{"output", opOutput, 1},
		5:  operation{"jump true", opJunpTrue, 2},
		6:  operation{"jump false", opJunpFalse, 2},
		7:  operation{"less", opLess, 3},
		8:  operation{"equal", opEquals, 3},
		9:  operation{"relative base", opRelativeBase, 1},
	}
}

type operation struct {
	name     string
	run      opCode
	argCount int
}

type opCode func(*Computer, []int)

func opDone(c *Computer, args []int) {
	c.done = true
}

func opAdd(c *Computer, args []int) {
	c.set(args[2], c.get(args[0])+c.get(args[1]))
	c.pos += 4
}

func opMul(c *Computer, args []int) {
	c.set(args[2], c.get(args[0])*c.get(args[1]))
	c.pos += 4
}

func opInput(c *Computer, args []int) {
	c.set(args[0], <-c.input)
	c.pos += 2
}

func opOutput(c *Computer, args []int) {
	c.output <- c.get(args[0])
	c.pos += 2
}

func opJunpTrue(c *Computer, args []int) {
	if c.get(args[0]) != 0 {
		c.pos = c.get(args[1])
	} else {
		c.pos += 3
	}
}

func opJunpFalse(c *Computer, args []int) {
	if c.get(args[0]) == 0 {
		c.pos = c.get(args[1])
	} else {
		c.pos += 3
	}
}

func opLess(c *Computer, args []int) {
	if c.get(args[0]) < c.get(args[1]) {
		c.set(args[2], 1)
	} else {
		c.set(args[2], 0)
	}
	c.pos += 4
}

func opEquals(c *Computer, args []int) {
	if c.get(args[0]) == c.get(args[1]) {
		c.set(args[2], 1)
	} else {
		c.set(args[2], 0)
	}
	c.pos += 4
}

func opRelativeBase(c *Computer, args []int) {
	c.relativeBase += c.get(args[0])
	c.pos += 2
}
