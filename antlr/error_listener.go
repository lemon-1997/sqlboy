package parser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type ErrorListener struct {
	*antlr.DefaultErrorListener
	errors []error
}

func NewErrorListener() *ErrorListener {
	return new(ErrorListener)
}

func (l *ErrorListener) HasError() bool {
	return len(l.errors) > 0
}

func (l *ErrorListener) Errors() []error {
	return l.errors
}

func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	p := recognizer.(antlr.Parser)
	stack := p.GetRuleInvocationStack(p.GetParserRuleContext())
	err := fmt.Errorf("rule: %v line %d: %d at %v : %s", stack[0], line, column, offendingSymbol, msg)
	l.errors = append(l.errors, err)
}
