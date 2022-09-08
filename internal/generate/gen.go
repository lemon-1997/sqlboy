package generate

import (
	"bytes"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"os"
)

type gen struct {
}

func New() *gen {
	return &gen{}
}

func (g *gen) Generate(in, out string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, in, nil, 0)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	if err = format.Node(&buf, fset, f); err != nil {
		panic(err)
	}

	file, err := os.Create(out)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err = io.Copy(file, &buf)
	if err != nil {
		panic(err)
	}
}
