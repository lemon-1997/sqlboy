package parser

import (
	"errors"
	"go/parser"
	"go/token"
	antrlParser "sqlboy/antrl"
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

type AntrlParseIn struct {
	Stmt string
}

type AntrlParseOut struct {
	antrlParser.TableAttr
}

type AntrlParser struct{}

func (*AntrlParser) Name() string {
	return "AntrlParser"
}

func (*AntrlParser) Parse(in interface{}) (interface{}, error) {
	parseIn, ok := in.(AntrlParseIn)
	if !ok {
		return nil, errors.New("parse in type error")
	}
	attr, errs := parseStmt(parseIn.Stmt)
	if len(errs) != 0 {
		return nil, errors.New("multi error")
	}
	return AntrlParseOut{TableAttr: attr}, nil
}
