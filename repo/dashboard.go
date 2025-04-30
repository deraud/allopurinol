package repo

import (
	"context"
	"healthcare-allopurinol/constant"

	"github.com/huandu/go-sqlbuilder"
)

// pharmacy pharmacist user

func (r *AdminRepoImpl) GetPharmacyCount(ctx context.Context) (int64, error) {
	var count int64

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("COUNT(id)").
		From("pharmacies").
		Where(sb.IsNull("deleted_at"))

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *AdminRepoImpl) GetPharmacistCount(ctx context.Context) (int64, error) {
	var count int64

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("COUNT(id)").
		From("pharmacists").
		Where(sb.IsNull("deleted_at"))

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *AdminRepoImpl) GetUserCount(ctx context.Context) (int64, error) {
	var count int64

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("COUNT(id)").
		From("users").
		Where(sb.IsNull("deleted_at"))

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
