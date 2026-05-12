package eval

func Parse(input string) (Expr, error) {

	return binary{
		op: '+',
		x:  literal(5),
		y: binary{
			op: '*',
			x: literal(2),
			y: literal(3),
		},
	}, nil
}
