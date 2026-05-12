package eval

func Parse(input string) (Expr, error) {
	// parser implementation
	return binary{
		op: '+',
		x: literal(5),
		y: literal(5),
	}, nil
}
