// Code generated by sqlBoy. DO NOT EDIT.
package sqlx

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type contextTxKey struct{}

type Dao struct {
	db *sqlx.DB
}

func NewDao(db *sqlx.DB) *Dao {
	return &Dao{
		db: db,
	}
}

func (d *Dao) InTx(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	err = fn(context.WithValue(ctx, contextTxKey{}, tx))

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (d *Dao) DB(ctx context.Context) DbTx {
	tx, ok := ctx.Value(contextTxKey{}).(*sqlx.Tx)
	if ok {
		return tx
	}
	return d.db
}
