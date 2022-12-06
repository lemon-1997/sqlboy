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
	data := struct {
		Package string
	}{
		Package: "render",
	}
	for _, tmpl := range tests {
		var buf bytes.Buffer
		if err := render(tmpl, &buf, &data); err != nil {
			t.Error(err)
		}
		t.Log(buf.String())
	}
}
