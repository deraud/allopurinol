package repo

import (
	"context"
	"github.com/shopspring/decimal"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
)

func (r *AddressRepoImpl) GetAddress(ctx context.Context, id int64) (*entity.Address, error) {
	var address entity.Address

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select(
		"id",
		"user_id",
		"pharmacy_id",
		"province",
		"city",
		"district",
		"subdistrict",
		"postal_code",
		"name",
		"phone_number",
		"is_active",
		"ST_XMax(location::geometry)",
		"ST_YMax(location::geometry)",
	).
		From("addresses").
		Where(
			sb.Equal("id", id),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(
		&address.Id,
		&address.UserId,
		&address.PharmacyId,
		&address.Province,
		&address.City,
		&address.District,
		&address.Subdistrict,
		&address.PostalCode,
		&address.Name,
		&address.PhoneNumber,
		&address.IsActive,
		&address.Longitude,
		&address.Latitude,
	)
	if err != nil {
		return nil, err
	}

	return &address, nil
}

func (r *AddressRepoImpl) GetAddressByUserId(ctx context.Context, userId int64) (*entity.Address, error) {
	var address entity.Address

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select(
		"id",
		"user_id",
		"pharmacy_id",
		"province",
		"city",
		"district",
		"subdistrict",
		"postal_code",
		"name",
		"phone_number",
		"is_active",
		"ST_XMax(location::geometry)",
		"ST_YMax(location::geometry)",
	).
		From("addresses").
		Where(
			sb.Equal("user_id", userId),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(
		&address.Id,
		&address.UserId,
		&address.PharmacyId,
		&address.Province,
		&address.City,
		&address.District,
		&address.Subdistrict,
		&address.PostalCode,
		&address.Name,
		&address.PhoneNumber,
		&address.IsActive,
		&address.Longitude,
		&address.Latitude,
	)
	if err != nil {
		return nil, err
	}

	return &address, nil
}

func (r *AddressRepoImpl) GetDistance(ctx context.Context, fromLat, fromLong, toLat, toLong float64) (decimal.Decimal, error) {
	var distance decimal.Decimal

	query := `SELECT ST_DistanceSphere(ST_MakePoint($1, $2),ST_MakePoint($3, $4))`

	err := r.db.QueryRowContext(ctx, query, fromLat, fromLong, toLat, toLong).Scan(&distance)
	if err != nil {
		return decimal.Zero, err
	}

	return distance, nil
}
