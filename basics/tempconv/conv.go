package tempconv

func CToF(c Celsisus) Farenheit {
	return Farenheit(c*9/5 + 32)
}

func FToC(f Farenheit) Celsisus {
	return Celsisus((f-32) * 5 / 9)
}
