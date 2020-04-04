package main

import "unicode/utf8"

type Scanner struct {
	source       string
	tokens       []Token
	start        int
	current      int
	line         int
	errorPrinter ErrorPrinter
}

func newScanner(source string, errorPrinter ErrorPrinter) Scanner {
	return Scanner{
		source:       source,
		tokens:       make([]Token, 0),
		start:        0,
		current:      0,
		line:         1,
		errorPrinter: errorPrinter,
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

	case '\n':
		break // Ignore
	case '\r':
		break // Ignore
	case ' ':
		break // Ignore

	default:
		scanner.errorPrinter.printError(scanner.line, "Unexpected character.")
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
