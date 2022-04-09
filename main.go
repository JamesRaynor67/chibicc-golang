package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Invalid number of args: %d. Expecting exactly 1 arg", len(os.Args)-1)
	}

	inputStr := args[1]
	operand, offset, _ := StringToInt64(inputStr, 10)

	fmt.Printf("  .globl main\n")
	fmt.Printf("main:\n")
	fmt.Printf("  mov $%d, %%rax\n", operand)

	for offset+1 < len(inputStr) {
		if inputStr[offset] == '+' {
			// assert len(inputStr) > offset + 1, same below
			operand, parsedOffset, _ := StringToInt64(inputStr[offset+1:], 10)
			offset++ // ++ due to symbol
			offset += parsedOffset
			fmt.Printf("  add $%d, %%rax\n", operand)
			continue
		}

		if inputStr[offset] == '-' {
			operand, parsedOffset, _ := StringToInt64(inputStr[offset+1:], 10)
			offset++ // ++ due to symbol
			offset += parsedOffset
			fmt.Printf("  sub $%d, %%rax\n", operand)
			continue
		}

		fmt.Printf("Unexpected char %c at %d\n", inputStr[offset], offset)
		os.Exit(1)
	}

	fmt.Printf("  ret\n")
}
