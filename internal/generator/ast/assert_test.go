package ast

import (
	"bytes"
	"go/format"
	"go/token"
	"testing"
)

func Test_buildAssertAST(t *testing.T) {
	file := buildAssertAST(
		"generate",
		map[string]string{
			"ast":   "\"go/ast\"",
			"token": "\"go/token\"",
		},
		map[string]string{
			"orderTable": "\"order ddl\"",
			"userTable":  "\"user ddl\"",
		},
	)

	var buf bytes.Buffer
	buf.WriteString("// DO NOT EDIT.\n")
	if err := format.Node(&buf, token.NewFileSet(), file); err != nil {
		t.Error(err)
	}
	t.Log(buf.String())
}
