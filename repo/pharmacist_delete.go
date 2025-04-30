package repo

import (
	"context"
	"healthcare-allopurinol/constant"

	"github.com/huandu/go-sqlbuilder"
)

func (r *PharmacistRepoImpl) DeletePharmacist(ctx context.Context, id int64) (int64, error) {
	var credId int64

	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("pharmacists").
		Set(
			ub.Assign("deleted_at", "NOW()"),
			ub.Assign("updated_at", "NOW()"),
		).
		Where(
			ub.Equal("id", id),
			ub.IsNull("deleted_at"),
		).
		SQL("RETURNING credential_id")

	tx := extractTx(ctx)
	query, args := ub.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := tx.QueryRowContext(ctx, query, args...).Scan(&credId)
	if err != nil {
		return 0, err
	}

	return credId, nil
}
