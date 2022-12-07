package ast

import (
	"fmt"
	"go/ast"
	"go/token"
	parser "sqlboy/antrl"
	"sqlboy/internal/generator"
	"strings"
)

type Tag int8

const (
	Gorm Tag = iota
	SqlX
	JSON
)

// Todo 重复列名 表名 invalid类型 错误返回(这个后面优化代码会做，目前先实现，怎么简单怎么来)
func buildModel(packageName string, tables map[string][]parser.ColumnDecl, tags []Tag) *ast.File {
	var importTime bool
	var decls []ast.Decl
	for name, columns := range tables {
		fields := make([]*ast.Field, 0)
		for _, item := range columns {
			if item.GoType == parser.Time {
				importTime = true
			}
			var tagVales []string
			for _, t := range tags {
				switch t {
				case Gorm:
					tagVales = append(tagVales, fmt.Sprintf(`gorm:"column:%s"`, item.Name))
				case SqlX:
					tagVales = append(tagVales, fmt.Sprintf(`db:"%s"`, item.Name))
				case JSON:
					tagVales = append(tagVales, fmt.Sprintf(`json:"%s"`, item.Name))
				}
			}
			var comment []*ast.Comment
			if item.Comment != "" {
				comment = append(comment, &ast.Comment{Text: fmt.Sprintf(" //%s", item.Comment)})
			}
			fields = append(fields, &ast.Field{
				Names: []*ast.Ident{ast.NewIdent(generator.CamelCase(item.Name, true))},
				Type:  ast.NewIdent(string(item.GoType)),
				Tag: &ast.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf("`%s`", strings.Join(tagVales, " ")),
				},
				Comment: &ast.CommentGroup{List: comment},
			})
		}
		genDecl := &ast.GenDecl{
			Tok:   token.TYPE,
			Specs: []ast.Spec{&ast.TypeSpec{Name: ast.NewIdent(generator.CamelCase(name, true)), Type: &ast.StructType{Fields: &ast.FieldList{List: fields}}}},
		}
		funcDecl := &ast.FuncDecl{
			Recv: &ast.FieldList{List: []*ast.Field{{Type: &ast.StarExpr{X: ast.NewIdent(generator.CamelCase(name, true))}}}},
			Name: ast.NewIdent("TableName"),
			Type: &ast.FuncType{Results: &ast.FieldList{List: []*ast.Field{{Type: ast.NewIdent("string")}}}},
			Body: &ast.BlockStmt{List: []ast.Stmt{&ast.ReturnStmt{Results: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("`%s`", name)}}}}},
		}
		decls = append(decls, genDecl, funcDecl)
	}
	if importTime {
		decls = append([]ast.Decl{&ast.GenDecl{
			Tok:   token.IMPORT,
			Specs: []ast.Spec{&ast.ImportSpec{Path: &ast.BasicLit{Kind: token.STRING, Value: `"time"`}}},
		}}, decls...)
	}
	return &ast.File{Name: ast.NewIdent(packageName), Decls: decls}
}
