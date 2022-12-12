package ast

import (
	"bytes"
	"go/format"
	"go/token"
	parser "sqlboy/antrl"
	"sqlboy/internal/generator"
	"testing"
)

func Test_buildModel(t *testing.T) {
	tests := []struct {
		packageName string
		tables      map[string][]parser.ColumnDecl
		mode        generator.Mode
	}{
		{
			packageName: "test",
			tables: map[string][]parser.ColumnDecl{
				"order_info": {
					{
						Name:      "id",
						Comment:   "主键",
						GoType:    parser.Uint64,
						IsNotNull: true,
					},
					{
						Name:      "product_name",
						Comment:   "商品名称",
						GoType:    parser.String,
						IsNotNull: true,
					},
					{
						Name:      "created_at",
						Comment:   "创建时间",
						GoType:    parser.Time,
						IsNotNull: true,
					},
				},
				"test_info123---": {
					{
						Name:      "id!@#$$",
						Comment:   "主键",
						GoType:    parser.Uint64,
						IsNotNull: true,
					},
					{
						Name:      "product_name",
						Comment:   "商品名称",
						GoType:    parser.String,
						IsNotNull: false,
					},
					{
						Name:      "created_at",
						Comment:   "创建时间",
						GoType:    parser.Time,
						IsNotNull: true,
					},
				},
			},
			mode: generator.ModeSqlx,
		},
	}
	for _, tt := range tests {
		file := buildModel(tt.packageName, tt.tables, tt.mode)
		var buf bytes.Buffer
		buf.WriteString("// Code generated by sqlBoy. DO NOT EDIT.\n")
		if err := format.Node(&buf, token.NewFileSet(), file); err != nil {
			t.Error(err)
		}
		t.Log(buf.String())
	}
}
