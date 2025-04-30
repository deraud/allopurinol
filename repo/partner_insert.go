package repo

import (
	"context"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
)

func (r *PartnerRepoImpl) InsertPartner(ctx context.Context, p *entity.Partner) (*entity.Partner, error) {
	partner := *p
	ib := sqlbuilder.PostgreSQL.NewInsertBuilder()
	ib.InsertInto("partners").
		Cols("name", "year_founded", "active_days", "operational_hour_start", "operational_hour_end", "is_active").
		Values(p.Name, p.YearFounded, p.ActiveDays, p.OperationalHourStart, p.OperationalHourEnd, p.IsActive).
		SQL("RETURNING id")

	query, args := ib.Build()
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&partner.Id)
	if err != nil {
		return nil, err
	}

	return &partner, nil
}
