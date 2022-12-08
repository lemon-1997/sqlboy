package render

import (
	"fmt"
	"io"
	"sqlboy/internal/generator"
	"strings"
	"text/template"
)

type Column struct {
	Name string
	Type string
}

type renderData struct {
	Package    string
	Table      string
	Columns    []Column
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

// TODO 特殊类型的查询会有问题
func renderQuery(tmpl string, wr io.Writer, data renderData) error {
	funcMap := template.FuncMap{
		"caseExport":      camelCaseExport,
		"caseInternal":    camelCaseInternal,
		"joinQuote":       joinColumnsWithQuote,
		"joinPlace":       joinColumnsWithPlace,
		"addQuote":        addQuote,
		"whereQuery":      whereQuery,
		"sqlxPlaceholder": sqlxPlaceholder,
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

func addQuote(str string) string {
	return fmt.Sprintf("`%s`", str)
}

func whereQuery(columns []Column) string {
	var s []string
	for _, item := range columns {
		s = append(s, fmt.Sprintf("`%s` = ?", item.Name))
	}
	return strings.Join(s, " AND ")
}

func sqlxPlaceholder(columns []Column) string {
	var s []string
	for _, item := range columns {
		s = append(s, fmt.Sprintf("`%s` = :%s", item.Name, item.Name))
	}
	return strings.Join(s, " AND ")
}

func joinColumnsWithQuote(columns []Column) string {
	var s []string
	for _, item := range columns {
		s = append(s, fmt.Sprintf("`%s`", item.Name))
	}
	return strings.Join(s, ",")
}

func joinColumnsWithPlace(columns []Column) string {
	var s []string
	for _, item := range columns {
		s = append(s, fmt.Sprintf(":%s", item.Name))
	}
	return strings.Join(s, ",")
}
