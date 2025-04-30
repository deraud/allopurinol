package repo

import (
	"context"
	"fmt"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
)

func (r *PartnerRepoImpl) GetPartners(ctx context.Context, options *entity.PartnerOptions) ([]*entity.Partner, error) {
	size := options.Limit
	remaining := options.TotalRows - (options.Page-1)*options.Limit
	if remaining <= 0 {
		size = 0
	} else if remaining < options.Limit {
		size = remaining
	}
	partners := make([]*entity.Partner, size)

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "name", "year_founded", "active_days", "operational_hour_start", "operational_hour_end", "is_active").
		From("partners").
		AddWhereClause(r.generateGetPartnersWhereClause(options)).
		OrderBy(fmt.Sprintf("%s %s", options.SortBy, options.SortOrder)).
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
			partner entity.Partner
		)
		err := rows.Scan(
			&partner.Id,
			&partner.Name,
			&partner.YearFounded,
			&partner.ActiveDays,
			&partner.OperationalHourStart,
			&partner.OperationalHourEnd,
			&partner.IsActive,
		)

		if err != nil {
			return nil, err
		}
		partners[i] = &partner
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return partners, nil
}

func (r *PartnerRepoImpl) GetPartnersCount(ctx context.Context, options *entity.PartnerOptions) (int, error) {
	var count int

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("COUNT(id)").
		From("partners").
		AddWhereClause(r.generateGetPartnersWhereClause(options))

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *PartnerRepoImpl) generateGetPartnersWhereClause(options *entity.PartnerOptions) *sqlbuilder.WhereClause {
	whereClause := sqlbuilder.NewWhereClause()
	cond := sqlbuilder.NewCond()

	whereClause.AddWhereExpr(
		cond.Args,
		cond.ILike(options.SearchBy, fmt.Sprintf("%%%s%%", options.SearchValue)),
		cond.IsNull("deleted_at"),
	)
	return whereClause
}

func (r *PartnerRepoImpl) GetPartner(ctx context.Context, id int64) (*entity.Partner, error) {
	var (
		partner entity.Partner
	)

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "name", "year_founded", "active_days", "operational_hour_start", "operational_hour_end", "is_active").
		From("partners").
		Where(
			sb.Equal("id", id),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(
		&partner.Id,
		&partner.Name,
		&partner.YearFounded,
		&partner.ActiveDays,
		&partner.OperationalHourStart,
		&partner.OperationalHourEnd,
		&partner.IsActive,
	)
	if err != nil {
		return nil, err
	}

	return &partner, nil
}
