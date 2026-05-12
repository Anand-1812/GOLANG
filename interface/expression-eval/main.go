package main

import (
	"fmt"

	"expression-eval/eval"
)

func main() {

	expr, err := eval.Parse("5 + 3 * 2")

	if err != nil {
		panic(err)
	}

	fmt.Println("Expression:")
	fmt.Println(expr.String())

	fmt.Println("Result:")
	fmt.Println(expr.Eval(eval.Env{}))
}
