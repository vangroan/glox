// Automatically generated on 2020-04-05 at 14:36:31 +02:00
//
// go run github.com/vangroan/glox/gen > expr.go
package main

type Expr interface{}

type BaseExpr struct {
}

func (base BaseExpr) accept(visitor ExprVisitor) interface{} {
	return visitor.visitBase(base)
}

type BinaryExpr struct {
	base     BaseExpr
	left     Expr
	operator Token
	right    Expr
}

func (binary BinaryExpr) accept(visitor ExprVisitor) interface{} {
	return visitor.visitBinary(binary)
}

type GroupingExpr struct {
	base       BaseExpr
	expression Expr
}

func (grouping GroupingExpr) accept(visitor ExprVisitor) interface{} {
	return visitor.visitGrouping(grouping)
}

type LiteralExpr struct {
	base  BaseExpr
	value TokenLiteral
}

func (literal LiteralExpr) accept(visitor ExprVisitor) interface{} {
	return visitor.visitLiteral(literal)
}

type UnaryExpr struct {
	base     BaseExpr
	operator Token
	right    Expr
}

func (unary UnaryExpr) accept(visitor ExprVisitor) interface{} {
	return visitor.visitUnary(unary)
}

type ExprVisitor interface {
	visitBase(BaseExpr) interface{}
	visitBinary(BinaryExpr) interface{}
	visitGrouping(GroupingExpr) interface{}
	visitLiteral(LiteralExpr) interface{}
	visitUnary(UnaryExpr) interface{}
}
