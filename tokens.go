package main

import "fmt"

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   TokenLiteral
	line      int
}

func newToken(tokenType TokenType, lexeme string, literal TokenLiteral, line int) Token {
	return Token{
		tokenType: tokenType,
		lexeme:    lexeme,
		literal:   literal,
		line:      line,
	}
}

func (token Token) String() string {
	if token.literal != nil {
		return token.tokenType.String() + " " + token.lexeme + " " + token.literal.String()
	}
	return token.tokenType.String() + " " + token.lexeme
}

type TokenType uint

const (
	tokenEOF TokenType = iota

	// Single-character tokens.

	tokenLeftParen  TokenType = iota
	tokenRightParen TokenType = iota
	tokenLeftBrace  TokenType = iota
	tokenRightBrace TokenType = iota
	tokenComma      TokenType = iota
	tokenDot        TokenType = iota
	tokenMinus      TokenType = iota
	tokenPlus       TokenType = iota
	tokenSemicolon  TokenType = iota
	tokenSlash      TokenType = iota
	tokenStar       TokenType = iota

	// One or two character tokens.

	tokenBang         TokenType = iota
	tokenBangEqual    TokenType = iota
	tokenEqual        TokenType = iota
	tokenEqualEqual   TokenType = iota
	tokenGreater      TokenType = iota
	tokenGreaterEqual TokenType = iota
	tokenLess         TokenType = iota
	tokenLessEqual    TokenType = iota

	// Literals.

	tokenIdent  TokenType = iota
	tokenString TokenType = iota
	tokenNumber TokenType = iota

	// Keywords.

	tokenAnd    TokenType = iota
	tokenClass  TokenType = iota
	tokenElse   TokenType = iota
	tokenFalse  TokenType = iota
	tokenFunc   TokenType = iota
	tokenFor    TokenType = iota
	tokenIf     TokenType = iota
	tokenNil    TokenType = iota
	tokenOr     TokenType = iota
	tokenPrint  TokenType = iota
	tokenReturn TokenType = iota
	tokenSuper  TokenType = iota
	tokenThis   TokenType = iota
	tokenTrue   TokenType = iota
	tokenVar    TokenType = iota
	tokenWhile  TokenType = iota
)

func (tt TokenType) String() string {
	switch tt {
	case tokenEOF:
		return "EOF"
	case tokenLeftParen:
		return "LEFT_PAREN"
	case tokenRightParen:
		return "RIGHT_PAREN"
	case tokenLeftBrace:
		return "LEFT_BRACE"
	case tokenRightBrace:
		return "RIGHT_BRACE"
	case tokenComma:
		return "COMMA"
	case tokenDot:
		return "DOT"
	case tokenMinus:
		return "MINUS"
	case tokenPlus:
		return "PLUS"
	case tokenSemicolon:
		return "SEMICOLON"
	case tokenSlash:
		return "SLASH"
	case tokenStar:
		return "STAR"

	case tokenBang:
		return "BANG"
	case tokenBangEqual:
		return "BANG_EQUAL"
	case tokenEqual:
		return "EQUAL"
	case tokenEqualEqual:
		return "EQUAL_EQUAL"
	case tokenGreater:
		return "GREATER"
	case tokenGreaterEqual:
		return "GREATER_EQUAL"
	case tokenLess:
		return "LESS"
	case tokenLessEqual:
		return "LESS_EQUAL"

	case tokenString:
		return "STRING"
	case tokenNumber:
		return "NUMBER"
	}

	return ""
}

type TokenLiteral interface {
	String() string
}

type StringLiteral struct {
	value string
}

func (s StringLiteral) String() string {
	return s.value
}

type NumberLiteral struct {
	value float64
}

func (s NumberLiteral) String() string {
	return fmt.Sprintf("%f", s.value)
}
