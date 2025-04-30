package repo

import (
	"context"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
)

func (r *PharmacistRepoImpl) InsertPharmacist(ctx context.Context, p *entity.Pharmacist) (*entity.Pharmacist, error) {
	pharmacist := *p

	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("pharmacists").
		Cols("credential_id", "name", "sipa_number", "phone_number", "years_of_experience").
		Values(p.CredId, p.Name, p.SipaNumber, p.PhoneNumber, p.YearsOfExperience).
		SQL("RETURNING id, pharmacy_id")

	tx := extractTx(ctx)
	query, args := ib.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := tx.QueryRowContext(ctx, query, args...).Scan(&pharmacist.Id, &pharmacist.PharmacyId)
	if err != nil {
		return nil, err
	}

	return &pharmacist, nil
}
