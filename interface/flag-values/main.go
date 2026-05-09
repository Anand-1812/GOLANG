package main

import (
	"fmt"
	"flag"
)

type Celsius float64
type Fahrenheit float64
type celsiusFlag struct{ Celsius }

func (c Celsius) String() string {
	return fmt.Sprintf("%.1fC", c)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f-32) * 5/9)
}

func (cf *celsiusFlag) Set(s string) error {
	var unit string
	var value float64

	// "222F"
	fmt.Sscanf(s, "%f%s", &value, &unit) // value = 222, unit = F (parsing)
	switch unit {
		case "C", "°C":
			cf.Celsius = Celsius(value)
			return nil
		case "F", "°F":
			cf.Celsius = FToC(Fahrenheit(value))
			return nil
	}

	return fmt.Errorf("Invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}

	flag.CommandLine.Var(&f, name, usage)

	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "temperature") // 20.0 -> Celsius type -> String() -> output: "20.0C"

func main() {
	flag.Parse()

	fmt.Println("Temperature:", *temp)
}
