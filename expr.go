// Automatically generated on 2020-04-06 at 11:03:28 +02:00
//
// go run github.com/vangroan/glox/gen > expr.go
package main

type Expr interface {
	accept(ExprVisitor) interface{}
}

type BaseExpr struct{}

type BinaryExpr struct {
	BaseExpr
	base     BaseExpr
	left     Expr
	operator Token
	right    Expr
}

func (binary BinaryExpr) accept(visitor ExprVisitor) interface{} {
	return visitor.visitBinary(binary)
}

type GroupingExpr struct {
	BaseExpr
	base       BaseExpr
	expression Expr
}

func (grouping GroupingExpr) accept(visitor ExprVisitor) interface{} {
	return visitor.visitGrouping(grouping)
}

type LiteralExpr struct {
	BaseExpr
	base  BaseExpr
	value TokenLiteral
}

func (literal LiteralExpr) accept(visitor ExprVisitor) interface{} {
	return visitor.visitLiteral(literal)
}

type UnaryExpr struct {
	BaseExpr
	base     BaseExpr
	operator Token
	right    Expr
}

func (unary UnaryExpr) accept(visitor ExprVisitor) interface{} {
	return visitor.visitUnary(unary)
}

type ExprVisitor interface {
	visitBinary(BinaryExpr) interface{}
	visitGrouping(GroupingExpr) interface{}
	visitLiteral(LiteralExpr) interface{}
	visitUnary(UnaryExpr) interface{}
}

type BaseExprVisitor struct{}

func (visitor BaseExprVisitor) visitBinary(expr BinaryExpr) interface{} {
	// Do nothing
	return nil
}

func (visitor BaseExprVisitor) visitGrouping(expr GroupingExpr) interface{} {
	// Do nothing
	return nil
}

func (visitor BaseExprVisitor) visitLiteral(expr LiteralExpr) interface{} {
	// Do nothing
	return nil
}

func (visitor BaseExprVisitor) visitUnary(expr UnaryExpr) interface{} {
	// Do nothing
	return nil
}
