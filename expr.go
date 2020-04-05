// Automatically generated on 2020-04-05 at 14:15:50 +02:00
//
// go run github.com/vangroan/glox/gen
package main

type Expr interface{}

type BaseExpr struct {
}

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
