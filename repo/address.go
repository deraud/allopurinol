package repo

import (
	"context"
	"database/sql"
	"errors"
	"github.com/shopspring/decimal"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
)

type AddressRepoItf interface {
	InsertAddress(ctx context.Context, a *entity.Address) (*entity.Address, error)
	UpdateAddressPharmacy(ctx context.Context, a *entity.Address) (*entity.Address, error)
	DeletePharmacyAddress(ctx context.Context, pharmacyId int64) error
	GetAddress(ctx context.Context, id int64) (*entity.Address, error)
	IsUserActiveAddressExist(ctx context.Context, userId int64) (bool, error)
	GetUserActiveAddressId(ctx context.Context, userId int64) (int64, error)
	IsAddressExist(ctx context.Context, id int64) (bool, error)
	GetUserId(ctx context.Context, id int64) (*int64, error)
	GetAddressByUserId(ctx context.Context, userId int64) (*entity.Address, error)
	GetDistance(ctx context.Context, fromLat, fromLong, toLat, toLong float64) (decimal.Decimal, error)
}

type AddressRepoImpl struct {
	db *sql.DB
}

func NewAddressRepo(database *sql.DB) AddressRepoItf {
	return &AddressRepoImpl{
		db: database,
	}
}

func (r *AddressRepoImpl) IsUserActiveAddressExist(ctx context.Context, userId int64) (bool, error) {
	var exist bool

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("1").
		From("addresses").
		Where(
			sb.Equal("user_id", userId),
			sb.Equal("is_active", true),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&exist)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (r *AddressRepoImpl) GetUserActiveAddressId(ctx context.Context, userId int64) (int64, error) {
	var id int64

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id").
		From("addresses").
		Where(
			sb.Equal("user_id", userId),
			sb.Equal("is_active", true),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AddressRepoImpl) IsAddressExist(ctx context.Context, id int64) (bool, error) {
	var exist bool

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("1").
		From("addresses").
		Where(
			sb.Equal("id", id),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&exist)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (r *AddressRepoImpl) GetUserId(ctx context.Context, id int64) (*int64, error) {
	var userId *int64

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("user_id").
		From("addresses").
		Where(
			sb.Equal("id", id),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&userId)
	if err != nil {
		return nil, err
	}

	return userId, nil
}
