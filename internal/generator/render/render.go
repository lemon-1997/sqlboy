package render

import (
	"fmt"
	"io"
	"sqlboy/internal/generator"
	genTmpl "sqlboy/internal/generator/template"
	"strings"
	"text/template"
)

type Column struct {
	Name      string
	Type      string
	IsNotNull bool
}

type QueryData struct {
	Package    string
	Table      string
	Columns    []Column
	PrimaryKey []Column
	UniqueKeys [][]Column
	ImportSqlx bool
}

func render(tmpl string, wr io.Writer, data interface{}) error {
	t, err := template.New(tmpl).Parse(tmpl)
	if err != nil {
		return err
	}
	return t.Execute(wr, data)
}

// TODO 特殊类型的查询会有问题
func renderQuery(wr io.Writer, data QueryData, mode generator.Mode) error {
	funcMap := template.FuncMap{
		"caseExport":      camelCaseExport,
		"caseInternal":    camelCaseInternal,
		"joinQuote":       joinColumnsWithQuote,
		"joinPlace":       joinColumnsWithPlace,
		"addQuote":        addQuote,
		"whereQuery":      whereQuery,
		"sqlxPlaceholder": sqlxPlaceholder,
		"sqlxType":        sqlxType,
	}
	var tmpl string
	switch mode {
	case generator.ModeGorm:
		tmpl = genTmpl.QueryGorm
		data.ImportSqlx = false
	case generator.ModeSqlx:
		tmpl = genTmpl.QuerySqlx
		data.ImportSqlx = importSqlx(data)
	default:
		// err ?
		tmpl = genTmpl.QueryGorm
		data.ImportSqlx = false
	}
	t, err := template.New(tmpl).Funcs(funcMap).Parse(tmpl)
	if err != nil {
		return err
	}
	return t.Execute(wr, data)
}

func importSqlx(data QueryData) bool {
	for _, item := range data.PrimaryKey {
		if !item.IsNotNull {
			return true
		}
	}
	for _, uk := range data.UniqueKeys {
		for _, item := range uk {
			if !item.IsNotNull {
				return true
			}
		}
	}
	return false
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

func sqlxType(column Column) string {
	if column.IsNotNull {
		return column.Type
	}
	return generator.ToSqlNullType(column.Type)
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
