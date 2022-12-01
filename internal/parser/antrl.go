package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	parser "sqlboy/antrl"
)

// ColumnDecl 还需要注意特殊列名不能作为结构体idName 去掉特殊符号idName重复？
func parseStmt(ddl string) (tableName string, columnDecls []parser.ColumnDecl, errors []error) {
	input := antlr.NewInputStream(ddl)
	lexer := parser.NewStmtLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewStmtParser(stream)

	el := parser.NewErrorListener()
	p.RemoveErrorListeners()
	p.AddErrorListener(el)

	p.BuildParseTrees = true
	tree := p.Prog()

	if el.HasError() {
		return "", nil, el.Errors()
	}

	l := parser.NewStmtListener()
	antlr.ParseTreeWalkerDefault.Walk(l, tree)
	return l.TableName, l.Columns, nil
}
