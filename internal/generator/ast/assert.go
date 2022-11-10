package ast

import (
	"go/ast"
	"go/token"
)

func buildAssertAST(packageName string, paths, asserts map[string]string) *ast.File {
	importSpecs := make([]ast.Spec, 0)
	for name, path := range paths {
		importSpecs = append(importSpecs, &ast.ImportSpec{
			Name: ast.NewIdent(name),
			Path: &ast.BasicLit{Kind: token.STRING, Value: path},
		})
	}

	maps := make([]*ast.CompositeLit, 0)
	for k, v := range asserts {
		elts := make([]ast.Expr, 0)
		elts = append(elts, &ast.KeyValueExpr{
			Key:   ast.NewIdent("false"),
			Value: &ast.CompositeLit{},
		})
		elts = append(elts, &ast.KeyValueExpr{
			Key:   &ast.BinaryExpr{X: ast.NewIdent(k), Op: token.EQL, Y: &ast.BasicLit{Kind: token.STRING, Value: v}},
			Value: &ast.CompositeLit{},
		})
		maps = append(maps, &ast.CompositeLit{
			Type: &ast.MapType{Key: ast.NewIdent("bool"), Value: &ast.StructType{Fields: &ast.FieldList{}}},
			Elts: elts,
		})
	}

	assignList := make([]ast.Stmt, 0)
	for _, item := range maps {
		assignList = append(assignList, &ast.AssignStmt{
			Lhs: []ast.Expr{ast.NewIdent("_")},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{item},
		})
	}

	decls := make([]ast.Decl, 0)
	if len(paths) != 0 {
		decls = append(decls, &ast.GenDecl{Tok: token.IMPORT, Specs: importSpecs})
	}
	if len(asserts) != 0 {
		decls = append(decls, &ast.FuncDecl{
			Doc:  &ast.CommentGroup{List: []*ast.Comment{{Text: "//compile-time assertion"}}},
			Name: ast.NewIdent("_"),
			Type: &ast.FuncType{},
			Body: &ast.BlockStmt{List: assignList},
		})
	}

	return &ast.File{Name: ast.NewIdent(packageName), Decls: decls}
}
