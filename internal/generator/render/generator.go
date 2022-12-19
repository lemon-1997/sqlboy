package render

import (
	"bytes"
	"errors"
	"github.com/lemon-1997/sqlboy/internal/generator"
	genTmpl "github.com/lemon-1997/sqlboy/internal/generator/template"
)

type DaoGenIn struct {
	PackageName string
	Mode        generator.Mode
}

type DaoGenerator struct{}

func (*DaoGenerator) Name() string {
	return "DaoGenerator"
}

func (*DaoGenerator) Generate(in interface{}) (*bytes.Buffer, error) {
	genIn, ok := in.(DaoGenIn)
	if !ok {
		return nil, errors.New("gen in type error")
	}
	var buf bytes.Buffer
	buf.WriteString(generator.Prefix)
	switch genIn.Mode {
	case generator.ModeGorm:
		if err := render(genTmpl.DaoGorm, &buf, genIn.PackageName); err != nil {
			return nil, err
		}
	case generator.ModeSqlx:
		if err := render(genTmpl.DaoSqlx, &buf, genIn.PackageName); err != nil {
			return nil, err
		}
	}
	return &buf, nil
}

type TxGenIn struct {
	PackageName string
	Mode        generator.Mode
}

type TxGenerator struct{}

func (*TxGenerator) Name() string {
	return "TxGenerator"
}

func (*TxGenerator) Generate(in interface{}) (*bytes.Buffer, error) {
	genIn, ok := in.(TxGenIn)
	if !ok {
		return nil, errors.New("gen in type error")
	}
	var buf bytes.Buffer
	buf.WriteString(generator.Prefix)
	switch genIn.Mode {
	case generator.ModeGorm:
		if err := render(genTmpl.TransactionGorm, &buf, genIn.PackageName); err != nil {
			return nil, err
		}
	case generator.ModeSqlx:
		if err := render(genTmpl.TransactionSqlx, &buf, genIn.PackageName); err != nil {
			return nil, err
		}
	}
	return &buf, nil
}

type QueryGenIn struct {
	Data QueryData
	Mode generator.Mode
}

type QueryGenerator struct{}

func (*QueryGenerator) Name() string {
	return "QueryGenerator"
}

func (*QueryGenerator) Generate(in interface{}) (*bytes.Buffer, error) {
	genIn, ok := in.(QueryGenIn)
	if !ok {
		return nil, errors.New("gen in type error")
	}
	var buf bytes.Buffer
	buf.WriteString(generator.Prefix)
	if err := renderQuery(&buf, genIn.Data, genIn.Mode); err != nil {
		return nil, err
	}
	return &buf, nil
}
