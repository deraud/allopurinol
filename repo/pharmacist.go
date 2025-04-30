package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
)

type PharmacistRepoItf interface {
	IsSipaNumberExist(ctx context.Context, sipaNumber string) (bool, error)
	IsPhoneNumberExist(ctx context.Context, phoneNumber string) (bool, error)
	GetIdByPhoneNumber(ctx context.Context, id int64, phoneNumber string) (int64, error)
	IsIdExist(ctx context.Context, id int64) (bool, error)
	IsAssigned(ctx context.Context, id int64) (bool, error)
	IsAssignedBulk(ctx context.Context, ids ...int64) ([]int64, error)
	InsertPharmacist(ctx context.Context, p *entity.Pharmacist) (*entity.Pharmacist, error)
	UpdatePharmacist(ctx context.Context, p *entity.Pharmacist) (*entity.Pharmacist, error)
	DeletePharmacist(ctx context.Context, id int64) (int64, error)
	GetPharmacists(ctx context.Context, options *entity.PharmacistOptions) ([]*entity.Pharmacist, []string, error)
	GetPharmacistsCount(ctx context.Context, options *entity.PharmacistOptions) (int, error)
	generateGetPharmacistsWhereClause(options *entity.PharmacistOptions) *sqlbuilder.WhereClause
	GetPharmacist(ctx context.Context, id int64) (*entity.Pharmacist, string, error)
	AssignPharmacistBulk(ctx context.Context, pharmacy_id int64, ids ...int64) error
	IsAssignedToPharmacyBulk(ctx context.Context, pharmacy_id int64, ids ...int64) ([]int64, error)
	UnassignPharmacistFromPharmacy(ctx context.Context, pharmacyId int64) error
	GetPharmacistsByPharmacy(ctx context.Context, pharmacyId int64) ([]*entity.Pharmacist, error)
	GetPharmacyIdByCredId(ctx context.Context, credId int64) (int64, error)
	IsAssignedToSpecificPharmacy(ctx context.Context, credId int64, pharmacyId int64) (bool, error)
	GetIdByCredId(ctx context.Context, credId int64) (int64, error)
}

type PharmacistRepoImpl struct {
	db *sql.DB
}

func NewPharmacistRepo(database *sql.DB) PharmacistRepoItf {
	return &PharmacistRepoImpl{
		db: database,
	}
}

func (r *PharmacistRepoImpl) IsSipaNumberExist(ctx context.Context, sipaNumber string) (bool, error) {
	var exist bool

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("1").
		From("pharmacists").
		Where(
			sb.Equal("sipa_number", sipaNumber),
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

func (r *PharmacistRepoImpl) IsPhoneNumberExist(ctx context.Context, phoneNumber string) (bool, error) {
	var exist bool

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("1").
		From("pharmacists").
		Where(
			sb.Equal("phone_number", phoneNumber),
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

func (r *PharmacistRepoImpl) GetIdByPhoneNumber(ctx context.Context, id int64, phoneNumber string) (int64, error) {
	var result int64

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id").
		From("pharmacists").
		Where(
			sb.NotEqual("id", id),
			sb.Equal("phone_number", phoneNumber),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&result)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (r *PharmacistRepoImpl) IsIdExist(ctx context.Context, id int64) (bool, error) {
	var exist bool

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("1").
		From("pharmacists").
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

func (r *PharmacistRepoImpl) IsAssigned(ctx context.Context, id int64) (bool, error) {
	var isAssigned bool

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select(sb.IsNotNull("pharmacy_id")).
		From("pharmacists").
		Where(
			sb.Equal("id", id),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&isAssigned)
	if err != nil {
		return false, err
	}

	return isAssigned, nil
}

func (r *PharmacistRepoImpl) IsAssignedBulk(ctx context.Context, ids ...int64) ([]int64, error) {
	whereClause := make([]string, len(ids))
	for _, id := range ids {
		whereClause = append(whereClause, fmt.Sprintf("id=%d", id))
	}

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id").
		From("pharmacists").
		Where(
			sb.Or(whereClause...),
			sb.IsNotNull("pharmacy_id"),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []int64
	for i := 0; rows.Next(); i++ {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		res = append(res, id)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *PharmacistRepoImpl) IsAssignedToPharmacyBulk(ctx context.Context, pharmacy_id int64, ids ...int64) ([]int64, error) {
	whereClause := make([]string, len(ids))
	for _, id := range ids {
		whereClause = append(whereClause, fmt.Sprintf("id=%d", id))
	}

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id").
		From("pharmacists").
		Where(
			sb.Or(whereClause...),
			sb.NotEqual("pharmacy_id", pharmacy_id),
			sb.IsNotNull("pharmacy_id"),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []int64
	for i := 0; rows.Next(); i++ {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		res = append(res, id)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *PharmacistRepoImpl) GetPharmacyIdByCredId(ctx context.Context, credId int64) (int64, error) {
	var pharmacyId int64

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("pharmacy_id").
		From("pharmacists").
		Where(
			sb.Equal("credential_id", credId),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&pharmacyId)
	if err != nil {
		return 0, err
	}

	return pharmacyId, nil
}

func (r *PharmacistRepoImpl) IsAssignedToSpecificPharmacy(ctx context.Context, credId int64, pharmacyId int64) (bool, error) {
	var assigned bool

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("1").
		From("pharmacists").
		Where(
			sb.Equal("credential_id", credId),
			sb.Equal("pharmacy_id", pharmacyId),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&assigned)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return assigned, nil
}

func (r *PharmacistRepoImpl) GetIdByCredId(ctx context.Context, credId int64) (int64, error) {
	var id int64

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id").
		From("pharmacists").
		Where(
			sb.Equal("credential_id", credId),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
