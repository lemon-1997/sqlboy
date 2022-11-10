package parser

import (
	"go/ast"
	"go/token"
)

func parse(file *ast.File) (packageName string, ddl map[string]string) {
	if file == nil {
		return
	}
	if file.Name != nil {
		packageName = file.Name.Name
	}
	ddl = make(map[string]string)
	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		if genDecl.Tok != token.CONST {
			continue
		}
		for _, spec := range genDecl.Specs {
			valueSpec, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}
			for i := range valueSpec.Names {
				value, ok := valueSpec.Values[i].(*ast.BasicLit)
				if !ok {
					continue
				}
				if value.Kind != token.STRING {
					continue
				}
				ddl[valueSpec.Names[i].Name] = value.Value
			}
		}
	}
	return
}
