package repo

import (
	"context"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
)

func (r *UserRepoImpl) UpdateAddress(ctx context.Context, a *entity.UserAddress) (*entity.UserAddress, error) {
	address := *a
	query := `
		UPDATE addresses
		SET 
			province = $3,
			city = $4,
			district = $5,
			subdistrict = $6,
			postal_code = $7,
			name = $8, 
			phone_number = $9,
			location = (ST_SetSRID(ST_MakePoint($10, $11), 4326)),
			updated_at = NOW()
		WHERE
			user_id = $1 and id = $2
		RETURNING 
			id
	`
	args := []any{a.UserId, *a.Id, a.Province, a.City, a.District, a.Subdistrict, a.PostalCode, a.Name, a.PhoneNumber, a.Longitude, a.Latitude}
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&address.Id)

	if err != nil {
		return nil, err
	}

	return &address, nil
}

func (r *UserRepoImpl) UpdateUserProfile(ctx context.Context, u *entity.UserProfile) (*entity.UserProfile, error) {
	user := *u

	ub := sqlbuilder.NewUpdateBuilder()

	var values = make([]string, 0)
	if u.Name != "" {
		values = append(values, ub.Assign("name", u.Name))
	}
	if u.ProfilePicture != "" {
		values = append(values, ub.Assign("profile_picture", u.ProfilePicture))
	}
	if u.Gender != nil {
		values = append(values, ub.Assign("gender", u.Gender))
	}
	values = append(values, ub.Assign("updated_at", "NOW()"))
	ub.Update("users").
		Set(values...).
		Where(ub.Equal("id", u.Id)).
		SQL("RETURNING id, name, profile_picture, gender, is_verified")

	query, args := ub.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&user.Id, &user.Name, &user.ProfilePicture, &user.Gender, &user.IsVerified)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepoImpl) UpdateUserActivateAddress(ctx context.Context, add *entity.UserAddress) (*entity.UserAddress, error) {
	address := *add
	ub := sqlbuilder.NewUpdateBuilder()

	ub.Update("addresses").
		Set(
			ub.Assign("is_active", true),
			ub.Assign("updated_at", "NOW()"),
		).
		Where(
			ub.Equal("id", *add.Id),
			ub.Equal("user_id", add.UserId),
			ub.IsNull("deleted_at")).
		SQL("RETURNING is_active")

	query, args := ub.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&address.IsActive)
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func (r *UserRepoImpl) UpdateUserRemovePicture(ctx context.Context, u *entity.UserProfile) (*entity.UserProfile, error) {
	user := *u
	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("users").
		Set(ub.Assign("profile_picture", constant.DEFAULT_PROFILE_PICTURE)).
		Where(ub.Equal("id", u.Id)).
		SQL("RETURNING id, name, profile_picture, gender, is_verified")

	query, args := ub.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&user.Id, &user.Name, &user.ProfilePicture, &user.Gender, &user.IsVerified)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
