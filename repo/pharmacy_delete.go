package repo

import (
	"context"
	"healthcare-allopurinol/constant"

	"github.com/huandu/go-sqlbuilder"
)

func (r *PharmacyRepoImpl) DeletePharmacy(ctx context.Context, id int64) error {
	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("pharmacies").
		Set(
			ub.Assign("deleted_at", "NOW()"),
			ub.Assign("updated_at", "NOW()"),
		).
		Where(
			ub.Equal("id", id),
			ub.IsNull("deleted_at"),
		)

	tx := extractTx(ctx)
	query, args := ub.BuildWithFlavor(constant.DATABASE_FLAVOR)
	_, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
