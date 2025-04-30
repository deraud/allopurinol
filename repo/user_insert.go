package repo

import (
	"context"
	"healthcare-allopurinol/entity"
)

func (r *UserRepoImpl) InsertAddress(ctx context.Context, address *entity.UserAddress) (*entity.UserAddress, error) {
	add := *address
	query := `
		INSERT INTO addresses
			(user_id, province, city, district, subdistrict, postal_code, name, phone_number, location)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, (ST_SetSRID(ST_MakePoint($9,$10), 4326)))
		RETURNING
			id
	`
	tx := extractTx(ctx)
	args := []any{add.UserId, add.Province, add.City, add.District, add.Subdistrict, add.PostalCode, add.Name, add.PhoneNumber, add.Longitude, add.Latitude}
	err := tx.QueryRowContext(ctx, query, args...).Scan(&add.Id)
	if err != nil {
		return nil, err
	}

	return &add, nil
}
