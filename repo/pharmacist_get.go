package repo

import (
	"context"
	"fmt"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
)

func (r *PharmacistRepoImpl) GetPharmacists(ctx context.Context, options *entity.PharmacistOptions) ([]*entity.Pharmacist, []string, error) {
	size := options.Limit
	remaining := options.TotalRows - (options.Page-1)*options.Limit
	if remaining <= 0 {
		size = 0
	} else if remaining < options.Limit {
		size = remaining
	}
	pharmacists := make([]*entity.Pharmacist, size)
	emails := make([]string, size)

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("p.id", "ph.id", "ph.name", "p.name", "c.email", "p.sipa_number", "p.phone_number", "p.years_of_experience").
		From("pharmacists p").
		Join("credentials c", "p.credential_id = c.id").
		JoinWithOption(sqlbuilder.LeftOuterJoin, "pharmacies ph", "p.pharmacy_id = ph.id").
		AddWhereClause(r.generateGetPharmacistsWhereClause(options)).
		OrderBy(fmt.Sprintf("p.%s %s", options.SortBy, options.SortOrder)).
		Offset((options.Page - 1) * options.Limit).
		Limit(options.Limit)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var (
			pharmacist entity.Pharmacist
			email      string
		)
		err := rows.Scan(
			&pharmacist.Id,
			&pharmacist.PharmacyId,
			&pharmacist.PharmacyName,
			&pharmacist.Name,
			&email,
			&pharmacist.SipaNumber,
			&pharmacist.PhoneNumber,
			&pharmacist.YearsOfExperience,
		)
		if err != nil {
			return nil, nil, err
		}
		pharmacists[i] = &pharmacist
		emails[i] = email
	}
	err = rows.Err()
	if err != nil {
		return nil, nil, err
	}
	return pharmacists, emails, nil
}

func (r *PharmacistRepoImpl) GetPharmacistsCount(ctx context.Context, options *entity.PharmacistOptions) (int, error) {
	var count int

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("COUNT(p.id)").
		From("pharmacists p").
		Join("credentials c", "p.credential_id = c.id").
		AddWhereClause(r.generateGetPharmacistsWhereClause(options))

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *PharmacistRepoImpl) generateGetPharmacistsWhereClause(options *entity.PharmacistOptions) *sqlbuilder.WhereClause {
	whereClause := sqlbuilder.NewWhereClause()
	cond := sqlbuilder.NewCond()

	whereClause.AddWhereExpr(
		cond.Args,
		cond.ILike(options.SearchBy, fmt.Sprintf("%%%s%%", options.SearchValue)),
		cond.IsNull("p.deleted_at"),
	)
	if options.Assigned != nil {
		if *options.Assigned {
			whereClause.AddWhereExpr(cond.Args, cond.IsNotNull("pharmacy_id"))
		} else {
			whereClause.AddWhereExpr(cond.Args, cond.IsNull("pharmacy_id"))
		}
	}
	if options.YearsExpStart != nil {
		whereClause.AddWhereExpr(cond.Args, cond.GTE("years_of_experience", *options.YearsExpStart))
	}
	if options.YearsExpEnd != nil {
		whereClause.AddWhereExpr(cond.Args, cond.LTE("years_of_experience", *options.YearsExpEnd))
	}
	return whereClause
}

func (r *PharmacistRepoImpl) GetPharmacist(ctx context.Context, id int64) (*entity.Pharmacist, string, error) {
	var (
		pharmacist entity.Pharmacist
		email      string
	)

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("p.id", "ph.id", "ph.name", "p.name", "c.email", "p.sipa_number", "p.phone_number", "p.years_of_experience").
		From("pharmacists p").
		Join("credentials c", "p.credential_id = c.id").
		JoinWithOption(sqlbuilder.LeftOuterJoin, "pharmacies ph", "p.pharmacy_id = ph.id").
		Where(
			sb.Equal("p.id", id),
			sb.IsNull("p.deleted_at"),
		)

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(
		&pharmacist.Id,
		&pharmacist.PharmacyId,
		&pharmacist.PharmacyName,
		&pharmacist.Name,
		&email,
		&pharmacist.SipaNumber,
		&pharmacist.PhoneNumber,
		&pharmacist.YearsOfExperience,
	)
	if err != nil {
		return nil, "", err
	}

	return &pharmacist, email, nil
}

func (r *PharmacistRepoImpl) GetPharmacistsByPharmacy(ctx context.Context, pharmacyId int64) ([]*entity.Pharmacist, error) {
	var pharmacists []*entity.Pharmacist

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("p.id", "p.pharmacy_id", "p.name", "p.sipa_number", "p.phone_number", "p.years_of_experience").
		From("pharmacists p").
		Where(sb.Equal("pharmacy_id", pharmacyId))

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var pharmacist entity.Pharmacist
		err := rows.Scan(
			&pharmacist.Id,
			&pharmacist.PharmacyId,
			&pharmacist.Name,
			&pharmacist.SipaNumber,
			&pharmacist.PhoneNumber,
			&pharmacist.YearsOfExperience,
		)
		if err != nil {
			return nil, err
		}
		pharmacists = append(pharmacists, &pharmacist)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return pharmacists, nil
}
