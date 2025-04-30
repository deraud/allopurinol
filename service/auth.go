package service

import (
	"context"
	"healthcare-allopurinol/bridge"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/repo"
)

type AuthServiceItf interface {
	LoginService(ctx context.Context, email, password string) (string, error)
	RegisterService(ctx context.Context, cred *entity.Credential, user *entity.User) error
	ForgotPasswordService(ctx context.Context, email string) error
	ResetPasswordService(ctx context.Context, token, password string) error
	VerifyService(ctx context.Context, credId int64) error
	VerifyTokenService(ctx context.Context, token string) error
}

type AuthServiceImpl struct {
	repo       repo.AuthRepoItf
	userRepo   repo.UserRepoItf
	transactor repo.TransactorItf
	redis      repo.RedisRepoItf
	mailBridge bridge.MailBridgeItf
}

func NewAuthService(repo repo.AuthRepoItf, userRepo repo.UserRepoItf, transactor repo.TransactorItf, redis repo.RedisRepoItf, mailBridge bridge.MailBridgeItf) AuthServiceItf {
	return &AuthServiceImpl{repo: repo, userRepo: userRepo, transactor: transactor, redis: redis, mailBridge: mailBridge}
}
