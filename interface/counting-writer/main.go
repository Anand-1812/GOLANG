package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// This is a pattern known as decorator/wrapper

type coutingWriter struct {
	w io.Writer
	n int64
}

func (cw *coutingWriter) Write(p []byte) (int, error) {
	n, err := cw.w.Write(p)

	cw.n += int64(n)

	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &coutingWriter{
		w: w,
	}

	return cw, &cw.n
}

func main() {

	input := strings.Join(os.Args[1:], " ")

	// we can't use input here, use something as io.Writer
	w, count := CountingWriter(os.Stdout)

	w.Write([]byte(input + "\n"))
	fmt.Println("Bytes written:", *count)
}

