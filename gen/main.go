// Abstract-syntax-tree generator tool
package main

import (
	"fmt"
	"strings"
)

const header string = `// Automatically generated on %s at %s
//
// go run github.com/vangroan/glox/gen
package main
`

const baseExpr string = `
type Expr interface{}

type BaseExpr struct{}
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
	}
}

func generate() string {
	var sb strings.Builder

	sb.WriteString(header)
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

		sb.WriteString(fmt.Sprintln("}"))
	}

	return sb.String()
}

func main() {
	fmt.Println(generate())
}
