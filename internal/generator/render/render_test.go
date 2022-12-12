package render

import (
	"bytes"
	"sqlboy/internal/generator"
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
	tests := []generator.Mode{
		generator.ModeGorm, generator.ModeSqlx,
	}
	data := renderData{
		Package: "render",
		Table:   "order_info",
		Columns: []Column{
			{Name: "id"},
			{Name: "product_name"},
			{Name: "c3"},
			{Name: "product_id"},
			{Name: "created_at"},
			{Name: "updated_at"},
		},
		PrimaryKey: []Column{
			{
				Name:      "id",
				Type:      "int64",
				IsNotNull: true,
			},
		},
		UniqueKeys: [][]Column{
			{
				{
					Name:      "uid",
					Type:      "int64",
					IsNotNull: false,
				},
				{
					Name:      "product_name",
					Type:      "string",
					IsNotNull: false,
				},
			},
			{
				{
					Name:      "product_id",
					Type:      "int64",
					IsNotNull: true,
				},
			},
		},
	}
	for _, mode := range tests {
		var buf bytes.Buffer
		if err := renderQuery(&buf, data, mode); err != nil {
			t.Error(err)
		}
		t.Log(buf.String())
	}
}

func Test_whereQuery(t *testing.T) {
	tests := []struct {
		in   []Column
		want string
	}{
		{
			in: []Column{
				{Name: "id"},
				{Name: "product_name"},
				{Name: "c3"},
			},
			want: "`id` = ? AND `product_name` = ? AND `c3` = ?",
		},
		{
			in:   nil,
			want: "",
		},
	}
	for index, tt := range tests {
		got := whereQuery(tt.in)
		if got != tt.want {
			t.Errorf("index(%d) got(%s) want(%s)", index, got, tt.want)
		}
	}
}

func Test_sqlxPlaceholder(t *testing.T) {
	tests := []struct {
		in   []Column
		want string
	}{
		{
			in: []Column{
				{Name: "id"},
				{Name: "product_name"},
				{Name: "c3"},
			},
			want: "`id` = :id AND `product_name` = :product_name AND `c3` = :c3",
		},
		{
			in:   nil,
			want: "",
		},
	}
	for index, tt := range tests {
		got := sqlxPlaceholder(tt.in)
		if got != tt.want {
			t.Errorf("index(%d) got(%s) want(%s)", index, got, tt.want)
		}
	}
}

func Test_joinColumnsWithQuote(t *testing.T) {
	tests := []struct {
		in   []Column
		want string
	}{
		{
			in: []Column{
				{Name: "id"},
				{Name: "product_name"},
				{Name: "c3"},
			},
			want: "`id`,`product_name`,`c3`",
		},
		{
			in:   nil,
			want: "",
		},
	}
	for index, tt := range tests {
		got := joinColumnsWithQuote(tt.in)
		if got != tt.want {
			t.Errorf("index(%d) got(%s) want(%s)", index, got, tt.want)
		}
	}
}

func Test_joinColumnsWithPlace(t *testing.T) {
	tests := []struct {
		in   []Column
		want string
	}{
		{
			in: []Column{
				{Name: "id"},
				{Name: "product_name"},
				{Name: "c3"},
			},
			want: ":id,:product_name,:c3",
		},
		{
			in:   nil,
			want: "",
		},
	}
	for index, tt := range tests {
		got := joinColumnsWithPlace(tt.in)
		if got != tt.want {
			t.Errorf("index(%d) got(%s) want(%s)", index, got, tt.want)
		}
	}
}
