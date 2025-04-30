package repo

import (
	"context"
	"database/sql"
	"errors"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
)

type PharmacyRepoItf interface {
	IsIdExist(ctx context.Context, id int64) (bool, error)
	GetPharmacistCount(ctx context.Context, id int64) (int, error)
	GetPharmacies(ctx context.Context, options *entity.PharmacyOptions) ([]*entity.Pharmacy, error)
	GetPharmaciesCount(ctx context.Context, options *entity.PharmacyOptions) (int, error)
	generateGetPharmaciesWhereClause(options *entity.PharmacyOptions) *sqlbuilder.WhereClause
	InsertPharmacy(ctx context.Context, p *entity.Pharmacy) (*entity.Pharmacy, error)
	UpdatePharmacy(ctx context.Context, p *entity.Pharmacy) (*entity.Pharmacy, error)
	GetPharmacy(ctx context.Context, id int64) (*entity.Pharmacy, error)
	DeletePharmacy(ctx context.Context, id int64) error
	GetPartnerId(ctx context.Context, id int64) (int64, error)
	HasOrders(ctx context.Context, id int64) (bool, error)
}

type PharmacyRepoImpl struct {
	db *sql.DB
}

func NewPharmacyRepo(database *sql.DB) PharmacyRepoItf {
	return &PharmacyRepoImpl{
		db: database,
	}
}

func (r *PharmacyRepoImpl) IsIdExist(ctx context.Context, id int64) (bool, error) {
	var exist bool

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("1").
		From("pharmacies").
		Where(
			sb.Equal("id", id),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&exist)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (r *PharmacyRepoImpl) GetPharmacistCount(ctx context.Context, id int64) (int, error) {
	var count int

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("COUNT(pharmacists.id)").
		From("pharmacies").
		Join("pharmacists", "pharmacies.id = pharmacists.pharmacy_id").
		Where(
			sb.Equal("pharmacies.id", id),
			sb.IsNull("pharmacies.deleted_at"),
			sb.IsNull("pharmacists.deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *PharmacyRepoImpl) GetPartnerId(ctx context.Context, id int64) (int64, error) {
	var partnerId int64

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("partner_id").
		From("pharmacies").
		Where(
			sb.Equal("id", id),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&partnerId)
	if err != nil {
		return 0, err
	}

	return partnerId, nil
}

func (r *PharmacyRepoImpl) HasOrders(ctx context.Context, id int64) (bool, error) {
	var count int

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("COUNT(oi.id)").
		From("catalogs c").
		Join("order_items oi", "c.id = oi.catalog_id").
		Join("pharmacies p", "c.pharmacy_id = p.id").
		Where(
			sb.Equal("p.id", id),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return false, err
	}

	return count != 0, nil
}
