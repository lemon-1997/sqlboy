package ast

import (
	"bytes"
	"errors"
	parser "github.com/lemon-1997/sqlboy/antrl"
	"github.com/lemon-1997/sqlboy/internal/generator"
	"go/format"
	"go/token"
)

type AssertGenIn struct {
	PackageName string
	Paths       map[string]string
	Asserts     map[string]string
}

type AssertGenerator struct{}

func (*AssertGenerator) Name() string {
	return "AssertGenerator"
}

func (*AssertGenerator) Generate(in interface{}) (*bytes.Buffer, error) {
	genIn, ok := in.(AssertGenIn)
	if !ok {
		return nil, errors.New("gen in type error")
	}
	file := buildAssertAST(genIn.PackageName, genIn.Paths, genIn.Asserts)
	var buf bytes.Buffer
	buf.WriteString(generator.Prefix)
	if err := format.Node(&buf, token.NewFileSet(), file); err != nil {
		return nil, err
	}
	return &buf, nil
}

type ModelGenIn struct {
	PackageName string
	Tables      map[string][]parser.ColumnDecl
	Mode        generator.Mode
}

type ModelGenerator struct{}

func (*ModelGenerator) Name() string {
	return "ModelGenerator"
}

func (*ModelGenerator) Generate(in interface{}) (*bytes.Buffer, error) {
	genIn, ok := in.(ModelGenIn)
	if !ok {
		return nil, errors.New("gen in type error")
	}
	file := buildModel(genIn.PackageName, genIn.Tables, genIn.Mode)
	var buf bytes.Buffer
	buf.WriteString(generator.Prefix)
	if err := format.Node(&buf, token.NewFileSet(), file); err != nil {
		return nil, err
	}
	return &buf, nil
}
