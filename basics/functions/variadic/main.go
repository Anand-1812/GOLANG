package main

import "fmt"

func maxMin(vals ...int) (int,int) {
	maxVal, minVal := vals[0], vals[0]

	for _, val := range vals {
		maxVal = max(maxVal, val)
		minVal = min(minVal, val)
	}

	return maxVal, minVal
}

func main() {
	val1, val2 := maxMin(1, 2, 3, 4, 5)
	fmt.Printf("Max: %d\nMin: %d\n", val1, val2)

}
