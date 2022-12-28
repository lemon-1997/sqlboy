package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	parser "github.com/lemon-1997/sqlboy/antlr"
)

func parseStmt(ddl string) (parser.TableAttr, []error) {
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
		return parser.TableAttr{}, el.Errors()
	}

	l := parser.NewStmtListener()
	antlr.ParseTreeWalkerDefault.Walk(l, tree)
	return l.TableAttr, nil
}
