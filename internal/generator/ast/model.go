package ast

import (
	"fmt"
	parser "github.com/lemon-1997/sqlboy/antlr"
	"github.com/lemon-1997/sqlboy/internal/generator"
	"go/ast"
	"go/token"
	"strings"
)

// Todo 重复列名 表名 invalid类型 错误返回(这个后面优化代码会做，目前先实现，怎么简单怎么来)
func buildModel(packageName string, tables map[string][]parser.ColumnDecl, mode generator.Mode) *ast.File {
	var importTime, importSqlx bool
	var decls []ast.Decl
	for name, columns := range tables {
		fields := make([]*ast.Field, 0)
		for _, item := range columns {
			if item.GoType == parser.Time {
				importTime = true
			}

			var tagVales []string
			switch mode {
			case generator.ModeGorm:
				tagVales = append(tagVales, fmt.Sprintf(`gorm:"column:%s"`, item.Name))
			case generator.ModeSqlx:
				tagVales = append(tagVales, fmt.Sprintf(`db:"%s"`, item.Name))
			default:
				// err ?
				tagVales = append(tagVales, fmt.Sprintf(`gorm:"column:%s"`, item.Name))
			}
			tagVales = append(tagVales, fmt.Sprintf(`json:"%s"`, item.Name))

			var comment []*ast.Comment
			if item.Comment != "" {
				comment = append(comment, &ast.Comment{Text: fmt.Sprintf(" //%s", item.Comment)})
			}

			fieldType := string(item.GoType)
			if mode == generator.ModeSqlx && !item.IsNotNull {
				importSqlx = true
				fieldType = generator.ToSqlNullType(string(item.GoType))
			}
			fields = append(fields, &ast.Field{
				Names: []*ast.Ident{ast.NewIdent(generator.CamelCase(item.Name, true))},
				Type:  ast.NewIdent(fieldType),
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
	if importTime || importSqlx {
		var specs []ast.Spec
		if importTime {
			specs = append(specs, &ast.ImportSpec{Path: &ast.BasicLit{Kind: token.STRING, Value: `"time"`}})
		}
		if importSqlx {
			specs = append(specs, &ast.ImportSpec{Path: &ast.BasicLit{Kind: token.STRING, Value: `"database/sql"`}})
		}
		decls = append([]ast.Decl{&ast.GenDecl{Tok: token.IMPORT, Specs: specs}}, decls...)
	}
	return &ast.File{Name: ast.NewIdent(packageName), Decls: decls}
}
