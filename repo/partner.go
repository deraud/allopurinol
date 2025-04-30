package repo

import (
	"context"
	"database/sql"
	"errors"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"
	"strconv"
	"strings"
	"time"

	"github.com/huandu/go-sqlbuilder"
)

type PartnerRepoItf interface {
	IsIdExist(ctx context.Context, id int64) (bool, error)
	InsertPartner(ctx context.Context, p *entity.Partner) (*entity.Partner, error)
	UpdatePartner(ctx context.Context, p *entity.Partner) (*entity.Partner, error)
	DeletePartner(ctx context.Context, id int64) error
	GetPartners(ctx context.Context, options *entity.PartnerOptions) ([]*entity.Partner, error)
	GetPartnersCount(ctx context.Context, options *entity.PartnerOptions) (int, error)
	GetPartner(ctx context.Context, id int64) (*entity.Partner, error)
	IsYearFoundedValid(ctx context.Context, p entity.Partner) bool
	IsPartnerIdExist(ctx context.Context, id int64) (bool, error)
	DeactivatePharmacies(ctx context.Context, p *entity.Partner) error
	generateGetPartnersWhereClause(options *entity.PartnerOptions) *sqlbuilder.WhereClause
	IsActive(ctx context.Context, id int64) (bool, error)
}

type PartnerRepoImpl struct {
	db *sql.DB
}

func NewPartnerRepo(database *sql.DB) PartnerRepoItf {
	return &PartnerRepoImpl{
		db: database,
	}
}

func (r *PartnerRepoImpl) IsIdExist(ctx context.Context, id int64) (bool, error) {
	var exist bool

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("1").
		From("partners").
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

func (t *PartnerRepoImpl) IsOperationalHourValid(ctx context.Context, p entity.Partner) (bool, error) {
	t1 := strings.Replace(p.OperationalHourStart, ":", "", -1)
	t2 := strings.Replace(p.OperationalHourEnd, ":", "", -1)
	t1int, err := strconv.Atoi(t1)
	if err != nil {
		return false, err
	}
	t2int, err := strconv.Atoi(t2)
	if err != nil {
		return false, err
	}
	if t1int > t2int {
		return false, err
	}
	return true, nil
}

func (t *PartnerRepoImpl) IsYearFoundedValid(ctx context.Context, p entity.Partner) bool {
	currentYear, _, _ := time.Now().Date()
	return currentYear >= int(p.YearFounded)
}

func (r *PartnerRepoImpl) IsPartnerIdExist(ctx context.Context, id int64) (bool, error) {
	var exist bool

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("1").
		From("partners").
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

func (r *PartnerRepoImpl) IsActive(ctx context.Context, id int64) (bool, error) {
	var active bool

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("is_active").
		From("partners").
		Where(
			sb.Equal("id", id),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&active)
	if err != nil {
		return false, err
	}

	return active, nil
}
