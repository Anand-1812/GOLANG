package main

import "fmt"

func sneaky() (result int) {
	defer func() {
		if r := recover(); r != nil {
			result = 42
		}
	}()
	panic("trigger")
}

func main() {
	fmt.Println(sneaky())
}
