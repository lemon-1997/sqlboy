package template

const TransactionGorm = `package {{.}}

import "context"

type Transaction interface {
	InTx(context.Context, func(ctx context.Context) error) error
}

func NewTransaction(d *Dao) Transaction {
	return d
}
`

const TransactionSqlx = `package {{.}}

import (
	"context"
	"database/sql"
)

type Transaction interface {
	InTx(context.Context, func(ctx context.Context) error) error
}

func NewTransaction(d *Dao) Transaction {
	return d
}

type DbTx interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}
`
