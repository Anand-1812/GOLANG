package internal

type Counter struct { n int }

func (c *Counter) Increment() { c.n++ }
func (c *Counter) Reset() { c.n = 0 }
// getter
func (c *Counter) Value() int { return c.n }
