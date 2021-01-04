// Abstract-syntax-tree generator tool
package main

import (
	"fmt"
	"strings"
	"time"
)

const header string = `// Automatically generated on %s at %s
//
// go run github.com/vangroan/glox/gen > expr.go
package main
`

const interfaceExpr string = `
type Expr interface{
	accept(ExprVisitor) interface{}
}

type BaseExpr struct {}
`

type exprDef struct {
	name   string
	fields string
}

func expressions() []exprDef {
	return []exprDef{
		exprDef{
			name:   "Binary",
			fields: "base: BaseExpr, left: Expr, operator: Token, right: Expr",
		},
		exprDef{
			name:   "Grouping",
			fields: "base: BaseExpr, expression: Expr",
		},
		exprDef{
			name:   "Literal",
			fields: "base: BaseExpr, value: TokenLiteral",
		},
		exprDef{
			name:   "Unary",
			fields: "base: BaseExpr, operator: Token, right: Expr",
		},
	}
}

func generateVisitor(sb *strings.Builder, exprs []exprDef, returnType string) {
	// =========
	// Interface
	sb.WriteString("\n")
	sb.WriteString("type ExprVisitor interface {\n")

	for _, expr := range exprs {
		sb.WriteString(fmt.Sprintf("	visit%s(%sExpr) %s\n", expr.name, expr.name, returnType))
	}

	sb.WriteString("}\n")

	// ====
	// Base
	sb.WriteString("\n")
	sb.WriteString("type BaseExprVisitor struct {}\n")

	for _, expr := range exprs {
		sb.WriteString("\n")
		sb.WriteString(fmt.Sprintf("func (visitor BaseExprVisitor) visit%s(expr %sExpr) %s {\n",
			expr.name, expr.name, returnType))
		sb.WriteString("	// Do nothing\n")
		sb.WriteString("	return nil\n")
		sb.WriteString("}\n")
	}
}

func generate() string {
	var sb strings.Builder

	visitorReturns := "interface{}"

	// ======
	// Header
	now := time.Now()

	sb.WriteString(fmt.Sprintf(header, now.Format("2006-01-02"), now.Format("15:04:05 -07:00")))
	sb.WriteString(interfaceExpr)

	// ===========
	// Expressions
	expr := expressions()
	for _, ex := range expr {
		// ------
		// Struct
		sb.WriteString("\n")
		sb.WriteString(fmt.Sprintf("type %sExpr struct {\n", ex.name))
		sb.WriteString("	BaseExpr\n")

		fields := strings.Split(ex.fields, ",")
		for _, part := range fields {
			pair := strings.Split(part, ":")
			if len(pair) == 2 {
				fieldName := strings.TrimSpace(pair[0])
				fieldType := strings.TrimSpace(pair[1])

				sb.WriteString(fmt.Sprintf("	%s	%s\n", fieldName, fieldType))
			}
		}

		sb.WriteString("}\n")

		// ------------
		// Visit Method
		self := strings.ToLower(ex.name)
		sb.WriteString("\n")
		sb.WriteString(
			fmt.Sprintf("func (%s %sExpr) accept(visitor ExprVisitor) %s {\n",
				self, ex.name, visitorReturns))
		sb.WriteString(fmt.Sprintf("	return visitor.visit%s(%s)\n", ex.name, self))
		sb.WriteString("}\n")
	}

	// =======
	// Visitor
	generateVisitor(&sb, expr, visitorReturns)

	return sb.String()
}

func main() {
	fmt.Println(generate())
}
