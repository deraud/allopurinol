package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
	"github.com/lib/pq"
)

type ShippingRepoItf interface {
	IsLogisticPartnerIdExistBulk(ctx context.Context, ids ...int64) ([]int64, error)
	GetLogisticPartners(ctx context.Context) ([]*entity.LogisticPartner, error)
	InsertShippingMethod(ctx context.Context, pharmacy_id int64, logisticPartnerIds ...int64) error
	IsShippingMethodExist(ctx context.Context, pharmacyId int64, logisticPartnerId int64) (bool, error)
	DeleteShippingMethodExceptIds(ctx context.Context, pharmacy_id int64, logisticPartnerIds ...int64) error
	GetLogisticPartnersByPharmacy(ctx context.Context, pharmacyId int64) ([]*entity.LogisticPartner, error)
	DeleteShippingMethod(ctx context.Context, pharmacyId int64) error
	GetLogisticPartnersByPharmacyBulk(ctx context.Context, pharmacyIds []int64) ([]*entity.Pharmacy, error)
}

type ShippingRepoImpl struct {
	db *sql.DB
}

func NewShippingRepo(database *sql.DB) ShippingRepoItf {
	return &ShippingRepoImpl{
		db: database,
	}
}

func (r *ShippingRepoImpl) IsLogisticPartnerIdExistBulk(ctx context.Context, ids ...int64) ([]int64, error) {
	whereClause := make([]string, len(ids))
	for _, id := range ids {
		whereClause = append(whereClause, fmt.Sprintf("id=%d", id))
	}

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id").
		From("logistic_partners").
		Where(
			sb.Or(whereClause...),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	rows, err := r.db.QueryContext(ctx, query, args...)
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

func (r *ShippingRepoImpl) GetLogisticPartners(ctx context.Context) ([]*entity.LogisticPartner, error) {
	var logisticPartners []*entity.LogisticPartner

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "name", "rate", "code").
		From("logistic_partners").
		Where(sb.IsNull("deleted_at"))

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var logisticPartner entity.LogisticPartner
		err := rows.Scan(
			&logisticPartner.Id,
			&logisticPartner.Name,
			&logisticPartner.Rate,
			&logisticPartner.Code,
		)
		if err != nil {
			return nil, err
		}
		logisticPartners = append(logisticPartners, &logisticPartner)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return logisticPartners, nil
}

func (r *ShippingRepoImpl) InsertShippingMethod(ctx context.Context, pharmacyId int64, logisticPartnerIds ...int64) error {
	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("shipping_methods").
		Cols("pharmacy_id", "logistic_partner_id")

	for _, id := range logisticPartnerIds {
		ib.Values(pharmacyId, id)
	}

	query, args := ib.BuildWithFlavor(constant.DATABASE_FLAVOR)
	tx := extractTx(ctx)
	_, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *ShippingRepoImpl) DeleteShippingMethodExceptIds(ctx context.Context, pharmacyId int64, logisticPartnerIds ...int64) error {
	whereClause := make([]string, len(logisticPartnerIds))
	for _, id := range logisticPartnerIds {
		whereClause = append(whereClause, fmt.Sprintf("logistic_partner_id!=%d", id))
	}

	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("shipping_methods").
		Set(
			ub.Assign("deleted_at", "NOW()"),
			ub.Assign("updated_at", "NOW()"),
		).
		Where(
			ub.Equal("pharmacy_id", pharmacyId),
			ub.And(whereClause...),
			ub.IsNull("deleted_at"),
		)

	query, args := ub.BuildWithFlavor(constant.DATABASE_FLAVOR)
	tx := extractTx(ctx)
	_, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *ShippingRepoImpl) DeleteShippingMethod(ctx context.Context, pharmacyId int64) error {
	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("shipping_methods").
		Set(
			ub.Assign("deleted_at", "NOW()"),
			ub.Assign("updated_at", "NOW()"),
		).
		Where(
			ub.Equal("pharmacy_id", pharmacyId),
			ub.IsNull("deleted_at"),
		)

	query, args := ub.BuildWithFlavor(constant.DATABASE_FLAVOR)
	tx := extractTx(ctx)
	_, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *ShippingRepoImpl) IsShippingMethodExist(ctx context.Context, pharmacyId int64, logisticPartnerId int64) (bool, error) {
	var exist bool

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("1").
		From("shipping_methods").
		Where(
			sb.Equal("pharmacy_id", pharmacyId),
			sb.Equal("logistic_partner_id", logisticPartnerId),
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

func (r *ShippingRepoImpl) GetLogisticPartnersByPharmacy(ctx context.Context, pharmacyId int64) ([]*entity.LogisticPartner, error) {
	var logisticPartners []*entity.LogisticPartner

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("lp.id", "lp.name", "lp.rate", "lp.code", "lp.courier").
		From("shipping_methods sm").
		Join("logistic_partners lp", "sm.logistic_partner_id = lp.id").
		Where(
			sb.Equal("pharmacy_id", pharmacyId),
			sb.IsNull("sm.deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var logisticPartner entity.LogisticPartner
		err := rows.Scan(
			&logisticPartner.Id,
			&logisticPartner.Name,
			&logisticPartner.Rate,
			&logisticPartner.Code,
			&logisticPartner.Courier,
		)
		if err != nil {
			return nil, err
		}
		logisticPartners = append(logisticPartners, &logisticPartner)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return logisticPartners, nil
}

func (r *ShippingRepoImpl) GetLogisticPartnersByPharmacyBulk(ctx context.Context, pharmacyIds []int64) ([]*entity.Pharmacy, error) {
	var pharmacies []*entity.Pharmacy

	whereClause := make([]string, len(pharmacyIds))
	for _, id := range pharmacyIds {
		whereClause = append(whereClause, fmt.Sprintf("sm.pharmacy_id=%d", id))
	}

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("sm.pharmacy_id", "ARRAY_AGG(lp.id)", "ARRAY_AGG(lp.name)").
		From("shipping_methods sm").
		Join("logistic_partners lp", "sm.logistic_partner_id = lp.id").
		Where(
			sb.Or(whereClause...),
			sb.IsNull("sm.deleted_at"),
		).
		GroupBy("sm.pharmacy_id")

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var (
			pharmacy             entity.Pharmacy
			logisticPartners     []*entity.LogisticPartner
			logisticPartnerIds   []int64
			logisticPartnerNames []string
		)
		err := rows.Scan(
			&pharmacy.Id,
			pq.Array(&logisticPartnerIds),
			pq.Array(&logisticPartnerNames),
		)
		if err != nil {
			return nil, err
		}

		for i = 0; i < len(logisticPartnerIds); i++ {
			logisticPartner := entity.LogisticPartner{
				Id:   logisticPartnerIds[i],
				Name: logisticPartnerNames[i],
			}
			logisticPartners = append(logisticPartners, &logisticPartner)
		}
		pharmacy.LogisticPartners = logisticPartners
		pharmacies = append(pharmacies, &pharmacy)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return pharmacies, nil
}
