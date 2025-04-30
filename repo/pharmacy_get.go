package repo

import (
	"context"
	"fmt"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
)

func (r *PharmacyRepoImpl) GetPharmacies(ctx context.Context, options *entity.PharmacyOptions) ([]*entity.Pharmacy, error) {
	size := options.Limit
	remaining := options.TotalRows - (options.Page-1)*options.Limit
	if remaining <= 0 {
		size = 0
	} else if remaining < options.Limit {
		size = remaining
	}
	pharmacies := make([]*entity.Pharmacy, size)

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("p.id", "pt.id", "pt.name", "p.name", "p.logo", "a.name", "p.is_active").
		From("pharmacies p").
		Join("addresses a", "p.id = a.pharmacy_id").
		Join("partners pt", "p.partner_id = pt.id").
		AddWhereClause(r.generateGetPharmaciesWhereClause(options)).
		Offset((options.Page - 1) * options.Limit).
		Limit(options.Limit)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var (
			pharmacy entity.Pharmacy
			partner  entity.Partner
			address  entity.Address
		)
		err := rows.Scan(
			&pharmacy.Id,
			&partner.Id,
			&partner.Name,
			&pharmacy.Name,
			&pharmacy.Logo,
			&address.Name,
			&pharmacy.IsActive,
		)
		if err != nil {
			return nil, err
		}
		pharmacy.Partner = &partner
		pharmacy.Address = &address
		pharmacies[i] = &pharmacy
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return pharmacies, nil
}

func (r *PharmacyRepoImpl) GetPharmaciesCount(ctx context.Context, options *entity.PharmacyOptions) (int, error) {
	var count int

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("COUNT(p.id)").
		From("pharmacies p").
		Join("addresses a", "p.id = a.pharmacy_id").
		AddWhereClause(r.generateGetPharmaciesWhereClause(options))

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *PharmacyRepoImpl) generateGetPharmaciesWhereClause(options *entity.PharmacyOptions) *sqlbuilder.WhereClause {
	whereClause := sqlbuilder.NewWhereClause()
	cond := sqlbuilder.NewCond()

	whereClause.AddWhereExpr(
		cond.Args,
		cond.ILike(fmt.Sprintf("p.%s", options.SearchBy), fmt.Sprintf("%%%s%%", options.SearchValue)),
		cond.IsNull("p.deleted_at"),
	)
	return whereClause
}

func (r *PharmacyRepoImpl) GetPharmacy(ctx context.Context, id int64) (*entity.Pharmacy, error) {
	var (
		pharmacy entity.Pharmacy
		address  entity.Address
		partner  entity.Partner
	)

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select(
		"p.id",
		"p.partner_id",
		"p.name",
		"p.logo",
		"pt.id",
		"pt.name",
		"a.id",
		"a.province",
		"a.city",
		"a.district",
		"a.subdistrict",
		"a.postal_code",
		"a.phone_number",
		"ST_XMax(a.location::geometry)",
		"ST_YMax(a.location::geometry)",
		"a.name",
		"p.is_active",
	).
		From("pharmacies p").
		Join("addresses a", "p.id = a.pharmacy_id").
		Join("partners pt", "p.partner_id = pt.id").
		Where(
			sb.Equal("p.id", id),
			sb.IsNull("p.deleted_at"),
		)

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(
		&pharmacy.Id,
		&pharmacy.PartnerId,
		&pharmacy.Name,
		&pharmacy.Logo,
		&partner.Id,
		&partner.Name,
		&address.Id,
		&address.Province,
		&address.City,
		&address.District,
		&address.Subdistrict,
		&address.PostalCode,
		&address.PhoneNumber,
		&address.Longitude,
		&address.Latitude,
		&address.Name,
		&pharmacy.IsActive,
	)
	if err != nil {
		return nil, err
	}
	pharmacy.Partner = &partner
	pharmacy.Address = &address
	return &pharmacy, nil
}
