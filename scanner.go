package main

import (
	"strconv"
	"unicode/utf8"
)

func keywords() map[string]TokenType {
	return map[string]TokenType{
		"and":    tokenAnd,
		"class":  tokenClass,
		"else":   tokenElse,
		"false":  tokenFalse,
		"for":    tokenFor,
		"fun":    tokenFunc,
		"if":     tokenIf,
		"nil":    tokenNil,
		"or":     tokenOr,
		"print":  tokenPrint,
		"return": tokenReturn,
		"super":  tokenSuper,
		"this":   tokenThis,
		"true":   tokenTrue,
		"var":    tokenVar,
		"while":  tokenWhile,
	}
}

type Scanner struct {
	source       string
	tokens       []Token
	start        int
	current      int
	line         int
	errorPrinter ErrorPrinter
	keywordMap   map[string]TokenType
}

func newScanner(source string, errorPrinter ErrorPrinter) Scanner {
	return Scanner{
		source:       source,
		tokens:       make([]Token, 0),
		start:        0,
		current:      0,
		line:         1,
		errorPrinter: errorPrinter,
		keywordMap:   keywords(),
	}
}

func (scanner Scanner) isAtEnd() bool {
	return scanner.current >= len(scanner.source)
}

func (scanner *Scanner) advance() rune {
	runeValue, width := utf8.DecodeRuneInString(scanner.source[scanner.current:])
	scanner.current += width
	return runeValue
}

func (scanner *Scanner) addToken(tokenType TokenType, literal TokenLiteral) {
	text := scanner.source[scanner.start:scanner.current]
	scanner.tokens = append(scanner.tokens, newToken(tokenType, text, literal, scanner.line))
}

// Look ahead at the next character.
func (scanner *Scanner) match(char rune) bool {
	if scanner.isAtEnd() {
		return false
	}

	runeValue, width := utf8.DecodeRuneInString(scanner.source[scanner.current:])
	if char != runeValue {
		return false
	}

	scanner.current += width
	return true
}

func (scanner Scanner) peek() rune {
	if scanner.isAtEnd() {
		return rune(0)
	}

	runeValue, _ := utf8.DecodeRuneInString(scanner.source[scanner.current:])
	return runeValue
}

func (scanner Scanner) peekNext() rune {
	// Current width
	_, width := utf8.DecodeRuneInString(scanner.source[scanner.current:])
	if scanner.current+width >= len(scanner.source) {
		return rune(0)
	}
	runeValue, _ := utf8.DecodeRuneInString(scanner.source[scanner.current+width:])
	return runeValue
}

func (scanner *Scanner) stringLiteral() {
	for scanner.peek() != '"' && !scanner.isAtEnd() {
		if scanner.peek() == '\n' {
			scanner.line++
		}
		scanner.advance()
	}

	// Unterminated string
	if scanner.isAtEnd() {
		scanner.errorPrinter.printError(scanner.line, "Unterminated string.")
		return
	}

	// Closing quote
	scanner.advance()

	// Trim the surrounding quotes
	val := scanner.source[scanner.start+1 : scanner.current-1]
	scanner.addToken(tokenString, StringLiteral{value: val})
}

func (scanner *Scanner) numberLiteral() {
	for isDigit(scanner.peek()) {
		scanner.advance()
	}

	// Look for a fractional part.
	if scanner.peek() == '.' && isDigit(scanner.peekNext()) {
		// Consume dot
		scanner.advance()

		for isDigit(scanner.peek()) {
			scanner.advance()
		}
	}

	value, err := strconv.ParseFloat(scanner.source[scanner.start:scanner.current], 64)
	if err != nil {
		panic(err)
	}

	scanner.addToken(tokenNumber, NumberLiteral{value: value})
}

func (scanner *Scanner) identifier() {
	for isAlphanumeric(scanner.peek()) {
		scanner.advance()
	}

	// See if the identifier is a reserved word.
	text := scanner.source[scanner.start:scanner.current]

	if tokenType, ok := scanner.keywordMap[text]; ok {
		// Keyword
		scanner.addToken(tokenType, nil)
	} else {
		// User defined identifier
		scanner.addToken(tokenIdent, nil)
	}
}

func (scanner *Scanner) scanToken() {
	c := scanner.advance()
	switch c {
	case '(':
		scanner.addToken(tokenLeftParen, nil)
	case ')':
		scanner.addToken(tokenRightParen, nil)
	case '{':
		scanner.addToken(tokenLeftBrace, nil)
	case '}':
		scanner.addToken(tokenRightBrace, nil)
	case ',':
		scanner.addToken(tokenComma, nil)
	case '.':
		scanner.addToken(tokenDot, nil)
	case '-':
		scanner.addToken(tokenMinus, nil)
	case '+':
		scanner.addToken(tokenPlus, nil)
	case ';':
		scanner.addToken(tokenSemicolon, nil)
	case '*':
		scanner.addToken(tokenStar, nil)

	case '!':
		if scanner.match('=') {
			scanner.addToken(tokenBangEqual, nil)
		} else {
			scanner.addToken(tokenBang, nil)
		}
	case '=':
		if scanner.match('=') {
			scanner.addToken(tokenEqualEqual, nil)
		} else {
			scanner.addToken(tokenEqual, nil)
		}
	case '<':
		if scanner.match('=') {
			scanner.addToken(tokenLessEqual, nil)
		} else {
			scanner.addToken(tokenLess, nil)
		}
	case '>':
		if scanner.match('=') {
			scanner.addToken(tokenGreaterEqual, nil)
		} else {
			scanner.addToken(tokenGreater, nil)
		}

	case '/':
		if scanner.match('/') {
			// A comment goes until the end of the line.
			for scanner.peek() != '\n' && !scanner.isAtEnd() {
				scanner.advance()
			}
		} else {
			scanner.addToken(tokenSlash, nil)
		}

	case '"':
		scanner.stringLiteral()

	case '\n':
		scanner.line++

	case '\r':
	case '\t':
	case ' ':
		break // Ignore

	default:
		if isDigit(c) {
			scanner.numberLiteral()
		} else if isAlpha(c) {
			scanner.identifier()
		} else {
			scanner.errorPrinter.printError(scanner.line, "Unexpected character.")
		}
	}
}

func (scanner *Scanner) scanTokens() []Token {
	for !scanner.isAtEnd() {
		// We are at the beginning of the next lexeme.
		scanner.start = scanner.current
		scanner.scanToken()
	}

	scanner.tokens = append(scanner.tokens, newToken(tokenEOF, "", nil, scanner.line))
	return scanner.tokens
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func isAlpha(char rune) bool {
	return (char >= 'a' && char <= 'z') ||
		(char >= 'A' && char <= 'Z') ||
		char == '_'
}

func isAlphanumeric(char rune) bool {
	return isAlpha(char) || isDigit(char)
}
