// Abstract-syntax-tree generator tool
package main

import (
	"fmt"
	"strings"
	"time"
)

const header string = `// Automatically generated on %s at %s
//
// go run github.com/vangroan/glox/gen
package main
`

const interfaceExpr string = `
type Expr interface{}
`

type exprDef struct {
	name   string
	fields string
}

func expressions() []exprDef {
	return []exprDef{
		exprDef{name: "Base", fields: ""},
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

func generateVisitor(sb *strings.Builder, exprs []exprDef) {
	sb.WriteString("\n")
	sb.WriteString("type ExprVisitor interface {\n")
	// visitBinary(BinaryExpr) interface{}
	returnType := "interface{}"

	for _, expr := range exprs {
		sb.WriteString(fmt.Sprintf("\tvisit%s(%sExpr) %s\n", expr.name, expr.name, returnType))
	}

	sb.WriteString("}\n")
}

func generate() string {
	var sb strings.Builder

	// ======
	// Header
	now := time.Now()

	sb.WriteString(fmt.Sprintf(header, now.Format("2006-01-02"), now.Format("15:04:05 -07:00")))
	sb.WriteString(interfaceExpr)

	// ===========
	// Expressions
	expr := expressions()
	for _, ex := range expr {
		sb.WriteString("\n")
		sb.WriteString(fmt.Sprintf("type %sExpr struct {\n", ex.name))

		fields := strings.Split(ex.fields, ",")
		for _, part := range fields {
			pair := strings.Split(part, ":")
			if len(pair) == 2 {
				fieldName := strings.TrimSpace(pair[0])
				fieldType := strings.TrimSpace(pair[1])

				sb.WriteString(fmt.Sprintf("\t%s\t%s\n", fieldName, fieldType))
			}
		}

		sb.WriteString("}\n")
	}

	// =======
	// Visitor
	generateVisitor(&sb, expr)

	return sb.String()
}

func main() {
	fmt.Println(generate())
}
