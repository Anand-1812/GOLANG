package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type WordCounter int
type LineCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	// splits the words
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		*w++
	}

	return len(p), scanner.Err()
}

func (l *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))

	for scanner.Scan() {
		*l++;
	}

	return len(p), scanner.Err()
}

func main() {
	fmt.Println("Enter text (ctrl+D to end)")

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var wc WordCounter
	var lc LineCounter

	_, err = wc.Write(data)
	_, err = lc.Write(data)

	fmt.Println("Words:", wc)
	fmt.Println("Lines:", lc)
}
