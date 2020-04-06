package main

import "strings"

type AstPrinter struct {
	BaseExprVisitor
}

func (pprint AstPrinter) Print(expr Expr) string {
	if s, ok := expr.accept(pprint).(string); ok {
		return s
	}

	return ""
}

func (pprint AstPrinter) parenthesize(name string, exprs ...Expr) string {
	var sb strings.Builder

	sb.WriteString("(")
	sb.WriteString(name)
	for _, expr := range exprs {
		sb.WriteString(" ")
		if p, ok := expr.accept(pprint).(string); ok {
			sb.WriteString(p)
		}
	}
	sb.WriteString(")")

	return sb.String()
}

func (pprint AstPrinter) visitBinary(expr BinaryExpr) interface{} {
	return pprint.parenthesize(expr.operator.lexeme, expr.left, expr.right)
}
