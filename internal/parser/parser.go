package parser

import (
	"errors"
	antlrParser "github.com/lemon-1997/sqlboy/antlr"
	"go/parser"
	"go/token"
)

type AstParseIn struct {
	Path string
}

type AstParseOut struct {
	PackageName string
	Stmt        map[string]string
}

type AstParser struct{}

func (*AstParser) Name() string {
	return "AstParser"
}

func (*AstParser) Parse(in interface{}) (interface{}, error) {
	parseIn, ok := in.(AstParseIn)
	if !ok {
		return nil, errors.New("parse in type error")
	}
	file, err := parser.ParseFile(token.NewFileSet(), parseIn.Path, nil, 0)
	if err != nil {
		return nil, err
	}
	packageName, ddl := parse(file)
	return AstParseOut{
		PackageName: packageName,
		Stmt:        ddl,
	}, nil
}

type AntlrParseIn struct {
	Stmt string
}

type AntlrParseOut struct {
	antlrParser.TableAttr
}

type AntlrParser struct{}

func (*AntlrParser) Name() string {
	return "AntlrParser"
}

func (*AntlrParser) Parse(in interface{}) (interface{}, error) {
	parseIn, ok := in.(AntlrParseIn)
	if !ok {
		return nil, errors.New("parse in type error")
	}
	attr, errs := parseStmt(parseIn.Stmt)
	if len(errs) != 0 {
		return nil, errs[0]
	}
	return AntlrParseOut{TableAttr: attr}, nil
}
