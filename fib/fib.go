package main

import (
	"fmt"
	"os"
	"strconv"
)

func fib(n int) int {
	if n < 0 {
		fmt.Println("The arg must be positive number.")
		os.Exit(1)
	}
	if n == 0 || n == 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: ./fib [integer]")
		os.Exit(1)
	}
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "argument must be integer: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(fib(i))
}
