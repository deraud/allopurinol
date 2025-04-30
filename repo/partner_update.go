package repo

import (
	"context"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
)

func (r *PartnerRepoImpl) UpdatePartner(ctx context.Context, p *entity.Partner) (*entity.Partner, error) {
	partner := *p

	ub := sqlbuilder.NewUpdateBuilder()
	var values = make([]string, 0)
	if p.ActiveDays != "" {
		values = append(values, ub.Assign("active_days", p.ActiveDays))
	}
	if p.OperationalHourStart != "" {
		values = append(values, ub.Assign("operational_hour_start", p.OperationalHourStart))
	}
	if p.OperationalHourStart != "" {
		values = append(values, ub.Assign("operational_hour_end", p.OperationalHourEnd))
	}
	values = append(values, ub.Assign("is_active", p.IsActive))
	values = append(values, ub.Assign("updated_at", "NOW()"))
	ub.Update("partners").
		Set(values...).
		Where(ub.Equal("id", p.Id)).
		SQL("RETURNING name, year_founded, active_days, operational_hour_start, operational_hour_end, is_active")

	query, args := ub.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(
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

func (r *PartnerRepoImpl) DeactivatePharmacies(ctx context.Context, p *entity.Partner) error {
	partner := *p
	ub := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	ub.Update("pharmacies").
		Set(
			ub.Assign("is_active", false),
			ub.Assign("updated_at", "NOW()"),
		).
		Where(ub.Equal("partner_id", partner.Id))
	query, args := ub.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(
		&partner.Name,
		&partner.YearFounded,
		&partner.ActiveDays,
		&partner.OperationalHourStart,
		&partner.OperationalHourEnd,
		&partner.IsActive,
	)
	if err != nil {
		return err
	}
	return nil
}
