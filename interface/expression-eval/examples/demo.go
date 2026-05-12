package main

import (
	"fmt"

	"expression-eval/eval"
)

func main() {
	expr, _ := eval.Parse("pow(x, 2) + 5")

	env := eval.Env{
		"x": 20,
	}

	fmt.Println(expr.Eval(env))
}
