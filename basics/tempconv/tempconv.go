package tempconv

import "fmt"

type Celsisus float32
type Farenheit float64

const (
	AbsoluteZeroC Celsisus = -273.15
	FreezingC Celsisus = 0
	BoilingC Celsisus = 100
)

func (c Celsisus) String() string {
	return fmt.Sprintf("%g C", c)
}

func (f Farenheit) String() string {
	return fmt.Sprintf("%g F", f)
}
