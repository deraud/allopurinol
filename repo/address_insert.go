package repo

import (
	"context"
	"healthcare-allopurinol/entity"
)

func (r *AddressRepoImpl) InsertAddress(ctx context.Context, a *entity.Address) (*entity.Address, error) {
	address := *a

	query := `
		INSERT INTO addresses
			(user_id, pharmacy_id, province, city, district, subdistrict, postal_code, name, phone_number, location)
		VALUES 
			($1, $2, $3, $4, $5, $6, $7, $8, $9, (ST_SetSRID(ST_MakePoint($10,$11), 4326)))
		RETURNING 
			id, is_active
	`
	args := []any{a.UserId, a.PharmacyId, a.Province, a.City, a.District, a.Subdistrict, a.PostalCode, a.Name, a.PhoneNumber, a.Longitude, a.Latitude}

	tx := extractTx(ctx)
	err := tx.QueryRowContext(ctx, query, args...).Scan(&address.Id, &address.IsActive)
	if err != nil {
		return nil, err
	}

	return &address, nil
}
