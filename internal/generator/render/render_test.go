package render

import (
	"bytes"
	"sqlboy/internal/generator/template"
	"testing"
)

func Test_render(t *testing.T) {
	tests := []string{
		template.DaoGorm, template.DaoSqlx,
		template.TransactionGorm, template.TransactionSqlx,
	}
	for _, tmpl := range tests {
		var buf bytes.Buffer
		if err := render(tmpl, &buf, "render"); err != nil {
			t.Error(err)
		}
		t.Log(buf.String())
	}
}

func Test_renderQuery(t *testing.T) {
	tests := []string{
		template.QueryGorm,
	}
	data := renderData{
		Package: "render",
		Table:   "order_info",
		PrimaryKey: []Column{
			{
				Name: "id",
				Type: "int64",
			},
		},
		UniqueKeys: [][]Column{
			{
				{
					Name: "uid",
					Type: "int64",
				},
				{
					Name: "product_name",
					Type: "string",
				},
			},
			{
				{
					Name: "product_id",
					Type: "int64",
				},
			},
		},
	}
	for _, tmpl := range tests {
		var buf bytes.Buffer
		if err := renderQuery(tmpl, &buf, data); err != nil {
			t.Error(err)
		}
		t.Log(buf.String())
	}
}
