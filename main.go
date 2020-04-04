package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func run(source string) {
	// scanner := newScanner(source)
	// tokens := scanner.scanTokens()

	// for token := range tokens {
	// fmt.Println(token)
	// }
}

func runFile(filename string) {
	fmt.Printf("running %s\n", filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	run(string(data))
}

func runRepl() {
	fmt.Println("REPL")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		run(input)
	}
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
