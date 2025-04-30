package repo

import (
	"context"
	"fmt"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
)

var (
	adminSb = sqlbuilder.NewSelectBuilder().
		Select("c.id", "c.role_id", "c.email", "a.name").
		From("credentials c").
		Join("admins a", "c.id = a.credential_id").
		Where("a.deleted_at IS NULL")

	userSb = sqlbuilder.NewSelectBuilder().
		Select("c.id", "c.role_id", "c.email", "u.name").
		From("credentials c").
		Join("users u", "c.id = u.credential_id").
		Where("u.deleted_at IS NULL")

	pharmacistSb = sqlbuilder.NewSelectBuilder().
			Select("c.id", "c.role_id", "c.email", "p.name").
			From("credentials c").
			Join("pharmacists p", "c.id = p.credential_id").
			Where("p.deleted_at IS NULL")

	ub = sqlbuilder.Union(adminSb, userSb, pharmacistSb)
)

func (r *AdminRepoImpl) GetUsers(ctx context.Context, options *entity.UserInfoOptions) ([]*entity.UserInfo, error) {
	size := options.Limit
	remaining := options.TotalRows - (options.Page-1)*options.Limit
	if remaining <= 0 {
		size = 0
	} else if remaining < options.Limit {
		size = remaining
	}
	users := make([]*entity.UserInfo, size)

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("user_data.id", "user_data.email", "user_data.name", "r.name AS role").
		From("roles r").
		Join(
			sb.BuilderAs(ub, "user_data"),
			"r.id = user_data.role_id",
		).
		OrderBy("user_data.id").
		Where(
			sb.Like(fmt.Sprintf("user_data.%s", options.SearchBy), fmt.Sprintf("%%%s%%", options.SearchValue)),
			sb.IsNull("r.deleted_at"),
		).
		Offset((options.Page - 1) * options.Limit).
		Limit(options.Limit)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var user entity.UserInfo
		err := rows.Scan(
			&user.Id,
			&user.Email,
			&user.Name,
			&user.Role,
		)
		if err != nil {
			return nil, err
		}
		users[i] = &user
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *AdminRepoImpl) GetUsersCount(ctx context.Context, options *entity.UserInfoOptions) (int, error) {
	var count int

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("COUNT(user_data.id)").
		From("roles r").
		Join(
			sb.BuilderAs(ub, "user_data"),
			"r.id = user_data.role_id",
		).
		Where(
			sb.Like(fmt.Sprintf("user_data.%s", options.SearchBy), fmt.Sprintf("%%%s%%", options.SearchValue)),
			sb.IsNull("r.deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
