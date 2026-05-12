package main

import (
	"fmt"
	"math"

	"expression-eval/eval"
)

func main() {
	expr, err := eval.Parse("sqrt(A / pi)")
	if err != nil {
		panic(err)
	}

	env := eval.Env{
		"A":  87616,
		"pi": math.Pi,
	}

	fmt.Println(expr.Eval(env))
}
