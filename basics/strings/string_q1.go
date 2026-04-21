package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer

	rem := n % 3
	if rem > 0 {
		buf.WriteString(s[:rem])
		if n > rem {
			buf.WriteByte(',')
		}
	}

	for i := rem;i < n;i += 3 {
		buf.WriteString(s[i:i+3])
		if i+3 < n {
			buf.WriteByte(',')
		}
	}

	return buf.String()
	
}

// Write a non-recursive version of comma, using bytes.Buffer instead of string concatenation
func main() {
	var s string

	fmt.Print("Enter the value: ")
	fmt.Scan(&s)

	dot := strings.LastIndex(s, ".")
	var s1, s2 string

	if dot >= 0 {
		s1 = s[:dot]
		s2 = s[dot:]
	} else {
		s1 = s
		s2 = ""
	}

	s1 = comma(s1)
	fmt.Println(s1+s2)

}
