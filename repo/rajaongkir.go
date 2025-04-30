package repo

import (
	"context"
	"database/sql"
	"errors"
	"github.com/huandu/go-sqlbuilder"
	"healthcare-allopurinol/constant"
)

type RORepoItf interface {
	GetROId(ctx context.Context, postalCode string) (int64, error)
	InsertROId(ctx context.Context, postalCode string, id int64) error
}

type RORepoImpl struct {
	db *sql.DB
}

func NewRORepo(database *sql.DB) RORepoItf {
	return &RORepoImpl{
		db: database,
	}
}

func (r *RORepoImpl) GetROId(ctx context.Context, postalCode string) (int64, error) {
	var id int64

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("raja_ongkir_id").
		From("raja_ongkir_ids").
		Where(
			sb.Equal("postal_code", postalCode),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *RORepoImpl) InsertROId(ctx context.Context, postalCode string, id int64) error {
	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("raja_ongkir_ids").
		Cols("postal_code", "raja_ongkir_id").
		Values(postalCode, id)

	query, args := ib.BuildWithFlavor(constant.DATABASE_FLAVOR)
	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
