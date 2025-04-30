package service

import (
	"context"
	"healthcare-allopurinol/bridge"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/repo"
)

type PharmacistServiceItf interface {
	CreatePharmacistService(ctx context.Context, c *entity.Credential, p *entity.Pharmacist) (*entity.Pharmacist, *entity.Credential, error)
	UpdatePharmacistService(ctx context.Context, p *entity.Pharmacist) (*entity.Pharmacist, error)
	DeletePharmacistService(ctx context.Context, id int64) error
	GetPharmacistsService(ctx context.Context, options *entity.PharmacistOptions) ([]*entity.Pharmacist, []string, *entity.PharmacistOptions, error)
	GetPharmacistService(ctx context.Context, id int64) (*entity.Pharmacist, string, error)
}

type PharmacistServiceImpl struct {
	transactorRepo repo.TransactorItf
	authRepo       repo.AuthRepoItf
	pharmacistRepo repo.PharmacistRepoItf
	pharmacyRepo   repo.PharmacyRepoItf
	mailBridge     bridge.MailBridgeItf
}

func NewPharmacistService(transactorRepo repo.TransactorItf, authRepo repo.AuthRepoItf, pharmacistRepo repo.PharmacistRepoItf, pharmacyRepo repo.PharmacyRepoItf, mailBridge bridge.MailBridgeItf) PharmacistServiceItf {
	return &PharmacistServiceImpl{
		transactorRepo: transactorRepo,
		authRepo:       authRepo,
		pharmacistRepo: pharmacistRepo,
		pharmacyRepo:   pharmacyRepo,
		mailBridge:     mailBridge,
	}
}
