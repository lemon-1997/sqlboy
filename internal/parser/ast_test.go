package parser

import (
	"go/parser"
	"go/token"
	"reflect"
	"testing"
)

const input = `package boy

type t1 struct{}
type t2 struct{}

const order = "ddl order"
const one = 1
const (
	a,b,c = "ddl a",1,"ddl c"
	d = "ddl d"
	e = false
)
`

func Test_parse(t *testing.T) {
	tests := []struct {
		fileString  string
		packageName string
		ddl         map[string]string
	}{
		{
			fileString:  input,
			packageName: "boy",
			ddl: map[string]string{
				"order": `"ddl order"`,
				"a":     `"ddl a"`,
				"c":     `"ddl c"`,
				"d":     `"ddl d"`,
			},
		},
	}
	for index, tt := range tests {
		file, err := parser.ParseFile(token.NewFileSet(), "", input, 0)
		if err != nil {
			t.Fatal(err)
		}
		name, ddl := parse(file)
		if name != tt.packageName {
			t.Errorf("index(%d) got(%s) want(%s)", index, name, tt.packageName)
		}
		if !reflect.DeepEqual(ddl, tt.ddl) {
			t.Errorf("index(%d) got(%v) want(%v)", index, ddl, tt.ddl)
		}
	}
}
