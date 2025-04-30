package repo

import (
	"context"
	"database/sql"
)

type TransactorItf interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

type TransactorImpl struct {
	db *sql.DB
}

func NewTransactor(db *sql.DB) *TransactorImpl {
	return &TransactorImpl{db: db}
}

func (t *TransactorImpl) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = fn(injectTx(ctx, tx))
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return errRollback
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

type TxKey struct{}

func injectTx(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, TxKey{}, tx)
}

func extractTx(ctx context.Context) *sql.Tx {
	if tx, ok := ctx.Value(TxKey{}).(*sql.Tx); ok {
		return tx
	}
	return nil
}
