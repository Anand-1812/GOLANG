package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"flag"
	"fmt"
	"math/bits"
	"os"
	"io"
)

func count(a, b []byte) int {
	hashA := sha256.Sum256(a)
	hashB := sha256.Sum256(b)

	return countBits(hashA[:], hashB[:])
}

func countBits(x, y []byte) int {
	if len(x) != len(y) {
		panic("inputs must have equal length")
	}

	cnt := 0
	for i := range x {
		cnt += bits.OnesCount8(x[i] ^ y[i]) // ^ will give diff in bytes
	}

	return cnt
}

func main() {
	a := []byte("rat")
	b := []byte("cat") 
	
	fmt.Println("Exercise: 4.1 solution")
	fmt.Println(count(a, b))

	fmt.Println("Exercise: 4.2 solution")
	useSHA384 := flag.Bool("sha384", false, "Use SHA384 hash algo")
	useSHA512 := flag.Bool("sha512", false, "Use SHA512 hash algo")
	flag.Parse()

	var hasher hash.Hash

	switch {
	case *useSHA384:
		hasher = sha512.New384()
	case *useSHA512:
		hasher = sha512.New()
	default:
		hasher = sha256.New()
	}

	if _, err := io.Copy(hasher, os.Stdin); err != nil {
		fmt.Println(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("%x\n", hasher.Sum(nil))

}
