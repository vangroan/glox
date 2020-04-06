package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// Lox virtual machine
type Lox struct {
	// hadError is to prevent the execution of invalid sourcecode.
	hadError bool
}

func newLox() Lox {
	return Lox{
		hadError: false,
	}
}

func (lox *Lox) printError(line int, message string) {
	lox.report(line, "", message)
}

func (lox *Lox) report(line int, where string, message string) {
	fmt.Printf("[line %d] Error %s: %s\n", line, where, message)
	lox.hadError = true
}

func (lox *Lox) run(source string) {
	scanner := newScanner(source, lox)
	tokens := scanner.scanTokens()

	// Token Print
	for _, token := range tokens {
		fmt.Println(token.String())
	}

	// AST Print
	root := BinaryExpr{
		operator: Token{tokenType: tokenStar, lexeme: "*"},
		left:     LiteralExpr{},
		right:    LiteralExpr{},
	}
	pprint := AstPrinter{}
	fmt.Println(pprint.Print(root))
}

func (lox Lox) runFile(filename string) {
	fmt.Printf("running %s\n", filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lox.run(string(data))
}

func (lox Lox) runRepl() {
	fmt.Println("REPL")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		lox.run(input)
		// Reset error state so user can continue using the REPL.
		lox.hadError = false
	}
}

type ErrorPrinter interface {
	printError(line int, message string)
}

func main() {
	args := os.Args[1:]
	lox := newLox()

	if len(args) > 1 {
		fmt.Println("Usage: glox <filename>.lox")
		os.Exit(64)
	} else if len(args) == 1 {
		lox.runFile(args[0])
	} else {
		lox.runRepl()
	}
}
