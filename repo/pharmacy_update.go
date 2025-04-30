package repo

import (
	"context"
	"fmt"
	"healthcare-allopurinol/entity"
)

func (r *PharmacyRepoImpl) UpdatePharmacy(ctx context.Context, p *entity.Pharmacy) (*entity.Pharmacy, error) {
	pharmacy := *p

	lpIds := ""
	for i, _ := range p.LogisticPartnerIds {
		if i != 0 {
			lpIds += ","
		}
		lpIds += fmt.Sprintf("$%d", i+7)
	}

	query := fmt.Sprintf(`
		UPDATE 
			pharmacies 
		SET 
			logo = $1, 
			name = $2, 
			is_active = $3, 
			updated_at = $4,
			location = (ST_SetSRID(ST_MakePoint($5, $6), 4326)),
			days = (SELECT MIN(days) FROM logistic_partners WHERE id IN (%s))
		WHERE 
			id = $%d AND 
			deleted_at IS NULL 
		RETURNING 
			partner_id
	`, lpIds, 7+len(p.LogisticPartnerIds))

	tx := extractTx(ctx)
	args := []any{p.Logo, p.Name, p.IsActive, "NOW()", p.Address.Longitude, p.Address.Latitude}
	for _, id := range p.LogisticPartnerIds {
		args = append(args, id)
	}
	args = append(args, p.Id)
	err := tx.QueryRowContext(ctx, query, args...).Scan(&pharmacy.PartnerId)
	if err != nil {
		return nil, err
	}

	return &pharmacy, nil
}
