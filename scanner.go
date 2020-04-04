package main

import "unicode/utf8"

type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func newScanner(source string) Scanner {
	return Scanner{
		source:  source,
		tokens:  make([]Token, 0),
		start:   0,
		current: 0,
		line:    1,
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
