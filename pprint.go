package main

type AstPrinter struct {
	BaseExprVisitor
}

func (pprint AstPrinter) Print(expr Expr) string {
	if s, ok := expr.accept(pprint).(string); ok {
		return s
	}

	return ""
}
