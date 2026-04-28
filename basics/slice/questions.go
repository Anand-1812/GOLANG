package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func reverse(arr *[]int) {
	n := len(*arr)

	for i, j := 0, n-1;i < j;i, j = i+1, j-1 {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
}

func squashSpaces(b []byte) []byte {
	w := 0
	i := 0
	inSpace := false

	for i < len(b) {
		r, size := utf8.DecodeRune(b[i:])

		if unicode.IsSpace(r) {
			if !inSpace {
				b[w] = ' '
				w++
				inSpace = true
			}
		} else {
			inSpace = false
			copy(b[w:], b[i:i+size])
			w += size
		}
		i += size
	}
	return b[:w]
}

func reverseUTF8(b []byte) {
	reverseBytes(b)

	for i := 0;i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reverseBytes(b[i : i+size])
		i += size
	}
}

func reverseBytes(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a)

	s := []byte("Hello\t\tworld\n\u00A0\u00A0Go  \t  lang")
	s = squashSpaces(s)
	fmt.Printf("%q\n", s)

	
	s1 := []byte("Hello, 世界")
	reverseUTF8(s1)
	fmt.Println(string(s1))
	
}
