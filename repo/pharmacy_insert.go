package repo

import (
	"context"
	"fmt"
	"healthcare-allopurinol/entity"
)

func (r *PharmacyRepoImpl) InsertPharmacy(ctx context.Context, p *entity.Pharmacy) (*entity.Pharmacy, error) {
	pharmacy := *p

	lpIds := ""
	for i, _ := range p.LogisticPartnerIds {
		if i != 0 {
			lpIds += ","
		}
		lpIds += fmt.Sprintf("$%d", i+6)
	}

	query := fmt.Sprintf(`
		INSERT INTO pharmacies
			(partner_id, logo, name, location, days)
		VALUES 
			($1, $2, $3, (ST_SetSRID(ST_MakePoint($4, $5), 4326)), (SELECT MIN(days) FROM logistic_partners WHERE id IN (%s)))
		RETURNING 
			id, is_active
	`, lpIds)

	tx := extractTx(ctx)
	args := []any{p.PartnerId, p.Logo, p.Name, p.Address.Longitude, p.Address.Latitude}
	for _, id := range p.LogisticPartnerIds {
		args = append(args, id)
	}
	err := tx.QueryRowContext(ctx, query, args...).Scan(&pharmacy.Id, &pharmacy.IsActive)
	if err != nil {
		return nil, err
	}

	return &pharmacy, nil
}
