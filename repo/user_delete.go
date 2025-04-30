package repo

import (
	"context"
	"healthcare-allopurinol/constant"

	"github.com/huandu/go-sqlbuilder"
)

func (r *UserRepoImpl) DeleteAddress(ctx context.Context, userId int64, addressId int64) error {

	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("addresses").
		Set(
			ub.Assign("deleted_at", "NOW()"),
			ub.Assign("updated_at", "NOW()"),
		).
		Where(
			ub.Equal("user_id", userId),
			ub.Equal("id", addressId),
			ub.IsNull("deleted_at"),
		)

	query, args := ub.BuildWithFlavor(constant.DATABASE_FLAVOR)
	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}
