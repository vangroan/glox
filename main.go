package main

import (
	"fmt"
	"os"
)

func runFile(filename string) {
	fmt.Printf("running %s\n", filename)
}

func runRepl() {
	fmt.Println("REPL")
}

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("Usage: glox <filename>.lox")
		os.Exit(64)
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runRepl()
	}
}
