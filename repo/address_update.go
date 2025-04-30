package repo

import (
	"context"
	"healthcare-allopurinol/entity"
)

func (r *AddressRepoImpl) UpdateAddressPharmacy(ctx context.Context, a *entity.Address) (*entity.Address, error) {
	address := *a

	query := `
		UPDATE addresses
		SET 
			pharmacy_id = $1,
			province = $2,
			city = $3,
			district = $4,
			subdistrict = $5,
			postal_code = $6,
			name = $7, 
			phone_number = $8,
			location = (ST_SetSRID(ST_MakePoint($9,$10), 4326)),
			updated_at = NOW()
		WHERE
			pharmacy_id = $11
		RETURNING 
			id, is_active
	`
	args := []any{*a.PharmacyId, a.Province, a.City, a.District, a.Subdistrict, a.PostalCode, a.Name, a.PhoneNumber, a.Longitude, a.Latitude, *a.PharmacyId}
	tx := extractTx(ctx)
	err := tx.QueryRowContext(ctx, query, args...).Scan(&address.Id, &address.IsActive)
	if err != nil {
		return nil, err
	}

	return &address, nil
}
