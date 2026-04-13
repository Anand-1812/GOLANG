package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// text of each lines appears more than ones
  counts := make(map[string]int)
  input := bufio.NewScanner(os.Stdin)

  for input.Scan() {
    counts[input.Text()]++
  }

  for line, n := range counts {
    if n > 1 {
      fmt.Printf("%d\t%s\n", n, line)
    }
  }

	// text of lines appear more than once
	counts1 := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts1)
	} else {
		for _, args := range files {
			f, err := os.Open(args)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts1)
			f.Close()
		}

		for line, n := range counts1 {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}

func countLines(f *os.File, count map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		count[input.Text()]++
	}
}
