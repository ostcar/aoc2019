package intcode

func (c *Computer) setOperations() {
	c.operations = map[int]operation{
		99: {"exit", opDone, 0},
		1:  {"add", opAdd, 3},
		2:  {"multi", opMul, 3},
		3:  {"input", opInput, 1},
		4:  {"output", opOutput, 1},
		5:  {"jump true", opJunpTrue, 2},
		6:  {"jump false", opJunpFalse, 2},
		7:  {"less", opLess, 3},
		8:  {"equal", opEquals, 3},
		9:  {"relative base", opRelativeBase, 1},
	}
}

type operation struct {
	name     string
	run      opCode
	argCount int
}

type opCode func(*Computer, []int)

func opDone(c *Computer, args []int) {
	if c.debug {
		c.log("Done")
	}
	c.done = true
}

func opAdd(c *Computer, args []int) {
	v1 := c.get(args[0])
	v2 := c.get(args[1])
	if c.debug {
		c.log("Add pos %d (%d) and pos %d (%d) in %d", args[0], v1, args[1], v2, args[2])
	}

	c.set(args[2], v1+v2)
	c.pos += 4
}

func opMul(c *Computer, args []int) {
	v1 := c.get(args[0])
	v2 := c.get(args[1])
	if c.debug {
		c.log("Mul pos %d (%d) and pos %d (%d) in %d", args[0], v1, args[1], v2, args[2])
	}

	c.set(args[2], v1*v2)
	c.pos += 4
}

func opInput(c *Computer, args []int) {
	var v int
	if c.debug {
		c.log("Input Start")
	}
	if c.inputFunc == nil {
		v = <-c.input
	} else {
		v = c.inputFunc()
	}
	if c.debug {
		c.log("Input Save %d into pos %d", v, args[0])
	}
	if c.ioDebug {
		c.log("Received %d", v)
	}

	c.set(args[0], v)
	c.pos += 2
}

func opOutput(c *Computer, args []int) {
	v := c.get(args[0])
	if c.debug {
		c.log("Output Send pos %d (%d)", args[0], v)
	}
	if c.ioDebug {
		c.log("Send %d", v)
	}

	c.output <- v
	if c.debug {
		c.log("Output Done")
	}
	c.pos += 2
}

func opJunpTrue(c *Computer, args []int) {
	v := c.get(args[0])
	if c.debug {
		c.log("JumpTrue on pos %d (%d) == %t, to pos %d (%d)", args[0], v, v != 0, args[1], c.get(args[1]))
	}
	if v != 0 {
		c.pos = c.get(args[1])
	} else {
		c.pos += 3
	}
}

func opJunpFalse(c *Computer, args []int) {
	v := c.get(args[0])
	if c.debug {
		c.log("JumpFalse on pos %d (%d) == %t, to pos %d %d", args[0], v, v != 0, args[1], c.get(args[1]))
	}
	if c.get(args[0]) == 0 {
		c.pos = c.get(args[1])
	} else {
		c.pos += 3
	}
}

func opLess(c *Computer, args []int) {
	v1 := c.get(args[0])
	v2 := c.get(args[1])
	if c.debug {
		c.log("Less pos %d (%d) and %d (%d) into pos %d", args[0], v1, args[1], v2, args[2])
	}
	if v1 < v2 {
		c.set(args[2], 1)
	} else {
		c.set(args[2], 0)
	}
	c.pos += 4
}

func opEquals(c *Computer, args []int) {
	v1 := c.get(args[0])
	v2 := c.get(args[1])
	if c.debug {
		c.log("Equals pos %d (%d) and %d (%d) into pos %d", args[0], v1, args[1], v2, args[2])
	}
	if v1 == v2 {
		c.set(args[2], 1)
	} else {
		c.set(args[2], 0)
	}
	c.pos += 4
}

func opRelativeBase(c *Computer, args []int) {
	v := c.get(args[0])
	c.relativeBase += v
	if c.debug {
		c.log("RelativeBase + pos %d (%d) == %d", args[0], v, c.relativeBase)
	}

	c.pos += 2
}
