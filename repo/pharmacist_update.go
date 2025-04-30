package repo

import (
	"context"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
)

func (r *PharmacistRepoImpl) UpdatePharmacist(ctx context.Context, p *entity.Pharmacist) (*entity.Pharmacist, error) {
	pharmacist := *p

	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("pharmacists").
		Set(
			ub.Assign("pharmacy_id", p.PharmacyId),
			ub.Assign("phone_number", p.PhoneNumber),
			ub.Assign("years_of_experience", p.YearsOfExperience),
			ub.Assign("updated_at", "NOW()"),
		).
		Where(
			ub.Equal("id", p.Id),
			ub.IsNull("deleted_at"),
		).
		SQL("RETURNING name, sipa_number")

	query, args := ub.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(
		&pharmacist.Name,
		&pharmacist.SipaNumber,
	)
	if err != nil {
		return nil, err
	}

	return &pharmacist, nil
}
