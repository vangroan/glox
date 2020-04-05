package main

type Expr interface{}

type BaseExpr struct{}

type BinaryExpr struct {
	base     BaseExpr
	left     Expr
	operator Token
	right    Expr
}

type GroupingExpr struct {
	base       BaseExpr
	expression Expr
}

type LiteralExpr struct {
	base  BaseExpr
	value TokenLiteral
}

type UnaryExpr struct {
	base     BaseExpr
	operator Token
	right    Expr
}
