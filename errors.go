package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Error_at(token Token, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "%s", token.wholeLine)
	fmt.Fprintf(os.Stderr, "%s^", strings.Repeat(" ", token.offset))
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Printf(format, args...)
	os.Exit(1)
}

func Log_panic(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func Error_panic(err error) {
	log.Panic(err)
}
