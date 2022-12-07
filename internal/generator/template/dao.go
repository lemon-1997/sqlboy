package template

const DaoGorm = `package {{.}}

import (
	"context"
	"gorm.io/gorm"
)

type contextTxKey struct{}

type Dao struct {
	db *gorm.DB
}

func NewDao(db *gorm.DB) *Dao {
	return &Dao{
		db: db,
	}
}

func (d *Dao) InTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func (d *Dao) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx.WithContext(ctx)
	}
	return d.db.WithContext(ctx)
}
`

const DaoSqlx = `package {{.}}

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
`
