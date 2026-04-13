package main

import (
	"fmt"
	"os"
)

// command line arguments
func main() {
	var s, sep string // var s and sep of type string
	for i := 1;i < len(os.Args);i++ {
		s += sep + os.Args[i]
		sep = " "
		fmt.Println(s, i)
	}

	s1, s2 := "", ""
	for _, arg := range os.Args[1:] {
		s1 += s2 + arg
		s2 = " "
	}
	fmt.Println(s1)

	fmt.Println(os.Args[0:])
}
