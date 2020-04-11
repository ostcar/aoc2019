package intcode

// Option is an option for the intcode.New() constructor.
type Option func(*Computer)

// WithInput sets an computer input
func WithInput(i ...int) Option {
	return func(c *Computer) {
		ch := make(chan int, len(i))
		for _, v := range i {
			ch <- v
		}
		c.input = ch
	}
}

// WithInputChan sets an computer input by a channel
func WithInputChan(ch <-chan int) Option {
	return func(c *Computer) {
		c.input = ch
	}
}

// WithOutputChan sets a chan where the output values can read from.
// Chan will be closed when the computer is finished
func WithOutputChan(ch chan int) Option {
	return func(c *Computer) {
		c.output = ch
	}
}

// WithName sets a name on the computer for debugging.
func WithName(name string) Option {
	return func(c *Computer) {
		c.name = name
	}
}

// WithInputFunc reads the next input value from an input function.
func WithInputFunc(f func() int) Option {
	return func(c *Computer) {
		c.inputFunc = f
	}
}

// WithDebug prints information on every operation.
func WithDebug() Option {
	return func(c *Computer) {
		c.debug = true
	}
}

// WithInputOutputDebug prints information on input and output operations
func WithInputOutputDebug() Option {
	return func(c *Computer) {
		c.ioDebug = true
	}
}
