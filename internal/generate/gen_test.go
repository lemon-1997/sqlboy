package generate

import "testing"

func TestGen_Generate(t *testing.T) {
	g := New()
	tests := []struct {
		in  string
		out string
	}{
		{
			in:  "gen.go",
			out: "test.go",
		},
	}
	for _, tt := range tests {
		g.Generate(tt.in, tt.out)
	}
}
