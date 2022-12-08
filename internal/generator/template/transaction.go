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
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}
`
