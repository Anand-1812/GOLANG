package eval

import (
	"fmt"
	"math"
	"testing"
)

func TestEval(t *testing.T) {

	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
	}

	for _, test := range tests {

		expr, err := Parse(test.expr)

		if err != nil {
			t.Fatal(err)
		}

		got := fmt.Sprintf("%.6g", expr.Eval(test.env))

		if got != test.want {
			t.Errorf("got %s want %s", got, test.want)
		}
	}
}
