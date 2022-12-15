package internal

import (
	"bytes"
)

type Parser interface {
	Name() string
	Parse(interface{}) (interface{}, error)
}

type Generator interface {
	Name() string
	Generate(interface{}) (*bytes.Buffer, error)
}
