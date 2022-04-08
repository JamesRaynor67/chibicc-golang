package main

import (
	// "bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Invalid number of args: %d. Expecting exactly 1 int args between 0 and 127.", len(os.Args)-1)
	}
	exitCode, err := strconv.Atoi(args[1])
	if err != nil || !(exitCode <= 127 && exitCode >= 0) {
		panic(err)
	}

	fmt.Printf("  .globl main\n")
	fmt.Printf("main:\n")
	fmt.Printf("  mov $%d, %%rax\n", exitCode)
	fmt.Printf("  ret\n")
}
