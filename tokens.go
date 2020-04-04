package main

type TokenType uint

const (
	tokenEOF TokenType = iota
)

// Single-character tokens.
const (
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
)

// One or two character tokens.
const (
	tokenBang         TokenType = iota
	tokenBangEqual    TokenType = iota
	tokenEqual        TokenType = iota
	tokenEqualEqual   TokenType = iota
	tokenGreater      TokenType = iota
	tokenGreaterEqual TokenType = iota
	tokenLess         TokenType = iota
	tokenLessEqual    TokenType = iota
)

// Literals.
const (
	tokenIdent  TokenType = iota
	tokenString TokenType = iota
	tokenNumber TokenType = iota
)

// Keywords.
const (
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
