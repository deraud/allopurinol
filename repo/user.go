package repo

import (
	"context"
	"database/sql"
	"errors"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
)

type UserRepoItf interface {
	InsertUser(ctx context.Context, u *entity.User) error
	GetUserByCredId(ctx context.Context, credId int64) (*entity.User, error)
	GetUserProfileByCredId(ctx context.Context, credId int64) (*entity.UserProfile, error)
	GetUserAddresses(ctx context.Context, id int64, count int64, options *entity.UserAddressOptions) ([]*entity.UserAddress, error)
	GetUserAddressesCount(ctx context.Context, id int64, options *entity.UserAddressOptions) (int64, error)
	generateGetUserAddressesWhereClause(options *entity.UserAddressOptions) *sqlbuilder.WhereClause
	GetUserProfile(ctx context.Context, id int64) (*entity.UserProfile, error)
	DeleteAddress(ctx context.Context, userId int64, addressId int64) error
	IsUserAddressExist(ctx context.Context, userId int64, addressId int64) (bool, error)
	IsUserAddressesExist(ctx context.Context, userId int64) (bool, error)
	IsUserAddressActive(ctx context.Context, userId int64, addressId int64) (bool, error)
	InsertAddress(ctx context.Context, address *entity.UserAddress) (*entity.UserAddress, error)
	UpdateAddress(ctx context.Context, a *entity.UserAddress) (*entity.UserAddress, error)
	UpdateUserProfile(ctx context.Context, u *entity.UserProfile) (*entity.UserProfile, error)
	UpdateUserRemovePicture(ctx context.Context, u *entity.UserProfile) (*entity.UserProfile, error)
	UpdateUserActivateAddress(ctx context.Context, add *entity.UserAddress) (*entity.UserAddress, error)
	DeactivateUserAddresses(ctx context.Context, add *entity.UserAddress) error
	IsUserVerified(ctx context.Context, userCredId int64) (bool, error)
}

type UserRepoImpl struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepoItf {
	return &UserRepoImpl{db: db}
}

func (r *UserRepoImpl) InsertUser(ctx context.Context, u *entity.User) error {
	user := *u

	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("users").
		Cols("credential_id", "name", "gender", "profile_picture").
		Values(u.CredId, u.Name, constant.DEFAULT_GENDER, constant.DEFAULT_PROFILE_PICTURE).
		SQL("RETURNING id")

	tx := extractTx(ctx)
	query, args := ib.BuildWithFlavor(sqlbuilder.PostgreSQL)
	err := tx.QueryRowContext(ctx, query, args...).Scan(&user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepoImpl) GetUserByCredId(ctx context.Context, credId int64) (*entity.User, error) {
	user := &entity.User{CredId: credId}

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "name", "gender", "profile_picture", "is_verified").
		From("users").
		Where(
			sb.Equal("credential_id", credId),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(
		&user.Id,
		&user.Name,
		&user.Gender,
		&user.ProfilePicture,
		&user.IsVerified,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepoImpl) GetUserProfileByCredId(ctx context.Context, credId int64) (*entity.UserProfile, error) {
	user := &entity.UserProfile{CredId: credId}

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "name", "gender", "profile_picture", "is_verified").
		From("users").
		Where(
			sb.Equal("credential_id", credId),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(
		&user.Id,
		&user.Name,
		&user.Gender,
		&user.ProfilePicture,
		&user.IsVerified,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepoImpl) IsUserAddressesExist(ctx context.Context, userId int64) (bool, error) {
	var exist bool
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("1").
		From("addresses").
		Where(
			sb.Equal("user_id", userId),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&exist)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (r *UserRepoImpl) IsUserAddressExist(ctx context.Context, userId int64, addressId int64) (bool, error) {
	var exist bool
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("1").
		From("addresses").
		Where(
			sb.Equal("user_id", userId),
			sb.Equal("id", addressId),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&exist)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (r *UserRepoImpl) IsUserAddressActive(ctx context.Context, userId int64, addressId int64) (bool, error) {
	var active bool
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("1").
		From("addresses").
		Where(
			sb.Equal("user_id", userId),
			sb.Equal("id", addressId),
			sb.Equal("is_active", true),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(active)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return active, nil
}

func (r *UserRepoImpl) DeactivateUserAddresses(ctx context.Context, add *entity.UserAddress) error {
	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("addresses").
		Set(
			ub.Assign("is_active", false),
			ub.Assign("updated_at", "NOW()"),
		).
		Where(
			ub.Equal("user_id", add.UserId),
			ub.IsNull("deleted_at"),
			ub.NotEqual("id", add.Id),
		)

	query, args := ub.BuildWithFlavor(constant.DATABASE_FLAVOR)
	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepoImpl) IsUserVerified(ctx context.Context, userCredId int64) (bool, error) {
	var verified bool
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("is_verified").
		From("users").
		Where(
			sb.Equal("credential_id", userCredId),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&verified)
	if err != nil {
		return false, err
	}

	return verified, nil
}
