package main

import (
	"os"
	"strconv"
)

func isPowerOf2(n int) bool {
	return (n != 0) && ((n & (n - 1)) == 0)
}

func main() {
	if len(os.Args) != 2 {
		return // Print nothing if there is no or more than one argument
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return // Print nothing if the argument is not a valid int
	}
	if isPowerOf2(n) {
		os.Stdout.Write([]byte("true\n"))
	} else {
		os.Stdout.Write([]byte("false\n"))
	}
}
