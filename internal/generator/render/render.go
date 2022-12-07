package render

import (
	"io"
	"sqlboy/internal/generator"
	"text/template"
)

type Column struct {
	Name string
	Type string
}

type renderData struct {
	Package    string
	Table      string
	PrimaryKey []Column
	UniqueKeys [][]Column
}

func render(tmpl string, wr io.Writer, data interface{}) error {
	t, err := template.New(tmpl).Parse(tmpl)
	if err != nil {
		return err
	}
	return t.Execute(wr, data)
}

func renderQuery(tmpl string, wr io.Writer, data renderData) error {
	funcMap := template.FuncMap{
		"caseExport":   camelCaseExport,
		"caseInternal": camelCaseInternal,
	}
	t, err := template.New(tmpl).Funcs(funcMap).Parse(tmpl)
	if err != nil {
		return err
	}
	return t.Execute(wr, data)
}

func camelCaseExport(str string) string {
	return generator.CamelCase(str, true)
}

func camelCaseInternal(str string) string {
	return generator.CamelCase(str, false)
}
