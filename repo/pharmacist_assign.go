package repo

import (
	"context"
	"fmt"
	"healthcare-allopurinol/constant"

	"github.com/huandu/go-sqlbuilder"
)

func (r *PharmacistRepoImpl) AssignPharmacistBulk(ctx context.Context, pharmacy_id int64, ids ...int64) error {
	whereClause := make([]string, len(ids))
	for _, id := range ids {
		whereClause = append(whereClause, fmt.Sprintf("id=%d", id))
	}

	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("pharmacists").
		Set(
			ub.Assign("pharmacy_id", pharmacy_id),
			ub.Assign("updated_at", "NOW()"),
		).
		Where(
			ub.Or(whereClause...),
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

func (r *PharmacistRepoImpl) UnassignPharmacistFromPharmacy(ctx context.Context, pharmacyId int64) error {
	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("pharmacists").
		Set(
			ub.Assign("pharmacy_id", nil),
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
