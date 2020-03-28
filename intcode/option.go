package intcode

// Option is an option for the intcode.New() constructor.
type Option func(*Computer)

// WithInput sets an computer input
func WithInput(i int) Option {
	return func(c *Computer) {
		c.input = i
	}
}
