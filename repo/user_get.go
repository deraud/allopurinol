package repo

import (
	"context"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
)

func (r *UserRepoImpl) GetUserProfile(ctx context.Context, id int64) (*entity.UserProfile, error) {
	var user entity.UserProfile

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("u.id", "u.name", "u.profile_picture", "c.email", "u.gender", "u.is_verified").
		From("users u").
		Join("credentials c", "u.credential_id = c.id").
		Where(sb.Equal("u.id", id))
	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(
		&user.Id,
		&user.Name,
		&user.ProfilePicture,
		&user.Email,
		&user.Gender,
		&user.IsVerified,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepoImpl) GetUserAddresses(ctx context.Context, id int64, count int64, options *entity.UserAddressOptions) ([]*entity.UserAddress, error) {
	userAddresses := make([]*entity.UserAddress, count)
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "province", "city", "district", "subdistrict", "postal_code", "name", "phone_number", "ST_XMax(location::geometry)", "ST_YMax(location::geometry)", "is_active").
		From("addresses").
		Where(
			sb.Equal("user_id", id),
		).
		AddWhereClause(r.generateGetUserAddressesWhereClause(options)).
		SQL("ORDER BY is_active desc")

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for i := 0; rows.Next(); i++ {
		var userAddress entity.UserAddress
		err := rows.Scan(
			&userAddress.Id,
			&userAddress.Province,
			&userAddress.City,
			&userAddress.District,
			&userAddress.Subdistrict,
			&userAddress.PostalCode,
			&userAddress.Name,
			&userAddress.PhoneNumber,
			&userAddress.Longitude,
			&userAddress.Latitude,
			&userAddress.IsActive,
		)

		if err != nil {
			return nil, err
		}
		userAddresses[i] = &userAddress
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return userAddresses, err
}

func (r *UserRepoImpl) GetUserAddressesCount(ctx context.Context, id int64, options *entity.UserAddressOptions) (int64, error) {
	var count int64

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("COUNT(id)").
		From("addresses").
		Where(
			sb.Equal("user_id", id),
		).
		AddWhereClause(r.generateGetUserAddressesWhereClause(options))
	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *UserRepoImpl) generateGetUserAddressesWhereClause(options *entity.UserAddressOptions) *sqlbuilder.WhereClause {
	whereClause := sqlbuilder.NewWhereClause()
	cond := sqlbuilder.NewCond()
	whereClause.AddWhereExpr(cond.Args, cond.IsNull("deleted_at"))
	if *options.IsActive == "true" {
		whereClause.AddWhereExpr(cond.Args, cond.Equal("is_active", *options.IsActive))
	}
	return whereClause
}
