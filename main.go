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

	inputStr := []byte(args[1])
	token := Tokenize(inputStr)

	fmt.Printf("  .globl main\n")
	fmt.Printf("main:\n")
	fmt.Printf("  mov $%d, %%rax\n", token.GetInt64Val())
	token = token.next

	for token.kind != TK_EOF {
		if token.Equals([]byte("+")) {
			token = token.next // Assume token.next is not nil
			fmt.Printf("  add $%d, %%rax\n", token.GetInt64Val())
			token = token.next
			continue
		}

		token = token.Skips([]byte("-"))
		fmt.Printf("  sub $%d, %%rax\n", token.GetInt64Val())
		token = token.next
	}

	fmt.Printf("  ret\n")
}
